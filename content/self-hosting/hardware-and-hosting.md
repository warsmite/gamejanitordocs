---
title: "Hardware & Hosting Options"
description: "Home server vs VPS vs dedicated hosting — what to pick, what specs matter, and what it costs."
order: 3
tags: ["hardware", "hosting", "vps"]
---

# Hardware & Hosting Options

## The Three Options

### 1. Home Server

Run the game server on hardware in your house.

**Pros:**
- No monthly cost beyond electricity
- Full physical control
- Can repurpose old hardware

**Cons:**
- Your home internet is the bottleneck (upload speed matters)
- You deal with port forwarding, dynamic DNS, and CGNAT
- Power outages take your server down
- Noise and heat in your living space

**Best for:** Small private servers for friends, testing, or if you already have spare hardware.

### 2. VPS (Virtual Private Server)

A virtual machine on someone else's hardware. You get a slice of a server.

**Pros:**
- Cheap ($5-40/month for most game servers)
- Static public IP, no port forwarding needed
- Runs 24/7 regardless of your home power
- Quick to set up and tear down

**Cons:**
- Shared hardware — noisy neighbors can affect performance
- CPU is often the weak point (shared cores, throttled burst)
- Limited RAM and storage for the price

**Best for:** Small to medium servers (1-20 players for most games). The sweet spot for most people.

**Providers:** Hetzner, OVH, Linode, DigitalOcean, Vultr. Avoid AWS/GCP/Azure for game servers — the pricing model punishes sustained workloads.

### 3. Dedicated Server

An entire physical server rented from a datacenter.

**Pros:**
- Dedicated CPU, RAM, and bandwidth — no sharing
- Best performance per dollar at scale
- Can run multiple game servers on one machine

**Cons:**
- More expensive ($40-150+/month)
- You manage everything (OS, updates, hardware failures)
- Overkill for small servers

**Best for:** Large community servers, running multiple games, or performance-critical setups (competitive Rust, large ARK clusters).

**Providers:** Hetzner Server Auction (great deals on older hardware), OVH Game, Vultr Bare Metal, ReliableSite.

## What Specs Actually Matter

### CPU: Single-Thread Performance Is King

Most game servers are single-threaded or lightly multi-threaded. The main game loop runs on one core. A CPU with 4 fast cores beats a CPU with 16 slow cores for game servers.

| Metric | Why it matters |
|--------|---------------|
| Single-thread speed | Directly impacts tick rate and server responsiveness |
| Core count | Matters if running multiple servers on one machine |
| Clock speed | Higher is better for game servers. 3.5GHz+ recommended |

**For VPS:** Look for "dedicated vCPU" or "compute-optimized" plans. Shared vCPUs get throttled under sustained load, which is exactly what a game server does.

### RAM: Depends Entirely on the Game

| Game | RAM per server |
|------|---------------|
| Factorio | 512MB - 2GB |
| Terraria | 512MB - 1GB |
| Minecraft (Paper) | 2GB - 8GB |
| Don't Starve Together | 1GB - 2GB |
| Valheim | 2GB - 4GB |
| CS2 | 2GB - 4GB |
| Project Zomboid | 4GB - 8GB |
| 7 Days to Die | 4GB - 8GB |
| Rust | 8GB - 16GB |
| Palworld | 8GB - 32GB |
| ARK: Survival Evolved | 8GB - 32GB |

Always leave 1-2GB for the OS on top of what the game needs.

### Storage: SSD Is Non-Negotiable

Game servers do constant small reads and writes — loading chunks, saving world data, writing logs. An SSD makes everything smoother. HDDs cause periodic freezes during auto-saves.

| Storage type | Game server verdict |
|-------------|-------------------|
| HDD | No. Save operations cause visible lag spikes |
| SATA SSD | Good enough for most servers |
| NVMe SSD | Best. Overkill for small servers but noticeable on large ones |

Storage size: 50GB is enough for most single-game servers. ARK and CS2 are the biggest at 15-35GB for the server alone.

### Bandwidth

Game servers use less bandwidth than you'd expect:

| Game | Bandwidth per player |
|------|---------------------|
| Factorio | ~5 KB/s |
| Minecraft | ~50-100 KB/s |
| Valheim | ~50-100 KB/s |
| Rust | ~100-200 KB/s |

A 100 Mbps connection handles most game servers comfortably. Upload matters more than download for home servers.

**Latency matters more than bandwidth.** A server with 10ms ping to your players on a 100 Mbps line will feel better than one with 80ms ping on a 1 Gbps line. Pick a server location close to your players.

## Cost Comparison

Rough monthly costs for running a medium game server (e.g., 10-player Minecraft or Valheim):

| Option | Monthly Cost | Notes |
|--------|-------------|-------|
| Home server | $5-15 electricity | Free if you already have hardware running |
| Cheap VPS (shared CPU) | $5-12 | Fine for light games (Terraria, Factorio, DST) |
| VPS (dedicated CPU) | $15-40 | Good for Minecraft, Valheim, PZ, 7DTD |
| Dedicated server | $40-80 | Can run 3-5 game servers simultaneously |
| Game-specific hosting | $10-30 | Easy but less flexible. Markup for convenience |

## Recommendations

**"I just want to play with 3 friends"** — Use a $10-20/month VPS with dedicated CPU. Hetzner or OVH.

**"I'm running a community server"** — Dedicated server. Hetzner Server Auction is unbeatable on price for the performance.

**"I want to learn and don't care about uptime"** — Home server on old hardware. Great learning experience.

**"I want it to just work"** — Game-specific hosting or [Game Janitor Hosting](https://gamejanitorhosting.com). You pay a premium but skip the Linux learning curve.

## Next Steps

- [Linux Basics](/self-hosting/linux-basics) — if this is your first time with Linux
- [Security Hardening](/self-hosting/security) — protect your server from the internet
- [Networking & Port Forwarding](/self-hosting/networking) — for home server setups
