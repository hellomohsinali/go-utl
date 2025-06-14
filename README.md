# Utility Tool (utl)

A simple command-line application that displays the current time when the 't' parameter is provided.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or later recommended)

## Running the Application

### Method 1: Direct Execution

1. Clone this repository
2. Navigate to the project directory
3. Run the application:

```bash
go run main.go
```

The application will randomly output either "head" or "tail" in the console.

To display the current time, run:

```bash
go run main.go t
```

### Method 2: Global Command Installation

You can install the application as a global command called `utl` that can be run from anywhere:

1. Clone this repository
2. Navigate to the project directory
3. Make the installation script executable:

```bash
chmod +x install.sh
```

4. Run the installation script:

```bash
./install.sh
```

5. Now you can use the `utl` command from anywhere:

```bash
utl t    # Displays the current time
```

The installation also creates an alias `tl` that you can use as a shortcut for `utl t` to quickly display the current time.

## Example Output

For current time:
```
Current time: Mon, 01 Jan 2024 12:34:56 UTC
```

## How It Works

The application has one functionality:

2. **Current Time Display**: When the 't' parameter is provided, the application uses Go's `time` package to get the current time and displays it in a human-readable format (RFC1123).

## Building an Executable

If you want to build an executable file that you can run without Go installed:

```bash
go build -o utl
```

Then run the executable:

```bash
./utl t     # Displays the current time
```

Note: This is similar to what the installation script does, but without installing the command globally.
