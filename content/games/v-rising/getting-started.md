---
title: "V Rising Server: Getting Started"
description: "Setting up a V Rising dedicated server with SteamCMD, including game settings and castle management."
order: 1
tags: ["v-rising", "steam", "setup"]
---

# V Rising Server: Getting Started

## Requirements

- **1-10 players:** 4GB RAM, 2-4 CPU cores
- **10-40 players:** 6-8GB RAM, 4 CPU cores
- **Storage:** 3GB for the server
- **OS:** Linux (via Wine/Proton) or Windows

The Linux version runs through Wine, which the included start script handles automatically. It works fine but occasionally lags behind the Windows version for updates.

## Installing via SteamCMD

```bash
# V Rising Dedicated Server App ID: 1829350
steamcmd +force_install_dir /opt/vrising +login anonymous +app_update 1829350 validate +quit
```

## Starting the Server

### Linux

```bash
cd /opt/vrising

# The included script handles Wine setup
./start_server_example.sh
```

Or run directly:

```bash
# Requires Wine/Proton
wine64 /opt/vrising/VRisingServer.exe \
  -persistentDataPath /opt/vrising/save-data \
  -serverName "My V Rising Server" \
  -saveName "world1" \
  -logFile /opt/vrising/logs/VRisingServer.log
```

### Launch Flags

| Flag | Description |
|------|-------------|
| `-persistentDataPath` | Where saves and settings are stored |
| `-serverName` | Server name in browser |
| `-saveName` | Save directory name |
| `-logFile` | Log file location |
| `-address` | Bind IP (default: all interfaces) |
| `-gamePort` | Game port (default: 9876) |
| `-queryPort` | Steam query port (default: 9877) |

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 9876 | UDP | Game traffic |
| 9877 | UDP | Steam query |

```bash
# UFW
sudo ufw allow 9876/udp
sudo ufw allow 9877/udp

# firewalld
sudo firewall-cmd --permanent --add-port=9876/udp
sudo firewall-cmd --permanent --add-port=9877/udp
sudo firewall-cmd --reload
```

## Configuration Files

V Rising uses two JSON files in the save data path:

- **`ServerHostSettings.json`** — Network, admin, RCON settings
- **`ServerGameSettings.json`** — Gameplay, difficulty, castle settings

### ServerHostSettings.json

```json
{
  "Name": "My V Rising Server",
  "Description": "",
  "Port": 9876,
  "QueryPort": 9877,
  "MaxConnectedUsers": 40,
  "MaxConnectedAdmins": 4,
  "ServerFps": 30,
  "SaveName": "world1",
  "Password": "",
  "Secure": true,
  "ListOnSteam": true,
  "ListOnEOS": true,
  "AutoSaveCount": 20,
  "AutoSaveInterval": 120,
  "CompressSaveFiles": true,
  "GameSettingsPreset": "",
  "AdminOnlyDebugEvents": true,
  "DisableDebugEvents": false,
  "Rcon": {
    "Enabled": false,
    "Port": 25575,
    "Password": ""
  }
}
```

## Running as a systemd Service

```ini
# /etc/systemd/system/vrising.service
[Unit]
Description=V Rising Dedicated Server
After=network.target

[Service]
Type=simple
User=vrising
WorkingDirectory=/opt/vrising
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/vrising +login anonymous +app_update 1829350 +quit
ExecStart=/usr/bin/wine64 /opt/vrising/VRisingServer.exe -persistentDataPath /opt/vrising/save-data -serverName "My V Rising Server" -saveName "world1" -logFile /opt/vrising/logs/VRisingServer.log
Restart=on-failure
RestartSec=30
Environment=WINEDLLOVERRIDES="version=n"

[Install]
WantedBy=multi-user.target
```

## Admin Setup

Add Steam IDs to `adminlist.txt` in the save data path:

```
76561198012345678
76561198087654321
```

Admins can use the admin console in-game.

## Backups

Save data is in the `persistentDataPath` under `Saves/v3/<savename>/`:

```bash
cp -r /opt/vrising/save-data/Saves /backups/vrising/$(date +%Y%m%d-%H%M)
```

## Next Steps

- [V Rising Configuration](/games/v-rising/configuration) — game settings, castle limits, PvP schedules, and difficulty
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
