package config

import (
	"fmt"

	"github.com/jgolang/sirius"
)

var sjson *sirius.SJson

func init() {
	var err error
	sjson, err = sirius.Open("config.json")
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s", err))
	}
}
