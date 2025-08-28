# CalendarN

CalendarN is calendar New :)
CLI tool to get the details for the calendar

[![codecov](https://codecov.io/gh/samit22/calendarN/branch/main/graph/badge.svg?token=A5XND1948Y)](https://codecov.io/gh/samit22/calendarN)
[![goreport](https://goreportcard.com/badge/github.com/samit22/calendarN)](https://goreportcard.com/report/github.com/samit22/calendarN)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=bugs)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)

### Installation

```
go install github.com/samit22/calendarN@latest
```

### Commands

#### Nep

This command is for Nepali Calendar
Available commands:

```bash
calendarN nep
```

To generate nepali calendar for specific year and month

```bash
calendarN nep 2079-05
```

#### Eng

This command is for English Calendar
Available commands:

- To generate english calendar

```bash
 calendarN eng
```

To generate english calendar for specific year and month

```bash
  calendarN eng 2022-08
```

#### Today

- To check today's date

```bash
  calendarN today
```

This supports extra flags

- --m for minified date (-m)
- --j for date in JSON (-j)

#### Countdown

Shows countdowns

- To create countdown for a date (supports english only for now)

```bash
  calendarN countdown 2022-08-18
```

This supports extra flags

- --name provide the name for the calendar, generates random characters if not provided
- --run to run the calendar for n seconds(default is 5), can be set to -1 for infinite
- --save to save the current countdown

To save a countdown with a name

```bash
calendarN countdown -n "My Birthday" 2024-08-18 -s
```

To list all the countdowns

```bash
calendarN countdown all
```

To get a specific countdown

```bash
 calendarN countdown show -n 'My Birthday'
```

To delete a countdown

```bash
calendarN countdown delete -n 'My Birthday'
```

### Convert

- Date converter

  ```bash
  calendarN convert etn '2022-08-18'
  ```

  Gives the converted date for the english to nepali date

### Requirement

- Go 1.18+

### Contributing

- For new feature/bug create an issue
- Check for the issues and assign yourself
- Make sure added code has 80% unit test coverage
