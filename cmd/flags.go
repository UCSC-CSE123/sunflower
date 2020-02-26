package cmd

import (
	"flag"
	"time"
)

type args struct {
	Host         string
	Port         string
	Duration     time.Duration
	Autos        int
	InitialCount int
	Delta        int
}

func getFlags() args {
	defaults := args{
		Host: "localhost",
		Port: "8080",
	}

	flag.IntVar(&defaults.Autos, "nAutos", 5, "number of autos to run during the simulation")
	flag.IntVar(&defaults.Delta, "delta", 25, "the amount of passengers to change during a stop [rand(-delta,delta)]")
	flag.IntVar(&defaults.InitialCount, "passengers", 50, "the amount of passengers autos start with")
	flag.DurationVar(&defaults.Duration, "period", 5*time.Second, "The periodicity of auto stops")
	flag.Parse()

	arguments := flag.Args()

	if len(arguments) >= 1 {
		defaults.Host = arguments[0]
	}

	if len(arguments) >= 2 {
		defaults.Port = arguments[1]
	}

	return defaults
}
