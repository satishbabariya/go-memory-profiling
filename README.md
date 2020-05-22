# go-memory-profiling

### add profiling into the main() function of your command
```go
package main

import (
   //...
   "github.com/pkg/profile"
)

func main() {
   // Memory profiling
   defer profile.Start(profile.MemProfile).Stop()

   //...
}
```

### build and run your program
This will generate a *.pprof file in a temp folder, and tell you where it’s located (will be needed later)

```log
2017/08/03 14:26:28 profile: cpu profiling enabled, /var/...../cpu.pprof
```

### install graphviz if you don’t have it installed yet
This is used to generate the graph on a pdf. On a Mac, it’s a simple brew install graphviz. Refer to https://www.graphviz.org for other platforms.


###  run go tool pprof
Pass your binary location, and the location of the cpu.pprof file as returned when running your program.

You can generate the analysis in various formats. The PDF one is pretty amazing:

```shell
go tool pprof --pdf ~/go/bin/yourbinary /var/path/to/cpu.pprof > file.pdf
```
