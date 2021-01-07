package main

import (
	"putil/log"
)

func main() {
	plog.Fatal("Fatal")
	plog.Fatalf("Fatalf:%d", 1)
	plog.Debug("Debug")
	plog.Debugf("Debug:%d", 2)
	plog.Info("Info")
	plog.Infof("Infof:%d", 3)
}
