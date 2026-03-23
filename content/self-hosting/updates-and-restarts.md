---
title: "Automated Updates & Restarts"
description: "Keep your game servers updated and running smoothly with SteamCMD automation, scheduled restarts, and graceful shutdowns."
order: 6
tags: ["updates", "steamcmd", "cron", "automation"]
---

# Automated Updates & Restarts

Game servers need regular updates and periodic restarts. Doing this manually gets old fast. Automate it.

## SteamCMD Update Scripts

Most game servers are installed via SteamCMD. A simple update script:

```bash
#!/bin/bash
# /opt/scripts/update-gameserver.sh

INSTALL_DIR="/opt/valheim"
APP_ID="896660"
SERVICE_NAME="valheim"

echo "[update] $(date): checking for updates..."

# Stop the server gracefully
sudo systemctl stop ${SERVICE_NAME}
sleep 5

# Run SteamCMD update
steamcmd +force_install_dir ${INSTALL_DIR} +login anonymous +app_update ${APP_ID} validate +quit

if [ $? -eq 0 ]; then
    echo "[update] ${SERVICE_NAME}: update complete"
else
    echo "[update] ${SERVICE_NAME}: update FAILED" >&2
fi

# Start the server
sudo systemctl start ${SERVICE_NAME}
echo "[update] ${SERVICE_NAME}: server started"
```

### Update on Boot (systemd ExecStartPre)

The simplest approach — update every time the service starts:

```ini
[Service]
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/valheim +login anonymous +app_update 896660 +quit
ExecStart=/opt/valheim/valheim_server.x86_64 ...
```

This adds 10-30 seconds to startup if there's no update, or a few minutes if there is one. For servers that don't restart often, this is the easiest approach.

### Scheduled Updates via Cron

For more control, schedule updates at a specific time:

```bash
# Update at 5 AM daily
0 5 * * * /opt/scripts/update-gameserver.sh >> /var/log/gameserver-update.log 2>&1
```

## Graceful Shutdowns

Don't just `kill` or `systemctl stop` your game server without saving first. Many games have RCON commands for graceful shutdown:

### RCON Save-Then-Stop

```bash
#!/bin/bash
# /opt/scripts/graceful-stop.sh

GAME="$1"  # minecraft, rust, etc.

case ${GAME} in
    minecraft)
        mcrcon -H localhost -P 25575 -p "rconpass" "say Server restarting in 30 seconds..."
        sleep 30
        mcrcon -H localhost -P 25575 -p "rconpass" "save-all" "stop"
        ;;
    rust)
        rcon -a localhost:28016 -p "rconpass" "say Server restarting in 30 seconds..."
        sleep 30
        rcon -a localhost:28016 -p "rconpass" "server.save"
        rcon -a localhost:28016 -p "rconpass" "quit"
        ;;
    factorio)
        mcrcon -H localhost -P 27015 -p "rconpass" "/save"
        sleep 5
        mcrcon -H localhost -P 27015 -p "rconpass" "/quit"
        ;;
    *)
        echo "Unknown game: ${GAME}"
        sudo systemctl stop "${GAME}"
        ;;
esac
```

### systemd ExecStop

You can integrate graceful shutdown directly into the service file:

```ini
[Service]
ExecStop=/opt/scripts/graceful-stop.sh minecraft
TimeoutStopSec=60
```

`TimeoutStopSec` gives the server time to save before systemd force-kills it.

## Scheduled Restarts

Some games (Palworld, Rust, Valheim) benefit from periodic restarts to clear memory leaks or accumulated state:

```bash
# Restart every 6 hours
0 */6 * * * sudo systemctl restart palworld

# Restart daily at 4 AM
0 4 * * * sudo systemctl restart minecraft

# Restart with warning (requires RCON)
0 4 * * * /opt/scripts/graceful-stop.sh minecraft && sleep 10 && sudo systemctl start minecraft
```

### Restart with Player Warning

Players hate surprise restarts. Warn them first:

```bash
#!/bin/bash
# /opt/scripts/scheduled-restart.sh

SERVICE="$1"
RCON_CMD="$2"  # Full rcon command for broadcasting

# 5 minute warning
${RCON_CMD} "Server restarting in 5 minutes"
sleep 240

# 1 minute warning
${RCON_CMD} "Server restarting in 60 seconds"
sleep 50

# 10 second warning
${RCON_CMD} "Server restarting in 10 seconds"
sleep 10

# Stop and restart
sudo systemctl restart ${SERVICE}
```

## Update + Backup + Restart (All-in-One)

The ideal maintenance script combines all three:

```bash
#!/bin/bash
# /opt/scripts/maintenance.sh

GAME="minecraft"
SERVICE="minecraft"
APP_ID="0"  # Set to 0 for non-Steam games
INSTALL_DIR="/opt/minecraft"
BACKUP_SCRIPT="/opt/scripts/backup-gameserver.sh"

echo "[maintenance] $(date): starting maintenance for ${GAME}"

# 1. Warn players (if RCON available)
# mcrcon -H localhost -P 25575 -p "pass" "say Server maintenance in 60 seconds"
# sleep 60

# 2. Stop the server
sudo systemctl stop ${SERVICE}
sleep 5

# 3. Backup
${BACKUP_SCRIPT}

# 4. Update (if Steam game)
if [ "${APP_ID}" != "0" ]; then
    steamcmd +force_install_dir ${INSTALL_DIR} +login anonymous +app_update ${APP_ID} +quit
fi

# 5. Start
sudo systemctl start ${SERVICE}

echo "[maintenance] $(date): maintenance complete for ${GAME}"
```

Schedule it:

```bash
# Daily maintenance at 4 AM
0 4 * * * /opt/scripts/maintenance.sh >> /var/log/maintenance.log 2>&1
```

## Monitoring Update Success

Check that your game server actually came back up after an update:

```bash
#!/bin/bash
# /opt/scripts/health-check.sh

SERVICE="minecraft"
PORT=25565

sleep 30  # Give it time to start

if systemctl is-active --quiet ${SERVICE}; then
    # Check if the port is actually listening
    if ss -tlnp | grep -q ":${PORT}"; then
        echo "[health] ${SERVICE}: running and listening on port ${PORT}"
    else
        echo "[health] ${SERVICE}: running but NOT listening on port ${PORT}" >&2
    fi
else
    echo "[health] ${SERVICE}: NOT running after update" >&2
fi
```

## Next Steps

- [Backups & Disaster Recovery](/self-hosting/backups) — always back up before updating
- [Performance Monitoring](/self-hosting/performance-monitoring) — make sure restarts are actually helping
