# String Calculator

This is a console application that implements a string-based calculator for Kata Academy's educational task.

## Description
The calculator performs the following operations on strings:

- Addition: Concatenates two strings.
- Subtraction: Removes a substring from the string.
- Multiplication: Repeats the string a specified number of times.
- Division: Divides the string into a specified number of parts and returns the first part.

## How to Use
To run the program, follow these steps:

- Navigate to the project directory:

```bash
cd calcFolder
```

- Start the application:

```bash
go run main.go
```

The program will launch a console interface where you can input expressions in the following format:

```bash
"string" operator "string" (or number)
```

- Example input:

```bash
"Hello" + "World"
"Bye-bye!" - "World!"
"Golang" * 3
"Example!!!" / 2
```

- Requirements:

Go version 1.22 or higher