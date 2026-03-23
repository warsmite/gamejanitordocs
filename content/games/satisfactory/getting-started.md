---
title: "Satisfactory Server: Getting Started"
description: "Setting up a Satisfactory dedicated server, including installation, save management, and multiplayer configuration."
order: 1
tags: ["satisfactory", "steam", "setup"]
---

# Satisfactory Server: Getting Started

## Requirements

- **1-4 players:** 6GB RAM, 4 CPU cores
- **4+ players with large factories:** 8-12GB RAM, 4-6 CPU cores
- **Storage:** 5GB for server install
- **OS:** Linux or Windows

Satisfactory's server resource usage scales with factory complexity, not player count. A small factory with 4 players uses less than a massive megabase with 1 player.

## Installing via SteamCMD

```bash
# Satisfactory Dedicated Server App ID: 1690800
steamcmd +force_install_dir /opt/satisfactory +login anonymous +app_update 1690800 validate +quit
```

## Starting the Server

```bash
cd /opt/satisfactory
./FactoryServer.sh -unattended -multihome=0.0.0.0
```

### Launch Flags

| Flag | Description |
|------|-------------|
| `-unattended` | Run without user prompts |
| `-multihome=0.0.0.0` | Bind to all network interfaces |
| `-ServerQueryPort=15777` | Query port override |
| `-BeaconPort=15000` | Beacon port override |
| `-Port=7777` | Game port override |
| `-log` | Enable logging |
| `-ini:Engine:[HTTPServer.Listeners]:DefaultBindAddress=any` | Bind web API to all interfaces |

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 7777 | UDP | Game traffic |
| 15000 | UDP | Beacon port |
| 15777 | UDP | Query port |

```bash
# UFW
sudo ufw allow 7777/udp
sudo ufw allow 15000/udp
sudo ufw allow 15777/udp

# firewalld
sudo firewall-cmd --permanent --add-port=7777/udp
sudo firewall-cmd --permanent --add-port=15000/udp
sudo firewall-cmd --permanent --add-port=15777/udp
sudo firewall-cmd --reload
```

## First-Time Setup

Satisfactory's dedicated server is managed through the Server Manager web interface, not config files.

1. Start the server
2. Connect to it from the game client: Multiplayer → Join Game → enter `<ip>:15777`
3. On first connection, you'll be prompted to create a server admin password and claim the server
4. Set the server name, create or upload a save

### Claiming the Server

The first player to connect becomes the admin. You set:
- **Server name** — visible in the server browser
- **Admin password** — for managing the server
- **Client password** — for players to join (optional)

## Running as a systemd Service

```ini
# /etc/systemd/system/satisfactory.service
[Unit]
Description=Satisfactory Dedicated Server
After=network.target

[Service]
Type=simple
User=satisfactory
LimitNOFILE=65535
WorkingDirectory=/opt/satisfactory
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/satisfactory +login anonymous +app_update 1690800 +quit
ExecStart=/opt/satisfactory/FactoryServer.sh -unattended -multihome=0.0.0.0
Restart=on-failure
RestartSec=30

[Install]
WantedBy=multi-user.target
```

## Save Management

Saves are stored in `~/.config/Epic/FactoryGame/Saved/SaveGames/server/`. The server auto-saves periodically.

### Uploading an Existing Save

You can upload a save from single-player via the in-game Server Manager. Or manually:

```bash
# Copy save file to server
cp MySave.sav ~/.config/Epic/FactoryGame/Saved/SaveGames/server/

# The server will detect it and make it available in the Server Manager
```

### Backups

```bash
cp -r ~/.config/Epic/FactoryGame/Saved/SaveGames/server/ /backups/satisfactory/$(date +%Y%m%d-%H%M)
```

## Server API

Satisfactory has a REST API for server management. It's available on port 7777 (same as game port) over HTTPS:

```bash
# Check server health
curl -k https://localhost:7777/api/v1/

# The API requires authentication via the admin password
```

The API can:
- List/load/create saves
- Start/stop the server
- Query server state
- Change settings

## Next Steps

- [Satisfactory Configuration](/games/satisfactory/configuration) — server settings, autosave, and performance tuning
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
