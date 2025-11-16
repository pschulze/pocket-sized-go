package main

import (
	"os"
	"time"

	"github.com/pschulze/pocket-sized-go/logger/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout))
	lgr.Infof("A little copying is better than a little dependency.")
	lgr.Errorf("Errors are values. Documention is for %s.", "users")
	lgr.Debugf("Make the zero (%d) value useful.", 0)

	lgr.Infof("Hello, %d %v", 2025, time.Now())
}
