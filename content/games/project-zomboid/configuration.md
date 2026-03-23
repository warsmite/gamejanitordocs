---
title: "Project Zomboid Configuration Reference"
description: "Sandbox settings, zombie population, loot respawn, and server.ini options for Project Zomboid servers."
order: 2
tags: ["project-zomboid", "configuration"]
---

# Project Zomboid Configuration Reference

## Config Files

Project Zomboid uses three config files in `~/Zomboid/Server/`:

- **`<servername>.ini`** — Server settings (network, admin, Steam)
- **`<servername>_SandboxVars.lua`** — Gameplay/sandbox settings (zombie behavior, loot, etc.)
- **`<servername>_spawnregions.lua`** — Spawn point definitions

## Server.ini — Key Settings

### Network

| Setting | Default | Description |
|---------|---------|-------------|
| `MaxPlayers` | `32` | Max connected players |
| `PingLimit` | `400` | Kick players above this ping (ms) |
| `Public` | `false` | List in server browser |
| `PublicName` | — | Server name in browser |
| `ServerPassword` | — | Password to join (empty = no password) |
| `PauseEmpty` | `true` | Pause game when no players online |
| `DefaultPort` | `16261` | Primary port |
| `UDPPort` | `16262` | Secondary UDP port |

### Gameplay

| Setting | Default | Description |
|---------|---------|-------------|
| `PVP` | `false` | Player vs player damage |
| `SafetySystem` | `true` | Allow players to toggle PvP individually |
| `SafetyCooldownTimer` | `3` | Hours before PvP flag takes effect |
| `SpawnItems` | — | Items given on spawn (comma-separated item IDs) |
| `NoFire` | `false` | Disable fire spread (anti-grief) |
| `AnnounceDeath` | `false` | Broadcast player deaths |
| `SteamScoreboard` | `true` | Show Steam names in scoreboard |

### Anti-Grief

| Setting | Default | Description |
|---------|---------|-------------|
| `AllowDestructionBySledgehammer` | `true` | Allow destroying player-built walls |
| `SledgehammerOnlyInSafehouse` | `false` | Restrict sledgehammer to safehouse owners |
| `SafehouseAllowTrepass` | `true` | Allow non-members to enter safehouses |
| `SafehouseAllowLoot` | `true` | Allow looting in others' safehouses |
| `SafehouseAllowRespawn` | `false` | Respawn in safehouse on death |
| `SafehouseDaySurvivedToClaim` | `0` | Days survived before claiming a safehouse |
| `SafeHouseRemovalTime` | `144` | Hours before abandoned safehouse unclaims |
| `PlayerSafehouse` | `true` | Enable safehouse system |

### Mod Configuration

```ini
WorkshopItems=2392987599;2478768005
Mods=eris_minimap;Arsenal(26)GunFighter
```

## SandboxVars.lua — Gameplay Settings

This is where the real customization happens. The file is Lua syntax.

### Zombie Population

```lua
SandboxVars = {
    ZombiePopulationModifier = 1.0,         -- Global zombie count multiplier (0.0-4.0)
    PopulationMultiplierStart = 1.0,         -- Starting population (day 0)
    PopulationMultiplierPeak = 1.5,          -- Peak population
    PopulationMultiplierPeakDay = 28,        -- Day peak is reached
    PopulationMultiplierAfterPeak = 0.5,     -- Population after peak
    RespawnHours = 72.0,                     -- Hours before zombies respawn in cleared areas
    RespawnUnseenHours = 16.0,               -- Hours unseen before respawn can happen
    RespawnMultiplier = 0.1,                 -- Respawn rate (0 = no respawn)
    RedistributeHours = 12.0,               -- Hours between zombie redistribution
    FollowSoundDistance = 100,               -- Tile distance zombies hear sounds
    Zombies = 4,                             -- 1=Insane, 2=Very High, 3=High, 4=Normal, 5=Low, 6=None
}
```

### Zombie Behavior

```lua
SandboxVars = {
    Speed = 2,                               -- 1=Sprinters, 2=Fast Shamblers, 3=Shamblers, 4=Random
    Strength = 2,                            -- 1=Superhuman, 2=Normal, 3=Weak
    Toughness = 2,                           -- 1=Tough, 2=Normal, 3=Fragile
    Transmission = 1,                        -- 1=Blood+Saliva, 2=Saliva Only, 3=Everyone's Infected
    Mortality = 5,                           -- 1=Instant, 2=0-30s, 3=0-1min, 4=0-12h, 5=2-3 days, 6=1-2 weeks, 7=Never
    Cognition = 1,                           -- 1=Navigate+Use Doors, 2=Navigate, 3=Basic Navigation
    Memory = 2,                              -- 1=Long, 2=Normal, 3=Short, 4=None
    Sight = 2,                               -- 1=Eagle, 2=Normal, 3=Poor
    Hearing = 2,                             -- 1=Pinpoint, 2=Normal, 3=Poor
}
```

### Popular Zombie Presets

**Romero (classic slow zombies):**
```lua
Speed = 3
Strength = 2
Toughness = 2
Cognition = 3
Memory = 3
Sight = 2
Hearing = 2
```

**28 Days Later (sprinters):**
```lua
Speed = 1
Strength = 1
Toughness = 3
Cognition = 1
Memory = 1
Sight = 1
Hearing = 1
```

**Mixed population (some slow, some fast):**
```lua
ActiveOnly = 1  -- Zombies are active only during night
Speed = 2       -- Base speed
-- Use mods like "Slow Shamblers + Sprinters" for true mixed populations
```

### Loot Settings

```lua
SandboxVars = {
    LootRarity = 3,                          -- 1=Extremely Rare, 2=Rare, 3=Normal, 4=Common, 5=Abundant
    HoursForLootRespawn = 0,                -- Hours for loot respawn (0 = never)
    MaxItemsForLootRespawn = 4,             -- Items in container to prevent respawn
    ConstructionBonusPoints = 0,            -- Bonus construction XP
    NightLength = 2,                        -- 1=Always Night, 2=Long, 3=Normal, 4=Short
    DayLength = 2,                          -- 1=15min, 2=30min, 3=1h, 4=2h, ..., 24=12h
    StartMonth = 7,                         -- Starting month (1-12). July is default
    StartDay = 9,                           -- Starting day
    WaterShutModifier = 14,                 -- Days until water shutoff (0 = instant, -1 = never)
    ElecShutModifier = 14,                  -- Days until power shutoff (0 = instant, -1 = never)
    FoodLoot = 2,                           -- 1=Very Rare through 5=Abundant
    WeaponLoot = 2,
    OtherLoot = 2,
}
```

### XP and Skills

```lua
SandboxVars = {
    XPMultiplier = 1.0,                     -- Global XP rate
    XPMultiplierAffectsPassive = true,      -- Whether multiplier affects passive skill gain
    StatsDecrease = 3,                      -- Stat drain speed (1=Very Fast...5=Very Slow)
    NutritionEnabled = true,                -- Enable nutrition system
    BonusFitnessPoints = 0,                -- Extra starting fitness
    BonusStrengthPoints = 0,               -- Extra starting strength
}
```

## Admin Console Commands

| Command | Description |
|---------|-------------|
| `/players` | List online players |
| `/adduser "user" "pass"` | Create account |
| `/addusertowhitelist "user"` | Whitelist player |
| `/removeuserfromwhitelist "user"` | Remove from whitelist |
| `/grantadmin "user"` | Grant admin |
| `/banuser "user"` | Ban player |
| `/unbanuser "user"` | Unban player |
| `/kickuser "user"` | Kick player |
| `/teleport "user" "user2"` | Teleport player to player |
| `/save` | Force world save |
| `/quit` | Graceful shutdown |
| `/changeoption <setting> <value>` | Change server.ini setting at runtime |

## Performance Notes

- **Map exploration drives RAM usage.** Each cell loaded takes memory. Large servers where players spread out need more RAM
- **Zombie count is the main CPU driver.** Reduce `ZombiePopulationModifier` if the server struggles
- **Fire spread is expensive.** `NoFire=true` prevents griefing AND improves performance
- **`PauseEmpty=true`** saves resources when nobody's online. Always enable for private servers
