# Go Todo Application

This is a simple Todo application written in Go. It provides functionality to store and retrieve data using a generic storage system.

## Features

- Generic storage implementation using Go generics.
- Save and load data to/from a file in JSON format.
- Easy to extend for different data types.

## File Structure

- `storage.go`: Contains the `Storage` struct and methods for saving and loading data.
- `command.go`: Handles command-line arguments and executes corresponding actions.
- `main.go`: Entry point of the application, orchestrates the flow of the program.
- `todo.go`: Defines the `Todo` struct and provides methods to manage todos.

## Usage

1. Clone the Repository

   ```bash
   git clone https://github.com/your-username/go-todo-app.git
   cd go-todo-app
   Build the Application
   ```

2. Make sure you have Go installed (version 1.18 or later). Then run:

   ```bash
   go build -o todo
   ```

   This will create an executable named todo in the current directory.

3. Run the Application

   You can now use the todo executable to interact with the app. Below are some basic usage examples (depending on how your command.go is set up):

   ```bash
   ./todo add "Buy groceries"
   ./todo list
   ./todo complete 1
   ./todo delete 1
   ```

   Replace the commands above with the actual ones implemented in your command.go. If you're not sure yet, this is a placeholder for how your CLI might look.

4. Data Storage

   Todos are saved to a local JSON file (e.g., todos.json) using the generic storage system. This makes it easy to persist your data across sessions.

## Requirements

- Go 1.18 or later (for generics support).

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes and push them to your fork.
4. Submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
