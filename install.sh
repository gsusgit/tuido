#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

# Run Go via mise when this repo ships mise.toml — avoids broken global `go` shims.
run_go() {
	if command -v mise >/dev/null 2>&1 && [[ -f "${ROOT}/mise.toml" ]]; then
		mise trust -y "${ROOT}" 2>/dev/null || mise trust "${ROOT}/mise.toml" 2>/dev/null || true
		MISE_YES=1 mise install -q
		mise exec -- go "$@"
		return
	fi
	go "$@"
}

if ! run_go version >/dev/null 2>&1; then
	echo "error: Go is not available." >&2
	echo "Install Go 1.26+ from https://go.dev/dl/ or install mise: https://mise.jdx.dev/" >&2
	exit 1
fi

run_go build -o tuido .

DEST="${HOME}/.local/bin"
mkdir -p "$DEST"
install -Dm755 tuido "${DEST}/tuido"

echo "Installed: ${DEST}/tuido"
echo "Run: tuido"

if ! command -v tuido >/dev/null 2>&1; then
	echo ""
	echo "If the command is not found, add ~/.local/bin to your PATH:"
	echo "  fish:  fish_add_path ~/.local/bin"
	echo "  bash:  echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ~/.bashrc"
fi
