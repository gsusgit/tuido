# Screenshots & media

Add these files so the main [README](../../README.md) renders correctly on GitHub.

## Task list (one per theme)

Use the **same tasks and terminal size** in all four shots so readers can compare palettes.

| File | Theme (`config.json` → `theme`) |
|------|----------------------------------|
| `task-list-catppuccin.png` | `catppuccin` |
| `task-list-tokyo-night.png` | `tokyo-night` |
| `task-list-one-dark.png` | `one-dark` |
| `task-list-monochrome.png` | `monochrome` |

Quick set theme before capture:

```bash
# example
jq '.theme = "tokyo-night"' ~/.config/tuido/config.json > /tmp/c.json && mv /tmp/c.json ~/.config/tuido/config.json
tuido
```

Or press `t` until the desired theme appears, then screenshot.

## Other views

| File | Description |
|------|-------------|
| `filters.png` | Filter panel (`f`) |
| `task-editor.png` | New or edit form (`n` / `e`) |

## Demo GIF

| File | Description |
|------|-------------|
| `../assets/demo.gif` | 10–15s loop — you create this locally |

## Capture tips

- Terminal: **≥ 60×20**, true color, dark background  
- PNG: `grim`, terminal export, or screenshot tool  
- GIF: `ffmpeg` or [vhs](https://github.com/charmbracelet/vhs)  

Keep filenames stable so README links do not break.
