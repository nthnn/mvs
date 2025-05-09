# Minimal Versioning System (MVS)

A lightweight version control system written in Go. MVS provides basic `init`, `add`, `remove`, `commit`, `log`, `branch`, `checkout`, `status`, and `tree` commands, with `msgpack` metadata, Ed25519 signatures for tamper-evident history, and global configuration via YAML.

- **Repository Initialization**: Create a new `.mvs` repository with `mvs init`.
- **Content Tracking**: Stage (`add`) and unstage (`remove`) files, then snapshot changes with `commit`.
- **History & Integrity**: Browse history with `log`; each commit is MsgPack-serialized, gzip-compressed, and signed with Ed25519.
- **Branching & Checkout**: Lightweight branches in `.mvs/refs/heads`; switch contexts via `branch` and `checkout`.
- **Status & Tree Views**: `status` shows staged, modified (including deletions), and untracked files; `tree` renders your branches in an ASCII UI.
- **Global Configuration**: Set your name and email in `~/.local/mvs/globals.yaml` for commits.

## Commands

| Command             | Description                          |
| ------------------- | ------------------------------------ |
| `init`              | Initialize a new repository          |
| `add <paths>`       | Stage file changes                   |
| `remove <paths>`    | Unstage or remove file changes       |
| `commit -m "<msg>"` | Commit staged changes                |
| `log`               | Show commit history                  |
| `branch [name]`     | List or create a branch              |
| `checkout <name>`   | Switch to branch or commit           |
| `status`            | Show staged/modified/untracked files |
| `tree`              | Render an ASCII branch tree          |

## To-do

- [ ] Complete `*.deb` installer generation on `build.sh`.
- [ ] Add more information to `status` command.
- [ ] One by one printing of commit history.
