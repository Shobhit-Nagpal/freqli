package main

import (
	"flag"
	"freqli/src/analysis"
)

func main() {
  shell := flag.String("shell", "bash", "The shell that is being used by the user")
  flag.Parse()

  analysis.Initiate(*shell)
}
