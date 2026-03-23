---
title: "Factorio Server: Getting Started"
description: "Setting up a Factorio headless server, including save management and multiplayer configuration."
order: 1
tags: ["factorio", "setup"]
---

# Factorio Server: Getting Started

## Why Factorio Servers Are Different

Factorio's multiplayer is deterministic lockstep — every client simulates the entire game, and the server just coordinates inputs. This means:

- The server doesn't need a beefy GPU or lots of RAM
- **But** the server CPU must keep up with the game simulation, which gets heavier as your factory grows
- Late-game megabases can bring any server to its knees regardless of hardware

## Requirements

- **Early-game (1-4 players):** 1 CPU core, 512MB RAM
- **Mid-game:** 1-2 cores, 1-2GB RAM
- **Megabase:** 2+ fast cores, 4GB+ RAM. Single-thread performance is king
- **Storage:** 100MB for the server, saves vary (10MB to 500MB+ for megabases)
- **Bandwidth:** Low. Factorio is efficient — ~5KB/s per player in normal play

Factorio is one of the most efficient dedicated servers out there. A $5/month VPS can run a small server fine.

## Installation

Factorio provides a headless Linux build. You don't need SteamCMD.

```bash
# Download the headless server
mkdir -p /opt/factorio && cd /opt/factorio
curl -Lo factorio.tar.xz "https://factorio.com/get-download/stable/headless/linux64"
tar xf factorio.tar.xz --strip-components=1

# Or via package manager on some distros
# The headless build has no GUI dependencies
```

The headless server is free to download — no Factorio account or purchase required.

## Creating a World

```bash
# Generate a new save with default settings
./bin/x64/factorio --create /opt/factorio/saves/myworld.zip

# Generate with a specific map-gen preset
./bin/x64/factorio --create /opt/factorio/saves/myworld.zip --preset rich-resources

# Generate with custom settings
./bin/x64/factorio --create /opt/factorio/saves/myworld.zip --map-gen-settings map-gen-settings.json --map-settings map-settings.json
```

### Map Generation Presets

| Preset | Description |
|--------|-------------|
| `default` | Standard generation |
| `rich-resources` | Larger, richer resource patches |
| `marathon` | Recipes cost more, patches are richer to compensate |
| `death-world` | Aggressive biters, less resources |
| `death-world-marathon` | Pain |
| `rail-world` | Resources far apart, biters don't expand |
| `ribbon-world` | Narrow horizontal strip |
| `island` | Starting area is an island |

## Starting the Server

```bash
./bin/x64/factorio --start-server /opt/factorio/saves/myworld.zip \
  --server-settings /opt/factorio/data/server-settings.json
```

That's it. Factorio's server is refreshingly simple.

## Server Settings

Copy the example config and edit it:

```bash
cp /opt/factorio/data/server-settings.example.json /opt/factorio/data/server-settings.json
```

Key fields in `server-settings.json`:

```json
{
  "name": "My Factorio Server",
  "description": "A multiplayer factory",
  "tags": ["game", "vanilla"],
  "max_players": 0,
  "visibility": {
    "public": false,
    "lan": true
  },
  "username": "",
  "password": "",
  "token": "",
  "game_password": "",
  "require_user_verification": true,
  "max_upload_in_kilobytes_per_second": 0,
  "max_upload_slots": 5,
  "minimum_latency_in_ticks": 0,
  "ignore_player_limit_for_returning_players": false,
  "allow_commands": "admins-only",
  "autosave_interval": 5,
  "autosave_slots": 5,
  "afk_autokick_interval": 0,
  "auto_pause": true,
  "only_admins_can_pause_the_game": true
}
```

### Settings Worth Changing

| Setting | Default | Recommendation |
|---------|---------|---------------|
| `visibility.public` | `false` | Set `true` + fill in `username`/`token` to list on the server browser |
| `game_password` | `""` | Set one for private servers |
| `auto_pause` | `true` | Pauses when no one is online. **Keep this on** unless you want your factory running idle and consuming biters' patience |
| `autosave_interval` | `5` | Minutes. 5 is good |
| `allow_commands` | `admins-only` | Keep this. `true` gives everyone console access |
| `require_user_verification` | `true` | Validates Factorio accounts. Keep on for public servers |

### Authentication Token

To list your server publicly, you need a token. Get it from your [Factorio profile page](https://factorio.com/profile) — look for the `token` field. You can also find it in `player-data.json` (under `service-token`) on any machine where you've logged into Factorio.

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 34197 | UDP | Game traffic |

```bash
# UFW
sudo ufw allow 34197/udp

# firewalld
sudo firewall-cmd --permanent --add-port=34197/udp
sudo firewall-cmd --reload
```

## Running as a systemd Service

```ini
# /etc/systemd/system/factorio.service
[Unit]
Description=Factorio Headless Server
After=network.target

[Service]
Type=simple
User=factorio
WorkingDirectory=/opt/factorio
ExecStart=/opt/factorio/bin/x64/factorio --start-server /opt/factorio/saves/myworld.zip --server-settings /opt/factorio/data/server-settings.json
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
```

## Console & Admin

While the server runs, you can type commands in the terminal:

```
/promote <player>     -- Make someone an admin
/ban <player>         -- Ban a player
/kick <player>        -- Kick a player
/save                 -- Force save
/quit                 -- Graceful shutdown
```

Admin list is stored in `server-adminlist.json`:

```json
["your_factorio_username"]
```

## Next Steps

- [Factorio Configuration](/games/factorio/configuration) — map generation settings, runtime tuning, and mod support
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
