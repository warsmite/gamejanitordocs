---
title: "Rust Server: Getting Started"
description: "Setting up a Rust dedicated server with SteamCMD, including map wipes, oxide/umod, and performance basics."
order: 1
tags: ["rust", "steam", "setup"]
---

# Rust Server: Getting Started

## The Wipe Cycle

Before you set up a Rust server, understand this: Rust has a mandatory wipe cycle. On the first Thursday of every month, Facepunch pushes a forced update that wipes the map. You can optionally do blueprint wipes too.

This isn't optional. Your server resets monthly whether you like it or not. Plan around it.

## Requirements

Rust servers are heavy. These are real numbers, not store-page minimums:

- **50-100 players:** 8GB RAM, 4 CPU cores, SSD required
- **100-200 players:** 16GB RAM, 6+ CPU cores
- **200+ players:** 32GB RAM, dedicated hardware
- **Storage:** 5-10GB for the server, map saves can reach 50MB+
- **Bandwidth:** 10-20Mbps upload for a busy server

Rust loads the entire map into memory and tracks every entity (buildings, items, NPCs). RAM is usually the bottleneck.

## Installing via SteamCMD

```bash
# Rust Dedicated Server App ID: 258550
steamcmd +force_install_dir /opt/rust +login anonymous +app_update 258550 validate +quit
```

## Starting the Server

```bash
cd /opt/rust

./RustDedicated -batchmode \
  +server.port 28015 \
  +server.queryport 28017 \
  +rcon.port 28016 \
  +rcon.password "yourrconpassword" \
  +rcon.web 1 \
  +server.hostname "My Rust Server" \
  +server.identity "myserver" \
  +server.maxplayers 100 \
  +server.worldsize 3500 \
  +server.seed 12345 \
  +server.saveinterval 300
```

### Key Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `server.port` | 28015 | Game port (UDP) |
| `server.queryport` | 28017 | Steam query port (UDP). Defaults to max(server.port, rcon.port) + 1 |
| `rcon.port` | 28016 | RCON port (TCP) |
| `rcon.password` | — | **Required.** No RCON password = no admin access |
| `rcon.web` | 1 | Use WebSocket RCON (1) vs legacy (0). Use 1 |
| `server.hostname` | — | Server name in browser |
| `server.identity` | — | Directory name for this server's data |
| `server.maxplayers` | 50 | Player cap |
| `server.worldsize` | 4500 | Map size in meters (1000-6000) |
| `server.seed` | random | Map seed. Same seed + same size = same map |
| `server.saveinterval` | 600 | Seconds between saves |
| `server.tickrate` | 30 | Server tick rate. Don't change unless you know what you're doing |

### World Size Guide

| Size | Feel | Good for |
|------|------|----------|
| 2000 | Tiny, constant PvP | 10-30 players |
| 3000 | Small, action-focused | 30-75 players |
| 3500 | Medium, balanced | 50-150 players |
| 4000 | Large | 100-200 players |
| 4500 | Default, spacious | 150-250 players |
| 6000 | Massive, lots of empty space | 250+ players |

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 28015 | UDP | Game traffic |
| 28016 | TCP | RCON |
| 28017 | UDP | Steam query |

```bash
# UFW
sudo ufw allow 28015/udp
sudo ufw allow 28016/tcp
sudo ufw allow 28017/udp

# firewalld
sudo firewall-cmd --permanent --add-port=28015/udp
sudo firewall-cmd --permanent --add-port=28016/tcp
sudo firewall-cmd --permanent --add-port=28017/udp
sudo firewall-cmd --reload
```

## Running as a systemd Service

```ini
# /etc/systemd/system/rust.service
[Unit]
Description=Rust Dedicated Server
After=network.target

[Service]
Type=simple
User=rust
LimitNOFILE=65535
WorkingDirectory=/opt/rust
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/rust +login anonymous +app_update 258550 +quit
ExecStart=/opt/rust/RustDedicated -batchmode +server.port 28015 +server.queryport 28017 +rcon.port 28016 +rcon.password "yourrconpassword" +rcon.web 1 +server.hostname "My Rust Server" +server.identity "myserver" +server.maxplayers 100 +server.worldsize 3500 +server.seed 12345 +server.saveinterval 300
Restart=on-failure
RestartSec=30

[Install]
WantedBy=multi-user.target
```

## Oxide / uMod (Plugin Framework)

Most Rust servers run Oxide (now uMod). It's the plugin framework that makes Rust servers actually manageable.

```bash
# Download latest Oxide for Rust
cd /opt/rust
curl -Lo oxide.zip "https://umod.org/games/rust/download"
unzip -o oxide.zip
```

Oxide overwrites some Rust DLLs — that's normal. After installing, plugins go in `oxide/plugins/`:

```bash
# Download a plugin
curl -Lo oxide/plugins/GatherManager.cs "https://umod.org/plugins/GatherManager.cs"
```

Plugins auto-compile on server start. Essential plugins:

- **GatherManager** — control resource gather rates
- **NTeleportation** — /home, /tpr teleport commands
- **Kits** — starter kits for new players
- **StackSizeController** — increase stack sizes
- **AdminRadar** — admin ESP for catching cheaters
- **RustIO / RustMap** — live web map

## Wipe Procedure

When it's wipe day:

```bash
# Stop the server
systemctl stop rust

# Update
steamcmd +force_install_dir /opt/rust +login anonymous +app_update 258550 +quit

# Delete map data (keep blueprints for map-only wipe)
rm /opt/rust/server/myserver/*.sav*
rm /opt/rust/server/myserver/*.map

# For full wipe (including blueprints)
rm /opt/rust/server/myserver/*.db

# Change the seed for a new map layout
# Update the seed in your start command

# Start
systemctl start rust
```

## Next Steps

- [Rust Configuration](/games/rust/configuration) — server.cfg, performance tuning, and Oxide config
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
