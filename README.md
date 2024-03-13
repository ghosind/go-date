# Date

![test](https://github.com/ghosind/go-date/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghosind/go-date)](https://goreportcard.com/report/github.com/ghosind/go-date)
[![codecov](https://codecov.io/gh/ghosind/go-date/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-date)
![Version Badge](https://img.shields.io/github/v/release/ghosind/go-date)
![License Badge](https://img.shields.io/github/license/ghosind/go-date)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/go-date.svg)](https://pkg.go.dev/github.com/ghosind/go-date)

The extends of the Golang built-in `time` package.

## Installation

Run the following command to install the library, which requires Go 1.20 and later versions.

```sh
go get -u github.com/ghosind/go-date
```

## Getting Started

This library overwrites the built-in format layouts, the following example is a simple implementation of parsing a string:

```go
str := "2024-01-10 23:59:30"
layout := "YYYY-MM-DD HH:mm:ss"
tm, err := date.Parse(layout, str)
if err != nil {
  // handle error
}
fmt.Print(tm) // 2024-01-10 23:59:30 +0000 CST
```

You can also use the `Format` method to format the `Time` to a string:

```go
tm := Date(2024, time.January, 1, 23, 59, 30, 0)
fmt.Print(fm.Format("YYYY-MM-DD HH:mm:ss")) // 2024-01-10 23:59:30
```

## Available Formats

| Format | Description                                 |       Example        |
| :----: | :------------------------------------------ | :------------------: |
| `YYYY` | 4-digits year                               |        `2023`        |
|  `YY`  | 2-digits year                               |         `23`         |
|  `MM`  | 2-digits month                              |      `01`-`12`       |
|  `M`   | Month, beginning at 1                       |       `1`-`12`       |
| `MMMM` | The month name                              | `January`-`December` |
| `MMM`  | The abbreviated month name                  |     `Jan`-`Dec`      |
|  `DD`  | The day of month, 2-digits                  |      `01`-`31`       |
|  `D`   | The day of month, beginning at 1            |       `1`-`31`       |
| `dddd` | The day of week                             |  `Sunday`-`Friday`   |
| `ddd`  | The abbreviated name of weekday             |     `Sun`-`Fri`      |
|  `d`   | The day of week, beginning at 0 (Sunday)    |       `0`-`6`        |
|  `HH`  | The hour of 24-hour clock, 2-digits         |      `00`-`23`       |
|  `H`   | The hour of 24-hour clock, beginning at 1   |       `0`-`23`       |
|  `hh`  | The hour of 12-hour clock, 2-digits         |      `01`-`12`       |
|  `h`   | The hour of 12-hour clock, beginning at 1   |       `1`-`12`       |
|  `mm`  | The minutes, 2-digits                       |      `00`-`59`       |
|  `m`   | The minutes                                 |       `0`-`59`       |
|  `ss`  | The seconds, 2-digits                       |      `00`-`59`       |
|  `s`   | The seconds                                 |       `0`-`59`       |
| `SSS`  | The milliseconds, 3-digits                  |     `000`-`999`      |
|  `SS`  | The tens of milliseconds, 2-digits          |      `00`-`99`       |
|  `S`   | The hundreds of milliseconds, 1-digit       |       `0`-`9`        |
|  `A`   | Post or ante meridiem, in upper case        |      `AM`, `PM`      |
|  `a`   | Post or ante meridiem, in lower case        |      `am`, `pm`      |
|  `Z`   | Timezone offset from UTC, separate by colon |       `-08:00`       |
|  `ZZ`  | Timezone offset from UTC                    |       `-0800`        |
