# Word Definition CLI

This command-line tool helps you quickly get the definition of a word using the Gemini AI API. It provides simple definitions and examples for the words you look up, right in your terminal.

## Features

- Get definitions and examples for words directly from your command line
- Powered by Gemini AI for accurate and context-rich definitions
- Easy to use as a CLI command

## Prerequisites

- Go programming language installed on your system
- Gemini AI API key (you can obtain one from [Google AI Studio](https://makersuite.google.com/app/apikey))

## Installation

1. Clone this repository or download the source code.

2. Navigate to the project directory.

3. Build the project:
   ```
   go build -o word-def
   ```

4. (Optional) Move the built executable to a directory in your system's PATH to use it as a global command. For example:
   ```
   mv word-def /usr/local/bin/
   ```

## Configuration

Before using the tool, you need to set up your Gemini AI API key. Open the `main.go` file and locate this line:

```go
apiKey := ""
```

Replace the empty string with your actual API key:

```go
apiKey := "your-actual-api-key-here"
```

**Note**: Make sure not to share your API key publicly if you plan to distribute this code.

## Usage

To use the tool, simply run it followed by the word you want to define:

```
word-def example
```

If you've added it to your PATH, you can run it from anywhere in your terminal.

If you haven't added it to your PATH, you'll need to run it from the directory where the executable is located or provide the full path to the executable.

## Adding to PATH (Windows)

To add the tool to your PATH on Windows:

1. Press Win + X and select "System".
2. Click on "Advanced system settings".
3. Click on "Environment Variables".
4. Under "System variables", find and select "Path", then click "Edit".
5. Click "New" and add the full path to the directory containing your `word-def.exe`.
6. Click "OK" to close all dialogs.

After adding to PATH, you may need to restart your command prompt for the changes to take effect.

## Adding to PATH (macOS/Linux)

To add the tool to your PATH on macOS or Linux:

1. Open your shell configuration file (`~/.bash_profile`, `~/.bashrc`, or `~/.zshrc`).
2. Add the following line, replacing `/path/to/word-def` with the actual path:
   ```
   export PATH=$PATH:/path/to/word-def
   ```
3. Save the file and run `source ~/.bash_profile` (or the appropriate file) to apply changes.
