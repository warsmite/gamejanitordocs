---
title: "Security Hardening"
description: "Essential security for game servers — SSH keys, fail2ban, firewalls, and not running as root."
order: 5
tags: ["security", "ssh", "firewall", "fail2ban"]
---

# Security Hardening

Game servers are constantly scanned and probed. Your server will receive SSH brute-force attempts within minutes of going online. This isn't hypothetical — it's guaranteed. Here's how to not be an easy target.

## 1. SSH Key Authentication

Passwords get brute-forced. SSH keys don't.

### Generate a Key Pair (on your local machine)

```bash
# Generate an Ed25519 key (modern, fast, secure)
ssh-keygen -t ed25519 -C "your-email@example.com"

# It creates two files:
# ~/.ssh/id_ed25519       (private key — NEVER share this)
# ~/.ssh/id_ed25519.pub   (public key — goes on the server)
```

### Copy the Public Key to Your Server

```bash
ssh-copy-id username@your-server-ip
```

### Disable Password Authentication

Once you've confirmed key login works:

```bash
sudo nano /etc/ssh/sshd_config
```

Set these values:

```
PasswordAuthentication no
PubkeyAuthentication yes
PermitRootLogin no
```

Restart SSH:

```bash
sudo systemctl restart sshd
```

**Test from a new terminal before closing your current session.** If you lock yourself out, you'll need console access from your hosting provider.

## 2. Fail2ban

Fail2ban watches log files and temporarily bans IPs that show malicious signs (like repeated failed SSH logins).

```bash
# Install
sudo apt install fail2ban

# Create local config (don't edit the default)
sudo cp /etc/fail2ban/jail.conf /etc/fail2ban/jail.local
sudo nano /etc/fail2ban/jail.local
```

Key settings in `jail.local`:

```ini
[DEFAULT]
bantime = 1h
findtime = 10m
maxretry = 5

[sshd]
enabled = true
port = ssh
logpath = %(sshd_log)s
backend = %(sshd_backend)s
```

This bans an IP for 1 hour after 5 failed SSH attempts in 10 minutes.

```bash
# Start and enable
sudo systemctl enable --now fail2ban

# Check banned IPs
sudo fail2ban-client status sshd
```

## 3. Firewall: Default Deny

Only open the ports you need. Everything else should be blocked.

### UFW (Ubuntu/Debian)

```bash
# Set default policies
sudo ufw default deny incoming
sudo ufw default allow outgoing

# Allow SSH (always do this BEFORE enabling)
sudo ufw allow ssh

# Allow your game port(s)
sudo ufw allow 25565/tcp    # Example: Minecraft

# Enable the firewall
sudo ufw enable

# Check status
sudo ufw status verbose
```

### firewalld (Fedora/RHEL)

```bash
# Allow SSH (should be default)
sudo firewall-cmd --permanent --add-service=ssh

# Allow game ports
sudo firewall-cmd --permanent --add-port=25565/tcp

# Reload
sudo firewall-cmd --reload

# Check
sudo firewall-cmd --list-all
```

### What NOT to Open

- **Telnet admin ports** (7DTD port 8081, etc.) — restrict to localhost or specific IPs
- **RCON ports** — if you must expose them, use strong passwords and consider IP whitelisting
- **Web panels** — put them behind a reverse proxy with HTTPS, or restrict to your IP

```bash
# Allow RCON only from your IP
sudo ufw allow from 203.0.113.50 to any port 25575 proto tcp
```

## 4. Don't Run as Root

This is covered in the [Linux Basics](/self-hosting/linux-basics) guide but it bears repeating. Create dedicated users:

```bash
sudo useradd -r -m -s /bin/bash minecraft
```

If a game server process gets exploited, the attacker only has access to that user's files — not the entire system.

## 5. Keep the System Updated

Unpatched vulnerabilities are how servers get compromised.

### Manual Updates

```bash
# Debian/Ubuntu
sudo apt update && sudo apt upgrade

# Fedora
sudo dnf upgrade
```

### Automatic Security Updates

On Debian/Ubuntu, install unattended-upgrades for automatic security patches:

```bash
sudo apt install unattended-upgrades
sudo dpkg-reconfigure unattended-upgrades
```

This applies security patches automatically without upgrading everything. It won't touch your game servers — just system packages.

## 6. RCON Security

RCON gives full server control. Treat it seriously:

- **Use strong, unique passwords** — not the same password as your server join password
- **Don't expose RCON to the internet** unless necessary. Bind to localhost and use SSH tunneling:

```bash
# SSH tunnel to access RCON locally
ssh -L 25575:localhost:25575 user@your-server

# Now connect to localhost:25575 on your machine
```

- **Use WebSocket RCON** where available (Rust) — it's more secure than legacy RCON
- **Monitor RCON access** — check logs for unexpected connections

## 7. File Permissions

Game server files should be owned by the game user and not world-readable:

```bash
# Set ownership
sudo chown -R minecraft:minecraft /opt/minecraft

# Remove world read/write (optional but good practice)
sudo chmod -R o-rwx /opt/minecraft
```

Config files with passwords should be especially restricted:

```bash
chmod 600 /opt/minecraft/server.properties
```

## Quick Checklist

- [ ] SSH keys enabled, password auth disabled
- [ ] Root login disabled
- [ ] fail2ban installed and running
- [ ] Firewall enabled with default deny
- [ ] Only necessary ports open
- [ ] Game servers run as dedicated users, not root
- [ ] System auto-updates enabled (at least security patches)
- [ ] RCON passwords are strong and unique
- [ ] RCON/admin ports restricted to specific IPs or localhost

## Next Steps

- [Networking & Port Forwarding](/self-hosting/networking) — open only what you need
- [Backups & Disaster Recovery](/self-hosting/backups) — recovery from the worst case
