---
title: "Enshrouded Configuration Reference"
description: "Game settings, difficulty tuning, time cycles, and enemy scaling for Enshrouded dedicated servers."
order: 2
tags: ["enshrouded", "configuration"]
---

# Enshrouded Configuration Reference

All settings are in `enshrouded_server.json`. Stop the server before editing.

## Game Settings

### Player Stats

| Setting | Default | Description |
|---------|---------|-------------|
| `playerHealthFactor` | `1` | Player health multiplier |
| `playerManaFactor` | `1` | Player mana multiplier |
| `playerStaminaFactor` | `1` | Player stamina multiplier |
| `enableDurability` | `true` | Equipment durability system |

### Enemy Settings

| Setting | Default | Description |
|---------|---------|-------------|
| `enemyDamageFactor` | `1` | Damage enemies deal |
| `enemyHealthFactor` | `1` | Enemy health multiplier |
| `enemyStaminaFactor` | `1` | Enemy stamina/stagger |
| `randomSpawnerAmount` | `"Normal"` | Enemy spawn density: `"None"`, `"Few"`, `"Normal"`, `"Many"`, `"Extreme"` |
| `aggroPoolAmount` | `"Normal"` | How many enemies aggro at once: `"None"`, `"Few"`, `"Normal"`, `"Many"`, `"Extreme"` |

### Resource & Crafting

| Setting | Default | Description |
|---------|---------|-------------|
| `miningDamageFactor` | `1` | Mining speed (higher = faster) |
| `resourceDropStackAmountFactor` | `1` | Resource drop multiplier |
| `factoryProductionSpeedFactor` | `1` | Crafting station speed |
| `plantGrowthSpeedFactor` | `1` | Farm crop growth speed |

### Death & Penalty

| Setting | Default | Description |
|---------|---------|-------------|
| `tombstoneMode` | `"AddBackpackMaterials"` | What happens on death |

Tombstone mode options:
- `"AddBackpackMaterials"` — Drop backpack contents as tombstone, keep equipped items
- `"Everything"` — Drop everything
- `"NoTombstone"` — Keep everything on death

### Time Cycle

| Setting | Default | Description |
|---------|---------|-------------|
| `dayTimeDuration` | `1800` | Day length in seconds (30 minutes) |
| `nightTimeDuration` | `720` | Night length in seconds (12 minutes) |
| `pacingOfTime` | `1` | Overall time speed multiplier |
| `shroudTimeFactor` | `1` | How fast shroud timer depletes |

The default day/night cycle is 42 minutes total (30 day + 12 night).

## Popular Presets

### Casual/Relaxed

```json
{
  "gameSettings": {
    "playerHealthFactor": 1.5,
    "enemyDamageFactor": 0.5,
    "enemyHealthFactor": 0.75,
    "tombstoneMode": "NoTombstone",
    "resourceDropStackAmountFactor": 2,
    "miningDamageFactor": 2,
    "factoryProductionSpeedFactor": 2,
    "enableDurability": false,
    "shroudTimeFactor": 0.5
  }
}
```

### Hardcore

```json
{
  "gameSettings": {
    "playerHealthFactor": 0.75,
    "enemyDamageFactor": 1.5,
    "enemyHealthFactor": 1.5,
    "tombstoneMode": "Everything",
    "resourceDropStackAmountFactor": 0.5,
    "randomSpawnerAmount": "Many",
    "aggroPoolAmount": "Many",
    "shroudTimeFactor": 1.5,
    "nightTimeDuration": 1200
  }
}
```

### Builder Focus

```json
{
  "gameSettings": {
    "enemyDamageFactor": 0.5,
    "resourceDropStackAmountFactor": 3,
    "miningDamageFactor": 3,
    "factoryProductionSpeedFactor": 3,
    "plantGrowthSpeedFactor": 3,
    "enableDurability": false,
    "tombstoneMode": "NoTombstone",
    "randomSpawnerAmount": "Few"
  }
}
```

## Performance Notes

- Enshrouded is still in Early Access — expect optimization improvements over time
- **Base complexity** is the primary performance driver, not player count
- **The Shroud** (fog areas) are more demanding to simulate than open areas
- **16 players max** is a hard limit — the game isn't designed for larger populations
- **Memory usage** typically 3-6GB depending on world exploration
- **SSD recommended** for save/load performance
