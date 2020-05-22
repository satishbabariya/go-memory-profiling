package main

import (
  //...
	"github.com/pkg/profile"
)

func main() {
	prof := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)

  //...
  
  prof.Stop()
}
