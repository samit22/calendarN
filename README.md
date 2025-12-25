# CalendarN

CalendarN is a powerful CLI tool for working with Nepali (Bikram Sambat) and English (Gregorian) calendars. Convert dates, view calendars, track countdowns, and more!

[![codecov](https://codecov.io/gh/samit22/calendarN/branch/main/graph/badge.svg?token=A5XND1948Y)](https://codecov.io/gh/samit22/calendarN)
[![goreport](https://goreportcard.com/badge/github.com/samit22/calendarN)](https://goreportcard.com/report/github.com/samit22/calendarN)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=bugs)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)

## Features

- 📅 View Nepali (BS) and English calendars
- 🔄 Convert dates between Nepali and English
- ⏱️ Create and manage countdowns
- 📊 Track year progress with visual progress bars
- 🎨 Beautiful terminal output with colors

## Installation

### Homebrew (macOS/Linux)

The easiest way to install calendarN on macOS or Linux:

```bash
# Add the tap
brew tap samit22/calendarN https://github.com/samit22/calendarN

# Install calendarN
brew install calendarn
```

To upgrade to the latest version:

```bash
brew upgrade calendarn
```

### Go Install

If you have Go installed (1.21+):

```bash
go install github.com/samit22/calendarN@latest
```

### Build from Source

```bash
git clone https://github.com/samit22/calendarN.git
cd calendarN
make build
./calendarN --help
```

### Manual Installation

Download the latest release from the [Releases page](https://github.com/samit22/calendarN/releases) for your platform.

## Commands

### Today

Get today's date in both Nepali and English with year progress:

```bash
calendarN today
```

**Flags:**
- `-m, --minified` - Minified date output
- `-j, --json` - Output in JSON format

```bash
# Minified output
calendarN today -m

# JSON output
calendarN today -j
```

### Nepali Calendar (nep)

View the Nepali calendar:

```bash
# Current month
calendarN nep

# Specific year and month (BS)
calendarN nep 2081-09
```

### English Calendar (eng)

View the English calendar:

```bash
# Current month
calendarN eng

# Specific year and month
calendarN eng 2024-12
```

### Date Conversion (convert)

Convert dates between English and Nepali:

```bash
# English to Nepali
calendarN convert etn 2024-12-25
# Output: Eng: 2024-12-25 => 2081-09-10 || २०८१-०९-१०
```

### Countdown

Create and manage countdowns to important dates:

```bash
# Create a countdown
calendarN countdown 2025-01-01

# Create with a name and save it
calendarN countdown -n "New Year" 2025-01-01 -s

# With specific time
calendarN countdown -n "Meeting" 2025-01-15 14:30:00 -s
```

**Flags:**
- `-n, --name` - Name for the countdown (random if not provided)
- `-r, --run` - Run countdown for n seconds (default: 5, use -1 for infinite)
- `-s, --save` - Save the countdown for later
- `-o, --overwrite` - Overwrite existing countdown with same name

**Managing saved countdowns:**

```bash
# List all saved countdowns
calendarN countdown all

# Show a specific countdown
calendarN countdown show -n "New Year"

# Delete a countdown
calendarN countdown delete -n "New Year"
```

### Version

Check the installed version:

```bash
calendarN version
```

## Examples

### View Today's Date

```bash
$ calendarN today -m
बुधवार, २५ पुष, २०८१ | Wednesday, 25 December, 2024
```

### Convert a Date

```bash
$ calendarN convert etn 2024-12-25
Eng: 2024-12-25 => 2081-09-10 || २०८१-०९-१०
```

### Create a Birthday Countdown

```bash
$ calendarN countdown -n "Birthday" 2025-08-15 -s -r 10
Countdown for: Birthday
45 days 12 hours 30 minutes 15 seconds
```

## Supported Date Range

- **Nepali (BS):** 2000 BS to 2090 BS
- **English:** 1944 AD onwards (corresponding to the BS range)

## Requirements

- Go 1.21+ (for building from source)
- macOS, Linux, or Windows (x86_64 or ARM64)

## Development

```bash
# Clone the repository
git clone https://github.com/samit22/calendarN.git
cd calendarN

# Run tests
make test

# Build
make build

# Cross-compile for all platforms
make release-all
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Make sure your code has adequate test coverage (80%+)
4. Commit your changes (`git commit -m 'Add amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

For bugs or feature requests, please [create an issue](https://github.com/samit22/calendarN/issues).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Samit Ghimire** - [info@samitghimire.com.np](mailto:info@samitghimire.com.np)
