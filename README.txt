# Fortune Reader

A simple command-line application that displays random quotes or "fortunes" from a text file.

## Overview

This Go application reads a file named `content.txt` containing quotes or sayings (one per line), and displays a random one each time the program is run.

## Features

- Reads quotes from a text file
- Randomly selects one quote to display
- Ignores empty lines
- Simple and lightweight

## Requirements

- Go 1.13 or later

## Installation

1. Clone this repository or download the source code:

```
git clone https://github.com/yourusername/fortune-reader.git
cd fortune-reader
```

2. Create a `content.txt` file in the same directory as the application with your quotes/fortunes (one per line).

Example `content.txt`:
```
The best way to predict the future is to create it.
Life is what happens when you're busy making other plans.
The journey of a thousand miles begins with one step.
```

## Usage

1. Build the application:

```
go build -o fortune-reader
```

2. Run the application:

```
./fortune-reader
```

The program will output a random quote from your `content.txt` file.

## Error Handling

The application will panic with appropriate error messages in these situations:
- If `content.txt` cannot be found or opened
- If `content.txt` is empty or contains only empty lines

## How It Works

1. The program opens and reads the `content.txt` file
2. It stores non-empty lines in an array
3. It generates a random number to select one of the lines
4. It displays the selected line to the console

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.