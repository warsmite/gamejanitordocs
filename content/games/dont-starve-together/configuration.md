---
title: "Don't Starve Together Configuration Reference"
description: "World generation, season settings, admin commands, and worldgenoverride.lua for DST servers."
order: 2
tags: ["dont-starve-together", "configuration"]
---

# Don't Starve Together Configuration Reference

## cluster.ini — Full Reference

### [GAMEPLAY]

| Setting | Default | Description |
|---------|---------|-------------|
| `game_mode` | `survival` | `survival`, `wilderness`, `endless` |
| `max_players` | `6` | Max player count |
| `pvp` | `false` | PvP damage |
| `pause_when_empty` | `true` | Pause when no players online |
| `vote_enabled` | `true` | Allow vote kick/restart |

### [NETWORK]

| Setting | Default | Description |
|---------|---------|-------------|
| `cluster_name` | — | Server name |
| `cluster_description` | — | Server description |
| `cluster_password` | — | Password (blank = open) |
| `cluster_intention` | — | `cooperative`, `competitive`, `social`, `madness` |
| `lan_only_cluster` | `false` | LAN only |
| `offline_cluster` | `false` | Offline mode |
| `autosaver_enabled` | `true` | Auto-save |
| `tick_rate` | `15` | Server tick rate (15 or 30). 30 = smoother but more CPU |
| `whitelist_slots` | `0` | Reserved slots for whitelisted players |

### [SHARD]

| Setting | Default | Description |
|---------|---------|-------------|
| `shard_enabled` | `false` | Enable multi-shard (required for caves) |
| `bind_ip` | `127.0.0.1` | Shard communication bind IP |
| `master_ip` | `127.0.0.1` | Master shard IP |
| `master_port` | `10889` | Master shard communication port |
| `cluster_key` | — | Shared secret between shards |

## World Generation (worldgenoverride.lua)

This file controls world generation. Placed in each shard directory. **Changing it after world creation requires regenerating the world.**

### Overworld Example

```lua
return {
  override_enabled = true,
  preset = "SURVIVAL_DEFAULT",
  overrides = {
    -- World size
    world_size = "default",         -- "small", "medium", "default", "huge"
    branching = "default",          -- Path branching

    -- Seasons
    autumn = "default",             -- "noseason", "veryshort", "short", "default", "long", "verylong"
    winter = "default",
    spring = "default",
    summer = "default",
    season_start = "default",       -- Starting season: "default", "winter", "spring", "summer", "autumnorspring", "winterorsummer", "random"

    -- Resources
    berrybush = "default",          -- "never", "rare", "uncommon", "default", "often", "mostly", "always"
    carrot = "default",
    flint = "default",
    grass = "default",
    sapling = "default",
    rock = "default",
    trees = "default",

    -- Creatures
    spiders = "default",
    hounds = "default",
    deerclops = "default",          -- Boss frequency
    bearger = "default",
    goosemoose = "default",
    antlion = "default",
    dragonfly = "default",
    beefalos = "default",
    pigs = "default",
    rabbits = "default",

    -- World
    day = "default",                -- "default", "longday", "longdusk", "longnight", "noday", "nodusk", "nonight", "onlyday", "onlydusk", "onlynight"
    regrowth = "default",           -- Resource regrowth speed
    frograin = "default",           -- Frog rain frequency
    wildfires = "default",          -- Summer wildfires
  }
}
```

### Caves Example

```lua
return {
  override_enabled = true,
  preset = "DST_CAVE",
  overrides = {
    world_size = "default",
    branching = "default",

    -- Cave-specific
    mushtree = "default",
    wormlights = "default",
    bunnymen = "default",
    slurtles = "default",
    rocky = "default",
    cave_spiders = "default",

    -- Ruins
    ancient_altar = "default",
    ancient_statues = "default",
  }
}
```

### Override Values

Most overrides use this scale:

| Value | Effect |
|-------|--------|
| `"never"` | Disabled |
| `"rare"` | Very uncommon |
| `"uncommon"` | Below normal |
| `"default"` | Standard |
| `"often"` | Above normal |
| `"mostly"` | Very common |
| `"always"` | Maximum |

## Admin Commands

Type in the server console or use `c_` prefix in the remote console:

### Player Management

| Command | Description |
|---------|-------------|
| `c_listallplayers()` | List all players with IDs |
| `c_kick("KU_xxxxxx")` | Kick by Klei ID |
| `c_ban("KU_xxxxxx")` | Ban by Klei ID |
| `c_rollback(count)` | Rollback `count` saves |
| `c_reset()` | Regenerate the world |
| `c_save()` | Force save |
| `c_shutdown()` | Graceful shutdown |

### Admin Lists

In the cluster directory:

**adminlist.txt** — One Klei User ID per line:
```
KU_abc12345
KU_def67890
```

**whitelist.txt** — Whitelisted players:
```
KU_abc12345
```

**blocklist.txt** — Banned players:
```
KU_banned123
```

Find Klei IDs in the server log when players connect, or from `c_listallplayers()`.

### Useful Admin Commands

```lua
-- Announce a message
c_announce("Server restarting in 5 minutes!")

-- Teleport to player
c_goto("playername")

-- Give item to nearest player
c_give("cutgrass", 40)

-- Skip to next day
c_skip(480)  -- seconds

-- Set season
TheWorld:PushEvent("ms_setseason", "winter")

-- Regenerate world
c_regenerateworld()
```

## Season Configuration

Seasons are a huge part of DST. Here's what each season brings:

| Season | Duration (default) | Key Threat |
|--------|-------------------|------------|
| Autumn | 20 days | Starting season, mild |
| Winter | 15 days | Freezing, Deerclops |
| Spring | 20 days | Rain, lightning, Moose/Goose |
| Summer | 15 days | Overheating, wildfires, Antlion |

To change season lengths, use `worldgenoverride.lua`:
- `"veryshort"` = 5 days
- `"short"` = 12 days
- `"default"` = normal
- `"long"` = 30 days
- `"verylong"` = 50 days
- `"noseason"` = skip entirely

## Performance Notes

- DST servers are lightweight — the game is 2D and relatively simple to simulate
- **Mod count** is the biggest performance factor. Each mod adds processing overhead
- **World age** — Older worlds accumulate more entities (dropped items, structures)
- **Tick rate** — 15 is default and fine for most servers. 30 is smoother but doubles CPU usage
- **Caves add ~50% resource overhead** — running two processes instead of one
