---
title: "Linux Basics for Game Servers"
description: "A crash course in Linux for people setting up their first game server — distros, users, systemd, and logs."
order: 2
tags: ["linux", "basics", "setup"]
---

# Linux Basics for Game Servers

If you're here, you probably know your way around a game but not around a terminal. That's fine — running a game server doesn't require deep Linux knowledge, just a handful of commands and concepts.

## Picking a Distro

For game servers, it doesn't matter much. Pick one and learn it.

| Distro | Why |
|--------|-----|
| **Debian 12** | Rock stable, minimal, widely supported. Most hosting guides assume Debian/Ubuntu |
| **Ubuntu Server 24.04 LTS** | Debian-based, slightly more batteries-included. Largest community for troubleshooting |
| **Fedora Server** | Newer packages, good systemd integration. Less common in game server guides |
| **Arch** | Rolling release, always up to date. Only if you already know Linux |

**If you're unsure: pick Ubuntu Server LTS or Debian.** You'll find the most help online.

Always use the **server** variant (no desktop environment). A GUI wastes RAM that your game server needs.

## Connecting to Your Server

You'll interact with your server over SSH:

```bash
# From your local machine
ssh username@your-server-ip

# If using a key (recommended — see the security guide)
ssh -i ~/.ssh/my_key username@your-server-ip
```

On Windows, use Windows Terminal with the built-in SSH client, or PuTTY.

## Essential Commands

### Navigation

```bash
pwd                     # Print current directory
ls                      # List files
ls -la                  # List all files with details
cd /opt/gameserver      # Change directory
cd ..                   # Go up one level
```

### File Operations

```bash
cp file.txt backup.txt          # Copy
mv old.txt new.txt              # Move/rename
rm file.txt                     # Delete (no undo!)
rm -r directory/                # Delete directory and contents
mkdir -p /opt/gameserver        # Create directory (and parents)
```

### Reading Files

```bash
cat file.txt                    # Print entire file
less file.txt                   # Scroll through file (q to quit)
tail -f server.log              # Follow a log file in real time
head -20 config.ini             # First 20 lines
```

### Editing Files

```bash
nano config.ini                 # Simple editor (Ctrl+O save, Ctrl+X exit)
vim config.ini                  # Powerful editor (press i to type, Esc then :wq to save and quit)
```

nano is fine. Don't let anyone tell you otherwise.

### Permissions

```bash
chmod +x start-server.sh        # Make a script executable
chown -R gameuser:gameuser /opt/gameserver  # Change ownership
```

## Users: Don't Run Servers as Root

This is the single most important thing in this guide. **Never run a game server as root.**

If your game server gets exploited (and they do — game servers are constantly probed), the attacker gets whatever permissions the server process has. Root means they own your entire machine.

Create a dedicated user for each game server:

```bash
# Create a user with no login shell and a home directory
sudo useradd -r -m -s /bin/bash gameserver

# Or for a specific game
sudo useradd -r -m -s /bin/bash minecraft
sudo useradd -r -m -s /bin/bash valheim
```

Then run the server as that user:

```bash
# Switch to the user
sudo -u minecraft bash

# Or run a single command as that user
sudo -u minecraft ./start-server.sh
```

## systemd: Managing Services

systemd is how Linux starts, stops, and monitors services. Every game server guide in this documentation uses systemd service files.

### Key Commands

```bash
# Start/stop/restart a service
sudo systemctl start minecraft
sudo systemctl stop minecraft
sudo systemctl restart minecraft

# Enable auto-start on boot
sudo systemctl enable minecraft

# Disable auto-start
sudo systemctl disable minecraft

# Check if it's running
systemctl status minecraft

# See recent logs
journalctl -u minecraft --no-pager -n 50

# Follow logs in real time
journalctl -u minecraft -f
```

### Writing a Service File

Service files live in `/etc/systemd/system/`. Here's the anatomy of a game server service:

```ini
# /etc/systemd/system/minecraft.service
[Unit]
Description=Minecraft Server          # Human-readable name
After=network.target                  # Wait for networking

[Service]
Type=simple                           # Process runs in foreground
User=minecraft                        # Run as this user (NOT root)
WorkingDirectory=/opt/minecraft       # cd here before running
ExecStart=/opt/minecraft/start.sh     # The actual command
Restart=on-failure                    # Restart if it crashes
RestartSec=10                         # Wait 10 seconds before restart

[Install]
WantedBy=multi-user.target            # Start at boot (when enabled)
```

After creating or editing a service file:

```bash
# Reload systemd to pick up changes
sudo systemctl daemon-reload

# Then start it
sudo systemctl start minecraft
```

### Common Service Issues

**"Service failed to start"** — Check the logs:
```bash
journalctl -u minecraft -n 30 --no-pager
```

**"Permission denied"** — The `User` in the service file doesn't have access to the files. Fix with:
```bash
sudo chown -R minecraft:minecraft /opt/minecraft
```

**"ExecStart path is not absolute"** — Use full paths, not relative ones. `/opt/minecraft/start.sh`, not `./start.sh`.

## Reading Logs with journalctl

journalctl is your best friend for debugging. Every service that runs through systemd logs here.

```bash
# All logs for a service
journalctl -u valheim

# Last 100 lines
journalctl -u valheim -n 100

# Since last boot
journalctl -u valheim -b

# Since a specific time
journalctl -u valheim --since "2024-01-15 14:00"

# Follow in real time (like tail -f)
journalctl -u valheim -f

# Only errors
journalctl -u valheim -p err
```

## Package Management

### Debian/Ubuntu (apt)

```bash
sudo apt update                   # Refresh package lists
sudo apt upgrade                  # Upgrade installed packages
sudo apt install steamcmd         # Install a package
sudo apt remove steamcmd          # Remove a package
```

### Fedora (dnf)

```bash
sudo dnf update
sudo dnf install steamcmd
```

### Arch (pacman)

```bash
sudo pacman -Syu                  # Update everything
sudo pacman -S steamcmd           # Install a package
```

## Disk Space

Game servers eat disk space. Check regularly:

```bash
# Overall disk usage
df -h

# What's using space in a directory
du -sh /opt/*

# Find large files
find /opt -type f -size +100M
```

## Processes

```bash
# See what's running
htop                              # Interactive process viewer (install with apt install htop)
ps aux | grep valheim             # Find a specific process

# Kill a stuck process
kill <pid>                        # Graceful
kill -9 <pid>                     # Force (last resort)
```

## Next Steps

- [Security Hardening](/self-hosting/security) — SSH keys, fail2ban, and keeping your server safe
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
