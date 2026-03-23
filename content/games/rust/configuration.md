---
title: "Rust Configuration Reference"
description: "Server.cfg, convars, performance tuning, and Oxide configuration for Rust dedicated servers."
order: 2
tags: ["rust", "configuration"]
---

# Rust Configuration Reference

## server.cfg

Rust reads `server/<identity>/cfg/server.cfg` on startup. Any command-line convar can go here instead:

```
server.hostname "My Rust Server"
server.description "A custom Rust server"
server.headerimage "https://example.com/banner.png"
server.url "https://example.com"
server.maxplayers 100
server.worldsize 3500
server.saveinterval 300

# Decay
decay.scale 1.0
decay.upkeep true

# Crafting (admin/moderator only — use a plugin for server-wide instant craft)
craft.instant false

# AI
ai.think true
ai.npc_enable true
halloween.enabled false
xmas.enabled false
```

## Important Convars

### Server Identity

| Convar | Default | Description |
|--------|---------|-------------|
| `server.hostname` | — | Server name |
| `server.description` | — | Description (supports `\n` for newlines) |
| `server.headerimage` | — | Banner image URL (512x256 PNG recommended) |
| `server.url` | — | Website link shown in server listing |
| `server.tags` | — | Comma-separated tags for filtering |
| `server.pve` | `false` | PvE mode (no player damage) |

### Gameplay

| Convar | Default | Description |
|--------|---------|-------------|
| `decay.scale` | `1.0` | Building decay speed multiplier. 0 = no decay |
| `decay.upkeep` | `true` | Require upkeep resources in TC |
| `server.radiation` | `true` | Radiation at monuments |
| `craft.instant` | `false` | Instant crafting (admin/moderator only — use a plugin for server-wide) |
| `spawn.min_rate` | `0.5` | Minimum animal spawn rate |
| `spawn.max_rate` | `1.0` | Maximum animal spawn rate |

### Performance

| Convar | Default | Description |
|--------|---------|-------------|
| `server.tickrate` | `30` | Server tick rate. Higher = smoother but more CPU. Don't go above 30 without serious hardware |
| `fps.limit` | `256` | Frame rate cap for the server process |
| `gc.buffer` | `256` | Garbage collection buffer in MB |
| `server.saveinterval` | `600` | Seconds between auto-saves. Lower = safer but more disk I/O |
| `server.entityrate` | `16` | Entity spawns per tick |
| `batching.colliders` | `true` | Batch collider processing |

### Network

| Convar | Default | Description |
|--------|---------|-------------|
| `server.maxpacketspersecond` | `500` | Max packets per second per client before kick |
| `server.maxpacketsize` | `4096` | Max packet size in bytes |
| `server.netcachesize` | `1024` | Network cache size |

## RCON

Rust uses WebSocket RCON by default. Connect with tools like:

- **RustAdmin** — GUI tool for Windows
- **webrcon** — web-based RCON panel
- **rcon-cli** — command line

```bash
# Example with rcon-cli
rcon -a localhost:28016 -p "yourrconpassword" "status"
```

### Useful RCON Commands

| Command | Description |
|---------|-------------|
| `status` | Show connected players |
| `say "message"` | Server-wide chat message |
| `kick <steamid> "reason"` | Kick player |
| `ban <steamid> "reason"` | Ban player |
| `unban <steamid>` | Unban player |
| `ownerid <steamid> "name" "reason"` | Grant admin (owner level) |
| `moderatorid <steamid> "name" "reason"` | Grant moderator |
| `server.save` | Force save |
| `server.writecfg` | Save current config |
| `quit` | Graceful shutdown |
| `env.time` | Show/set time of day |
| `weather.fog` | Set fog level |
| `weather.rain` | Set rain level |
| `global.teleport <player> <x y z>` | Teleport player |

## Oxide / uMod Configuration

Plugin configs live in `oxide/config/<pluginname>.json`. Edit while the server is running and use `oxide.reload <plugin>` to apply.

### Common Plugin Configs

#### GatherManager

```json
{
  "Options": {
    "Default Multiplier": 2.0,
    "Pickup Multiplier": 2.0,
    "Quarry Multiplier": 2.0,
    "Survey Multiplier": 2.0
  }
}
```

#### NTeleportation

```json
{
  "Settings": {
    "HomeMaximumAmount": 3,
    "HomeCooldown": 120,
    "TprCooldown": 60,
    "TprCountDown": 10
  }
}
```

## Performance Tuning

### What Kills Rust Server Performance

1. **Entity count** — Every building block, deployable, and item on the ground is an entity. Servers typically die around 300k-500k entities
2. **Player count at peak** — More players = more network traffic, more entity interactions
3. **AI** — Scientists, animals, and NPCs are CPU-expensive
4. **Save operations** — Large maps pause the game thread during saves

### What You Can Do

- **Wipe regularly.** Entity count grows every day. Monthly forced wipes exist for a reason
- **Use `decay.scale`** — Don't disable decay. It's the primary entity cleanup mechanism
- **Set `server.saveinterval`** appropriately — 300s for busy servers, 600s for quiet ones
- **Run on an SSD** — Save operations write hundreds of MB. Spinning disks cause server freezes during saves
- **Monitor entity count** — `ent count` in RCON shows the total. Above 300k, expect degradation

### Automated Cleanup

Many admins run plugins to manage entity count:

- **EntityCleanup** — Removes abandoned entities
- **Decay Manager** — Fine-tune decay rates
- **ServerRewards** — Auto-wipe scheduler
