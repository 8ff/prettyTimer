# PrettyTimer

![demo](media/demo.svg)

## Example code
```go
package main

import (
	"time"

	"github.com/8ff/prettyTimer"
)

func main() {
	ts := prettyTimer.NewTimingStats()
	ts.RecordTiming(1 * time.Second)
	ts.RecordTiming(2 * time.Second)
	ts.RecordTiming(3 * time.Second)
	ts.PrintStats()
}
```

## Example from [examples/main.go](examples/main.go)
```bash
$ examples % go run main.go 
Min Time: 1s, Max Time: 3s, Avg Time: 2s, Count: 3
50th: 2s, 90th: 3s, 99th: 3s
```
