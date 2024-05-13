package main

import (
	"flag"
	"freqli/pkg/analysis"
)

func main() {
  shell := flag.String("shell", "bash", "The shell that is being used by the user")
  flag.Parse()

  analysis.Initiate(*shell)
}
