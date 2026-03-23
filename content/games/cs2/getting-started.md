---
title: "Counter-Strike 2 Server: Getting Started"
description: "Setting up a CS2 dedicated server with SteamCMD, including GSLT tokens, game modes, and workshop maps."
order: 1
tags: ["cs2", "counter-strike", "steam", "setup"]
---

# Counter-Strike 2 Server: Getting Started

## Requirements

CS2 servers are relatively lightweight:

- **10-player competitive:** 2GB RAM, 2 CPU cores
- **12-16 player casual:** 2-4GB RAM, 2 CPU cores
- **24+ player community:** 4GB RAM, 4 CPU cores
- **Storage:** 35GB for the server
- **OS:** Linux (recommended) or Windows

CS2's dedicated server is efficient. The main resource concern is CPU — tick rate matters.

## Game Server Login Token (GSLT)

You need a GSLT to run a public CS2 server. Without one, the server won't be listed and clients may not connect.

1. Go to https://steamcommunity.com/dev/managegameservers
2. Log in with your Steam account
3. Create a new token for App ID **730** (yes, 730, not 2000 — CS2 uses the same App ID as CS:GO)
4. Save the token

One GSLT per server instance. Don't share them.

## Installing via SteamCMD

```bash
# CS2 Dedicated Server App ID: 730
steamcmd +force_install_dir /opt/cs2 +login anonymous +app_update 730 validate +quit
```

## Starting the Server

```bash
cd /opt/cs2

./game/bin/linuxsteamrt64/cs2 -dedicated \
  +map de_dust2 \
  -port 27015 \
  +game_type 0 \
  +game_mode 1 \
  -maxplayers 10 \
  +sv_setsteamaccount "YOUR_GSLT_TOKEN"
```

### Game Types and Modes

The `game_type` and `game_mode` combination determines what kind of server you're running:

| game_type | game_mode | Mode |
|-----------|-----------|------|
| 0 | 0 | Casual |
| 0 | 1 | Competitive |
| 0 | 2 | Wingman |
| 1 | 0 | Arms Race |
| 1 | 1 | Demolition |
| 1 | 2 | Deathmatch |
| 3 | 0 | Custom (community servers) |

### Key Launch Parameters

| Parameter | Description |
|-----------|-------------|
| `-dedicated` | Run as dedicated server |
| `-port 27015` | Server port |
| `-maxplayers 10` | Max player count |
| `+map <map>` | Starting map |
| `+sv_setsteamaccount` | GSLT token |
| `+mapgroup mg_active` | Map group for rotation |
| `-ip 0.0.0.0` | Bind to specific IP |

### Sub-Tick

CS2 uses a sub-tick system running at 64 Hz. Unlike CS:GO, there is no 128-tick option — Valve replaced the traditional tick model with sub-tick interpolation that aims to make actions feel tick-independent. All servers (official, community, and third-party) run at 64 Hz.

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 27015 | UDP + TCP | Game traffic + RCON |
| 27020 | UDP | SourceTV (if enabled) |

```bash
# UFW
sudo ufw allow 27015
sudo ufw allow 27020/udp

# firewalld
sudo firewall-cmd --permanent --add-port=27015/tcp
sudo firewall-cmd --permanent --add-port=27015/udp
sudo firewall-cmd --permanent --add-port=27020/udp
sudo firewall-cmd --reload
```

## Running as a systemd Service

```ini
# /etc/systemd/system/cs2.service
[Unit]
Description=Counter-Strike 2 Dedicated Server
After=network.target

[Service]
Type=simple
User=cs2
WorkingDirectory=/opt/cs2
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/cs2 +login anonymous +app_update 730 +quit
ExecStart=/opt/cs2/game/bin/linuxsteamrt64/cs2 -dedicated +map de_dust2 -port 27015 +game_type 0 +game_mode 1 -maxplayers 10 +sv_setsteamaccount "YOUR_GSLT_TOKEN"
Restart=on-failure
RestartSec=15

[Install]
WantedBy=multi-user.target
```

## Workshop Maps

To use Steam Workshop maps, you need a Steam Web API key:

1. Get a key from https://steamcommunity.com/dev/apikey
2. Add to launch: `+host_workshop_collection <collection_id> +workshop_start_map <map_id>`
3. Or use `host_workshop_map <id>` in the console

```bash
./game/bin/linuxsteamrt64/cs2 -dedicated \
  +map de_dust2 \
  -port 27015 \
  +game_type 0 +game_mode 1 \
  +sv_setsteamaccount "YOUR_GSLT" \
  -authkey "YOUR_WEB_API_KEY" \
  +host_workshop_collection 123456789
```

## Connecting

Players connect via:
- **Server browser:** if the GSLT is valid and the server is public
- **Console:** `connect <ip>:27015`
- **Favorites:** add the server IP in Steam

## Next Steps

- [CS2 Configuration](/games/cs2/configuration) — server.cfg, competitive settings, and plugin setup
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
