# Go Learning Projects

A collection of Go programming practice projects covering fundamental concepts, data structures, and practical applications.

## 📁 Project Structure

This repository contains several Go projects organized by topic:

### 🎯 Core Learning Projects

#### `array/`

- **Purpose**: Array and slice operations in Go
- **Features**:
  - Static array manipulation
  - Dynamic slice operations
  - Array slicing and reslicing
  - Product struct with dynamic lists
- **Run**: `go run main.go`

#### `structs/`

- **Purpose**: Struct definition and usage
- **Features**:
  - User struct with validation
  - Admin user creation
  - Method implementation
- **Run**: `go run main.go`

#### `struct-practice/`

- **Purpose**: Advanced struct and interface practice
- **Features**:
  - Note-taking application with JSON persistence
  - Todo list management
  - Interface implementation (`saver`, `outputtable`)
  - File I/O operations
- **Run**: `go run main.go`

### 🏦 Practical Applications

#### `bank-acct/`

- **Purpose**: Bank account simulation
- **Features**:
  - Account balance management
  - Credit/debit operations
  - Transaction history tracking
  - Comprehensive unit tests
- **Run**: `go run main.go`
- **Test**: `go test ./bank_account/`

#### `starter/`

- **Purpose**: Investment calculator
- **Features**:
  - Future value calculations
  - Inflation-adjusted returns
  - Interactive user input
  - Financial planning tools
- **Run**: `go run investment_calculator.go`

#### `assesment/`

- **Purpose**: Financial analysis tool
- **Features**:
  - Revenue and expense calculations
  - Tax rate computations
  - EBT (Earnings Before Tax) calculation
  - Profit margin analysis
  - File output for results
  - Prompts for `Revenue`, `Expenses`, and `Tax Rate`
  - Writes consolidated results to `finance.txt` with formatted values:
    - `EBT: <value>` (one decimal)
    - `Profit: <value>` (one decimal)
    - `Ratio: <value>` (one decimal)
- **Run**: `go run profit_calculator.go`

  Example:

  ```text
  Revenue: 1000
  Expenses: 200
  Tax Rate: 10
  800.0
  720.0
  1.111
  # Creates finance.txt with:
  # EBT: 800.0
  # Profit: 720.0
  # Ratio: 1.1
  ```

### 📊 Newer Projects

#### `price-calculator/`

- **Purpose**: Calculate tax-included prices from an input list and persist results
- **Features**:
  - Reads base prices from `prices.txt`
  - Converts lines to `float64` safely with error handling
  - Computes tax-inclusive totals for multiple tax rates (0%, 7%, 10%, 15%)
  - Persists results as JSON (`result_0.json`, `result_7.json`, `result_10.json`, `result_15.json`)
  - Modular packages: `conversions`, `filemanager`, `prices`
- **Run**:
  - `go run main.go`
  - Outputs JSON files in the project directory

#### `functions/`

- **Purpose**: Practice higher-order functions and function-returning-functions
- **Features**:
  - Demonstrates passing functions as arguments
  - Shows returning functions based on runtime conditions
  - Example transforms: double and triple number operations
- **Run**: `go run main.go`

#### `array/maps/`

- **Purpose**: Explore Go maps (create, read, update, delete)
- **Features**:
  - Initialize maps with string keys and values
  - Read by key, add entries, and delete entries
  - Prints state after each mutation
- **Run**: `go run maps/map.go`

## 🚀 Getting Started

### Prerequisites

- Go 1.19 or later
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd go
   ```

2. Navigate to any project directory:

   ```bash
   cd array  # or any other project
   ```

3. Run the project:
   ```bash
   go run main.go
   ```

## 📚 Learning Concepts Covered

- **Basic Go Syntax**: Variables, functions, control structures
- **Data Types**: Arrays, slices, structs, interfaces
- **Memory Management**: Pointers, value vs reference types
- **Error Handling**: Custom errors, error propagation
- **File I/O**: Reading and writing files, JSON serialization
- **Testing**: Unit tests, test coverage
- **Package Management**: Go modules, package organization
- **Concurrency**: Goroutines (in some projects)
- **Interface Design**: Polymorphism, method sets

## 🛠️ Project Features

### Data Structures

- Arrays and slices manipulation
- Struct definitions with methods
- Interface implementations
- JSON serialization/deserialization

### File Operations

- Reading user input
- Writing to files
- JSON data persistence
- Error handling for file operations

### Financial Calculations

- Investment future value calculations
- Tax and profit calculations
- Inflation adjustments
- Financial ratio analysis

### User Management

- User registration and validation
- Admin user creation
- Data persistence and retrieval

## 📝 Usage Examples

### Array Operations

```go
hobbies := [3]string{"Coding", "Anime", "Manga"}
lastHobbies := hobbies[1:]
courseGoals := []string{"be better at go", "land a new job"}
```

### Struct and Interface

```go
type saver interface {
    Save() error
}

type outputtable interface {
    saver
    Display()
}
```

### Financial Calculations

```go
futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
profit := ebt * (1 - taxRate/100)
```

## 🧪 Testing

Some projects include comprehensive unit tests:

```bash
# Run tests for bank account project
cd bank-acct
go test ./bank_account/
```

## 📖 Learning Path

1. Start with `array/` for basic Go concepts
2. Move to `structs/` for object-oriented concepts
3. Practice with `struct-practice/` for interfaces
4. Build applications with `bank-acct/`, `starter/`, and `assesment/`

## 🤝 Contributing

This is a learning repository. Feel free to:

- Add new practice projects
- Improve existing code
- Add more comprehensive tests
- Enhance documentation

## 📄 License

This project is for educational purposes and learning Go programming.
