---
title: "Enshrouded Server: Getting Started"
description: "Setting up an Enshrouded dedicated server with SteamCMD, including world setup and multiplayer configuration."
order: 1
tags: ["enshrouded", "steam", "setup"]
---

# Enshrouded Server: Getting Started

## Requirements

- **1-4 players:** 4GB RAM, 4 CPU cores
- **4-16 players:** 8GB RAM, 4-6 CPU cores
- **Storage:** 3-4GB for the server
- **OS:** Linux (via Proton) or Windows

Enshrouded is still in Early Access and server requirements may change with updates.

## Installing via SteamCMD

```bash
# Enshrouded Dedicated Server App ID: 2278520
steamcmd +force_install_dir /opt/enshrouded +login anonymous +app_update 2278520 validate +quit
```

## Starting the Server

### Linux (via Proton)

The Enshrouded server is a Windows binary. On Linux, use Proton:

```bash
# Using steamcmd's built-in Proton
cd /opt/enshrouded

# Set up Proton compatibility
export STEAM_COMPAT_DATA_PATH=/opt/enshrouded/steamapps/compatdata/2278520
export STEAM_COMPAT_CLIENT_INSTALL_PATH=/home/enshrouded/.steam/steam

# Run via Proton
/home/enshrouded/.steam/steam/compatibilitytools.d/GE-Proton*/proton run ./enshrouded_server.exe
```

### Windows

```
cd C:\enshrouded
enshrouded_server.exe
```

## Configuration

Enshrouded uses `enshrouded_server.json` in the server directory:

```json
{
  "name": "My Enshrouded Server",
  "password": "",
  "saveDirectory": "./savegame",
  "logDirectory": "./logs",
  "ip": "0.0.0.0",
  "gamePort": 15636,
  "queryPort": 15637,
  "slotCount": 16,
  "gameSettings": {
    "playerHealthFactor": 1,
    "playerManaFactor": 1,
    "playerStaminaFactor": 1,
    "enableDurability": true,
    "enemyDamageFactor": 1,
    "enemyHealthFactor": 1,
    "enemyStaminaFactor": 1,
    "tombstoneMode": "AddBackpackMaterials",
    "miningDamageFactor": 1,
    "plantGrowthSpeedFactor": 1,
    "resourceDropStackAmountFactor": 1,
    "factoryProductionSpeedFactor": 1,
    "pacingOfTime": 1,
    "dayTimeDuration": 1800,
    "nightTimeDuration": 720,
    "shroudTimeFactor": 1,
    "randomSpawnerAmount": "Normal",
    "aggroPoolAmount": "Normal"
  }
}
```

### Key Settings

| Setting | Default | Description |
|---------|---------|-------------|
| `name` | — | Server name |
| `password` | `""` | Join password (empty = open) |
| `slotCount` | `16` | Max players (up to 16) |
| `gamePort` | `15636` | Game port (UDP) |
| `queryPort` | `15637` | Query port (UDP) |

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 15636 | UDP | Game traffic |
| 15637 | UDP | Steam query |

```bash
# UFW
sudo ufw allow 15636/udp
sudo ufw allow 15637/udp

# firewalld
sudo firewall-cmd --permanent --add-port=15636/udp
sudo firewall-cmd --permanent --add-port=15637/udp
sudo firewall-cmd --reload
```

## Running as a systemd Service

```ini
# /etc/systemd/system/enshrouded.service
[Unit]
Description=Enshrouded Dedicated Server
After=network.target

[Service]
Type=simple
User=enshrouded
WorkingDirectory=/opt/enshrouded
Environment=STEAM_COMPAT_DATA_PATH=/opt/enshrouded/steamapps/compatdata/2278520
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/enshrouded +login anonymous +app_update 2278520 +quit
ExecStart=/usr/bin/wine64 /opt/enshrouded/enshrouded_server.exe
Restart=on-failure
RestartSec=30

[Install]
WantedBy=multi-user.target
```

## Connecting

1. Launch Enshrouded
2. Play → Join → Add Server
3. Enter `<ip>:15636`
4. Enter password if set

## Backups

Save data is in the `saveDirectory` path (default: `./savegame`):

```bash
cp -r /opt/enshrouded/savegame /backups/enshrouded/$(date +%Y%m%d-%H%M)
```

## Next Steps

- [Enshrouded Configuration](/games/enshrouded/configuration) — game settings, difficulty tuning, and time cycle
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
