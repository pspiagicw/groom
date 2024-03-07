# `groom`

`groom` is a yet another os-agnostic task runner.

## Features

- Task runner with a simple yet powerful toml syntax
- Single static binary without dependencies
- Fast without any runner overhead

## Functionality

### Config file
- Requires a simple `groom.toml` file in the current directory
- Tasks can have commands, description, environment variables and other task dependencies.
- A global `[variables]` section for user variables.
- They support inline variable substition.

```toml
name = "groom"

[variables]
version = "0.1.0"
build-dir = "build"
main-file = "main.go"
ldflags = "-X main.VERSION='$version'"

[task.build]
description = "Build the project"
command = 'go build -ldflags "$ldflags" -o $name ${main-file}'
environment = [ "CGO_ENABLED=0" ]
depends = [
    "format"
]

[task.format]
description = "Format the project"
command = "go fmt ./..."
```

![help](./gifs/help.gif)

### List

Run `groom` without any arguments to list all declared tasks.
- List all tasks with their dependencies, and description.
- Use `--simple` to list all tasks without any beautification. Useful with scripts.

![list](./gifs/list.gif)

### Executing tasks

Provide a list of tasks to execute and watch `groom` execute them!

- Will automatically run dependencies.

![build](./gifs/build.gif)

### Neovim Plugin

A `neovim` plugin is in the works for integrating `groom` with Neovim.
It automatically lists all the tasks and adds the output to the quickfix list.

![plugin](./gifs/plugin.gif)

