---
title: "Satisfactory Configuration Reference"
description: "Server settings, API usage, auto-save configuration, and performance notes for Satisfactory dedicated servers."
order: 2
tags: ["satisfactory", "configuration"]
---

# Satisfactory Configuration Reference

## Configuration Method

Unlike most game servers, Satisfactory manages most settings through the in-game Server Manager UI and the Server API rather than config files. There are limited file-based settings.

## Engine Configuration

Engine settings can be tweaked in `FactoryGame/Saved/Config/LinuxServer/Engine.ini`:

```ini
[/Script/Engine.Player]
ConfiguredInternetSpeed=104857600
ConfiguredLanSpeed=104857600

[/Script/OnlineSubsystemUtils.IpNetDriver]
MaxClientRate=104857600
MaxInternetClientRate=104857600

[/Script/SocketSubsystemEpic.EpicNetDriver]
MaxClientRate=104857600
MaxInternetClientRate=104857600
```

These network rate settings help with larger factories where lots of data needs to sync.

## Game Configuration

`FactoryGame/Saved/Config/LinuxServer/Game.ini`:

```ini
[/Script/Engine.GameSession]
MaxPlayers=4
```

## Server Settings (via Server Manager)

These are set in-game through the Server Manager:

| Setting | Description |
|---------|-------------|
| Server Name | Displayed in server browser |
| Admin Password | Required for server management |
| Client Password | Required to join (blank = no password) |
| Auto-Save Interval | How often the server auto-saves (in seconds) |
| Auto-Pause | Pause when no players are connected |
| Auto-Save on Disconnect | Save when the last player leaves |

## Server API

The REST API is the primary way to manage the server programmatically.

### Authentication

```bash
# Get an auth token
curl -k -X POST https://localhost:7777/api/v1 \
  -H "Content-Type: application/json" \
  -d '{"function":"PasswordLogin","data":{"MinimumPrivilegeLevel":"Administrator","Password":"youradminpassword"}}'
```

### Common API Calls

```bash
# Query server state
curl -k -X POST https://localhost:7777/api/v1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"function":"QueryServerState"}'

# Save the game
curl -k -X POST https://localhost:7777/api/v1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"function":"SaveGame","data":{"SaveName":"ManualSave"}}'

# Load a save
curl -k -X POST https://localhost:7777/api/v1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"function":"LoadGame","data":{"SaveName":"MySave","EnableAdvancedGameSettings":false}}'

# Shutdown
curl -k -X POST https://localhost:7777/api/v1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"function":"Shutdown"}'

# Get server options
curl -k -X POST https://localhost:7777/api/v1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"function":"GetServerOptions"}'
```

### Available API Functions

| Function | Description |
|----------|-------------|
| `HealthCheck` | Server health status (no auth required) |
| `QueryServerState` | Current server state, players, tick rate |
| `GetServerOptions` | Current settings |
| `ApplyServerOptions` | Update settings |
| `CreateNewGame` | Start a new game |
| `SaveGame` | Save with a given name |
| `LoadGame` | Load a specific save |
| `DeleteSaveFile` | Delete a save |
| `EnumerateSessions` | List available saves |
| `Shutdown` | Graceful shutdown |

## Advanced Game Settings

Satisfactory supports "Advanced Game Settings" that change gameplay rules. These are set per-save:

| Setting | Description |
|---------|-------------|
| No Build Cost | Buildings don't require resources |
| No Power | Buildings don't need power |
| No Fuel | Vehicles don't need fuel |
| Set Starting Tier | Start with specific tiers unlocked |
| Set Game Phase | Start at a specific space elevator phase |
| Flight Mode | Allow flying |

These are set when creating/loading a save through the Server Manager.

## Performance Notes

### What Causes Server Lag

1. **Factory complexity** — More machines, conveyors, and logistics = more simulation. This is the #1 factor
2. **Vehicles** — Autonomous vehicles (trucks, trains) are expensive to simulate, especially many trucks
3. **Conveyor belt count** — Each belt segment is a simulated entity
4. **Player count** — Each player needs their own replication stream

### What You Can Do

- **Network rate settings** in `Engine.ini` help with sync lag in large factories
- **Auto-save interval** — Large saves can cause hitches during auto-save. Increase the interval for megabases
- **SSD storage** — Save/load operations are much faster on SSD
- **The server does NOT need a GPU** — it's entirely headless
- **Memory grows with factory size** — Monitor and allocate accordingly

### Monitoring

Use the API's `QueryServerState` to monitor:
- Average tick rate (target: 30)
- Connected player count
- Server uptime
