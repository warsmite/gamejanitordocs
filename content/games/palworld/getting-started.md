---
title: "Palworld Server: Getting Started"
description: "How to set up a Palworld dedicated server, including SteamCMD installation and initial configuration."
order: 1
tags: ["palworld", "steam", "setup"]
---

# Palworld Server: Getting Started

## Requirements

Palworld's dedicated server is resource-hungry. Don't underestimate these numbers:

- **1-4 players:** 8GB RAM minimum, 4 CPU cores
- **5-16 players:** 16GB RAM, 4-6 CPU cores
- **16-32 players:** 16-32GB RAM, 6+ CPU cores
- **Storage:** 5GB+ for the server, SSD strongly recommended
- **OS:** Linux or Windows

Yes, 8GB minimum for even a small server. Palworld loads the entire map into memory and tracks all Pal entities server-side. This is not Minecraft — you cannot run it on a 2GB VPS.

## Installing via SteamCMD

```bash
# Install the dedicated server (App ID 2394010)
steamcmd +force_install_dir /opt/palworld +login anonymous +app_update 2394010 validate +quit
```

The dedicated server App ID is **2394010**.

## Starting the Server

```bash
cd /opt/palworld

# Linux
./PalServer.sh -port=8211 -players=16 -useperfthreads -NoAsyncLoadingThread -UseMultithreadForDS
```

### Launch Flags

| Flag | What it does |
|------|-------------|
| `-port=8211` | Game port (UDP) |
| `-players=16` | Max player count |
| `-useperfthreads` | Better thread scheduling |
| `-NoAsyncLoadingThread` | Reduces stuttering during asset loading |
| `-UseMultithreadForDS` | Enables multithreaded processing for the dedicated server |
| `-publiclobby` | List as a community server in the server browser |

### The Performance Flags

`-useperfthreads -NoAsyncLoadingThread -UseMultithreadForDS` are not optional — they make a meaningful difference in server performance. Always include them.

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 8211 | UDP | Game traffic |
| 27015 | TCP + UDP | Steam query (server browser) |

```bash
# UFW
sudo ufw allow 8211/udp
sudo ufw allow 27015

# firewalld
sudo firewall-cmd --permanent --add-port=8211/udp
sudo firewall-cmd --permanent --add-port=27015/tcp
sudo firewall-cmd --permanent --add-port=27015/udp
sudo firewall-cmd --reload
```

## Running as a systemd Service

```ini
# /etc/systemd/system/palworld.service
[Unit]
Description=Palworld Dedicated Server
After=network.target

[Service]
Type=simple
User=palworld
LimitNOFILE=65535
WorkingDirectory=/opt/palworld
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/palworld +login anonymous +app_update 2394010 +quit
ExecStart=/opt/palworld/PalServer.sh -port=8211 -players=16 -useperfthreads -NoAsyncLoadingThread -UseMultithreadForDS
Restart=on-failure
RestartSec=30

[Install]
WantedBy=multi-user.target
```

Note `LimitNOFILE=65535` — Palworld opens a lot of file descriptors. Without this, you'll get cryptic crashes on busy servers.

## First Connection

1. Launch Palworld
2. Join Multiplayer Game
3. Enter `<your-ip>:8211` in the connection box
4. Enter the server password if set

Community servers appear in the browser if `bIsUseBackupSaveData` is configured. For private servers, direct connect is simpler.

## Scheduled Restarts

Palworld has a known memory leak — RAM usage grows over time. Schedule restarts:

```bash
# Restart every 6 hours via cron
0 */6 * * * systemctl restart palworld
```

This is not a workaround — it's standard practice for Palworld servers. Even official servers restart periodically.

## Backups

Save data lives in `Pal/Saved/SaveGames/0/<server-hash>/`:

```bash
# Backup the entire save directory
cp -r /opt/palworld/Pal/Saved/SaveGames /backups/palworld/$(date +%Y%m%d-%H%M)
```

Back up before every server update — Palworld updates have broken save compatibility before.

## Next Steps

- [Palworld Configuration](/games/palworld/configuration) — rates, difficulty, Pal settings, and all PalWorldSettings.ini options
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
