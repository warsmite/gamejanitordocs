---
title: "Terraria Configuration Reference"
description: "TShock configuration, permissions system, and server settings for Terraria dedicated servers."
order: 2
tags: ["terraria", "tshock", "configuration"]
---

# Terraria Configuration Reference

## TShock Configuration

TShock stores its config at `tshock/config.json`. This is generated on first run — edit it while the server is stopped.

### Key Settings

```json
{
  "ServerPassword": "",
  "MaxSlots": 16,
  "SpawnProtection": true,
  "SpawnProtectionRadius": 10,
  "RangeChecks": true,
  "ServerSideCharacter": false,
  "AutoSave": true,
  "AutoSaveInterval": 5,
  "HardcoreOnly": false,
  "MediumcoreOnly": false,
  "DisableBuild": false,
  "DisableInvisPvP": false,
  "MaxDamage": 1175,
  "MaxProjectileDamage": 1175,
  "RestApiEnabled": false,
  "RestApiPort": 7878
}
```

### Settings That Actually Matter

| Setting | Default | Recommendation |
|---------|---------|---------------|
| `ServerSideCharacter` | `false` | **Set to `true`** for public servers. Without this, players can cheat items via inventory editors |
| `SpawnProtection` | `true` | Keep on. Prevents griefing the spawn area |
| `SpawnProtectionRadius` | `10` | Tiles around spawn that are protected. 10-20 is reasonable |
| `AutoSaveInterval` | `5` | Minutes between auto-saves. 5 is fine, lower if you're paranoid about crashes |
| `RangeChecks` | `true` | Keep on for public servers. Blocks abnormal tile placement and out-of-range interactions |
| `MaxDamage` | `1175` | Max damage per hit before anti-cheat flags it. Increase if using mods that add stronger weapons |
| `RestApiEnabled` | `false` | Enable if you want remote management via HTTP. Set a strong `RestApiPort` password |

## Permission System

TShock uses a group-based permission system. Groups are hierarchical.

### Default Groups

| Group | Purpose |
|-------|---------|
| `guest` | Unregistered players |
| `default` | Registered players |
| `vip` | Extra permissions (skip queue, etc.) |
| `newadmin` | Junior admins |
| `admin` | Standard admins |
| `trustedadmin` | Senior admins |
| `owner` | Full server control |

Note: `superadmin` exists as a special internal group but isn't assigned directly — use `owner` for full-access accounts.

### Common Permission Commands

```
# In-game commands (requires admin)
/group addperm <group> <permission>
/group delperm <group> <permission>
/group list
/user group <player> <group>
```

### Useful Permissions

| Permission | What it allows |
|------------|---------------|
| `tshock.world.modify` | Place/break blocks |
| `tshock.tp.self` | Teleport to coordinates |
| `tshock.tp.others` | Teleport to other players |
| `tshock.kick` | Kick players |
| `tshock.ban.manage` | Ban/unban players |
| `tshock.world.time.set` | Change time of day |
| `tshock.npc.butcher` | Kill all NPCs |
| `tshock.item.give` | Spawn items |
| `tshock.item.spawn` | Spawn items for others |

For public servers, `default` group should have minimal permissions — `tshock.world.modify` and basic chat. Add permissions as trust is earned.

## Region Protection

Regions let you protect specific areas of the world:

```
# Define a region (stand at one corner, note coordinates with /pos)
/region define <name> <x1> <y1> <x2> <y2>

# Protect it
/region protect <name> true

# Allow a user to build in it
/region allow <name> <player>
```

## Server-Side Characters (SSC)

When `ServerSideCharacter` is `true`, character data (inventory, health, mana) is stored on the server instead of the client. This is critical for public servers because:

- Players can't dupe items via save manipulation
- Players can't bring in end-game items from singleplayer
- Character progress is tied to the server

The downside: players lose their existing characters when SSC is enabled. It's a fresh start for everyone.

### SSC Settings

```json
{
  "ServerSideCharacter": true,
  "StartingHealth": 100,
  "StartingMana": 20,
  "StartingInventory": [
    {
      "netID": -15,
      "prefix": 0,
      "stack": 1
    }
  ]
}
```

## Vanilla Server Config (serverconfig.txt)

If running the vanilla server, the config file uses a simpler format:

```
world=/opt/terraria/worlds/MyWorld.wld
autocreate=3
worldname=MyWorld
difficulty=0
maxplayers=16
port=7777
password=secretpassword
motd=Welcome to the server!
worldpath=/opt/terraria/worlds/
banlist=banlist.txt
secure=1
language=en-US
upnp=0
npcstream=60
priority=1
```

### Difficulty Values

| Value | Mode | Description |
|-------|------|-------------|
| 0 | Classic | Standard mode |
| 1 | Expert | Harder enemies, exclusive drops |
| 2 | Master | Even harder, exclusive vanity |
| 3 | Journey | Creative-ish, adjustable difficulty |

Difficulty is set at world creation and **cannot be changed after**. Choose carefully.
