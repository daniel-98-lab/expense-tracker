# GitHub Expense Tracker

Application to manage your finances

## Feature

- Add expense
- Remove expense
- Update expense
- View all expenses with optional criteria(category)
- View a summary of expenses with optional criteria(month or category)
- Export expenses in csv or json file

## Installation

To get started, you will need to have Go installed on your machine. Follow the instructions [here](https://golang.org/doc/install) if you don't have Go set up.

1. Clone this repository:
```bash
git clone https://github.com/daniel-98-lab/expense-tracker.git
    ```

2. Navigate to the project Repository
```bash
cd expense-tracker
```

3. Install dependencies:
```bash
go mod tidy
```

4. Build the project:
```bash
go build -o expense-tracker ./cmd
```

## Usage
To use the tool, run the following command from the terminal:
```bash
./expense-tracker <command> -t "desc"
```

## Examples:
1. Add Expense: (-a Amount, -d Description, -D date, -c category)
```bash
./expense-tracker add -a 38.23 -d "esto es la descripción del gasto 4"  -D 2024-10-15 -c viaje
```

2. List Expenses: (-c Category)
```bash
./expense-tracker list
./expense-tracker list -c viaje
```

3. Delete Expense: (-i ID)
```bash
./expense-tracker delete -i 3
```

4. Export Expenses: (default json)
```bash
./expense-tracker export
./expense-tracker export -t json
./expense-tracker export -t csv
```

5. Get Expense Summary: (-m Month, -c Category)
```bash
./expense-tracker summary
./expense-tracker summary -m 10
./expense-tracker summary -c viaje
```

6. Update Expense: (-i ID, -a Amount, -d Description, -D date, -c category)
```bash
./expense-tracker update -i 2 -a 50.00 -d "Nueva descripción" -D 2024-10-20 -c comida
```