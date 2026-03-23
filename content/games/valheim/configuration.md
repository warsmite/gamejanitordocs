---
title: "Valheim Configuration Reference"
description: "World modifiers, server settings, and performance tuning for Valheim dedicated servers."
order: 2
tags: ["valheim", "configuration"]
---

# Valheim Configuration Reference

Valheim doesn't have a traditional config file like most servers. Settings are split between launch flags and in-game world modifiers.

## Launch Flags

These are set when starting the server binary:

| Flag | Default | Description |
|------|---------|-------------|
| `-name` | — | Server name in browser |
| `-port` | 2456 | Base UDP port |
| `-world` | — | World file name |
| `-password` | — | Server password (min 5 chars) |
| `-savedir` | `~/.config/unity3d/IronGate/Valheim` | Save location |
| `-public` | 1 | List in server browser |
| `-logFile` | — | Log output path |
| `-crossplay` | off | Enable Steam/Xbox crossplay |
| `-preset` | — | Difficulty preset (see below) |
| `-modifier` | — | Individual world modifier (repeatable) |
| `-setkey` | — | Set a world key (boss progression) |
| `-resetkeys` | — | Reset all boss keys |

## World Modifiers

Added in the Mistlands update, these let you tune difficulty without mods. Set via `-modifier <key> <value>` or `-preset`:

### Presets

| Preset | Description |
|--------|-------------|
| `casual` | Reduced damage, lower resource costs, portals allow all items |
| `easy` | Slightly reduced difficulty |
| `normal` | Default Valheim experience |
| `hard` | More enemies, tougher bosses |
| `hardcore` | Significantly harder. Permadeath-adjacent |
| `immersive` | No map, no boss markers |
| `hammer` | Creative-ish. Reduced costs, no raids |

### Individual Modifiers

Use these to fine-tune instead of presets:

| Modifier | Values | What it does |
|----------|--------|-------------|
| `combat` | `veryeasy`, `easy`, `hard`, `veryhard` | Enemy damage and health |
| `deathpenalty` | `casual`, `veryeasy`, `easy`, `hard`, `hardcore` | Skill loss on death |
| `resources` | `muchless`, `less`, `more`, `muchmore`, `most` | Drop rates from nodes and enemies |
| `raids` | `none`, `muchless`, `less`, `more`, `muchmore` | Frequency of base raids |
| `portals` | `casual`, `hard`, `veryhard` | What items can go through portals. `casual` = everything, `hard` = nothing |

### World Keys (via -setkey)

These are boolean toggles set with `-setkey <key>`:

| Key | What it does |
|-----|-------------|
| `playerevents` | Events (raids) only trigger when a player is nearby |
| `nobuildcost` | Building costs no resources |
| `passivemobs` | Creatures don't attack unless provoked |
| `nomap` | Disables the map |

Example with multiple modifiers:

```bash
./valheim_server.x86_64 \
  -name "Custom Server" \
  -port 2456 \
  -world "CustomWorld" \
  -password "mypassword" \
  -modifier combat hard \
  -modifier resources more \
  -modifier portals casual \
  -modifier deathpenalty easy
```

## Adminlist, Bannedlist, Permittedlist

Three plain-text files in your save directory control access:

- **adminlist.txt** — Steam IDs with admin commands (kick, ban, etc.)
- **bannedlist.txt** — Steam IDs that cannot connect
- **permittedlist.txt** — If non-empty, acts as a whitelist. Only listed IDs can join

Format is one Steam ID per line:

```
76561198012345678
76561198087654321
```

Find Steam IDs at steamid.io or by pressing F2 in-game.

## Performance Tuning

Valheim servers don't have many performance knobs — most of the load comes from world complexity:

### What Actually Causes Lag

1. **Massive builds** — Every placed piece is tracked. A 5000-piece castle will lag harder than 50 small huts
2. **Lots of active instances** — Smelters, kilns, fires, windmills all tick independently
3. **Unexplored vs explored terrain** — Loading new terrain chunks is expensive. Once explored, it's cached
4. **Enemy spawns during events** — Raids can spike CPU usage, especially in dense bases

### What You Can Do

- **Limit build complexity in concentrated areas.** This is the #1 cause of server lag. There's no config for this — it's a "talk to your players" situation
- **Use `-savedir` on an SSD.** World saves happen periodically and can stutter the server on slow disks
- **Restart the server daily.** Memory usage creeps up over time. A cron job at 4 AM solves this
- **Keep the server updated.** Iron Gate regularly ships server-side performance improvements

### Memory Behavior

Valheim's dedicated server typically uses 1-2GB at start and grows as more of the world is explored. A fully-explored world with large builds can use 4GB+. There are no JVM-style memory flags — it's a Unity game, memory management is internal.

## BepInEx / Server-Side Mods

For mod support, install BepInEx:

```bash
# Download and extract to the Valheim server directory
# Mods go in BepInEx/plugins/
```

Popular mods:
- **Valheim Plus** — comprehensive config overhaul (shared map, build tweaks, stamina changes)
- **Server Devcommands** — admin tools
- **World Size Unlocker** — larger world generation

**Important:** Most mods require clients to have the same mods installed. Valheim does not automatically sync mods — all players must manually install matching versions.
