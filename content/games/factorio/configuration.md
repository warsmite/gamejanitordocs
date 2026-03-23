---
title: "Factorio Configuration Reference"
description: "Map generation settings, runtime configuration, RCON, and mod management for Factorio servers."
order: 2
tags: ["factorio", "configuration"]
---

# Factorio Configuration Reference

## Map Generation Settings (map-gen-settings.json)

These are set at world creation and **cannot be changed after**. Get them right before generating.

```bash
# Export default settings to edit
./bin/x64/factorio --map-gen-settings /opt/factorio/map-gen-settings.json
```

### Key Settings

```json
{
  "terrain_segmentation": 1,
  "water": 1,
  "width": 0,
  "height": 0,
  "starting_area": 1,
  "peaceful_mode": false,
  "autoplace_controls": {
    "coal": { "frequency": 1, "size": 1, "richness": 1 },
    "stone": { "frequency": 1, "size": 1, "richness": 1 },
    "copper-ore": { "frequency": 1, "size": 1, "richness": 1 },
    "iron-ore": { "frequency": 1, "size": 1, "richness": 1 },
    "uranium-ore": { "frequency": 1, "size": 1, "richness": 1 },
    "crude-oil": { "frequency": 1, "size": 1, "richness": 1 },
    "enemy-base": { "frequency": 1, "size": 1, "richness": 1 },
    "trees": { "frequency": 1, "size": 1, "richness": 1 }
  },
  "seed": null
}
```

### Frequency/Size/Richness Scale

Values are multipliers. The scale works like this:

| Value | Meaning |
|-------|---------|
| 0 | None (disabled) |
| 0.17 | Very low |
| 0.33 | Low |
| 0.5 | Below normal |
| 1 | Normal |
| 1.5 | Above normal |
| 2 | High |
| 3 | Very high |
| 6 | Maximum |

For a chill server: set `enemy-base` frequency/size to 0.5, increase resource richness to 2.

For a challenge: max out enemy settings, reduce starting area, lower resources.

### Map Size

`width` and `height` control map dimensions in tiles. `0` means infinite. Setting these limits exploration and keeps save sizes manageable:

```json
{
  "width": 5000,
  "height": 5000
}
```

## Map Settings (map-settings.json)

These **can** be changed at runtime via console commands, unlike map-gen.

```json
{
  "difficulty_settings": {
    "recipe_difficulty": 0,
    "technology_difficulty": 0,
    "technology_price_multiplier": 1,
    "research_queue_setting": "always"
  },
  "pollution": {
    "enabled": true,
    "diffusion_ratio": 0.02,
    "min_to_diffuse": 15,
    "ageing": 1,
    "expected_max_per_chunk": 150,
    "min_to_show_per_chunk": 50
  },
  "enemy_evolution": {
    "enabled": true,
    "time_factor": 0.000004,
    "destroy_factor": 0.002,
    "pollution_factor": 0.0000009
  },
  "enemy_expansion": {
    "enabled": true,
    "max_expansion_distance": 7,
    "settler_group_min_size": 5,
    "settler_group_max_size": 20,
    "min_expansion_cooldown": 14400,
    "max_expansion_cooldown": 216000
  }
}
```

### Recipe/Technology Difficulty

| Value | Meaning |
|-------|---------|
| 0 | Normal |
| 1 | Expensive (marathon mode) |

`technology_price_multiplier` multiplies all research costs. Setting this to 4 with expensive recipes gives the marathon experience.

## RCON

Factorio has built-in RCON support:

```bash
./bin/x64/factorio --start-server /opt/factorio/saves/myworld.zip \
  --server-settings server-settings.json \
  --rcon-port 27015 \
  --rcon-password "yoursecretpassword"
```

Connect with any RCON client (mcrcon, rcon-cli, etc.):

```bash
mcrcon -H localhost -P 27015 -p yoursecretpassword "/players"
```

## Mod Support

Factorio's mod system works on servers. Mods go in the `mods/` directory.

```bash
# Install mods
cp my-mod_1.0.0.zip /opt/factorio/mods/

# mod-list.json controls which mods are active
cat /opt/factorio/mods/mod-list.json
```

`mod-list.json`:

```json
{
  "mods": [
    { "name": "base", "enabled": true },
    { "name": "some-mod", "enabled": true }
  ]
}
```

### Important Mod Notes

- All clients must have the same mods installed at the same versions. Factorio enforces this strictly
- Mod updates can break saves. Always backup before updating mods
- Some mods are server-side only (admin tools, statistics) but most require client matching
- Download mods manually from the [mod portal](https://mods.factorio.com/) and place the `.zip` files in the `mods/` directory

## Console Commands

Useful runtime commands (type in server console or via RCON):

| Command | Description |
|---------|-------------|
| `/players` | List connected players |
| `/promote <name>` | Give admin to player |
| `/demote <name>` | Remove admin from player |
| `/ban <name>` | Ban player |
| `/unban <name>` | Unban player |
| `/kick <name>` | Kick player |
| `/evolution` | Show current biter evolution factor |
| `/save <name>` | Force save |
| `/quit` | Graceful shutdown and save |
| `/config set afk-auto-kick-interval <minutes>` | Change AFK kick at runtime |

## Performance Notes

- **UPS (Updates Per Second)** is the server's tick rate. Target is 60. Below 60 means the factory is too complex for the hardware
- Factorio is single-threaded for game logic. Buy the fastest single-core CPU you can
- Belts are cheaper than bots for UPS in megabases
- Circuit networks with complex conditions can tank UPS
- Monitor with `/performance` in the console
