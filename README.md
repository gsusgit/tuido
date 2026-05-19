```
 ________  ______    ___  ____ 
/_  __/ / / /  _/___/ _ \/ __ \
 / / / /_/ // //___/ // / /_/ /
/_/  \____/___/   /____/\____/ 
```

# Tui-do

**Minimal terminal-first todo manager for tiling WM users.**

Built with Go for fast, keyboard-driven workflows — no mouse required.

<p align="center">
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go 1.26+" /></a>
  <a href="https://github.com/charmbracelet/bubbletea"><img src="https://img.shields.io/badge/Bubble%20Tea-TUI-ff75ab?style=for-the-badge" alt="Bubble Tea" /></a>
  <a href="#installation"><img src="https://img.shields.io/badge/Install-local--first-89b4fa?style=for-the-badge" alt="Local-first" /></a>
</p>

---

## Why?

**Tui-do** (`tuido`) was built for developers and terminal users who live inside **Hyprland**, **Neovim**, and keyboard-driven workflows.

The goal is simple: a **clean, minimal, distraction-free** task manager that feels native inside a terminal workspace — not another bloated productivity app fighting for your attention.

- **Local-first** — your tasks stay on disk under XDG config paths  
- **Fast** — single Go binary, instant startup  
- **Aesthetic** — Catppuccin, Tokyo Night, and more; built for dark terminals  
- **Respectful** — fixed header/footer, scrollable list, compact logo on small panes  

---

## Features

| | |
|---|---|
| ⌨️ | **Keyboard-first** — navigate, create, edit, delete without leaving home row |
| 📐 | **Responsive TUI** — adapts logo tier, layout, and scroll hints to terminal size |
| ● | **Task priorities** — alta / media / baja with color-coded indicators |
| 🏷️ | **Categories & filters** — status, category, sort order in a dedicated filter panel |
| 🔍 | **Live search** — filter the list by title with `/` |
| 🌍 | **Multi-language** — English, Español, Français, Deutsch, Italiano, Português |
| 🎨 | **Theme cycling** — Catppuccin Mocha, Tokyo Night, One Dark, Monochrome |
| 💾 | **Local persistence** — JSON on disk, autosave after every change |
| 🪶 | **Minimal footprint** — no database, no daemon, no cloud account |
| 🧱 | **Tiling-friendly** — designed for floating/split terminals (Hyprland, i3, sway…) |

---

## Screenshots

> Drop PNGs into `docs/screenshots/` — see [docs/screenshots/README.md](docs/screenshots/README.md) for filenames and capture tips.

### Task list · themes

Press `t` in the app to cycle themes, or set `theme` in `~/.config/tuido/config.json`.

#### Catppuccin Mocha

<p align="center">
  <img src="docs/screenshots/task-list-catppuccin.png" alt="Task list — Catppuccin Mocha" width="800" />
</p>

#### Tokyo Night

<p align="center">
  <img src="docs/screenshots/task-list-tokyo-night.png" alt="Task list — Tokyo Night" width="800" />
</p>

#### One Dark

<p align="center">
  <img src="docs/screenshots/task-list-one-dark.png" alt="Task list — One Dark" width="800" />
</p>

#### Monochrome

<p align="center">
  <img src="docs/screenshots/task-list-monochrome.png" alt="Task list — Monochrome" width="800" />
</p>

### Filters

<p align="center">
  <img src="docs/screenshots/filters.png" alt="Filter panel: status, category, sort field and direction" width="800" />
</p>

### Task editor

<p align="center">
  <img src="docs/screenshots/task-editor.png" alt="New or edit task: title, category, priority" width="800" />
</p>

### Demo

<p align="center">
  <img src="docs/assets/demo.gif" alt="Short demo: create task, toggle complete, filter, change theme" width="800" />
</p>

---

## Installation

### From source (recommended)

```bash
git clone https://github.com/gsus/todo-app.git
cd todo-app/tui
go build -o tuido .
install -Dm755 tuido ~/.local/bin/tuido   # or: cp tuido ~/.local/bin/
tuido
```

### Go install

```bash
go install github.com/gsus/todo-app/tui@latest
```

The installed binary is named **`tui`** (module path basename). Rename or symlink if you prefer `tuido`:

```bash
ln -sf "$(go env GOPATH)/bin/tui" ~/.local/bin/tuido
```

### Requirements

- Go **1.26+**
- A true-color terminal (`TERM=xterm-256color` or similar)
- Minimum terminal size: **60×20** columns×lines

### CLI commands

```bash
tuido              # Launch the TUI
tuido lang         # Pick language interactively
tuido lang es      # Set language (es, en, fr, de, it, pt)
tuido reset        # Delete all tasks (with confirmation)
tuido reset -f     # Delete without prompt
tuido --version
```

---

## Controls

### List view

| Key | Action |
|-----|--------|
| `↑` / `k`, `↓` / `j` | Move selection |
| `n` | New task |
| `e` | Edit selected task |
| `d` / `x` | Delete (confirm with `Enter`, cancel with `Esc`) |
| `Space` | Toggle completed |
| `f` | Open filters |
| `/` | Search by title |
| `r` | Reset active filters (when any filter is applied) |
| `t` | Cycle theme |
| `c` / `?` / `F1` | Controls help |
| `Esc` / `q` / `Ctrl+C` | Quit |

### Filter panel

| Key | Action |
|-----|--------|
| `Tab` | Next filter field |
| `←` / `→` | Change option (status, category, sort…) |
| `Enter` | Apply and return to list |
| `r` | Reset all filters |
| `Esc` | Back without applying |

### New / edit task

| Key | Action |
|-----|--------|
| `Tab` | Cycle title → category → priority |
| `←` / `→` | Change category or priority |
| `Enter` | Save |
| `Esc` | Cancel |

### Search (list)

| Key | Action |
|-----|--------|
| `/` | Focus search input |
| `Enter` | Apply search |
| `Esc` | Close search |

---

## Configuration

Tui-do follows **XDG-style paths** under `~/.config/tuido/`:

| File | Purpose |
|------|---------|
| `config.json` | Language (`lang`) and theme (`theme`) |
| `data.json` | Task storage (JSON) |

Example `config.json`:

```json
{
  "lang": "es",
  "theme": "tokyo-night"
}
```

**Themes** (`theme` field): `catppuccin` · `tokyo-night` · `one-dark` · `monochrome`  
Press `t` in the app to cycle without editing the file.

**Languages**: `en`, `es`, `fr`, `de`, `it`, `pt` — detected from `$LANG` on first run, or set via `tuido lang <code>`.

Legacy data at `~/.config/todotui/` is migrated automatically on first launch.

---

## Philosophy

### Non-goals

Tui-do is **intentionally simple**.

It is **not**:

- a team collaboration platform  
- a cloud-first SaaS with accounts and sync locks  
- a bloated “life OS” with calendars, habits, and widgets  

The focus is **speed**, **simplicity**, and **terminal-native** workflows.

---

## Roadmap

- [x] Live search (`/`)
- [x] Multi-theme support
- [x] i18n (6 languages)
- [x] Filter panel (status, category, sort)
- [ ] Custom user-defined categories
- [ ] JSON import / export
- [ ] Deeper Vim-style motions (`gg`, `G`, visual…)
- [ ] Optional cloud sync (opt-in, never default)
- [ ] Ultra-compact UI mode for laptop + tiling splits
- [ ] GitHub Releases with prebuilt binaries

Contributions welcome — especially docs, screenshots, and themes.

---

## Built with

- [Go](https://go.dev/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — styling & layout
- [Bubbles](https://github.com/charmbracelet/bubbles) — text inputs, help, viewport

---

## Project layout

```
tui/
├── main.go              # CLI entry (tuido / lang / reset)
├── cmd/                 # Subcommands
└── internal/
    ├── config/          # ~/.config/tuido/config.json
    ├── storage/         # Task persistence
    ├── model/           # App state
    ├── theme/           # Color palettes
    ├── i18n/            # Translations
    ├── tuiapp/          # Bubble Tea program
    └── view/            # Rendering
```

---

## Capturing screenshots & GIFs

See **[docs/screenshots/README.md](docs/screenshots/README.md)** for filenames and a short capture guide.

Suggested workflow aesthetic:

1. Dark wallpaper, rounded terminal (padding in Hyprland rules)  
2. `tuido` in the center tile — same task list in every theme shot for easy comparison  
3. `btop` or `cava` in a side tile — subtle, not cluttered  
4. Record `docs/assets/demo.gif`: add task → complete → filter → `t` theme  

---

<p align="center">
  <sub>Made for people who <code>$</code> live in the terminal.</sub>
</p>
