# Butterfly

**Butterfly** is a command-line (CLI) tool written in Go that converts **Markdown (.md)** files into **PDF** documents, preserving basic formatting such as headings, paragraphs, and Unicode text.

The generated binary's name (`butterfly`) also appears as a subtle watermark in the footer of every page of the generated PDF.

---

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
  - [Cloning the repository](#cloning-the-repository)
  - [Building the binary](#building-the-binary)
- [Adding Butterfly to the Environment Variables](#adding-butterfly-to-the-environment-variables)
  - [Linux / macOS](#linux--macos)
  - [Windows](#windows)
- [Usage](#usage)
- [Example](#example)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [How It Works](#how-it-works)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- Converts `.md` files to `.pdf` via the command line.
- Supports headings (`#`, `##`, `###`, ...) with distinct font sizes.
- Supports paragraphs with automatic spacing.
- Supports Unicode/accented characters (via `fpdf`'s font translator).
- Automatic footer with a "butterfly" watermark on every page.
- Automatically appends the `.pdf` extension to the output file if the user forgets to include it.

---

## Requirements

- [Go](https://go.dev/dl/) version **1.21** or later installed.
- Internet access to download the project's dependencies (`go mod download`).
- Git (to clone the repository).

---

## Installation

### Cloning the repository

```bash
git clone https://github.com/YOUR_USERNAME/butterfly.git
cd butterfly
```

### Building the binary

Download the dependencies and build the project:

```bash
go mod tidy
go build -o butterfly main.go
```

- On **Linux/macOS**, this generates an executable binary called `butterfly`.
- On **Windows**, build the binary with the `.exe` extension:

```powershell
go build -o butterfly.exe main.go
```

---

## Adding Butterfly to the Environment Variables

To run the `butterfly` command from any directory in the terminal, you need to add the binary's path to your operating system's `PATH` environment variable.

### Linux / macOS

1. Move the compiled binary to a directory of your choice (e.g., `/usr/local/bin` or `$HOME/bin`):

```bash
mkdir -p $HOME/bin
mv butterfly $HOME/bin/
```

2. Add the directory to `PATH` by editing your shell's configuration file (`~/.bashrc`, `~/.zshrc`, or `~/.profile`, depending on the shell you use):

```bash
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

> If you use **Zsh** (default on macOS), replace `~/.bashrc` with `~/.zshrc`:

```bash
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

3. Verify that the command is available:

```bash
butterfly
```

You can also add the variable **temporarily** (valid only for the current terminal session):

```bash
export PATH="$PATH:/path/to/the/binary"
```

---

### Windows

#### Option 1 — Via PowerShell (permanent, current user)

```powershell
[Environment]::SetEnvironmentVariable("Path", $Env:Path + ";C:\path\to\the\binary", "User")
```

Afterward, close and reopen PowerShell/CMD for the change to take effect.

#### Option 2 — Via CMD (permanent, current user)

```cmd
setx PATH "%PATH%;C:\path\to\the\binary"
```

#### Option 3 — Via PowerShell (temporary, current session only)

```powershell
$Env:Path += ";C:\path\to\the\binary"
```

#### Option 4 — Graphical interface

1. Press `Win + R`, type `sysdm.cpl`, and press Enter.
2. Go to the **Advanced** tab -> **Environment Variables**.
3. Under **User variables** (or **System variables**, for all users), select the `Path` variable and click **Edit**.
4. Click **New** and add the full path to the folder containing `butterfly.exe` (e.g., `C:\Users\YourUser\bin`).
5. Click **OK** on all windows to save.
6. Open a new terminal (CMD or PowerShell) and test:

```powershell
butterfly
```

---

## Usage

Once installed and configured in `PATH`, the command can be run from anywhere:

```bash
butterfly <input_file.md> <output_file.pdf>
```

### Parameters

| Parameter     | Description                                                                 |
|---------------|-------------------------------------------------------------------------------|
| `input_file`  | Path to the Markdown (`.md`) file to be converted.                            |
| `output_file` | Name/path of the output PDF file. If the `.pdf` extension is not provided, it will be added automatically. |

If the arguments are not provided correctly, the program will display a usage message:

```
Usage: butterfly <input_file.md> <output_file.pdf>
```

---

## Example

Assuming a file `example_readme.md`:

```markdown
# Main Title

This is an example paragraph with **text** in Markdown.

## Subtitle

Another paragraph here.
```

Run:

```bash
butterfly example_readme.md output
```

Expected terminal output:

```
Success! PDF generated at: output.pdf
```

The `output.pdf` file will be created in the current directory, containing the formatted content and the "butterfly" watermark in the footer of every page.

---

## Project Structure

```
butterfly/
├── main.go       # Main application source code
├── go.mod        # Module definition and dependencies
├── go.sum        # Dependency checksums
└── README.md     # This file
```

---

## Dependencies

This project uses the following Go libraries:

- [`github.com/go-pdf/fpdf`](https://github.com/go-pdf/fpdf) — PDF document generation.
- [`github.com/yuin/goldmark`](https://github.com/yuin/goldmark) — CommonMark-compliant Markdown parser.

Install/update the dependencies with:

```bash
go mod tidy
```

---

## How It Works

1. The program reads the command-line arguments (`parseArgs`) and validates whether the output file has a `.pdf` extension.
2. The Markdown file's content is read from disk (`os.ReadFile`).
3. A PDF document is initialized in A4 format (`initPDF`), with a footer function configured to display the "butterfly" watermark.
4. The Markdown content is parsed by `goldmark`, generating a syntax tree (AST).
5. The tree is traversed (`ast.Walk`) and, depending on the node type found (heading, paragraph, text), the content is written to the PDF with the appropriate formatting.
6. The final PDF is saved to the path specified by the user.

---

## Contributing

Contributions are welcome! To contribute:

1. Fork this repository.
2. Create a branch for your feature/fix: `git checkout -b my-feature`.
3. Commit your changes: `git commit -m 'Add my feature'`.
4. Push to your fork: `git push origin my-feature`.
5. Open a Pull Request.

---

## License

This project is licensed under the terms of the **MIT** license. Feel free to use, modify, and distribute it as needed.
