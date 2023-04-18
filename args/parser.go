package args

import (
	"log"
	"strconv"
)

type Args struct {
	Port int
}

func (a *Args) Parse(argList []string) {
	a.Port = 80
	var err error

	if len(argList) > 1 {
		a.Port, err = strconv.Atoi(argList[1])
		if err != nil {
			log.Fatalf("Invalid port: %s\n", argList[1])
		}
	}
}

func ParseArgs(argList []string) Args {
	args := Args{}
	args.Parse(argList)
	return args
}
