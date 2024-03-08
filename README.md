# 🧹`groom`

`groom` is a yet another os-agnostic task runner.

## ✨ Features

- 🔥 Task runner with a simple yet powerful toml syntax
- 🧹 Single static binary without dependencies
- 💻 Neovim plugin for integration with `neovim`.

## Functionality

### Config file

- Requires a simple `groom.toml` file in the project root.
- `groom` automatically recurses to the project root to find a `groom.toml` file.
- A global `[variables]` section for user variables.
- They support inline variable substition.

```toml
name = "example-project"

[variables]
version = "0.0.1"

# Tasks start with '[task.<task-name>]'
# They should contain, 'command' property.
# Other fields are optional.
[task.build]
description = "Build the project."
command = "go build ."

# Tasks can contain 'commands' as a list of commands.
[task.run]
commands = [
    "go run main.go",
    "python -m exaple-project",
]

# Tasks can contain dependencies, and environment variables defined
[task.test]
environment = [ "TESTS=1" ]
command = "python -m unittest"
depends = [
    "format"
]

[task.format]
commands = "go fmt ./..."
```

> You can run `groom --example-config` to get a working example config.

![help](./gifs/help.gif)

### List

Run `groom` without any arguments to list all configured tasks.

> Use `--simple` to list all tasks without any fancy printing. Useful with scripts.

![list](./gifs/list.gif)

### Executing tasks

Provide a list of tasks to execute and watch `groom` execute them!

>  Use the `--dry-run` argument to show the log without actually running anything.

![build](./gifs/build.gif)

### Neovim Plugin

A `neovim` plugin is in the works for integrating `groom` with Neovim.

It allows you to run tasks without leaving your editor.

> Find it [here](https://github.com/pspiagicw/groom.nvim)

![plugin](./gifs/plugin.gif)

## Contributing

⭐ Star the project if you like it!

Feel free to contribute to the project, by either raising a issue or opening a PR.
