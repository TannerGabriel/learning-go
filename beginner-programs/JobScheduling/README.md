# Job scheduling using Cron

Cron helps you to execute a job at a specific time interval or date. This can be a beneficial feature to run a process or task in the background, for example, pulling data from an API server or syncing data between servers.

## Installation

To install the current latest release.

```bash
go get github.com/robfig/cron/v3@v3.0.0
```

## Getting started

You can then register functions to be invoked at a specific time/schedule. Cron will then repeatedly run them in their own goroutines.

```Go
package main

import (
    "fmt"
    "time"

    "github.com/robfig/cron/v3"
)

func main() {

    c := cron.New()

    // Funcs are automatically invoked in their own goroutine, asynchronously
    c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
    c.AddFunc("TZ=America/New_York 30 04 * * * *", func() { fmt.Println("Runs at 04:30 New York time every day") })
    c.AddFunc("@hourly", func() { fmt.Println("Runs every hour") })
    c.AddFunc("@every 0h0m1s", func() { sayHelloTo("Someone!") })
    c.Start()

    // Funcs may also be added to a running Cron
    c.AddFunc("@daily", func() { fmt.Println("Every day") })

    // Added time to see output
    time.Sleep(10 * time.Second)

    c.Stop() // Stop the scheduler (does not stop any jobs already running)

}

func sayHelloTo(name string) {
    fmt.Println("Hello ", name)
}

```

## CRON Expression Format

A CRON expression represents a set of time composed of 5 space-separated fields.

```bash
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
```

Here is a example of using the CRON expression format:

```Go
# Runs at 6am in time.Local
cron.New().AddFunc("0 6 * * ?", ...)
```

A CRON expression can also be created using a parser:

```Go
specParser := NewParser(Minute | Hour | Dom | Month | Dow)
sched, err := specParser.Parse("0 0 15 */3 *")
```

## Predefined schedules

The CRON library provides a few predefined schedules that can be used instead of the expression formats.

```Go
Entry                  | Description                                | Equivalent To
-----                  | -----------                                | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 1 * *
@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 * * * *
```

Here is an example of how to use a predefined schedule:

```Go
c.AddFunc("@hourly", func() { fmt.Println("Runs every hour") })
```

## Intervals

You may also schedule a job to execute at a fixed interval. This can be done using the `@every` keyword, followed by the interval you want to run the job in.

```Go
c.AddFunc("@every 0h0m1s", func() { sayHelloTo("Someone!") })
```