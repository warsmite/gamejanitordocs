---
title: "CS2 Configuration Reference"
description: "Server.cfg, competitive settings, RCON, SourceTV, and plugin setup for Counter-Strike 2 servers."
order: 2
tags: ["cs2", "counter-strike", "configuration"]
---

# CS2 Configuration Reference

## Config Files

CS2 reads configs from `game/csgo/cfg/`. The main file is `server.cfg`:

```
// server.cfg - loaded on every map change

hostname "My CS2 Server"
rcon_password "yourrconpassword"
sv_password ""
sv_cheats 0

// Gameplay
mp_autoteambalance 1
mp_limitteams 1
mp_friendlyfire 0
mp_autokick 1

// Round settings
mp_maxrounds 24
mp_roundtime 1.92
mp_roundtime_defuse 1.92
mp_roundtime_hostage 1.92
mp_freezetime 15
mp_buytime 20

// Economy
mp_startmoney 800
mp_afterroundmoney 0
cash_team_bonus_shorthanded 0

// Warmup
mp_warmuptime 60
mp_do_warmup_period 1

// Communication
sv_alltalk 0
sv_deadtalk 1
sv_voiceenable 1

// Server performance
sv_maxrate 0
sv_minrate 128000
sv_maxupdaterate 128
sv_mincmdrate 128
```

## Important ConVars

### Server Settings

| ConVar | Default | Description |
|--------|---------|-------------|
| `hostname` | — | Server name |
| `rcon_password` | — | RCON password. Leave empty to disable |
| `sv_password` | `""` | Server join password |
| `sv_cheats` | `0` | Enable cheat commands |
| `sv_lan` | `0` | LAN-only mode |
| `sv_maxrate` | `0` | Max bytes/sec per client. 0 = unlimited |
| `sv_pure` | `1` | File consistency check. 1 = validate, 2 = strict |

### Match Settings

| ConVar | Default | Description |
|--------|---------|-------------|
| `mp_maxrounds` | `24` | Rounds per half |
| `mp_roundtime` | `1.92` | Round time in minutes |
| `mp_freezetime` | `15` | Freeze time in seconds |
| `mp_buytime` | `20` | Buy time in seconds |
| `mp_startmoney` | `800` | Starting money |
| `mp_halftime` | `1` | Enable halftime |
| `mp_overtime_enable` | `0` | Enable overtime |
| `mp_overtime_maxrounds` | `6` | Overtime rounds |
| `mp_overtime_startmoney` | `10000` | Overtime starting money |
| `mp_warmuptime` | `60` | Warmup duration in seconds |

### Competitive Config

For a standard competitive match:

```
// competitive.cfg
game_type 0
game_mode 1
mp_maxrounds 24
mp_roundtime_defuse 1.92
mp_freezetime 15
mp_buytime 20
mp_startmoney 800
mp_overtime_enable 1
mp_overtime_maxrounds 6
mp_overtime_startmoney 10000
mp_halftime 1
mp_match_end_restart 0
sv_alltalk 0
sv_deadtalk 0
mp_friendlyfire 1
mp_autoteambalance 0
mp_limitteams 0
```

### Practice Mode Config

```
// practice.cfg
sv_cheats 1
mp_limitteams 0
mp_autoteambalance 0
mp_roundtime_defuse 60
mp_maxmoney 60000
mp_startmoney 60000
mp_buytime 9999
mp_freezetime 0
mp_warmuptime 0
sv_infinite_ammo 1
ammo_grenade_limit_total 5
sv_grenade_trajectory_prac_pipreview 1
mp_restartgame 1
```

## RCON

Set `rcon_password` in `server.cfg`, then connect remotely:

```bash
# Using any Source RCON tool
rcon -a <server-ip>:27015 -p "yourrconpassword" "status"
```

Or from within the game console:

```
rcon_address <ip>:27015
rcon_password yourrconpassword
rcon status
```

## SourceTV (GOTV)

SourceTV lets spectators watch matches with a delay:

```
tv_enable 1
tv_port 27020
tv_delay 90
tv_maxclients 128
tv_name "GOTV"
tv_autorecord 1
```

| Setting | Default | Description |
|---------|---------|-------------|
| `tv_enable` | `0` | Enable SourceTV |
| `tv_port` | `27020` | SourceTV port |
| `tv_delay` | `90` | Spectator delay in seconds |
| `tv_maxclients` | `128` | Max spectators |
| `tv_autorecord` | `0` | Auto-record demos |

## Map Rotation

Edit `game/csgo/gamemodes_server.txt` to define map pools:

```
"gamemodes_server.txt"
{
    "mapgroups"
    {
        "mg_custom"
        {
            "name" "mg_custom"
            "maps"
            {
                "de_dust2" ""
                "de_mirage" ""
                "de_inferno" ""
                "de_nuke" ""
                "de_anubis" ""
                "de_ancient" ""
                "de_overpass" ""
            }
        }
    }
}
```

Then launch with `+mapgroup mg_custom`.

## Plugins (MetaMod / CounterStrikeSharp)

CS2 uses **CounterStrikeSharp** (CSS) as the modern plugin framework, running on MetaMod:Source:

```bash
# Install MetaMod
# Download from https://www.sourcemm.net/downloads.php?branch=master
# Extract to game/csgo/

# Install CounterStrikeSharp
# Download from https://github.com/roflmuffin/CounterStrikeSharp/releases
# Extract to game/csgo/addons/counterstrikesharp/
```

Plugins are C# DLLs placed in `game/csgo/addons/counterstrikesharp/plugins/`.

### Essential Plugins

- **MatchZy** — competitive match management (ready-up, knife round, veto)
- **CS2-SimpleAdmin** — admin commands (kick, ban, mute)
- **SharpTimer** — surf/KZ timer
- **Deathmatch** — FFA deathmatch mode
- **RetakesPlugin** — retakes game mode
