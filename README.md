
# Channel Manager
## Easily await concurrent tasks in Go with or without a deadline


### Await tasks with a deadline

```go
package main

import (
    "log"
    "testing"
    "time"

    "github.com/joshring/channelmgr"
)

func main()
{
    mgr := channelmgr.NewManager()

    mgr.AddTask(
        func(args []any) {
            log.Printf("Sleeping for 1sec\n")
            time.Sleep(1 * time.Second)
            log.Printf("finished 1sec sleep\n")
        },
        []any{},
    )

    mgr.AddTask(
        func(args []any) {
            log.Printf("Sleeping for 2sec\n")
            time.Sleep(2 * time.Second)
            log.Printf("finished 2sec sleep\n")
        },
        []any{},
    )

    timeout := 3 * time.Second
    log.Printf("waiting for jobs to complete for %v\n", timeout)
    timeStart := time.Now()
    timeTaken := time.Hour

    ok := mgr.WaitWithDeadline(timeout)
    if ok {
        timeTaken = time.Since(timeStart)
        log.Printf("Complete in: %vsec", time.Since(timeStart))
    } else {
        log.Fatal("unable to wait for deadline correctly")
    }
}
```


### Await tasks without a deadline

```go
package main

import (
    "log"
    "testing"
    "time"

    "github.com/joshring/channelmgr"
)

func main()
{
    mgr := channelmgr.NewManager()

    mgr.AddTask(
        func(args []any) {
            log.Printf("Sleeping for 1sec\n")
            time.Sleep(1 * time.Second)
            log.Printf("finished 1sec sleep\n")
        },
        []any{},
    )

    mgr.AddTask(
        func(args []any) {
            log.Printf("Sleeping for 2sec\n")
            time.Sleep(2 * time.Second)
            log.Printf("finished 2sec sleep\n")
        },
        []any{},
    )

    log.Printf("waiting for jobs to complete without timeout\n")
    timeStart := time.Now()

    mgr.WaitWithoutDeadline()

    timeTaken := time.Since(timeStart)
    log.Printf("Complete in: %vsec", timeTaken)

    if timeTaken.Truncate(time.Second) != 2*time.Second {
        log.Fatal("jobs did not complete within expected timeframe")
    }
}
```