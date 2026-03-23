---
title: "server.properties Reference"
description: "Every Minecraft server.properties setting explained with actual recommendations, not just defaults."
order: 2
tags: ["minecraft", "configuration"]
---

# server.properties Reference

The settings that actually matter, with real recommendations instead of just listing defaults.

## Performance Settings

### view-distance

Default: `10`. **Set this to 6-8** unless you have RAM and CPU to spare. This is the single biggest performance knob you have. Each increment adds exponentially more chunks to track.

- 10 chunks: ~440 chunks loaded per player
- 6 chunks: ~168 chunks loaded per player

That's a 60% reduction in chunk work.

### simulation-distance

Default: `10`. Controls how far from the player entities and block ticks are processed. **Set to 4-6**. This is separate from view-distance on 1.18+ — you can have high view distance for visuals but low simulation distance for performance.

### max-tick-time

Default: `60000` (60 seconds). Time in ms before the server watchdog kills the server for an unresponsive tick. Set to `-1` to disable if you're debugging performance issues, but keep it enabled in production.

## Gameplay Settings

### difficulty

`peaceful` / `easy` / `normal` / `hard`

This affects mob spawning rates and damage. On `hard`, zombies can break doors and villagers can die. Most survival servers use `normal` or `hard`.

### pvp

Default: `true`. Whether players can damage each other. Note: this doesn't prevent lava/TNT griefing — you need plugins for that.

### spawn-protection

Default: `16`. Radius of blocks around spawn that non-ops can't modify. Set to `0` if you use a protection plugin (WorldGuard, etc) instead.

## Network Settings

### network-compression-threshold

Default: `256`. Packets larger than this (in bytes) get compressed. Lower values = more CPU, less bandwidth. Higher values = less CPU, more bandwidth. **256 is fine for most servers.** Set to `-1` to disable compression if your server and players are on the same network.

### rate-limit

Default: `0` (disabled). Packets per second before a player gets kicked. Set to `500` as a basic anti-flood measure.
