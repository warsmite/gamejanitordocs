---
title: "Minecraft Server: Getting Started"
description: "How to actually set up a Minecraft server from scratch, including which JAR to pick and why."
order: 1
tags: ["minecraft", "java", "setup"]
---

# Minecraft Server: Getting Started

## Picking a Server JAR

This is where most guides fail you. "Just download the server jar from minecraft.net" — sure, if you want garbage performance with zero configuration.

Here's what actually matters:

| JAR | Use Case | Performance |
|-----|----------|-------------|
| Vanilla | You want the pure experience, no plugins | Baseline |
| Paper | Most servers. Plugin support, major performance gains | Excellent |
| Purpur | Paper + extra configuration options | Excellent |
| Fabric + Lithium | Modded server, performance-focused | Very Good |
| Forge | Modded server, massive mod ecosystem | Varies |

**For most people: use Paper.** It's a drop-in replacement for vanilla with significantly better performance and plugin support.

## Minimum Requirements

Forget the "512MB RAM" nonsense you see everywhere. Real numbers:

- **1-5 players:** 2GB RAM minimum, 4GB recommended
- **5-20 players:** 4-6GB RAM
- **20+ players:** 8GB+ RAM, SSD storage is non-negotiable
- **CPU:** Single-thread performance matters more than core count. Minecraft's main game loop is single-threaded.

## Java Version

Minecraft 1.17+: **Java 17 or later**. Don't use Java 8 for modern versions — you'll get worse performance and missing security patches.

```bash
# Check your java version
java -version

# On Debian/Ubuntu
sudo apt install openjdk-21-jre-headless

# On Arch
sudo pacman -S jre-openjdk-headless
```

## Starting the Server

```bash
# Download Paper (replace version as needed)
mkdir minecraft-server && cd minecraft-server

# First run generates eula.txt and server.properties
java -Xms2G -Xmx4G -jar paper.jar --nogui

# Accept the EULA
sed -i 's/eula=false/eula=true/' eula.txt

# Run for real
java -Xms2G -Xmx4G -jar paper.jar --nogui
```

### JVM Flags That Actually Matter

Don't just copy-paste Aikar's flags without understanding them. Here's what matters:

```bash
java -Xms4G -Xmx4G \
  -XX:+UseG1GC \
  -XX:+ParallelRefProcEnabled \
  -XX:MaxGCPauseMillis=200 \
  -XX:+UnlockExperimentalVMOptions \
  -XX:+DisableExplicitGC \
  -XX:G1NewSizePercent=30 \
  -XX:G1MaxNewSizePercent=40 \
  -XX:G1HeapRegionSize=8M \
  -jar paper.jar --nogui
```

`-Xms` and `-Xmx` should be **equal** — you've already allocated the RAM, let Java use it from the start instead of growing the heap at runtime.

## Next Steps

- [server.properties reference](/games/minecraft/server-properties) — every setting explained
- [Networking & Port Forwarding](/self-hosting/networking) — making your server reachable
