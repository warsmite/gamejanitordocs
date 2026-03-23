---
title: "ARK Configuration Reference"
description: "GameUserSettings.ini and Game.ini reference for ARK: Survival Evolved servers — rates, taming, breeding, and more."
order: 2
tags: ["ark", "configuration"]
---

# ARK Configuration Reference

ARK has two main config files in `ShooterGame/Saved/Config/LinuxServer/`:

- **GameUserSettings.ini** — Most server settings
- **Game.ini** — Advanced overrides (per-dino stats, engrams, etc.)

## GameUserSettings.ini

### [ServerSettings]

#### Rates

| Setting | Default | Description |
|---------|---------|-------------|
| `TamingSpeedMultiplier` | `1.0` | Taming speed. 5.0+ is common for unofficial |
| `HarvestAmountMultiplier` | `1.0` | Resource gathering rate |
| `XPMultiplier` | `1.0` | Experience gain rate |
| `MatingIntervalMultiplier` | `1.0` | Time between breeding. Lower = faster |
| `BabyMatureSpeedMultiplier` | `1.0` | How fast babies grow. 10+ is common |
| `EggHatchSpeedMultiplier` | `1.0` | Egg hatching speed |
| `BabyCuddleIntervalMultiplier` | `1.0` | Time between imprint cuddles. Lower = faster |
| `CropGrowthSpeedMultiplier` | `1.0` | Crop growing speed |
| `LayEggIntervalMultiplier` | `1.0` | How often dinos lay eggs |
| `HairGrowthSpeedMultiplier` | `1.0` | Hair/wool growth |

#### Popular Unofficial Rate Presets

**Casual PvE:**
```ini
TamingSpeedMultiplier=5.0
HarvestAmountMultiplier=3.0
XPMultiplier=3.0
BabyMatureSpeedMultiplier=20.0
EggHatchSpeedMultiplier=20.0
MatingIntervalMultiplier=0.1
BabyCuddleIntervalMultiplier=0.1
```

**Boosted PvP:**
```ini
TamingSpeedMultiplier=10.0
HarvestAmountMultiplier=5.0
XPMultiplier=5.0
BabyMatureSpeedMultiplier=50.0
EggHatchSpeedMultiplier=50.0
MatingIntervalMultiplier=0.01
```

#### Player & Dino Stats

| Setting | Default | Description |
|---------|---------|-------------|
| `PlayerCharacterWaterDrainMultiplier` | `1.0` | Water consumption rate |
| `PlayerCharacterFoodDrainMultiplier` | `1.0` | Food consumption rate |
| `PlayerCharacterStaminaDrainMultiplier` | `1.0` | Stamina drain rate |
| `DinoCharacterFoodDrainMultiplier` | `1.0` | Dino food consumption |
| `PlayerCharacterHealthRecoveryMultiplier` | `1.0` | Health regen rate |
| `DinoDamageMultiplier` | `1.0` | Dino damage output |
| `PlayerDamageMultiplier` | `1.0` | Player damage output |
| `PlayerResistanceMultiplier` | `1.0` | Player damage taken (lower = tankier) |
| `DinoResistanceMultiplier` | `1.0` | Dino damage taken |

#### Structure Settings

| Setting | Default | Description |
|---------|---------|-------------|
| `StructureResistanceMultiplier` | `1.0` | Structure damage taken |
| `PvEStructureDecayPeriodMultiplier` | `1.0` | PvE decay timer |
| `PerPlatformMaxStructuresMultiplier` | `1.0` | Max structures on platform saddles |
| `MaxStructuresInRange` | `6000` | Max structures within render range |
| `AutoDestroyOldStructuresMultiplier` | `0.0` | Auto-demolish timer. 0 = disabled |
| `RCONPort` | `27020` | RCON port |
| `RCONEnabled` | `false` | Enable RCON |

#### PvP/PvE Settings

| Setting | Default | Description |
|---------|---------|-------------|
| `ServerPVE` | `false` | PvE mode |
| `AllowCaveBuildingPvE` | `false` | Allow building in caves (PvE) |
| `EnablePvPGamma` | `false` | Allow gamma changes in PvP |
| `DisableFriendlyFire` | `false` | Prevent tribe damage |
| `PreventTribeAlliances` | `false` | Disable alliances |
| `MaxTribeLogs` | `100` | Tribe log entries |

### [SessionSettings]

```ini
[SessionSettings]
SessionName=My ARK Server
```

### [/Script/ShooterGame.ShooterGameMode]

Advanced settings in `GameUserSettings.ini` or `Game.ini`:

```ini
[/Script/ShooterGame.ShooterGameMode]
bDisableDinoDecayPvE=false
bAllowFlyerCarryPvE=true
MaxNumberOfPlayersInTribe=10
bDisableStructureDecayPvE=false
```

## Game.ini

This file handles per-dino and per-engram overrides.

### Level Caps

```ini
[/Script/ShooterGame.ShooterGameMode]
LevelExperienceRampOverrides=(ExperiencePointsForLevel[0]=5,ExperiencePointsForLevel[1]=20,...)
OverrideMaxExperiencePointsPlayer=500000
OverrideMaxExperiencePointsDino=500000
DestroyTamesOverLevel=450
```

### Disable Specific Dinos

```ini
[/Script/ShooterGame.ShooterGameMode]
NPCReplacements=(FromClassName="Gigant_Character_BP_C",ToClassName="")
```

### Custom Engram Points

```ini
[/Script/ShooterGame.ShooterGameMode]
OverridePlayerLevelEngramPoints=10
OverridePlayerLevelEngramPoints=12
OverridePlayerLevelEngramPoints=14
```

Each line corresponds to a level (level 1, 2, 3...).

## Admin Commands

Use `enablecheats <AdminPassword>` in-game console, then:

| Command | Description |
|---------|-------------|
| `admincheat SaveWorld` | Force save |
| `admincheat DestroyWildDinos` | Kill all wild dinos (forces respawn) |
| `admincheat ListPlayers` | List connected players |
| `admincheat KickPlayer <steamid>` | Kick player |
| `admincheat BanPlayer <steamid>` | Ban player |
| `admincheat Fly` | Toggle flight |
| `admincheat God` | Toggle invincibility |
| `admincheat GiveItemNum <id> <qty> <quality> <bp>` | Give items |
| `admincheat SetTimeOfDay <HH:MM>` | Set time |
| `admincheat Slomo <rate>` | Time speed (1.0 = normal) |

## Performance Tips

- **`DestroyWildDinos`** regularly — wild dino count builds up and tanks performance. Many admins schedule this weekly
- **Limit max structures** — `MaxStructuresInRange=6000` is default. Lower it if bases are causing lag
- **Disable BattlEye** — `-NoBattlEye` saves meaningful CPU. Use if you trust your players
- **Lower `ViewDistanceMultiplier`** — Clients can set this, but capping it server-side helps
- **SSD is mandatory** — ARK saves are large and frequent. HDDs cause freeze spikes during saves
