---
title: "V Rising Configuration Reference"
description: "ServerGameSettings.json explained: PvP schedules, castle limits, difficulty, and gameplay settings."
order: 2
tags: ["v-rising", "configuration"]
---

# V Rising Configuration Reference

## ServerGameSettings.json

This file controls all gameplay mechanics. Located in your save data path.

### Game Mode Presets

V Rising supports presets that set many values at once:

| Preset | Description |
|--------|-------------|
| `StandardPvP` | Full PvP with raiding |
| `StandardPvE` | PvE, no raiding |
| `PvPEasy` | Relaxed PvP rules |
| `PvPHard` | Full loot PvP, harder settings |
| `DuoPvP` | Optimized for 2-player clans |

Set `GameSettingsPreset` in `ServerHostSettings.json`, or leave blank and customize everything.

### Key Gameplay Settings

```json
{
  "GameModeType": "PvP",
  "ClanSize": 4,
  "PlayerDamageMode": "Always",
  "CastleDamageMode": "TimeRestricted",
  "SiegeWeaponHealth": "Normal",
  "PlayerInteractionSettings": {
    "TimeZone": "Local",
    "VSPlayerWeekdayTime": { "StartHour": 17, "StartMinute": 0, "EndHour": 23, "EndMinute": 0 },
    "VSPlayerWeekendTime": { "StartHour": 17, "StartMinute": 0, "EndHour": 23, "EndMinute": 0 },
    "VSCastleWeekdayTime": { "StartHour": 17, "StartMinute": 0, "EndHour": 23, "EndMinute": 0 },
    "VSCastleWeekendTime": { "StartHour": 17, "StartMinute": 0, "EndHour": 23, "EndMinute": 0 }
  }
}
```

### PvP Settings

| Setting | Values | Description |
|---------|--------|-------------|
| `GameModeType` | `PvP`, `PvE` | Base game mode |
| `PlayerDamageMode` | `Always`, `TimeRestricted`, `Never` | When players can damage each other |
| `CastleDamageMode` | `Always`, `TimeRestricted`, `Never` | When castles can be raided |
| `ClanSize` | 1-10 | Max players per clan |

### Time-Restricted PvP

The `PlayerInteractionSettings` block lets you set PvP and raiding windows. This is huge for servers where players have jobs — nobody wants their castle raided at 3 AM.

Common setups:
- **Evenings only:** 17:00-23:00
- **Weekends extended:** 14:00-01:00
- **Always on:** Set `PlayerDamageMode` to `Always`
- **Never:** PvE mode

### Castle Settings

```json
{
  "CastleLimit": 2,
  "CastleHeartDamageMode": "CanBeDestroyedByPlayers",
  "CastleUnderAttackTimer": 60.0,
  "CastleDecayRateModifier": 1.0,
  "CastleBloodEssenceDrainModifier": 1.0,
  "CastleSiegeTimer": 420.0,
  "FreeCastleClaim": false,
  "FreeCastleDestroy": false
}
```

| Setting | Default | Description |
|---------|---------|-------------|
| `CastleLimit` | `2` | Max castles per player |
| `CastleDecayRateModifier` | `1.0` | How fast castles decay without blood essence |
| `CastleBloodEssenceDrainModifier` | `1.0` | Blood essence consumption rate. Lower = less upkeep |
| `CastleSiegeTimer` | `420` | Seconds for a siege to complete |

### Difficulty & Rates

```json
{
  "UnitStatModifiers_Global": {
    "MaxHealthModifier": 1.0,
    "PowerModifier": 1.0,
    "ResourceYieldModifier": 1.0
  },
  "UnitStatModifiers_VBlood": {
    "MaxHealthModifier": 1.0,
    "PowerModifier": 1.0
  },
  "EquipmentStatModifiers_Global": {
    "MaxEnergyModifier": 1.0,
    "MaxHealthModifier": 1.0,
    "PhysicalPowerModifier": 1.0,
    "SpellPowerModifier": 1.0,
    "ResourceYieldModifier": 1.0
  },
  "CraftRateModifier": 1.0,
  "ResearchCostModifier": 1.0,
  "RefinementCostModifier": 1.0,
  "RefinementRateModifier": 1.0,
  "ResearchTimeModifier": 1.0,
  "CraftingTimeModifier": 1.0,
  "DismantleResourceModifier": 0.75,
  "BloodDrainModifier": 1.0,
  "DurabilityDrainModifier": 1.0,
  "GarlicAreaStrengthModifier": 1.0,
  "HolyAreaStrengthModifier": 1.0,
  "SilverStrengthModifier": 1.0,
  "SunDamageModifier": 1.0,
  "DropTableModifier_General": 1.0,
  "DropTableModifier_Missions": 1.0,
  "MaterialYieldModifier_Global": 1.0,
  "InventoryStacksModifier": 1.0
}
```

### Popular Presets

**Casual PvE:**
```json
"GameModeType": "PvE",
"ResourceYieldModifier": 2.0,
"CraftingTimeModifier": 0.5,
"CastleBloodEssenceDrainModifier": 0.5,
"InventoryStacksModifier": 2.0
```

**Competitive PvP:**
```json
"GameModeType": "PvP",
"CastleDamageMode": "TimeRestricted",
"PlayerDamageMode": "Always",
"ClanSize": 4,
"CastleLimit": 1,
"ResourceYieldModifier": 1.5
```

**Solo-Friendly:**
```json
"ClanSize": 1,
"CastleLimit": 3,
"UnitStatModifiers_VBlood": { "MaxHealthModifier": 0.5, "PowerModifier": 0.5 },
"ResourceYieldModifier": 2.0,
"CraftingTimeModifier": 0.5
```

### Day/Night Cycle

```json
{
  "DayDurationInSeconds": 1080,
  "BloodMoonFrequency": 10,
  "VSPlayerWeekdayTime": {
    "StartHour": 17,
    "StartMinute": 0,
    "EndHour": 23,
    "EndMinute": 0
  }
}
```

## RCON

Enable in `ServerHostSettings.json`:

```json
"Rcon": {
  "Enabled": true,
  "Port": 25575,
  "Password": "yourrconpassword"
}
```

### RCON Commands

| Command | Description |
|---------|-------------|
| `adminauth` | Authenticate as admin |
| `listusers` | List connected users |
| `kick <steamid>` | Kick player |
| `ban <steamid>` | Ban player |
| `unban <steamid>` | Unban player |
| `save` | Force save |
| `shutdown` | Graceful shutdown |
| `announce <message>` | Server-wide message |

## Performance Notes

- **ServerFps** in `ServerHostSettings.json` defaults to 30. This is the server tick rate — don't increase it without good hardware
- **Castle complexity** drives server load more than player count
- **Auto-save causes brief hitches** — adjust `AutoSaveInterval` based on world size
- **Memory usage** is typically 2-4GB and scales with explored map and castle count
