---
title: "Performance Monitoring"
description: "How to tell if your game server is struggling — CPU, RAM, disk I/O, network, and game-specific metrics."
order: 7
tags: ["performance", "monitoring", "htop", "diagnostics"]
---

# Performance Monitoring

Your players say "the server is laggy." Is it the server, their internet, or the game itself? Here's how to actually diagnose performance issues.

## The Quick Check: htop

```bash
# Install if not present
sudo apt install htop

# Run it
htop
```

htop gives you a live view of CPU, RAM, and processes. Look for:

- **CPU bars maxed out** — your game server is CPU-bottlenecked
- **Memory bar near full** — you need more RAM or the game has a memory leak
- **A single core pegged at 100%** — normal for single-threaded game servers, but means you've hit the ceiling
- **High system (kernel) CPU** — could indicate disk I/O problems

### What to Look For Per Game

| Symptom | Likely Cause |
|---------|-------------|
| One CPU core at 100%, others idle | Game server is single-threaded and maxed out |
| RAM steadily climbing over hours | Memory leak (common in Palworld, Valheim) |
| Periodic CPU spikes | Auto-save or world generation |
| High I/O wait | Disk is too slow (HDD or overloaded SSD) |

## CPU Monitoring

```bash
# Real-time CPU usage per core
mpstat -P ALL 1

# Average load over time
uptime
# Output: load average: 0.95, 1.02, 0.87
# These are 1, 5, and 15-minute averages
# For a 4-core machine, a load of 4.0 means fully loaded
```

### Is My Server CPU-Bottlenecked?

If the game server process consistently uses 95-100% of a single core in htop, you're CPU-bound. Solutions:

- Reduce simulation complexity (lower view distance, fewer entities, smaller world)
- Get a server with faster single-thread CPU performance
- There's no way to make a single-threaded game use more cores — that's a game engine limitation

## RAM Monitoring

```bash
# Current memory usage
free -h

# Output:
#               total   used   free   shared  buff/cache  available
# Mem:           16G    8.2G   1.1G    24M      6.7G       7.4G
```

**Don't panic about low "free" memory.** Linux uses spare RAM for disk cache (`buff/cache`). The number that matters is `available` — that's what's actually free for new processes.

```bash
# Watch memory over time (every 5 seconds)
watch -n 5 free -h

# Memory usage of your game server specifically
ps aux | grep -i valheim
# The RSS column shows actual RAM usage in KB
```

### When to Worry

- `available` drops below 500MB — you're running tight
- The game server's RSS keeps growing without plateau — memory leak, schedule restarts
- You see `oom-killer` in `dmesg` — the kernel killed something because RAM ran out

```bash
# Check if OOM killer has fired
dmesg | grep -i "oom\|killed"
```

## Disk I/O

Disk performance matters for auto-saves, world loading, and chunk streaming.

```bash
# Install iostat
sudo apt install sysstat

# Watch disk I/O in real time
iostat -x 1

# Key columns:
# %util  — percentage of time the disk is busy. >80% means it's saturated
# await  — average time (ms) for I/O requests. >10ms on SSD is concerning
# r/s, w/s — reads and writes per second
```

### Auto-Save Lag

If your server freezes briefly every few minutes, it's probably auto-saving to a slow disk:

```bash
# Watch I/O during an auto-save
# Run iostat, then trigger a save (via RCON or wait for auto-save)
iostat -x 1
```

Solutions:
- Move to an SSD
- Increase the auto-save interval
- Some games support async saving (Factorio, Paper Minecraft)

## Network Monitoring

```bash
# Bandwidth usage per interface
# Install nload or iftop
sudo apt install iftop

# Watch network traffic
sudo iftop -i eth0

# Or for a simpler view
ip -s link show eth0
```

### Network Issues vs Server Issues

If players report lag but your CPU/RAM/disk look fine, it's probably network:

```bash
# Check packet loss and latency from the server to a player
ping player-ip

# Check for dropped packets
cat /proc/net/dev
# Look for the drop and errs columns
```

Common network issues:
- **High ping** — server is geographically far from players. Can't fix without moving the server
- **Packet loss** — check your hosting provider's status page. Could be your ISP
- **Bandwidth saturation** — check with `iftop`. Unlikely unless you have many players on a home connection

## Game-Specific Metrics

### Minecraft (Paper)

```
# In-game or via RCON
/tps
# Shows ticks per second. Target: 20.0
# Below 18: noticeable lag
# Below 15: serious problems

/timings
# Detailed performance report
```

### Factorio

```
# Server console
/performance
# Shows UPS (Updates Per Second). Target: 60
```

### Rust

```bash
# Via RCON
status
# Shows FPS (server frame rate). Target: 30
# Also shows entity count — main performance driver

perf 1
# Detailed performance stats
```

### Valheim / Unity Games

No built-in metrics. Monitor externally with htop and watch for:
- Memory growth over time
- CPU spikes during events or player activity

## Setting Up Alerts

For serious setups, get notified when something goes wrong:

```bash
#!/bin/bash
# /opt/scripts/alert-check.sh
# Run via cron every 5 minutes

SERVICE="minecraft"

if ! systemctl is-active --quiet ${SERVICE}; then
    echo "[ALERT] ${SERVICE} is DOWN as of $(date)" | \
        mail -s "Game Server Down: ${SERVICE}" you@example.com
fi

# Check RAM usage
AVAILABLE_MB=$(free -m | awk '/^Mem:/{print $7}')
if [ "${AVAILABLE_MB}" -lt 500 ]; then
    echo "[ALERT] Low memory: ${AVAILABLE_MB}MB available" | \
        mail -s "Low Memory on Game Server" you@example.com
fi
```

```bash
# Cron: check every 5 minutes
*/5 * * * * /opt/scripts/alert-check.sh 2>/dev/null
```

## Next Steps

- [Hardware & Hosting Options](/self-hosting/hardware-and-hosting) — if you need to upgrade
- [Automated Updates & Restarts](/self-hosting/updates-and-restarts) — scheduled restarts help memory leaks
