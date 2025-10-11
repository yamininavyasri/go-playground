# go-playground
A personal playground for exploring and practicing Go (Golang) concepts, snippets, and experiments

📂 Project Structure
go-playground/
├── collections/
├── conditions/
├── context_examples/
├── defer_example/
├── env_config/
├── error_handling/
├── fileio /
├── functions_example/
├── go.mod
├── go.sum
├── go_routines/
├── http_example/
├── interfaces/
├── loop_examples/
├── mutex/
├── packages/
├── pointers/
├── string_manipulations/
├── structs/
├── switch_example/
├── time_example/
├── unit_testing/
└── variable_types/
Each folder contains Go files demonstrating specific concepts or features of the language.

✅ Prerequisites
Ensure you have the following installed:
Go 1.25.1 or higher
A terminal or command prompt
A text editor (e.g., VSCode, GoLand)

Setup Instructions
1. Clone the Repository
git clone https://github.com/yamininavyasri/go-playground.git
cd go-playground
2. Initialize Go Modules
If you haven't already initialized Go modules:
go mod tidy
This command will download and install any dependencies specified in the go.mod file.

▶️ Running the Programs
Each folder contains Go files that can be run individually. For example, to run the functions_example:
cd functions_example
go run functions.go // run the respective .go file.
Repeat this process for other folders as needed.

🧪 Running Unit Tests
Some folders include unit tests. To run them:
cd unit_testing
go test

🌐 Running HTTP Servers
If a folder contains an HTTP server (e.g., http_example), you can start it by running:
cd http_example
go run main.go
Then, open your browser and navigate to http://localhost:8080 to interact with the server.

🧠 Understanding the Examples

Each folder demonstrates a specific Go concept:

build_tags/: Examples of custom build tags for conditional compilation (dev, prod, etc.).

channels/: Working with Go channels for concurrency, synchronization, and communication between goroutines.

collections/: Working with slices, maps, and arrays

conditions/: Using if, else, and switch statements

context_examples/: Managing request contexts

defer_example/: Understanding defer statements

env_config/: Examples of environment configuration using os.Getenv and os.LookupEnv.

error_handling/: Demonstrating idiomatic error checking and handling in Go.

fileio/: Performing file input/output operations, including creating, reading, writing, and deleting files.

functions_example/: Defining and using functions

go_routines/: Concurrency with goroutines

http_example/: Building simple HTTP servers

interfaces/: Implementing and using interfaces

loop_examples/: Iterating with for loops

mutex/: Synchronization with mutexes

packages/: Organizing code into packages

pointers/: Working with pointers

string_manipulations/: String operations

structs/: Defining and using structs

switch_example/: Using switch statements

time_example/: Working with time and dates

unit_testing/: Writing and running tests

variable_types/: Understanding variable declarations