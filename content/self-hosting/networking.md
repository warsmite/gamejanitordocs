---
title: "Networking & Port Forwarding"
description: "How to make your gameserver reachable from the internet, explained without hand-waving."
order: 1
tags: ["networking", "port-forwarding", "firewall"]
---

# Networking & Port Forwarding

This is the part where most people get stuck. Here's how networking actually works for gameservers.

## The Problem

Your gameserver listens on a port (e.g., 25565 for Minecraft). But your home network sits behind a router that does NAT — it has one public IP, and your server has a private IP (like 192.168.1.x). Players on the internet can reach your public IP, but not your private one.

You need to tell your router: "when traffic comes in on port 25565, send it to 192.168.1.x."

## Step 1: Static IP for Your Server

Your server needs the same local IP every time. If DHCP gives it a new address, your port forwarding rule breaks.

```bash
# Find your current IP and gateway
ip addr show
ip route show default

# On most Linux systems, configure via netplan, networkd, or your distro's tool
# Example: /etc/systemd/network/10-static.network
[Match]
Name=eth0

[Network]
Address=192.168.1.100/24
Gateway=192.168.1.1
DNS=1.1.1.1
```

Or just use a DHCP reservation in your router — find your server's MAC address and pin it to an IP. This is usually easier.

## Step 2: Port Forwarding

Every router is different, but the process is the same:

1. Log in to your router (usually 192.168.1.1 or 192.168.0.1)
2. Find "Port Forwarding", "NAT", or "Virtual Servers"
3. Add a rule:
   - **External port:** the game port (e.g., 25565)
   - **Internal IP:** your server's static IP
   - **Internal port:** same as external
   - **Protocol:** TCP, UDP, or both (game-dependent)

### Common Game Ports

| Game | Port | Protocol |
|------|------|----------|
| Minecraft | 25565 | TCP |
| Valheim | 2456-2458 | UDP |
| Terraria | 7777 | TCP |
| Palworld | 8211 | UDP |
| Factorio | 34197 | UDP |

## Step 3: Firewall

If your server runs Linux, you probably have a firewall. Make sure the port is open:

```bash
# UFW (Ubuntu/Debian)
sudo ufw allow 25565/tcp

# firewalld (Fedora/RHEL)
sudo firewall-cmd --permanent --add-port=25565/tcp
sudo firewall-cmd --reload

# iptables (if you're old school)
sudo iptables -A INPUT -p tcp --dport 25565 -j ACCEPT
```

## Step 4: Test It

From **outside your network** (use your phone on mobile data, or a service):

```bash
# Replace with your public IP
nc -zv YOUR_PUBLIC_IP 25565
```

Find your public IP:
```bash
curl ifconfig.me
```

## Dynamic DNS

If your ISP changes your public IP (most residential connections do), you need Dynamic DNS. Options:

- **DuckDNS** — free, works well
- **Cloudflare** — if you own a domain, use their API to update A records
- **No-IP** — free tier available

> Don't want to deal with port forwarding, dynamic DNS, and firewall rules? This is one of the things a hosting provider handles for you. [Game Janitor Hosting](https://gamejanitorhosting.com) gives you a public IP out of the box.

## CGNAT: When Port Forwarding Doesn't Work

If port forwarding doesn't work even though you did everything right, your ISP might be using CGNAT (Carrier-Grade NAT). This means you don't have a real public IP — you share one with other customers.

How to check:
```bash
# Compare your router's WAN IP to your public IP
# Router WAN IP: check your router admin page
# Public IP:
curl ifconfig.me
```

If they're different, you're behind CGNAT. Your options:
- Call your ISP and ask for a public IP (some provide this for free or a small fee)
- Use a VPN/tunnel service (Tailscale, WireGuard, Cloudflare Tunnel)
- Use a hosting provider
