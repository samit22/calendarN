# CalendarN

A powerful CLI tool for working with Nepali (Bikram Sambat) and English (Gregorian) calendars.

[![codecov](https://codecov.io/gh/samit22/calendarN/branch/main/graph/badge.svg?token=A5XND1948Y)](https://codecov.io/gh/samit22/calendarN)
[![goreport](https://goreportcard.com/badge/github.com/samit22/calendarN)](https://goreportcard.com/report/github.com/samit22/calendarN)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=bugs)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)

## Quick Start

```bash
# Install via Homebrew
brew tap samit22/calendarN https://github.com/samit22/calendarN
brew install calendarn

# View today's date
calendarN today

# Convert English to Nepali
calendarN convert etn 2024-12-25

# Convert Nepali to English
calendarN convert nte 2081-09-10
```

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Today](#today)
  - [Calendars](#calendars)
  - [Date Conversion](#date-conversion)
  - [Countdown](#countdown)
- [Supported Date Range](#supported-date-range)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Features

| Feature | Description |
|---------|-------------|
| View Calendars | Display Nepali (BS) and English calendars |
| Date Conversion | Convert dates between Nepali and English (both directions) |
| Countdowns | Create and manage countdowns to important dates |
| Year Progress | Track year progress with visual progress bars |
| Beautiful Output | Colorful terminal output with Devanagari support |

## Installation

### Homebrew (Recommended)

```bash
brew tap samit22/calendarN https://github.com/samit22/calendarN
brew install calendarn

# Upgrade to latest
brew upgrade calendarn
```

### Go Install

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

### Manual Download

Download the latest release from the [Releases page](https://github.com/samit22/calendarN/releases).

## Usage

### Today

Display today's date in both Nepali and English with year progress.

```bash
calendarN today           # Full output with progress bar
calendarN today -m        # Minified output
calendarN today -j        # JSON format
```

**Example Output:**
```
$ calendarN today -m
बुधवार, २५ पुष, २०८१ | Wednesday, 25 December, 2024
```

### Calendars

#### Nepali Calendar

```bash
calendarN nep             # Current month
calendarN nep 2081-09     # Specific year and month (BS)
```

#### English Calendar

```bash
calendarN eng             # Current month
calendarN eng 2024-12     # Specific year and month
```

### Date Conversion

Convert dates between English (Gregorian) and Nepali (Bikram Sambat) calendars.

#### English to Nepali (etn)

```bash
$ calendarN convert etn 2024-12-25
Eng: 2024-12-25 => 2081-09-10 || २०८१-०९-१०
```

#### Nepali to English (nte)

```bash
$ calendarN convert nte 2081-09-10
Nep: 2081-09-10 => 2024-12-25
```

### Countdown

Create and manage countdowns to important dates.

#### Create Countdown

```bash
calendarN countdown 2025-01-01                       # Quick countdown
calendarN countdown -n "New Year" 2025-01-01 -s      # Named and saved
calendarN countdown -n "Meeting" 2025-01-15 14:30:00 -s  # With time
```

**Flags:**

| Flag | Description |
|------|-------------|
| `-n, --name` | Name for the countdown |
| `-r, --run` | Run for n seconds (default: 5, use -1 for infinite) |
| `-s, --save` | Save the countdown |
| `-o, --overwrite` | Overwrite existing countdown |

#### Manage Countdowns

```bash
calendarN countdown all              # List all saved
calendarN countdown show -n "Name"   # Show specific
calendarN countdown delete -n "Name" # Delete
```

### Version

```bash
calendarN version
```

## Supported Date Range

| Calendar | Range |
|----------|-------|
| Nepali (BS) | 2000 BS - 2090 BS |
| English (AD) | 1944 AD onwards |

## Development

```bash
git clone https://github.com/samit22/calendarN.git
cd calendarN

make test         # Run tests
make build        # Build binary
make release-all  # Cross-compile for all platforms
```

**Requirements:** Go 1.21+

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Ensure test coverage is adequate (80%+)
4. Commit your changes (`git commit -m 'Add amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

For bugs or feature requests, please [create an issue](https://github.com/samit22/calendarN/issues).

## License

MIT License - see [LICENSE](LICENSE) for details.

## Author

**Samit Ghimire** - [info@samitghimire.com.np](mailto:info@samitghimire.com.np)
