# Task Manager CLI

A simple and efficient command-line task manager built with Go. Manage your tasks with ease using intuitive commands to add, update, delete, and track your to-do items.

**Project Source:** [roadmap.sh - Task Tracker](https://roadmap.sh/projects/task-tracker)

## Features

- ✅ Add new tasks
- ✏️ Update existing tasks
- 🗑️ Delete tasks
- 📊 Track task status (Todo, In-Progress, Done)
- 📋 List all tasks or filter by status
- 💾 Persistent storage using JSON

## Prerequisites

- Go 1.16 or higher installed on your system
- Basic familiarity with command-line interfaces

## Installation

### Clone or Download

Download the source code or clone the repository to your local machine.

### Build the Application

Navigate to the project directory and build the executable:

```bash
go build -o task main.go
```

This will create an executable file named `task` (or `task.exe` on Windows) in your current directory.

### Optional: Add to PATH

To use the task manager from anywhere on your system, you can add the executable to your PATH:

**Linux/macOS:**
```bash
# Move the executable to a directory in your PATH
sudo mv task /usr/local/bin/

# Or add the current directory to your PATH
export PATH=$PATH:$(pwd)
```

**Windows:**
```cmd
# Add the current directory to your PATH environment variable
# Or move task.exe to a directory already in your PATH
```

## Usage

### Basic Syntax

```bash
task <command> [arguments]
```

### Commands

#### 1. Add a Task

Add a new task to your task list.

```bash
task add "task description"
```

**Example:**
```bash
task add "Complete project documentation"
task add "Buy groceries"
```

#### 2. Update a Task

Update the description of an existing task by its ID.

```bash
task update <id> "updated task description"
```

**Example:**
```bash
task update 1 "Complete project documentation and README"
```

#### 3. Delete a Task

Delete a task by its ID.

```bash
task delete <id>
```

**Example:**
```bash
task delete 2
```

#### 4. Mark Task Status

Update the status of a task to track progress.

```bash
task mark <status> <id>
```

**Available statuses:**
- `in-progress` or `inprogress`
- `done`

**Examples:**
```bash
task mark in-progress 1
task mark done 3
```

#### 5. List Tasks

Display all tasks or filter by status.

**List all tasks:**
```bash
task list
```

**Filter by status:**
```bash
task list <status>
```

**Available filters:**
- `todo` - Show only tasks with Todo status
- `in-progress` - Show only tasks in progress
- `done` - Show only completed tasks

**Examples:**
```bash
task list
task list todo
task list in-progress
task list done
```

## Data Storage

Tasks are automatically saved to a `data.json` file in the same directory as the executable. This file is created automatically on first run and updated with each operation.

**Note:** Do not manually edit the `data.json` file as it may cause data corruption.

## Example Workflow

```bash
# Add some tasks
task add "Write project proposal"
task add "Review code changes"
task add "Prepare presentation"

# List all tasks
task list

# Mark a task as in-progress
task mark in-progress 1

# Mark a task as done
task mark done 2

# Update a task description
task update 3 "Prepare and rehearse presentation"

# List only completed tasks
task list done

# Delete a task
task delete 3
```

## Task Properties

Each task contains the following information:
- **ID**: Unique identifier (auto-incremented)
- **Description**: Task description
- **Status**: Current status (Todo, In-Progress, or Done)
- **CreatedAt**: Timestamp when the task was created
- **UpdatedAt**: Timestamp when the task was last modified

## Troubleshooting

### "command not found: task"

The executable is not in your PATH. Either:
- Run it with `./task` (Linux/macOS) or `task.exe` (Windows) from the project directory
- Add the executable to your PATH as described in the Installation section

### "Failed to read tasks" or "Failed to save tasks"

Ensure you have read/write permissions in the directory where you're running the application.

### Invalid ID error

Make sure you're using a valid task ID. Use `task list` to see all available task IDs.

## Development

### Running in Development

You can run the application without building:

```bash
go run main.go <command> [arguments]
```

**Example:**
```bash
go run main.go add "Test task"
go run main.go list
```

### Testing

To test the application, run a series of commands:

```bash
go build -o task main.go
./task add "Test task 1"
./task add "Test task 2"
./task list
./task mark in-progress 1
./task list in-progress
./task delete 2
./task list
```

## License

This project is open source and available for personal and commercial use.

## Contributing

Feel free to fork this project and submit pull requests for any improvements.

---

**Happy Task Managing! 🚀**