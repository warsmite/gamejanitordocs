---
title: "7 Days to Die Server: Getting Started"
description: "Setting up a 7 Days to Die dedicated server, including SteamCMD installation, world generation, and basic config."
order: 1
tags: ["7-days-to-die", "steam", "setup"]
---

# 7 Days to Die Server: Getting Started

## Requirements

- **1-4 players:** 4GB RAM, 2-4 CPU cores
- **5-8 players:** 8GB RAM, 4 CPU cores
- **8-16 players:** 8-12GB RAM, 4-6 CPU cores
- **Storage:** 3-5GB for server, worlds are 1-3GB
- **OS:** Linux or Windows

7DTD generates the world procedurally, which is CPU-intensive at first. Once generated, CPU usage drops to reasonable levels.

## Installing via SteamCMD

```bash
# 7 Days to Die Dedicated Server App ID: 294420
steamcmd +force_install_dir /opt/7dtd +login anonymous +app_update 294420 validate +quit
```

## Starting the Server

```bash
cd /opt/7dtd
./startserver.sh -configfile=serverconfig.xml
```

On first run, the server generates the world. This can take 5-30 minutes depending on world size and CPU. Don't panic — it's normal.

## Server Configuration

The main config is `serverconfig.xml`. A template ships with the server. The important sections:

```xml
<?xml version="1.0"?>
<ServerSettings>
    <!-- Server identity -->
    <property name="ServerName" value="My 7DTD Server"/>
    <property name="ServerDescription" value="A 7 Days to Die server"/>
    <property name="ServerWebsiteURL" value=""/>
    <property name="ServerPassword" value=""/>
    <property name="ServerMaxPlayerCount" value="8"/>
    <property name="ServerPort" value="26900"/>

    <!-- Admin -->
    <property name="ControlPanelEnabled" value="false"/>
    <property name="ControlPanelPort" value="8080"/>
    <property name="ControlPanelPassword" value=""/>
    <property name="TelnetEnabled" value="true"/>
    <property name="TelnetPort" value="8081"/>
    <property name="TelnetPassword" value="youradminpassword"/>

    <!-- World -->
    <property name="GameWorld" value="Navezgane"/>
    <property name="WorldGenSeed" value="my seed"/>
    <property name="WorldGenSize" value="6144"/>
    <property name="GameName" value="MyGame"/>
    <property name="GameMode" value="GameModeSurvival"/>

    <!-- Difficulty -->
    <property name="GameDifficulty" value="2"/>
    <property name="DayNightLength" value="60"/>
    <property name="DayLightLength" value="18"/>
    <property name="BloodMoonFrequency" value="7"/>
    <property name="BloodMoonRange" value="0"/>
    <property name="BloodMoonEnemyCount" value="8"/>
</ServerSettings>
```

### World Options

| Value | Description |
|-------|-------------|
| `Navezgane` | Pre-made map. Consistent, well-designed. Good for small groups |
| `RWG` (Random World Gen) | Procedurally generated. Set `WorldGenSeed` and `WorldGenSize` |

### World Gen Sizes

| Size | Description |
|------|-------------|
| 4096 | Small. Good for 1-4 players |
| 6144 | Medium. Default, works for most servers |
| 8192 | Large. 8+ players |
| 10240 | Very large. Lots of empty space |

## Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 26900 | TCP + UDP | Game traffic |
| 26901 | UDP | Additional game traffic |
| 26902 | UDP | Additional game traffic |
| 8080 | TCP | Web control panel (if enabled) |
| 8081 | TCP | Telnet (if enabled) |

```bash
# UFW
sudo ufw allow 26900:26902/tcp
sudo ufw allow 26900:26902/udp
sudo ufw allow 8081/tcp  # telnet admin

# firewalld
sudo firewall-cmd --permanent --add-port=26900-26902/tcp
sudo firewall-cmd --permanent --add-port=26900-26902/udp
sudo firewall-cmd --permanent --add-port=8081/tcp
sudo firewall-cmd --reload
```

Don't expose telnet (8081) to the internet unless you've changed the default password and ideally restricted it to specific IPs.

## Running as a systemd Service

```ini
# /etc/systemd/system/7dtd.service
[Unit]
Description=7 Days to Die Dedicated Server
After=network.target

[Service]
Type=simple
User=sdtd
WorkingDirectory=/opt/7dtd
ExecStartPre=/usr/games/steamcmd +force_install_dir /opt/7dtd +login anonymous +app_update 294420 +quit
ExecStart=/opt/7dtd/startserver.sh -configfile=serverconfig.xml
Restart=on-failure
RestartSec=30

[Install]
WantedBy=multi-user.target
```

## Admin Access

### Telnet

Connect via telnet for server management:

```bash
telnet localhost 8081
# Enter TelnetPassword
```

### Web Control Panel

Enable in config with `ControlPanelEnabled=true`. Provides a browser-based admin interface.

### Admin File

Grant admin via `serveradmin.xml` in the saves directory:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<adminTools>
    <admins>
        <admin steamID="76561198012345678" permission_level="0" />
    </admins>
</adminTools>
```

Permission level 0 = full admin.

## Next Steps

- [7 Days to Die Configuration](/games/7-days-to-die/configuration) — difficulty tuning, blood moon settings, loot, and performance
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
