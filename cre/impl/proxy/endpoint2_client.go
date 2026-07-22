package proxy

import (
	"context"
	"fmt"
	"io"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smartcontractkit/libocr/commontypes"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// ProxyEndpoint2Factory is an ocr2types.BinaryNetworkEndpoint2Factory (OCR3.1)
// that delegates to a remote Endpoint2Proxy server.
type ProxyEndpoint2Factory struct {
	peerID string
	client Endpoint2ProxyClient
	conn   *grpc.ClientConn
}

// ClosableBinaryNetworkEndpoint2Factory is a BinaryNetworkEndpoint2Factory
// whose underlying connection can be released.
type ClosableBinaryNetworkEndpoint2Factory interface {
	ocr2types.BinaryNetworkEndpoint2Factory
	io.Closer
}

var _ ClosableBinaryNetworkEndpoint2Factory = (*ProxyEndpoint2Factory)(nil)

// NewProxyEndpoint2Factory dials the proxy at proxyAddr and returns a factory
// that creates remote-backed OCR3.1 endpoints.
func NewProxyEndpoint2Factory(peerID, proxyAddr string) (ClosableBinaryNetworkEndpoint2Factory, error) {
	conn, err := grpc.NewClient(proxyAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to proxy: %w", err)
	}
	return &ProxyEndpoint2Factory{
		peerID: peerID,
		client: NewEndpoint2ProxyClient(conn),
		conn:   conn,
	}, nil
}

func (f *ProxyEndpoint2Factory) PeerID() string {
	return f.peerID
}

func (f *ProxyEndpoint2Factory) Close() error {
	return f.conn.Close()
}

func (f *ProxyEndpoint2Factory) NewEndpoint(
	cd ocr2types.ConfigDigest,
	peerIDs []string,
	v2bootstrappers []commontypes.BootstrapperLocator,
	defaultPriorityConfig ocr2types.BinaryNetworkEndpoint2Config,
	lowPriorityConfig ocr2types.BinaryNetworkEndpoint2Config,
) (ocr2types.BinaryNetworkEndpoint2, error) {
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := f.client.Connect(ctx)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to open endpoint2 connection: %w", err)
	}

	pbBootstrappers := make([]*BootstrapperLocator, len(v2bootstrappers))
	for i, b := range v2bootstrappers {
		pbBootstrappers[i] = &BootstrapperLocator{PeerId: b.PeerID, Addrs: b.Addrs}
	}

	if err := stream.Send(&Endpoint2ClientRequest{
		Message: &Endpoint2ClientRequest_NewEndpoint{
			NewEndpoint: &NewEndpoint2Request{
				ConfigDigest:          cd[:],
				PeerIds:               peerIDs,
				V2Bootstrappers:       pbBootstrappers,
				DefaultPriorityConfig: endpoint2ConfigToPB(defaultPriorityConfig),
				LowPriorityConfig:     endpoint2ConfigToPB(lowPriorityConfig),
			},
		},
	}); err != nil {
		cancel()
		return nil, fmt.Errorf("failed to send new endpoint2 request: %w", err)
	}

	e := &proxyEndpoint2{
		stream:   stream,
		cancel:   cancel,
		recvChan: make(chan ocr2types.InboundBinaryMessageWithSender, defaultRecvBufferSize),
		sendChan: make(chan *Endpoint2ClientRequest, defaultRecvBufferSize),
	}
	e.wg.Add(2)
	go e.receiveLoop()
	go e.sendLoop()
	return e, nil
}

func endpoint2ConfigToPB(c ocr2types.BinaryNetworkEndpoint2Config) *Endpoint2Config {
	pb := &Endpoint2Config{
		Limits: &BinaryNetworkEndpointLimits{
			MaxMessageLength:          int32(c.MaxMessageLength),
			MessagesRatePerOracle:     c.MessagesRatePerOracle,
			MessagesCapacityPerOracle: int32(c.MessagesCapacityPerOracle),
			BytesRatePerOracle:        c.BytesRatePerOracle,
			BytesCapacityPerOracle:    int32(c.BytesCapacityPerOracle),
		},
	}
	if c.OverrideIncomingMessageBufferSize != nil {
		v := int32(*c.OverrideIncomingMessageBufferSize)
		pb.OverrideIncomingMessageBufferSize = &v
	}
	if c.OverrideOutgoingMessageBufferSize != nil {
		v := int32(*c.OverrideOutgoingMessageBufferSize)
		pb.OverrideOutgoingMessageBufferSize = &v
	}
	return pb
}

// proxyEndpoint2 implements ocr2types.BinaryNetworkEndpoint2 over a Connect stream.
type proxyEndpoint2 struct {
	stream Endpoint2Proxy_ConnectClient
	cancel context.CancelFunc

	recvChan chan ocr2types.InboundBinaryMessageWithSender
	sendChan chan *Endpoint2ClientRequest

	closeOnce sync.Once
	wg        sync.WaitGroup
}

var _ ocr2types.BinaryNetworkEndpoint2 = (*proxyEndpoint2)(nil)

func (e *proxyEndpoint2) SendTo(msg ocr2types.OutboundBinaryMessage, to commontypes.OracleID) {
	pb, ok := outboundToPB(msg)
	if !ok {
		return
	}
	e.enqueue(&Endpoint2ClientRequest{
		Message: &Endpoint2ClientRequest_SendTo{
			SendTo: &Endpoint2Send{ToOracleId: uint32(to), Msg: pb},
		},
	})
}

func (e *proxyEndpoint2) Broadcast(msg ocr2types.OutboundBinaryMessage) {
	pb, ok := outboundToPB(msg)
	if !ok {
		return
	}
	e.enqueue(&Endpoint2ClientRequest{
		Message: &Endpoint2ClientRequest_Broadcast{
			Broadcast: &Endpoint2Send{Msg: pb},
		},
	})
}

func (e *proxyEndpoint2) enqueue(req *Endpoint2ClientRequest) {
	select {
	case e.sendChan <- req:
	default:
		// Buffer full: drop, matching the lenient best-effort delivery of the
		// underlying networking stack.
	}
}

func (e *proxyEndpoint2) Receive() <-chan ocr2types.InboundBinaryMessageWithSender {
	return e.recvChan
}

func (e *proxyEndpoint2) Close() error {
	e.closeOnce.Do(func() {
		_ = e.stream.CloseSend()
		e.cancel()
		e.wg.Wait()
		close(e.recvChan)
	})
	return nil
}

func (e *proxyEndpoint2) sendLoop() {
	defer e.wg.Done()
	for {
		select {
		case <-e.stream.Context().Done():
			return
		case req := <-e.sendChan:
			if err := e.stream.Send(req); err != nil {
				return
			}
		}
	}
}

func (e *proxyEndpoint2) receiveLoop() {
	defer e.wg.Done()
	for {
		msg, err := e.stream.Recv()
		if err != nil {
			return
		}
		inbound := pbToInbound(e, msg)
		if inbound == nil {
			continue
		}
		select {
		case e.recvChan <- ocr2types.InboundBinaryMessageWithSender{
			InboundBinaryMessage: inbound,
			Sender:               commontypes.OracleID(msg.Sender),
		}:
		case <-e.stream.Context().Done():
			return
		default:
			// Buffer full: drop.
		}
	}
}

// outboundToPB converts an ocr2types.OutboundBinaryMessage to its wire form.
func outboundToPB(msg ocr2types.OutboundBinaryMessage) (*OutboundMessage2, bool) {
	switch m := msg.(type) {
	case ocr2types.OutboundBinaryMessagePlain:
		return &OutboundMessage2{
			Payload:  m.Payload,
			Priority: uint32(m.Priority),
			Kind:     &OutboundMessage2_Plain{Plain: &OutboundPlain2{}},
		}, true
	case ocr2types.OutboundBinaryMessageRequest:
		req := &OutboundRequest2{}
		if p, ok := m.ResponsePolicy.(ocr2types.SingleUseSizedLimitedResponsePolicy); ok {
			req.PolicyMaxSize = int64(p.MaxSize)
			req.PolicyExpiryUnixMs = p.ExpiryTimestamp.UnixMilli()
		}
		return &OutboundMessage2{
			Payload:  m.Payload,
			Priority: uint32(m.Priority),
			Kind:     &OutboundMessage2_Request{Request: req},
		}, true
	case ocr2types.OutboundBinaryMessageResponse:
		h, ok := ocr2types.MustGetOutboundBinaryMessageResponseRequestHandle(m).(*proxyRequestHandle2)
		if !ok {
			// A response built from a handle we didn't issue cannot be routed.
			return nil, false
		}
		return &OutboundMessage2{
			Payload:  m.Payload,
			Priority: uint32(m.Priority),
			Kind:     &OutboundMessage2_Response{Response: &OutboundResponse2{RequestId: h.requestID}},
		}, true
	default:
		return nil, false
	}
}

// pbToInbound converts a wire message to an ocr2types.InboundBinaryMessage.
func pbToInbound(e *proxyEndpoint2, msg *Endpoint2ServerMessage) ocr2types.InboundBinaryMessage {
	priority := ocr2types.BinaryMessageOutboundPriority(msg.Priority)
	switch k := msg.Kind.(type) {
	case *Endpoint2ServerMessage_Plain:
		return ocr2types.InboundBinaryMessagePlain{Payload: msg.Payload, Priority: priority}
	case *Endpoint2ServerMessage_Request:
		return ocr2types.InboundBinaryMessageRequest{
			RequestHandle: &proxyRequestHandle2{requestID: k.Request.RequestId, priority: priority},
			Payload:       msg.Payload,
			Priority:      priority,
		}
	case *Endpoint2ServerMessage_Response:
		return ocr2types.InboundBinaryMessageResponse{Payload: msg.Payload, Priority: priority}
	default:
		return nil
	}
}

// proxyRequestHandle2 is the client-side stand-in for a libocr RequestHandle
// held by the proxy server; responses are routed back by requestID. The
// response priority must match the request priority (a ragep2p requirement),
// so it is captured here.
type proxyRequestHandle2 struct {
	requestID uint64
	priority  ocr2types.BinaryMessageOutboundPriority
}

var _ ocr2types.RequestHandle = (*proxyRequestHandle2)(nil)

func (h *proxyRequestHandle2) MakeResponse(payload []byte) ocr2types.OutboundBinaryMessageResponse {
	return ocr2types.MustMakeOutboundBinaryMessageResponse(h, payload, h.priority)
}
