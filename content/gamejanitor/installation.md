---
title: "Installing Game Janitor"
description: "Get Game Janitor up and running on your server."
order: 1
tags: ["gamejanitor", "installation", "setup"]
---

# Installing Game Janitor

Game Janitor is an open source gameserver control panel and orchestrator. This guide gets you from zero to managing gameservers.

## Requirements

- Linux (Debian 12+, Ubuntu 22.04+, Fedora 39+, Arch)
- 1GB RAM minimum (for the panel itself — gameservers need their own resources)
- Docker (recommended) or systemd

## Quick Start

```bash
# Install via the install script
curl -fsSL https://get.gamejanitor.net | bash

# Or manually with Go
go install github.com/warsmite/gamejanitor@latest
```

## Configuration

Game Janitor uses a single config file at `/etc/gamejanitor/config.yaml`:

```yaml
server:
  addr: ":8080"

storage:
  path: /var/lib/gamejanitor

# Game server definitions
games:
  - name: minecraft
    enabled: true
```

## Next Steps

- Configure your first gameserver
- Set up users and permissions
- Connect to Game Janitor Hosting for remote management
