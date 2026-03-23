---
title: "Running Multiple Game Servers"
description: "How to run several game servers on one machine — port planning, resource isolation, user separation, and Docker."
order: 8
tags: ["multiple-servers", "docker", "isolation", "ports"]
---

# Running Multiple Game Servers

One machine can run several game servers simultaneously. The key is keeping them isolated — separate ports, separate users, separate resources.

## Port Planning

Every game server needs its own set of ports. Plan ahead to avoid conflicts.

### Example Port Layout

| Game | Game Port | Query Port | RCON Port | Notes |
|------|-----------|-----------|-----------|-------|
| Minecraft | 25565 | — | 25575 | TCP |
| Valheim | 2456 | 2457 | — | UDP |
| Factorio | 34197 | — | 27015 | UDP |
| Terraria | 7777 | — | — | TCP |

Most games use unique default ports, so conflicts are rare. Problems arise when running **multiple instances of the same game** (e.g., two Minecraft servers):

| Instance | Game Port | RCON Port |
|----------|-----------|-----------|
| Minecraft (survival) | 25565 | 25575 |
| Minecraft (creative) | 25566 | 25576 |

Document your port layout somewhere. Future you will thank present you.

## User Separation

Create a dedicated Linux user for each game server:

```bash
sudo useradd -r -m -s /bin/bash minecraft
sudo useradd -r -m -s /bin/bash valheim
sudo useradd -r -m -s /bin/bash factorio
```

Each user owns only its game's files:

```bash
sudo chown -R minecraft:minecraft /opt/minecraft
sudo chown -R valheim:valheim /opt/valheim
sudo chown -R factorio:factorio /opt/factorio
```

Benefits:
- If one game server is exploited, the attacker can't access other game data
- Easy to see which server is using which resources in htop
- Clean permission boundaries

## Separate systemd Services

Each game server gets its own service file and can be managed independently:

```bash
# Manage them individually
sudo systemctl start minecraft
sudo systemctl stop valheim
sudo systemctl restart factorio

# Check what's running
systemctl list-units --type=service | grep -E "minecraft|valheim|factorio"
```

## Resource Limits with systemd

Prevent one game server from starving the others by setting resource limits in the service file:

```ini
[Service]
# Limit RAM usage
MemoryMax=8G
MemoryHigh=6G

# Limit CPU (100% = one full core)
CPUQuota=200%

# Limit file descriptors
LimitNOFILE=65535
```

`MemoryHigh` is a soft limit (the system tries to reclaim memory), `MemoryMax` is hard (the process gets killed if it exceeds this).

### Example: Sharing a 32GB Server

| Game | MemoryMax | CPUQuota | Notes |
|------|-----------|----------|-------|
| OS + overhead | — | — | Reserve 2GB |
| Minecraft | 8G | 200% | Needs lots of RAM |
| Rust | 12G | 200% | Memory hungry |
| Factorio | 2G | 100% | Lightweight |
| Valheim | 4G | 150% | Moderate |

Total allocated: 26GB of 32GB, leaving 6GB buffer.

## Shared SteamCMD

Don't install SteamCMD separately for each game. Install it once:

```bash
# System-wide install
sudo apt install steamcmd

# Or install to a shared location
mkdir -p /opt/steamcmd
cd /opt/steamcmd
curl -sqL "https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz" | tar zxvf -
```

Then reference it from each service's `ExecStartPre`:

```ini
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/valheim +login anonymous +app_update 896660 +quit
```

## Docker Approach

Docker containers provide stronger isolation than Linux users alone. Each container has its own filesystem, network namespace, and resource limits.

### When Docker Makes Sense

- Running many game servers with clean isolation
- Easy teardown and recreation (wipe a server by deleting a container)
- Reproducible setups (Dockerfiles document the entire install)

### When Docker Is Overkill

- Running 1-2 game servers
- You're still learning Linux basics
- The game has complex Proton/Wine requirements

### Basic Docker Setup

```bash
# Install Docker
curl -fsSL https://get.docker.com | sh
sudo usermod -aG docker $USER
```

### Example: Minecraft in Docker

```bash
# Using a community image
docker run -d \
  --name minecraft \
  -p 25565:25565 \
  -v /opt/minecraft-data:/data \
  -e EULA=TRUE \
  -e MEMORY=4G \
  --restart unless-stopped \
  itzg/minecraft-server
```

### Example: docker-compose for Multiple Servers

```yaml
# docker-compose.yml
services:
  minecraft:
    image: itzg/minecraft-server
    ports:
      - "25565:25565"
    volumes:
      - /opt/minecraft-data:/data
    environment:
      EULA: "TRUE"
      MEMORY: "4G"
    restart: unless-stopped
    deploy:
      resources:
        limits:
          memory: 6G

  valheim:
    image: lloesche/valheim-server
    ports:
      - "2456-2457:2456-2457/udp"
    volumes:
      - /opt/valheim-data:/config
    environment:
      SERVER_NAME: "My Valheim Server"
      WORLD_NAME: "MyWorld"
      SERVER_PASS: "secretpassword"
    restart: unless-stopped
    deploy:
      resources:
        limits:
          memory: 4G

  factorio:
    image: factoriotools/factorio
    ports:
      - "34197:34197/udp"
      - "27015:27015/tcp"
    volumes:
      - /opt/factorio-data:/factorio
    restart: unless-stopped
    deploy:
      resources:
        limits:
          memory: 2G
```

```bash
# Start everything
docker compose up -d

# Check status
docker compose ps

# View logs for one service
docker compose logs -f minecraft

# Stop everything
docker compose down
```

### Docker Caveats for Game Servers

- **UDP port mapping** — Always specify `/udp` for UDP ports. Docker defaults to TCP
- **Performance** — Docker adds minimal overhead. Network-intensive servers might see 1-2ms extra latency from NAT
- **Storage** — Always use bind mounts (`-v /host/path:/container/path`) not Docker volumes, so backups are straightforward
- **Updates** — Pull new images or rebuild containers. Don't `apt upgrade` inside containers

## Monitoring Multiple Servers

When running several servers, you need at-a-glance status:

```bash
#!/bin/bash
# /opt/scripts/status-all.sh

echo "=== Game Server Status ==="
echo ""

for service in minecraft valheim factorio terraria; do
    if systemctl is-active --quiet ${service} 2>/dev/null; then
        MEM=$(ps aux | grep -i ${service} | grep -v grep | awk '{sum+=$6} END {printf "%.0fMB", sum/1024}')
        echo "  ${service}: RUNNING (${MEM})"
    else
        echo "  ${service}: STOPPED"
    fi
done
```

## Next Steps

- [Performance Monitoring](/self-hosting/performance-monitoring) — diagnose which server is causing problems
- [Backups & Disaster Recovery](/self-hosting/backups) — back up all your servers
- [Security Hardening](/self-hosting/security) — isolation is part of security
