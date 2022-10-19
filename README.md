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

- To generate neplai calendar

  ```
  calendarN nep
  ```

  To Generete nepali calendar for specific year and month

  ```
  calendarN nep 2079-05
  ```

- To generate english calendar

  ```
  calendarN eng
  ```

      To Generete english calendar for specific year and month

  ```
  calendarN nep 2022-08
  ```

- To check today's date

  ```
  calendarN today
  ```

- To create countdown for a date (supports english only for now)

  ```
  calendarN coutdown 2022-08-18
  ```

  This supports extra flags

  - --name provide the name for the calendar, generates random characters if not provided
  - --run to run the calendar for n seconds(default is 5), can be set to -1 for infinite
  - --save to save the current coutdown (to be implemented)

- Date converter
  ```
  calendarN convert etn '2022-08-18'
  ```
  Gives the converted date for the english to nepali date

### Requirement

- Go 1.18+

### Contributing
