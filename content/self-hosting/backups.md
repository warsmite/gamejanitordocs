---
title: "Backups & Disaster Recovery"
description: "Automated backup strategies for game servers — cron scripts, off-site copies, and how to actually restore."
order: 4
tags: ["backups", "cron", "disaster-recovery"]
---

# Backups & Disaster Recovery

Your world file is hundreds of hours of player effort. A corrupted save, a bad update, or an `rm -rf` in the wrong directory can erase all of it. Backups are not optional.

## The Backup Strategy

A good backup setup has three properties:

1. **Automated** — If you have to remember to do it, it won't happen
2. **Versioned** — Keep multiple backups so you can go back further than the last save
3. **Off-site** — At least one copy lives on a different machine or cloud storage

## What to Back Up

Every game stores world data differently, but the pattern is the same: find the save directory and copy it.

| Game | Save Location |
|------|--------------|
| Minecraft | `<server>/world/` |
| Valheim | `-savedir` path or `~/.config/unity3d/IronGate/Valheim/worlds_local/` |
| Terraria | Path from `-world` flag |
| Palworld | `Pal/Saved/SaveGames/` |
| Factorio | `saves/` directory |
| Rust | `server/<identity>/` |
| ARK | `ShooterGame/Saved/` |
| Project Zomboid | `~/Zomboid/Saves/Multiplayer/<servername>/` |
| 7 Days to Die | Saves directory |
| Satisfactory | `~/.config/Epic/FactoryGame/Saved/SaveGames/server/` |
| V Rising | `Saves/v3/<savename>/` |
| Enshrouded | `savegame/` directory |
| DST | `~/.klei/DoNotStarveTogether/<cluster>/` |

## Basic Backup Script

```bash
#!/bin/bash
# /opt/scripts/backup-gameserver.sh

GAME_NAME="minecraft"
SAVE_DIR="/opt/minecraft/world"
BACKUP_DIR="/backups/${GAME_NAME}"
MAX_BACKUPS=48  # Keep 48 backups (2 days at every-hour schedule)

TIMESTAMP=$(date +%Y%m%d-%H%M%S)
BACKUP_FILE="${BACKUP_DIR}/${GAME_NAME}-${TIMESTAMP}.tar.gz"

mkdir -p "${BACKUP_DIR}"

# Create compressed backup
tar -czf "${BACKUP_FILE}" -C "$(dirname "${SAVE_DIR}")" "$(basename "${SAVE_DIR}")"

# Check if backup was created successfully
if [ $? -eq 0 ]; then
    echo "[backup] ${GAME_NAME}: created ${BACKUP_FILE} ($(du -sh "${BACKUP_FILE}" | cut -f1))"
else
    echo "[backup] ${GAME_NAME}: FAILED to create backup" >&2
    exit 1
fi

# Prune old backups — keep only the most recent $MAX_BACKUPS
cd "${BACKUP_DIR}"
ls -t ${GAME_NAME}-*.tar.gz 2>/dev/null | tail -n +$((MAX_BACKUPS + 1)) | xargs -r rm
echo "[backup] ${GAME_NAME}: pruned to ${MAX_BACKUPS} backups"
```

Make it executable:

```bash
chmod +x /opt/scripts/backup-gameserver.sh
```

## Scheduling with Cron

```bash
# Edit crontab
crontab -e

# Back up every hour
0 * * * * /opt/scripts/backup-gameserver.sh >> /var/log/gameserver-backup.log 2>&1

# Back up every 6 hours
0 */6 * * * /opt/scripts/backup-gameserver.sh >> /var/log/gameserver-backup.log 2>&1

# Back up daily at 4 AM
0 4 * * * /opt/scripts/backup-gameserver.sh >> /var/log/gameserver-backup.log 2>&1
```

Pick a frequency based on how much progress you're willing to lose. For active servers, every 1-2 hours is reasonable.

## Pre-Save Backup (RCON Integration)

Some games should be saved before backing up to avoid copying a half-written file. If the game supports RCON:

```bash
#!/bin/bash
# Save before backup (Minecraft example using rcon-cli)

RCON_PASS="yourrconpassword"
RCON_PORT=25575

# Tell the server to save and pause auto-save
mcrcon -H localhost -P ${RCON_PORT} -p ${RCON_PASS} "save-all" "save-off"
sleep 5  # Wait for save to complete

# Run the backup
/opt/scripts/backup-gameserver.sh

# Re-enable auto-save
mcrcon -H localhost -P ${RCON_PORT} -p ${RCON_PASS} "save-on"
```

## Off-Site Backups with rsync

Local backups protect against game corruption. Off-site backups protect against disk failure, accidental deletion, or the machine dying entirely.

### To Another Server

```bash
# Sync backups to a remote machine
rsync -avz --delete /backups/minecraft/ backup-user@remote-server:/backups/minecraft/
```

### To Cloud Storage with rclone

rclone works with S3, Backblaze B2, Google Drive, and dozens of other providers:

```bash
# Install rclone
sudo apt install rclone

# Configure (interactive)
rclone config

# Sync backups to cloud
rclone sync /backups/minecraft/ remote:gameserver-backups/minecraft/
```

Backblaze B2 is the cheapest option for backup storage (~$0.005/GB/month). For a typical game server generating 500MB of backups, that's essentially free.

### Cron for Off-Site

```bash
# Sync to remote daily at 5 AM (after the 4 AM local backup)
0 5 * * * rclone sync /backups/ remote:gameserver-backups/ >> /var/log/offsite-backup.log 2>&1
```

## Restoring from Backup

This is the part people forget to practice. **Test your restore process before you need it.**

### Basic Restore

```bash
# Stop the server
sudo systemctl stop minecraft

# Move the broken world out of the way (don't delete it yet)
mv /opt/minecraft/world /opt/minecraft/world-broken

# Extract the backup
tar -xzf /backups/minecraft/minecraft-20240115-040000.tar.gz -C /opt/minecraft/

# Fix permissions
sudo chown -R minecraft:minecraft /opt/minecraft/world

# Start the server
sudo systemctl start minecraft

# Verify it's working
journalctl -u minecraft -f
```

### Listing Available Backups

```bash
# See what backups exist, newest first
ls -lht /backups/minecraft/ | head -20
```

## How Much Space Do Backups Use?

Compressed backup sizes vary by game and world age:

| Game | Typical Backup Size |
|------|-------------------|
| Terraria | 5-20 MB |
| Factorio | 10-100 MB |
| DST | 10-50 MB |
| Minecraft | 50-500 MB |
| Valheim | 50-200 MB |
| Project Zomboid | 100-500 MB |
| Rust | 100-500 MB |
| ARK | 500 MB - 2 GB |

With hourly backups and 48 retained copies, plan for 10-50x a single backup in total storage.

## Pre-Update Backups

Always back up before updating the game server. Updates can break save compatibility:

```bash
# Manual pre-update backup
/opt/scripts/backup-gameserver.sh

# Then update
steamcmd +force_install_dir /opt/gameserver +login anonymous +app_update <appid> +quit
```

## Next Steps

- [Automated Updates & Restarts](/self-hosting/updates-and-restarts) — automate the update process too
- [Security Hardening](/self-hosting/security) — protect your backup data
