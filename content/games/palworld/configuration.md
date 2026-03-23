---
title: "Palworld Configuration Reference"
description: "PalWorldSettings.ini explained: XP rates, Pal capture rates, difficulty, and all server settings."
order: 2
tags: ["palworld", "configuration"]
---

# Palworld Configuration Reference

## Config File Location

Settings live in `Pal/Saved/Config/LinuxServer/PalWorldSettings.ini` (or `WindowsServer` on Windows).

On first run, this file is mostly empty. Copy the defaults from `DefaultPalWorldSettings.ini` in the root directory:

```bash
cp /opt/palworld/DefaultPalWorldSettings.ini /opt/palworld/Pal/Saved/Config/LinuxServer/PalWorldSettings.ini
```

**Important:** The entire settings block must be on a single line under `[/Script/Pal.PalGameWorldSettings]`. Yes, one enormous line. That's how Unreal Engine does it.

## Key Settings

The format is `OptionSettings=(Setting1=Value1,Setting2=Value2,...)`. Here are the settings that matter:

### Server Identity

| Setting | Default | Description |
|---------|---------|-------------|
| `ServerName` | `"Default Palworld Server"` | Shown in server browser |
| `ServerDescription` | `""` | Server description |
| `AdminPassword` | `""` | Admin password for in-game commands |
| `ServerPassword` | `""` | Password to join |
| `ServerPlayerMaxNum` | `32` | Max players |
| `PublicPort` | `8211` | Game port |

### Difficulty & Rates

| Setting | Default | Description |
|---------|---------|-------------|
| `DayTimeSpeedRate` | `1.0` | How fast daytime passes |
| `NightTimeSpeedRate` | `1.0` | How fast nighttime passes |
| `ExpRate` | `1.0` | XP multiplier. 2.0 = double XP |
| `PalCaptureRate` | `1.0` | Pal catch rate. Higher = easier catches |
| `PalSpawnNumRate` | `1.0` | Wild Pal density. Higher = more Pals |
| `EnemyDropItemRate` | `1.0` | Drop rate from enemies |
| `DeathPenalty` | `All` | What you lose on death: `None`, `Item`, `ItemAndEquipment`, `All` |
| `GuildPlayerMaxNum` | `20` | Max players per guild |

### Damage Rates

| Setting | Default | Description |
|---------|---------|-------------|
| `PalDamageRateAttack` | `1.0` | Damage Pals deal |
| `PalDamageRateDefense` | `1.0` | Damage Pals receive (lower = tankier) |
| `PlayerDamageRateAttack` | `1.0` | Damage players deal |
| `PlayerDamageRateDefense` | `1.0` | Damage players receive |
| `PlayerStomachDecreaceRate` | `1.0` | Hunger drain speed. 0.5 = half as fast |
| `PlayerStaminaDecreaceRate` | `1.0` | Stamina drain speed |

### Base & Building

| Setting | Default | Description |
|---------|---------|-------------|
| `BaseCampMaxNum` | `128` | Total bases allowed on server |
| `BaseCampMaxNumInGuild` | `3` | Bases per guild |
| `BaseCampWorkerMaxNum` | `15` | Pals per base |
| `BuildObjectDeteriorationDamageRate` | `1.0` | Structure decay rate. 0 = no decay |
| `BuildObjectDamageRate` | `1.0` | Damage structures take |
| `bIsMultiplay` | `true` | Multiplayer mode |
| `bCanPickupOtherGuildDeathPenaltyDrop` | `false` | Whether others can loot your death drops |

### Raid Settings

| Setting | Default | Description |
|---------|---------|-------------|
| `bEnableInvaderEnemy` | `true` | Base raids enabled |
| `bEnableDefenseOtherGuildPlayer` | `false` | PvP base raiding |
| `bEnableNonLoginPenalty` | `true` | Whether offline players' bases can be damaged |

## Popular Presets

### Casual PvE

```
ExpRate=2.0,PalCaptureRate=1.5,PalSpawnNumRate=1.5,PlayerStomachDecreaceRate=0.5,PlayerStaminaDecreaceRate=0.5,DeathPenalty=None,BuildObjectDeteriorationDamageRate=0.0,BaseCampWorkerMaxNum=20
```

### Hardcore PvP

```
ExpRate=1.0,PalCaptureRate=0.8,DeathPenalty=All,bEnableDefenseOtherGuildPlayer=true,bCanPickupOtherGuildDeathPenaltyDrop=true,PlayerDamageRateAttack=1.0,PlayerDamageRateDefense=1.0
```

### Boosted Rates (Catch up with friends)

```
ExpRate=5.0,PalCaptureRate=2.0,EnemyDropItemRate=3.0,PalSpawnNumRate=2.0
```

## Admin Commands

Set `AdminPassword` in the config, then in-game use `/AdminPassword <password>` to authenticate. Key commands:

| Command | What it does |
|---------|-------------|
| `/Shutdown <seconds> <message>` | Graceful shutdown with warning |
| `/Save` | Force save |
| `/BanPlayer <steamid>` | Ban a player |
| `/KickPlayer <steamid>` | Kick a player |
| `/Broadcast <message>` | Send message to all players |
| `/TeleportToPlayer <steamid>` | Teleport to a player |
| `/ShowPlayers` | List connected players |

## RCON

Palworld supports RCON for remote administration:

| Setting | Default | Description |
|---------|---------|-------------|
| `RCONEnabled` | `false` | Enable RCON |
| `RCONPort` | `25575` | RCON port |

Enable RCON if you want to run commands without being logged in.
