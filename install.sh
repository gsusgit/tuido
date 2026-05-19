#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

if command -v mise >/dev/null 2>&1; then
	mise install
fi

if ! command -v go >/dev/null 2>&1; then
	echo "error: Go is not available." >&2
	echo "Install Go 1.26+ from https://go.dev/dl/ or run: mise use -g go@1.26.3" >&2
	exit 1
fi

go build -o tuido .

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
