# Todo CLI Go

**todo-cli-go** is a simple command-line interface for creating and managing a list of tasks.  
This is my first application written in Go, with a focus on learning structs, pointers, and basic language syntax.

## Features

The CLI offers all core functionality expected from a basic Todo app:

- ✅ **Add items**: Quickly add new tasks. They are created with an "unsolved" status by default.
- 🔁 **Change item status**: Mark tasks as "solved" once they're completed.
- ❌ **Remove items**: Delete tasks you no longer need.
- 💾 **Persistent storage**: Your list is saved locally and updated automatically after each change.

## How to use

Since this is a learning project, the easiest way to run it is directly from the command line using:

```
go run .
```

Alternatively, you can build the binary and use it as a standalone CLI tool:

```
go build -o todo
./todo
```

## Tech Stack

This project is written in Go, using only Go's standard libraries.
It's focused on exploring the core language features without external dependencies.
