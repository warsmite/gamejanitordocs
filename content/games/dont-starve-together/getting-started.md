---
title: "Don't Starve Together Server: Getting Started"
description: "Setting up a Don't Starve Together dedicated server, including cluster configuration, caves, and mod support."
order: 1
tags: ["dont-starve-together", "setup"]
---

# Don't Starve Together Server: Getting Started

## How DST Servers Work

Don't Starve Together uses a **cluster** system. Each world (overworld, caves) is a separate **shard** running its own process. A typical setup runs two shards:

- **Master shard** — The overworld
- **Caves shard** — The underground caves

Both shards run simultaneously and players can travel between them. You need enough resources to run both.

## Requirements

- **Overworld only:** 1GB RAM, 1 CPU core
- **Overworld + Caves:** 2GB RAM, 2 CPU cores
- **With mods:** 2-4GB RAM depending on mod count
- **Storage:** 500MB for the server

DST is lightweight. It's one of the easier dedicated servers to run.

## Installation

DST provides a Linux dedicated server through SteamCMD:

```bash
# Don't Starve Together Dedicated Server App ID: 343050
steamcmd +force_install_dir /opt/dst +login anonymous +app_update 343050 validate +quit
```

## Cluster Token

You need a cluster token from Klei to run a server:

1. Launch Don't Starve Together on your gaming PC
2. Press `~` to open the console
3. Type `TheNet:GenerateClusterToken()`
4. The token is saved to your DST data folder
5. Copy the token to your server

Or get it from https://accounts.klei.com/account/game/servers?game=DontStarveTogether

## Directory Structure

```
~/.klei/DoNotStarveTogether/
└── MyCluster/
    ├── cluster.ini              # Cluster-wide settings
    ├── cluster_token.txt        # Your Klei auth token
    ├── Master/
    │   ├── server.ini           # Overworld shard settings
    │   └── worldgenoverride.lua # World gen settings
    └── Caves/
        ├── server.ini           # Caves shard settings
        └── worldgenoverride.lua # Caves gen settings
```

Create this structure manually:

```bash
mkdir -p ~/.klei/DoNotStarveTogether/MyCluster/{Master,Caves}
```

## Configuration

### cluster.ini

```ini
[GAMEPLAY]
game_mode = survival
max_players = 6
pvp = false
pause_when_empty = true

[NETWORK]
cluster_name = My DST Server
cluster_description = A Don't Starve Together server
cluster_password =
cluster_intention = cooperative
lan_only_cluster = false

[MISC]
console_enabled = true

[SHARD]
shard_enabled = true
bind_ip = 127.0.0.1
master_ip = 127.0.0.1
master_port = 10889
cluster_key = supersecretkey
```

### cluster_token.txt

Paste your Klei cluster token (just the token string, nothing else).

### Master/server.ini (Overworld)

```ini
[NETWORK]
server_port = 10999

[SHARD]
is_master = true

[STEAM]
master_server_port = 27018
authentication_port = 8768
```

### Caves/server.ini

```ini
[NETWORK]
server_port = 10998

[SHARD]
is_master = false
name = Caves

[STEAM]
master_server_port = 27019
authentication_port = 8769
```

### Game Modes

| Mode | Description |
|------|-------------|
| `survival` | Standard. Ghosts on death, can be resurrected |
| `wilderness` | Random spawn, no resurrection |
| `endless` | Ghosts can self-resurrect at portal, no rollback on death |

## Starting the Server

```bash
cd /opt/dst/bin64

# Start the overworld (Master)
./dontstarve_dedicated_server_nullrenderer_x64 -console -cluster MyCluster -shard Master &

# Start the caves
./dontstarve_dedicated_server_nullrenderer_x64 -console -cluster MyCluster -shard Caves &
```

Both processes need to be running simultaneously.

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 10999 | UDP | Master shard (game traffic) |
| 10998 | UDP | Caves shard |
| 10889 | UDP | Inter-shard communication |
| 27018-27019 | UDP | Steam master server |
| 8768-8769 | UDP | Steam authentication |

```bash
# UFW — open all needed ports
sudo ufw allow 10998:10999/udp
sudo ufw allow 27018:27019/udp
sudo ufw allow 8768:8769/udp
```

Only the Master shard port (10999) needs to be reachable from the internet. Caves and inter-shard ports can stay local if both shards run on the same machine.

## Running as systemd Services

Create two services, one per shard:

```ini
# /etc/systemd/system/dst-master.service
[Unit]
Description=DST Overworld
After=network.target

[Service]
Type=simple
User=dst
WorkingDirectory=/opt/dst/bin64
ExecStart=/opt/dst/bin64/dontstarve_dedicated_server_nullrenderer_x64 -console -cluster MyCluster -shard Master
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
```

```ini
# /etc/systemd/system/dst-caves.service
[Unit]
Description=DST Caves
After=dst-master.service
Requires=dst-master.service

[Service]
Type=simple
User=dst
WorkingDirectory=/opt/dst/bin64
ExecStart=/opt/dst/bin64/dontstarve_dedicated_server_nullrenderer_x64 -console -cluster MyCluster -shard Caves
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
```

The caves service depends on the master service.

## Mod Support

Mods are configured per-cluster in `dedicated_server_mods_setup.lua` and per-shard in `modoverrides.lua`.

### Installing Mods

Edit `/opt/dst/mods/dedicated_server_mods_setup.lua`:

```lua
-- Workshop mod IDs
ServerModSetup("378160763")  -- Global Positions
ServerModSetup("375850593")  -- Global Pause
ServerModSetup("462434129")  -- Wormhole Marks
```

### Enabling Mods

Create `modoverrides.lua` in each shard directory (`Master/` and `Caves/`):

```lua
return {
  ["workshop-378160763"] = { enabled = true },
  ["workshop-375850593"] = { enabled = true },
  ["workshop-462434129"] = { enabled = true, configuration_options = {} }
}
```

Mods need to be enabled in **both** shard directories if they should run on both shards.

## Next Steps

- [DST Configuration](/games/dont-starve-together/configuration) — world generation, season tuning, and admin commands
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
