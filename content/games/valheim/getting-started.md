---
title: "Valheim Server: Getting Started"
description: "How to set up a Valheim dedicated server, including SteamCMD setup, world configuration, and crossplay."
order: 1
tags: ["valheim", "steam", "setup"]
---

# Valheim Server: Getting Started

## Why a Dedicated Server?

Valheim's built-in "Start Server" option ties the world to whoever is hosting. They log off, everyone gets kicked. A dedicated server runs independently — your world stays up 24/7 and you don't sacrifice a player slot to the host.

## Requirements

Real minimum specs, not what Steam says:

- **1-4 players:** 2GB RAM, 2 CPU cores
- **5-10 players:** 4GB RAM, 2-4 CPU cores
- **10+ players:** 8GB+ RAM — world size and build complexity matter more than player count here
- **Storage:** 1-2GB for the server, worlds grow over time (explored terrain = more data)
- **OS:** Linux or Windows. Linux is recommended for headless setups.

Valheim is surprisingly CPU-hungry. Large builds with lots of instances (torches, hearths, smelters) tank the tick rate.

## Installing via SteamCMD

SteamCMD is the standard way to install Valheim's dedicated server. No Steam account purchase required — the dedicated server is free.

```bash
# Install SteamCMD (Debian/Ubuntu)
sudo apt install steamcmd

# Or manually
mkdir -p ~/steamcmd && cd ~/steamcmd
curl -sqL "https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz" | tar zxvf -

# Install Valheim Dedicated Server (App ID 896660)
steamcmd +force_install_dir /opt/valheim +login anonymous +app_update 896660 validate +quit
```

The dedicated server App ID is **896660** — not the game's App ID (892970).

## Starting the Server

Valheim ships with a `start_server.sh` script, but it's easier to understand what's happening if you run it directly:

```bash
cd /opt/valheim

./valheim_server.x86_64 \
  -name "My Valheim Server" \
  -port 2456 \
  -world "MyWorld" \
  -password "secretpassword" \
  -savedir "/opt/valheim/saves" \
  -public 0
```

### Key Flags

| Flag | What it does |
|------|-------------|
| `-name` | Server name shown in the browser |
| `-port` | Base port. Valheim uses port and port+1 (2456-2457) |
| `-world` | World name. Creates a new one if it doesn't exist |
| `-password` | Must be at least 5 characters. The password and server name must not contain each other |
| `-savedir` | Where world files are stored |
| `-public 0` | Hide from server browser. Players connect via IP directly |
| `-crossplay` | Enable crossplay between Steam and Xbox/Microsoft Store |

### The Password Gotcha

Valheim checks that the password and server name don't contain each other. If your server is called "Viking Server" and your password is "Viking", players will get "wrong password" errors. Make them completely different.

## Ports

Valheim needs **two consecutive UDP ports** open:

| Port | Purpose |
|------|---------|
| 2456 | Game traffic |
| 2457 | Query port (Steam server browser) |

```bash
# UFW
sudo ufw allow 2456:2457/udp

# firewalld
sudo firewall-cmd --permanent --add-port=2456-2457/udp
sudo firewall-cmd --reload
```

## Running as a systemd Service

```ini
# /etc/systemd/system/valheim.service
[Unit]
Description=Valheim Dedicated Server
After=network.target

[Service]
Type=simple
User=valheim
WorkingDirectory=/opt/valheim
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/valheim +login anonymous +app_update 896660 +quit
ExecStart=/opt/valheim/valheim_server.x86_64 -name "My Server" -port 2456 -world "MyWorld" -password "secretpassword" -savedir "/opt/valheim/saves" -public 0
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
```

The `ExecStartPre` line auto-updates the server every time it starts. Remove it if you want to control update timing.

## Connecting

Players connect via:
- **Steam Server Browser:** if `-public 1` is set
- **Direct connect:** In-game, Join Game → Add Server → `<your-ip>:2456`
- **Steam overlay:** View → Servers → Favorites → Add → `<your-ip>:2457` (Steam uses the query port)

## World Backups

Valheim worlds consist of two files:
- `<worldname>.fwl` — world metadata
- `<worldname>.db` — actual world data

Back these up regularly. A bad crash or power outage can corrupt the `.db` file. A simple cron job works:

```bash
# Backup every 6 hours
0 */6 * * * cp /opt/valheim/saves/worlds_local/MyWorld.{fwl,db} /backups/valheim/$(date +\%Y\%m\%d-\%H\%M)/
```

## Next Steps

- [Valheim Configuration](/games/valheim/configuration) — world modifiers, combat settings, and performance tuning
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
