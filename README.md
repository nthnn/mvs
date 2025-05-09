<p align="center">
    <img src="https://raw.githubusercontent.com/nthnn/mvs/refs/heads/main/assets/mvs-logo.png" width="250" />
</p>
<h1 align="center">Minimal Versioning System</h1>

<p align="center">
    <img alt="Build MVS" src="https://github.com/nthnn/mvs/actions/workflows/build_ci.yml/badge.svg" />
</p>

A lightweight version control system written in Go. MVS provides basic `init`, `add`, `remove`, `commit`, `amend`, `log`, `branch`, `checkout`, `status`, and `tree` commands, with `msgpack` metadata, Ed25519 signatures for tamper-evident history, and global configuration via YAML.

<p align="center">
    <img src="https://raw.githubusercontent.com/nthnn/mvs/refs/heads/main/assets/screenshot.png" width="60%" />
</p>

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
| `branch [name]`     | List or create a branch              |
| `checkout <name>`   | Switch to branch or commit           |
| `add <paths>`       | Stage file changes                   |
| `remove <paths>`    | Unstage or remove file changes       |
| `commit -m <msg>`   | Commit staged changes                |
| `amend -m <msg>`    | Amend the message of previous commit |
| `log`               | Show commit history                  |
| `status`            | Show staged/modified/untracked files |
| `tree`              | Render an ASCII branch tree          |

## License

This project is licensed under the **MIT License**. See the [LICENSE](https://github.com/nthnn/mvs/blob/main/LICENSE) file for details.
