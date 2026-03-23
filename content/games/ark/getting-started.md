---
title: "ARK: Survival Evolved Server: Getting Started"
description: "Setting up an ARK dedicated server, including the massive install size, cluster setup, and basic configuration."
order: 1
tags: ["ark", "steam", "setup"]
---

# ARK: Survival Evolved Server: Getting Started

## Fair Warning

ARK's dedicated server is one of the most resource-hungry game servers you can run. The install alone is 15-20GB, and a running server can easily eat 10-16GB of RAM. If you're looking for something lightweight, this isn't it.

## Requirements

- **Small server (1-10 players):** 8GB RAM minimum, 4 CPU cores
- **Medium server (10-30 players):** 16GB RAM, 4-6 CPU cores
- **Large server (30-70 players):** 32GB RAM, 8 CPU cores
- **Storage:** 15-20GB for base install. Each map is 5-10GB more. Mods can add another 10-20GB
- **OS:** Linux or Windows. The Linux version uses Proton/Wine under the hood and can be finicky

RAM is the killer here. ARK loads the entire map plus all tamed dinos and structures into memory.

## Installing via SteamCMD

```bash
# ARK Dedicated Server App ID: 376030
steamcmd +force_install_dir /opt/ark +login anonymous +app_update 376030 validate +quit
```

This will take a while. 15-20GB download.

## Starting the Server

```bash
cd /opt/ark/ShooterGame/Binaries/Linux

./ShooterGameServer \
  "TheIsland?listen?SessionName=My ARK Server?ServerPassword=secret?ServerAdminPassword=adminpass?MaxPlayers=30?Port=7777?QueryPort=27015" \
  -server -log -NoBattlEye
```

Yes, the config format is a URL query string. Welcome to Unreal Engine.

### Map Names

| Map | ID | Notes |
|-----|----|-------|
| The Island | `TheIsland` | Default map, most stable |
| Scorched Earth | `ScorchedEarth_P` | DLC |
| Aberration | `Aberration_P` | DLC |
| Extinction | `Extinction` | DLC |
| Ragnarok | `Ragnarok` | Free community map |
| Valguero | `Valguero_P` | Free community map |
| Crystal Isles | `CrystalIsles` | Free community map |
| The Center | `TheCenter` | Free community map |
| Lost Island | `LostIsland` | Free community map |
| Fjordur | `Fjordur` | Free community map |

### Key URL Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `SessionName` | — | Server name in browser |
| `ServerPassword` | — | Password to join |
| `ServerAdminPassword` | — | Admin password (use `enablecheats <password>` in-game) |
| `MaxPlayers` | 70 | Player cap |
| `Port` | 7777 | Game port (UDP) |
| `QueryPort` | 27015 | Steam query port (UDP) |
| `?Multihome=<IP>` | — | Bind to specific IP (for multi-home servers) |
| `?RCONEnabled=True` | — | Enable RCON |
| `?RCONPort=27020` | — | RCON port |

### Launch Flags

| Flag | What it does |
|------|-------------|
| `-server` | Run as dedicated server |
| `-log` | Enable console logging |
| `-NoBattlEye` | Disable BattlEye anti-cheat (reduces CPU load) |
| `-automanagedmods` | Auto-download and update mods |
| `-crossplay` | Enable crossplay |
| `-ClusterId=<name>` | Cluster ID for server transfers |

## Ports

ARK uses multiple ports:

| Port | Protocol | Purpose |
|------|----------|---------|
| 7777 | UDP | Game traffic |
| 7778 | UDP | Raw UDP socket (auto: game port + 1) |
| 27015 | UDP | Steam query |
| 27020 | TCP | RCON (if enabled) |

```bash
# UFW
sudo ufw allow 7777:7778/udp
sudo ufw allow 27015/udp
sudo ufw allow 27020/tcp

# firewalld
sudo firewall-cmd --permanent --add-port=7777-7778/udp
sudo firewall-cmd --permanent --add-port=27015/udp
sudo firewall-cmd --permanent --add-port=27020/tcp
sudo firewall-cmd --reload
```

## Running as a systemd Service

```ini
# /etc/systemd/system/ark.service
[Unit]
Description=ARK Dedicated Server
After=network.target

[Service]
Type=simple
User=ark
LimitNOFILE=100000
WorkingDirectory=/opt/ark/ShooterGame/Binaries/Linux
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/ark +login anonymous +app_update 376030 +quit
ExecStart=/opt/ark/ShooterGame/Binaries/Linux/ShooterGameServer "TheIsland?listen?SessionName=My ARK Server?ServerAdminPassword=adminpass?MaxPlayers=30?Port=7777?QueryPort=27015?RCONEnabled=True?RCONPort=27020" -server -log -NoBattlEye
Restart=on-failure
RestartSec=60

[Install]
WantedBy=multi-user.target
```

Note `LimitNOFILE=100000` — ARK opens an absurd number of file descriptors.

## Cluster Setup (Server Transfers)

To allow players to transfer characters and dinos between maps, run multiple servers with the same `-ClusterId` and a shared cluster directory:

```bash
# Server 1 - The Island
./ShooterGameServer "TheIsland?listen?..." -server -log -ClusterId=mycluster -ClusterDirOverride=/opt/ark/cluster

# Server 2 - Ragnarok
./ShooterGameServer "Ragnarok?listen?..." -server -log -ClusterId=mycluster -ClusterDirOverride=/opt/ark/cluster
```

Each server needs its own port set.

## Mods

Add mods via the `-automanagedmods` flag and `GameModIds` in `GameUserSettings.ini`:

```ini
[ServerSettings]
ActiveMods=731604991,889745138
```

Mod IDs come from the Steam Workshop. Mods are downloaded automatically on server start.

## Next Steps

- [ARK Configuration](/games/ark/configuration) — rates, taming, breeding, and GameUserSettings.ini reference
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
