---
title: "7 Days to Die Configuration Reference"
description: "serverconfig.xml settings, blood moon tuning, loot and difficulty, and performance optimization."
order: 2
tags: ["7-days-to-die", "configuration"]
---

# 7 Days to Die Configuration Reference

All settings are in `serverconfig.xml`. Stop the server before editing.

## Difficulty Settings

| Property | Default | Description |
|----------|---------|-------------|
| `GameDifficulty` | `2` | 0=Scavenger, 1=Adventurer, 2=Nomad, 3=Warrior, 4=Survivalist, 5=Insane |
| `ZombieMove` | `0` | 0=Walk, 1=Jog, 2=Run, 3=Sprint, 4=Nightmare |
| `ZombieMoveNight` | `3` | Same scale as above, but at night |
| `ZombieFeralMove` | `3` | Feral zombie speed |
| `ZombieBMMove` | `3` | Blood moon zombie speed |
| `PlayerKillingMode` | `3` | 0=No PvP, 1=Allies only, 2=Strangers only, 3=Everyone |

### Blood Moon

The blood moon horde is the defining feature of 7DTD. Tune it carefully:

| Property | Default | Description |
|----------|---------|-------------|
| `BloodMoonFrequency` | `7` | Days between blood moons |
| `BloodMoonRange` | `0` | Random variance in days. 0 = exact frequency |
| `BloodMoonWarning` | `8` | Hour the red sky warning starts |
| `BloodMoonEnemyCount` | `8` | Max alive zombies per player during horde. **This is the big one** |
| `AirDropFrequency` | `72` | Hours between airdrops. 0 = disabled |

`BloodMoonEnemyCount` is per-player and directly impacts server performance. For 8 players at the default of 8, that's potentially 64 zombies active simultaneously. On weaker hardware, lower this to 4-6.

## Day/Night Cycle

| Property | Default | Description |
|----------|---------|-------------|
| `DayNightLength` | `60` | Real minutes for a full 24-hour day |
| `DayLightLength` | `18` | In-game hours of daylight |

Common setups:
- **Shorter days (40 min):** More blood moons, faster pace
- **Longer days (90 min):** More time to build and explore
- **18 hours daylight:** Standard. 6 hours of dangerous night
- **12 hours daylight:** Harder. Half the day is night

## Loot Settings

| Property | Default | Description |
|----------|---------|-------------|
| `LootAbundance` | `100` | Loot amount percentage. 200 = double loot |
| `LootRespawnDays` | `7` | Days before looted containers refill |
| `BlockDamagePlayer` | `100` | Percentage of block damage by players |
| `BlockDamageAI` | `100` | Percentage of block damage by zombies |
| `BlockDamageAIBM` | `100` | Block damage during blood moon |
| `XPMultiplier` | `100` | XP gain percentage |
| `PartySharedKillRange` | `100` | Range (meters) for shared kill XP |

## Land Claim

| Property | Default | Description |
|----------|---------|-------------|
| `LandClaimSize` | `41` | Land claim block protection radius |
| `LandClaimDeadZone` | `30` | Min distance between claims |
| `LandClaimExpiryTime` | `7` | Days before inactive claims expire |
| `LandClaimDecayMode` | `0` | 0=Linear, 1=Exponential, 2=Full protection until expiry |
| `LandClaimOnlineDurabilityModifier` | `4` | Block durability multiplier when owner is online |
| `LandClaimOfflineDurabilityModifier` | `4` | Block durability when owner is offline |

For PvP servers, `LandClaimOfflineDurabilityModifier` is critical — it determines how hard it is to raid bases when the owner is offline. Lower = easier raids.

## Crafting & Progression

| Property | Default | Description |
|----------|---------|-------------|
| `CraftTimer` | `1.0` | Crafting speed multiplier (lower = faster) |
| `LootTimer` | `1.0` | Loot speed multiplier |
| `PlayerSafeZoneLevel` | `5` | Zombie-free zone around new players until this gamestage |
| `PlayerSafeZoneHours` | `5` | Hours of safe zone protection |

## Performance Settings

| Property | Default | Description |
|----------|---------|-------------|
| `MaxSpawnedZombies` | `64` | Total zombies alive on the server |
| `MaxSpawnedAnimals` | `50` | Total animals alive on the server |
| `EnemySpawnMode` | `true` | Spawn enemies. `false` = peaceful |
| `EnemySenseMemory` | `60` | Seconds zombies remember a player |

### Performance Recommendations

- **`MaxSpawnedZombies`** — Lower this first if performance is bad. 40-50 for weaker hardware
- **`BloodMoonEnemyCount`** — Per-player count. Lower for large player counts
- **World size** — Smaller worlds = less to simulate. 6144 is the sweet spot
- **SSD storage** — World saves are frequent and large. HDDs cause server freezes

## Console Commands (via Telnet)

| Command | Description |
|---------|-------------|
| `listplayers` | List connected players |
| `kick <name/id> "reason"` | Kick player |
| `ban add <name/id> <duration> <unit> "reason"` | Ban player (duration: number, unit: minutes/hours/days/years) |
| `ban remove <name/id>` | Unban |
| `shutdown` | Graceful shutdown |
| `saveworld` | Force save |
| `settime day` | Set to daytime |
| `settime night` | Set to nighttime |
| `weather <type>` | Change weather |
| `spawnsupply` | Trigger airdrop |
| `bloodmoon` | Check days until next blood moon |
| `admin add <steamid> <level>` | Grant admin (0=highest) |

## Mod Support

7DTD supports XML modding and has a modding API:

- Mods go in the `Mods/` directory
- Server-side mods (XML overrides) don't require client installation
- Overhaul mods (Darkness Falls, Undead Legacy) require matching client mods
- EAC (Easy Anti-Cheat) must be disabled for most mods: set `EACEnabled` to `false` in config
