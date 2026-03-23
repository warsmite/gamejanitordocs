---
title: "Terraria Server: Getting Started"
description: "Setting up a Terraria dedicated server with TShock, including world generation and player management."
order: 1
tags: ["terraria", "tshock", "setup"]
---

# Terraria Server: Getting Started

## Vanilla vs TShock

Terraria ships with a built-in dedicated server (`TerrariaServer.exe` / `TerrariaServer.bin.x86_64`). It works, but it's barebones — no permissions, no anti-cheat, no grief protection.

**For most servers, use TShock.** It's a drop-in replacement that adds:
- User groups and permissions
- Anti-cheat and anti-grief
- Region protection
- Server-side characters (prevents inventory cheating)
- REST API for remote management

## Requirements

- **1-4 players:** 512MB RAM, 1 CPU core
- **5-16 players:** 1-2GB RAM, 2 CPU cores
- **Large worlds:** More RAM. A large world with heavy exploration can use 1GB+ on its own
- **Storage:** 200MB for the server, worlds are tiny (5-20MB)

Terraria servers are lightweight compared to most games.

## Installing TShock

```bash
# Download the latest TShock release
mkdir -p /opt/terraria && cd /opt/terraria

# Get the latest release from GitHub
# Check https://github.com/Pryaxis/TShock/releases for the current version
# Example (update version as needed):
curl -Lo tshock.zip "https://github.com/Pryaxis/TShock/releases/latest/download/TShock-Beta-linux-x64-Release.zip"
unzip tshock.zip

# Make executable
chmod +x TShock.Server
```

### Vanilla Server (Alternative)

If you don't want TShock:

```bash
# Download from Terraria's site
# The Linux server binary is included in the PC dedicated server package
mkdir -p /opt/terraria && cd /opt/terraria
# Extract the Linux directory from the server zip
chmod +x TerrariaServer.bin.x86_64
```

## First Run & World Generation

```bash
cd /opt/terraria
./TShock.Server -autocreate 3 -worldname "MyWorld" -world "/opt/terraria/worlds/MyWorld.wld" -port 7777 -maxplayers 16 -password "secretpassword"
```

### Launch Flags

| Flag | What it does |
|------|-------------|
| `-autocreate <1-3>` | Auto-generate world. 1=small, 2=medium, 3=large |
| `-worldname` | Name of the world |
| `-world` | Full path to the world file |
| `-port` | Server port (default: 7777) |
| `-maxplayers` | Player cap (default: 16, max: 255) |
| `-password` | Server password. Omit for no password |
| `-secure` | Enable anti-cheat (vanilla server) |
| `-noupnp` | Disable automatic UPnP port forwarding |
| `-lang <id>` | Language (1=English, 2=German, 3=Italian, etc.) |

### World Size Matters

| Size | Dimensions | Generation Time | Use Case |
|------|-----------|----------------|----------|
| Small | 4200 x 1200 | ~15 seconds | Quick games, testing |
| Medium | 6400 x 1800 | ~30 seconds | Small groups |
| Large | 8400 x 2400 | ~60 seconds | Most servers |

**Large is usually the right choice.** More space means less player conflict over resources and more biome variety. The extra RAM cost is negligible.

## Ports

Terraria uses a single TCP port:

| Port | Protocol | Purpose |
|------|----------|---------|
| 7777 | TCP | Game traffic |

```bash
# UFW
sudo ufw allow 7777/tcp

# firewalld
sudo firewall-cmd --permanent --add-port=7777/tcp
sudo firewall-cmd --reload
```

## Running as a systemd Service

```ini
# /etc/systemd/system/terraria.service
[Unit]
Description=Terraria Dedicated Server (TShock)
After=network.target

[Service]
Type=simple
User=terraria
WorkingDirectory=/opt/terraria
ExecStart=/opt/terraria/TShock.Server -autocreate 3 -worldname "MyWorld" -world "/opt/terraria/worlds/MyWorld.wld" -port 7777 -maxplayers 16
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
```

## TShock First-Time Setup

On first launch, TShock creates an admin token printed to the console. In-game, use:

```
/setup <token>
```

This grants you temporary elevated privileges so you can create an admin account:

```
/user add <username> <password> owner
/group addperm default tshock.world.modify
```

## World Backups

World files are at the path specified by `-world`. TShock also auto-creates backups in `/opt/terraria/tshock/backups/`.

For manual backups:

```bash
# Backup before maintenance
cp /opt/terraria/worlds/MyWorld.wld /backups/terraria/MyWorld-$(date +%Y%m%d-%H%M).wld
```

## Next Steps

- [Terraria Configuration](/games/terraria/configuration) — TShock config, permissions, and gameplay settings
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
