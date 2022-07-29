# Calendar

CLI tool to get the details for the calendar

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

### Requirement

- Go 1.18+

### Contributing
