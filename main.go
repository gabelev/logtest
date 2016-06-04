package main

import (
	"fmt"
	"github.com/Magnetic/go-logging"
	"github.com/Magnetic/go-logging/example"
	"gopkg.in/inconshreveable/log15.v2"
	"os"
)

var Log = logging.RootInstance.New("modile", "log.test")

func main() {
	handler := log15.StreamHandler(os.Stdout, log15.LogfmtFormat())
	Log.SetHandler(handler)
	fmt.Printf("%t\n", example.IsPrime(13))
}
