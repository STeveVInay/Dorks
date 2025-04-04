# Dorks
This Python script automates Google Dorks generation for a user-specified domain. It replaces "{domain}" in a predefined list of 52 dorks with the input domain, displays the results, and optionally saves them to a file named "&lt;domain>_dorks.txt". Useful for security researchers to find exposed data ethically.

# Google Dorks Automation Tool
This tool generates Google Dorks for a user-specified domain, reading from `dorks.txt` and saving to `<domain>_dorks.txt`. Available in Python and Go.

## Python Version
### Dependencies
- **Python**: 3.5+
- **Dependencies**: None (standard library: `os`, `concurrent.futures`)

### Installation
1. Install Python 3.5+ (`python3 --version`).
2. Run: `python dork_generator.py`

## Go Version
### Dependencies
- **Go**: 1.11+
- **Dependencies**: None (standard library: `bufio`, `fmt`, `os`, `strings`, `sync`)

### Installation
1. Install Go 1.11+ (`go version`).
2. Run: `go run dork_generator.go`
   - Or build: `go build dork_generator.go` and run `./dork_generator`

## Usage
1. Create/edit `dorks.txt` with dorks (e.g., `site:{domain} inurl:"config.php"`).
2. Run the script, enter a domain (e.g., `google.com`), and choose to save results.

## Notes
- Ensure `dorks.txt` exists in the same directory.
- Generated files are named `<domain>_dorks.txt`.
