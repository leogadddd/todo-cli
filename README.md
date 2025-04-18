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

1. Create a new storage instance:

   ```go
   storage := NewStorage[YourType]("filename.json")
   ```

2. Save data:

   ```go
   err := storage.Save(yourData)
   if err != nil {
       log.Fatal(err)
   }
   ```

3. Load data:
   ```go
   var yourData YourType
   err := storage.Load(&yourData)
   if err != nil {
       log.Fatal(err)
   }
   ```

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
