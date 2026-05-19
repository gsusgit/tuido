# TUI-DO

**Minimal terminal-first todo manager for tiling WM users.**

[github.com/gsusgit/tuido](https://github.com/gsusgit/tuido)

Built with Go for fast, keyboard-driven workflows — no mouse required.

**TUI** (terminal UI) + **DO** (your to-do list). The banner yells **TUIDO** in slant ASCII; you run **`tuido`** in lowercase — same tradition as `vim`, `git`, and other tools that let the work speak louder than the name.

<p align="center">
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go 1.26+" /></a>
  <a href="https://github.com/charmbracelet/bubbletea"><img src="https://img.shields.io/badge/Bubble%20Tea-TUI-ff75ab?style=for-the-badge" alt="Bubble Tea" /></a>
  <a href="#installation"><img src="https://img.shields.io/badge/Install-local--first-89b4fa?style=for-the-badge" alt="Local-first" /></a>
</p>

<p align="center">
  <img src="docs/assets/demo.gif" alt="TUI-DO demo" width="800" />
</p>

---

## Why?

I've always liked having **one simple place** for what's pending — not a productivity suite with hundreds of toggles I'll never use. I just want to see **what I've done** and **what's still left**.

My setup is a Hyprland workspace with the essentials always on screen: music, processes, terminal… and it started to feel like **something was missing** there. Riding the Hyprland / terminal-aesthetic wave, I thought: *let's experiment* — build something **for myself**, that lives in that same tile and stays out of the way.

That's how **TUI-DO** was born: out of a personal need first. If it fits your workflow too, welcome.

---

## Features

- ⌨️ **Keyboard-first** — navigate, create, edit, delete without leaving home row
- 📐 **Responsive TUI** — adapts layout and scroll to your terminal size
- ● **Task priorities** — high / medium / low with color-coded indicators
- 🏷️ **Categories & filters** — status, category, sort in a dedicated panel
- 🔍 **Live search** — filter the list by title with `/`
- 🌍 **Multi-language** — English, Español, Français, Deutsch, Italiano, Português
- 🎨 **Four themes** — Catppuccin Mocha, Tokyo Night, One Dark, Monochrome
- 💾 **Local-first** — JSON on disk under `~/.config/tuido/`, autosave
- 🪶 **Lightweight** — single binary, no daemon, no account
- 🧱 **Tiling-friendly** — made for Hyprland, i3, sway and split terminals

---

## Screenshots

### Task list · themes

<table>
  <tr>
    <td width="25%" align="center" valign="top">
      <img src="docs/screenshots/task-list-catppuccin.png" alt="Catppuccin Mocha" width="100%" />
      <br /><sub><b>Catppuccin Mocha</b></sub>
    </td>
    <td width="25%" align="center" valign="top">
      <img src="docs/screenshots/task-list-tokyo-night.png" alt="Tokyo Night" width="100%" />
      <br /><sub><b>Tokyo Night</b></sub>
    </td>
    <td width="25%" align="center" valign="top">
      <img src="docs/screenshots/task-list-one-dark.png" alt="One Dark" width="100%" />
      <br /><sub><b>One Dark</b></sub>
    </td>
    <td width="25%" align="center" valign="top">
      <img src="docs/screenshots/task-list-monochrome.png" alt="Monochrome" width="100%" />
      <br /><sub><b>Monochrome</b></sub>
    </td>
  </tr>
</table>

### Views

<table>
  <tr>
    <td width="25%" align="center" valign="top">
      <img src="docs/screenshots/filters.png" alt="Filters" width="100%" />
      <br /><sub><b>Filters</b></sub>
    </td>
    <td width="25%" align="center" valign="top">
      <img src="docs/screenshots/task-editor.png" alt="Task editor" width="100%" />
      <br /><sub><b>Task editor</b></sub>
    </td>
    <td width="25%" align="center" valign="top">
      <img src="docs/screenshots/controls.png" alt="Controls" width="100%" />
      <br /><sub><b>Controls</b></sub>
    </td>
    <td width="25%"></td>
  </tr>
</table>

---

## Requirements

| | |
|---|---|
| **Go** | [1.26+](https://go.dev/dl/) — or [mise](https://mise.jdx.dev/) (repo includes `mise.toml`) |
| **git** | to clone the repository |
| **Terminal** | true color, at least **60×20** columns×lines |
| **PATH** | `~/.local/bin` on your `$PATH` (installer uses it by default) |

---

## Installation

```bash
git clone https://github.com/gsusgit/tuido.git
cd tuido
./install.sh
tuido
```

`install.sh` installs Go via `mise.toml` when needed (`mise exec`) so a broken global `go` shim does not block the build.

**Command not found?** Add the install directory to your shell:

```bash
fish_add_path ~/.local/bin    # fish
```

### Manual build

```bash
go build -o tuido .
install -Dm755 tuido ~/.local/bin/tuido
```

### CLI

```bash
tuido              # run
tuido lang es      # set language
tuido reset -f     # wipe all tasks
```

---

## Controls

| Key | Action |
|-----|--------|
| `↑` `↓` / `k` `j` | Navigate |
| `n` | New task |
| `e` | Edit |
| `d` | Delete → `Enter` confirm, `Esc` cancel |
| `Space` | Toggle done |
| `f` | Filters |
| `r` | Reset filters (when active) |
| `/` | Search |
| `t` | Cycle theme |
| `c` `?` | Help |
| `Esc` `q` | Quit |

In **filters**: `Tab` / `←` `→` to change options, `Enter` to apply.  
In **new/edit**: `Tab` fields, `←` `→` category/priority, `Enter` save.

---

## Configuration

| Path | Content |
|------|---------|
| `~/.config/tuido/config.json` | `lang`, `theme` |
| `~/.config/tuido/data.json` | tasks |

Themes: `catppuccin` · `tokyo-night` · `one-dark` · `monochrome` — or press `t` in-app.

Languages: `en` · `es` · `fr` · `de` · `it` · `pt`

---

## Non-goals

TUI-DO is **not** a team suite, a cloud SaaS, or a bloated life OS.

Just **speed**, **simplicity**, and **terminal-native** focus.

---

## Built with

[Go](https://go.dev/) · [Bubble Tea](https://github.com/charmbracelet/bubbletea) · [Lip Gloss](https://github.com/charmbracelet/lipgloss) · [Bubbles](https://github.com/charmbracelet/bubbles)

---

<p align="center">
  <sub>Made for people who <code>$</code> live in the terminal.</sub>
</p>
