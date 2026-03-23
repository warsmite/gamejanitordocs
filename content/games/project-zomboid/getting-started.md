---
title: "Project Zomboid Server: Getting Started"
description: "Setting up a Project Zomboid dedicated server, including world setup, Steam integration, and mod support."
order: 1
tags: ["project-zomboid", "steam", "setup"]
---

# Project Zomboid Server: Getting Started

## Requirements

- **1-4 players:** 4GB RAM, 2 CPU cores
- **5-16 players:** 6-8GB RAM, 4 CPU cores
- **16-32 players:** 8-16GB RAM, 4-6 CPU cores
- **Storage:** 3GB for server, worlds grow as players explore
- **Java:** Project Zomboid bundles its own JRE — you don't need to install Java separately

RAM usage scales with explored map area. A server where players spread out across the entire Knox County map will use significantly more RAM than one where everyone stays in Muldraugh.

## Installing via SteamCMD

```bash
# Project Zomboid Dedicated Server App ID: 380870
steamcmd +force_install_dir /opt/zomboid +login anonymous +app_update 380870 validate +quit
```

## First Run

```bash
cd /opt/zomboid

./start-server.sh -servername MyServer
```

On first run, it will:
1. Ask you to set an admin password
2. Generate world files
3. Create the config at `~/Zomboid/Server/MyServer.ini`

The server stores data in `~/Zomboid/` by default, not the install directory.

### Changing the Save Location

Use the `-cachedir` launch option to change where server data is stored:

```bash
./start-server.sh -servername MyServer -cachedir=/opt/zomboid-data
```

## Server Config

The main config is at `~/Zomboid/Server/<servername>.ini`. It's generated on first run — stop the server and edit it.

Key settings for getting started:

```ini
# Server identity
PublicName=My PZ Server
PublicDescription=A Project Zomboid server
Public=false
MaxPlayers=16
ServerPassword=

# Admin
AdminPassword=youradminpassword

# Gameplay
PVP=false
PauseEmpty=true
SpawnPoint=0,0,0

# Map
Map=Muldraugh, KY;West Point, KY;Rosewood, KY;Riverside, KY;Valley Station, KY;Louisville, KY
```

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 16261 | UDP | Game traffic |
| 16262 | UDP | Direct connection |

Some setups also use 16263-16264 for additional connections. Open a range to be safe:

```bash
# UFW
sudo ufw allow 16261:16262/udp

# firewalld
sudo firewall-cmd --permanent --add-port=16261-16262/udp
sudo firewall-cmd --reload
```

## Running as a systemd Service

```ini
# /etc/systemd/system/zomboid.service
[Unit]
Description=Project Zomboid Dedicated Server
After=network.target

[Service]
Type=simple
User=zomboid


WorkingDirectory=/opt/zomboid
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/zomboid +login anonymous +app_update 380870 +quit
ExecStart=/opt/zomboid/start-server.sh -servername MyServer
Restart=on-failure
RestartSec=30

[Install]
WantedBy=multi-user.target
```

## In-Game Admin

Connect to the server and use the admin panel (press Esc → Admin Panel) or console commands:

```
/adduser "username" "password"
/addusertowhitelist "username"
/grantadmin "username"
/removeadmin "username"
/kickuser "username"
/banuser "username"
/save
/quit
```

## Mod Support

Project Zomboid has excellent Steam Workshop mod support. Add mods in the server config:

```ini
# In MyServer.ini
WorkshopItems=2392987599;2478768005;2169435993
Mods=eris_minimap;Arsenal(26)GunFighter;MyOtherMod
```

- `WorkshopItems` — semicolon-separated Workshop IDs. The server downloads these automatically
- `Mods` — semicolon-separated mod folder names (from the mod's `mod.info` file)

Both fields are required. `WorkshopItems` handles download, `Mods` handles activation.

### Finding Mod IDs

The Workshop ID is in the URL: `https://steamcommunity.com/sharedfiles/filedetails/?id=2392987599`

The mod folder name is in the mod's `mod.info` file, field `id=`.

## Backups

World data is in `~/Zomboid/Saves/Multiplayer/<servername>/`:

```bash
# Backup the world
cp -r ~/Zomboid/Saves/Multiplayer/MyServer /backups/zomboid/$(date +%Y%m%d-%H%M)

# Also backup player data
cp -r ~/Zomboid/db /backups/zomboid/db-$(date +%Y%m%d-%H%M)
```

## Next Steps

- [Project Zomboid Configuration](/games/project-zomboid/configuration) — sandbox settings, zombie behavior, loot respawn, and all server.ini options
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
