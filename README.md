# Task Tracker

This program is a task tracker implemented in Go that allows you to manage tasks through operations such as adding, listing, updating, deleting, and changing their status. Data is saved in a JSON file (`Task.json`) to maintain persistence between executions.

## Prerequisites

1. Have [Go](https://go.dev/dl/) installed.
2. Set up the `main.go` file in your environment.
3. Create an empty `Task.json` file if it doesn't exist yet:

```bash
touch Task.json
```

## Usage Instructions

Compile and run the program using the commands below. The program accepts various flags to perform different operations.

### Compilation

```bash
go build main.go
```

### Execution

Below are the available operations:

#### 1. Add a Task

```bash
./main -add "Task description"
```

Creates a new task with the initial status `todo`.

#### 2. List All Tasks

```bash
./main -list
```

Displays all registered tasks along with their details (ID, description, status, creation date, and update date).

#### 3. Filter Tasks by Status

```bash
./main -list <status>
```

Filters tasks by their status (`todo`, `in-progress`, `done`). Example:

```bash
./main -list done
```

#### 4. Update a Task's Description

```bash
./main -update <id> "New description"
```

Updates the description of the task specified by its ID.

#### 5. Delete a Task

```bash
./main -delete <id>
```

Deletes the task specified by its ID.

#### 6. Mark a Task as "In Progress"

```bash
./main -mark-in-progress <id>
```

Updates the status of the task to `in-progress`.

#### 7. Mark a Task as "Completed"

```bash
./main -mark-done <id>
```

Updates the status of the task to `done`.

## Error Handling

- If you provide an invalid or non-existent ID, the program will display an error message.
- If the `Task.json` file is not found, the program will automatically create it.

## Example Workflow

1. Add a task:

```bash
./main -add "Prepare Monday's meeting"
```

2. List tasks:

```bash
./main -list
```

Output:

```plaintext
Task id: 23, description: Prepare Monday's meeting, status: todo, create_at: 2025-01-22 10:00:00, update_at: 2025-01-22 10:00:00
```

3. Update the task's description:

```bash
./main -update 23 "Review presentation for the meeting"
```

4. Mark the task as "Completed":

```bash
./main -mark-done 23
```

5. Delete the task:

```bash
./main -delete 23
```

## Additional Notes

- This program uses a map (`map[int]Task`) to manage tasks in memory before persisting them in `Task.json`.
- Task IDs are generated randomly, so collisions may occur if multiple tasks are created quickly.

You can view the project online at [Task Tracker Project](https://roadmap.sh/projects/task-tracker).

Enjoy using your task tracker! If you encounter any issues, review the code or contact the developer.
