#!/usr/bin/env bash
# Install tuido from the latest GitHub release (Linux amd64/arm64).
set -euo pipefail

REPO="${TUIDO_REPO:-gsusgit/tuido}"
INSTALL_DIR="${INSTALL_DIR:-${HOME}/.local/bin}"

if [[ "$(uname -s)" != Linux ]]; then
	echo "error: install-release.sh supports Linux only" >&2
	exit 1
fi

case "$(uname -m)" in
x86_64 | amd64) GOARCH=amd64 ;;
aarch64 | arm64) GOARCH=arm64 ;;
*)
	echo "error: unsupported architecture: $(uname -m)" >&2
	exit 1
	;;
esac

api="https://api.github.com/repos/${REPO}/releases/latest"
if command -v jq >/dev/null 2>&1; then
	TAG=$(curl -fsSL "${api}" | jq -r .tag_name)
else
	TAG=$(curl -fsSL "${api}" | grep -m1 '"tag_name"' | sed -E 's/.*"tag_name"[[:space:]]*:[[:space:]]*"([^"]+)".*/\1/')
fi

if [[ -z "${TAG}" || "${TAG}" == "null" ]]; then
	echo "error: could not resolve latest release for ${REPO}" >&2
	exit 1
fi

VERSION="${TAG#v}"
ASSET="tuido_${VERSION}_linux_${GOARCH}.tar.gz"
BASE="https://github.com/${REPO}/releases/download/${TAG}"
PKG_DIR="tuido_${VERSION}_linux_${GOARCH}"

tmpdir=$(mktemp -d)
trap 'rm -rf "${tmpdir}"' EXIT

curl -fsSL -o "${tmpdir}/checksums.txt" "${BASE}/checksums.txt"
curl -fsSL -o "${tmpdir}/${ASSET}" "${BASE}/${ASSET}"

(
	cd "${tmpdir}"
	grep -F "${ASSET}" checksums.txt | sha256sum -c -
)

tar -xzf "${tmpdir}/${ASSET}" -C "${tmpdir}"
mkdir -p "${INSTALL_DIR}"
install -Dm755 "${tmpdir}/${PKG_DIR}/tuido" "${INSTALL_DIR}/tuido"

echo "Installed: ${INSTALL_DIR}/tuido (${TAG}, linux/${GOARCH})"
echo "Run: tuido"

if ! command -v tuido >/dev/null 2>&1; then
	echo ""
	echo "If the command is not found, add ${INSTALL_DIR} to your PATH:"
	echo "  fish:  fish_add_path ${INSTALL_DIR}"
fi
