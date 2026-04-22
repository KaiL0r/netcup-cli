# netcup-cli

A small CLI for the [Netcup Server Control Panel API (SCP)](https://www.netcup.com/en/helpcenter/documentation/servercontrolpanel/api).

This repository implements a subset of SCP API endpoints and provides a simple CLI to interact with them.

## Quickstart

1. Install Go (1.20+ recommended).
2. to use the pre-compiled binary run

   ```bash
   cd netcup-cli
   ./netcup-cli --help
   ```

   or build it yourself:

   ```bash
   cd netcup-cli
   go build -o netcup-cli
   ./netcup-cli --help
   ```

## Available top-level commands

The CLI uses cobra and exposes these top-level commands:

- `servers` — Manage servers (list, get, update, logs, rescue-system, storage-optimize, disks, snapshots, metrics)
- `users` — Manage user details and SSH API keys
- `tasks` — Manage long-running tasks

See `./netcup-cli servers --help` for subcommands and flags.

## Implemented API endpoints

Currently the only implemented Categories from the Documentation Page are:

- Server Disks
- Server Metrics
- Server Snapshots
- Servers
- Tasks
- Users

ToDO:

- Server Firewalls
- Server ISO
- Server Images
- Server Networking

## Contributing

Add missing endpoints by following the existing pattern:

- Add API methods in `internal/api/<category>.go`.
- Add CLI commands in `cmd/<category>.go` and register under server command.

You can also refactor code or correct naming of things. Nothing is set in stone right now. Just send me a pull request and i'm gonna look into it, when i have time
