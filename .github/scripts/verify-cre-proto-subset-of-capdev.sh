#!/usr/bin/env bash
# Verify PR cre/ proto changes are already present on capabilities-development.
#
# Usage (from repo root, on a PR branch):
#   bash ./.github/scripts/verify-cre-proto-subset-of-capdev.sh

set -euo pipefail

BASE_BRANCH="${BASE_BRANCH:-main}"
CAP_DEV_BRANCH="${CAP_DEV_BRANCH:-capabilities-development}"
# When set to 1 (CI fallback after patch-id failure), fail if the PR has no proto changes.
REQUIRE_PROTO_CHANGES="${REQUIRE_PROTO_CHANGES:-0}"

git fetch origin "${BASE_BRANCH}" "${CAP_DEV_BRANCH}" --quiet

PROTO_FILES=$(git diff --name-only "origin/${BASE_BRANCH}...HEAD" -- 'cre/**/*.proto' || true)

if [[ -z "${PROTO_FILES}" ]]; then
  if [[ "${REQUIRE_PROTO_CHANGES}" == "1" ]]; then
    echo "::error::Patch-id check failed and PR has no cre/**/*.proto changes for subset fallback."
    exit 1
  fi
  echo "No cre/ proto files changed in PR. Subset check skipped."
  exit 0
fi

echo "Checking proto file(s) against origin/${CAP_DEV_BRANCH}..."
echo ""

WORKDIR=$(mktemp -d)
trap 'rm -rf "$WORKDIR"' EXIT

fail() {
  echo "::error::$1"
  exit 1
}

extract_block() {
  local file=$1 kind=$2 name=$3
  awk -v kind="$kind" -v name="$name" '
    $1 == kind && $2 == name {
      depth = 0
      in_block = 1
    }
    in_block {
      print
      if ($0 ~ /{/) depth++
      if ($0 ~ /}/) {
        depth--
        if (depth == 0) { exit }
      }
    }
  ' "$file"
}

field_lines() {
  grep -E '^\s*(repeated\s+)?[A-Za-z0-9_.]+\s+[A-Za-z0-9_]+\s*=\s*[0-9]+;' || true
}

check_proto_subset() {
  local rel=$1
  local pr_file=$2
  local cap_file=$3

  echo "  file: ${rel}"

  while IFS= read -r rpc_line; do
    [[ -z "$rpc_line" ]] && continue
    grep -qF "$rpc_line" "$cap_file" || fail "${rel}: RPC missing on ${CAP_DEV_BRANCH}: ${rpc_line}"
  done <<EOF
$(grep -E '^\s+rpc ' "$pr_file" || true)
EOF

  while IFS= read -r enum_name; do
    [[ -z "$enum_name" ]] && continue
    extract_block "$cap_file" enum "$enum_name" >/dev/null || \
      fail "${rel}: enum missing on ${CAP_DEV_BRANCH}: ${enum_name}"

    while IFS= read -r val_line; do
      [[ -z "$val_line" ]] && continue
      grep -qF "$val_line" "$cap_file" || \
        fail "${rel}: enum value missing on ${CAP_DEV_BRANCH} (${enum_name}): ${val_line}"
    done <<EOF
$(extract_block "$pr_file" enum "$enum_name" | grep -E '^\s*[A-Z0-9_]+\s*=' || true)
EOF
  done <<EOF
$(grep -E '^enum ' "$pr_file" | awk '{print $2}' || true)
EOF

  while IFS= read -r msg_name; do
    [[ -z "$msg_name" ]] && continue
    extract_block "$cap_file" message "$msg_name" >/dev/null || \
      fail "${rel}: message missing on ${CAP_DEV_BRANCH}: ${msg_name}"

    while IFS= read -r field_line; do
      [[ -z "$field_line" ]] && continue
      grep -qF "$field_line" "$cap_file" || \
        fail "${rel}: field missing on ${CAP_DEV_BRANCH} (${msg_name}): ${field_line}"
    done <<EOF
$(extract_block "$pr_file" message "$msg_name" | field_lines)
EOF
  done <<EOF
$(grep -E '^message ' "$pr_file" | awk '{print $2}' || true)
EOF

  echo "  OK"
}

while IFS= read -r rel; do
  [[ -z "$rel" ]] && continue

  if ! git cat-file -e "origin/${CAP_DEV_BRANCH}:${rel}" 2>/dev/null; then
    fail "${rel} does not exist on origin/${CAP_DEV_BRANCH}"
  fi

  git show "HEAD:${rel}" > "${WORKDIR}/pr.proto"
  git show "origin/${CAP_DEV_BRANCH}:${rel}" > "${WORKDIR}/cap.proto"

  check_proto_subset "$rel" "${WORKDIR}/pr.proto" "${WORKDIR}/cap.proto"
  echo ""
done <<EOF
${PROTO_FILES}
EOF

echo "All PR proto changes are present on origin/${CAP_DEV_BRANCH}."
exit 0
