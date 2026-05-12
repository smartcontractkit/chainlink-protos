package data_feeds

//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./bridge_status/v1/bridge_status_event.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./bridge_status/v1/job_info.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./bridge_status/v1/runtime_info.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./bridge_status/v1/metrics_info.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./bridge_status/v1/endpoint_info.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./bridge_status/v1/configuration_item.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./job_spec/v1/job_spec_event.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./job_spec/v1/ocr2_oracle_spec_info.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./job_spec/v1/ocr2_evm_relay_config.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./job_spec/v1/ocr2_median_plugin_config.proto
