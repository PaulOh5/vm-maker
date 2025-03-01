# My Go Project

This is a simple Go web application that uses the Fiber framework to set up a web server.

## Prerequisites

- Go 1.16 or later
- A working Go environment

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   cd my-vm-maker
   ```

2. **Install dependencies:**

   Navigate to the project directory and run:

   ```bash
   go mod tidy
   ```

3. **Run the application:**

   Execute the following command to start the server:

   ```bash
   go run src/main.go
   ```

4. **Access the application:**

   Open your web browser and go to `http://localhost:3000`. You should see "Hello, World!" displayed on the page.

## Project Structure

```
my-vm-maker
├── src
│   ├── main.go        # Entry point of the application
├── go.mod             # Module definition and dependencies
└── README.md          # Project documentation
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.