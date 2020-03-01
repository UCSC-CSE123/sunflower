package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type args struct {
	Host            string
	Port            string
	StopPeriod      time.Duration
	StopDuration    time.Duration
	StopProbability int
	Autos           int
	InitialCount    int
	Delta           int
	Seed            int64
}

func getFlags() args {
	defaults := args{
		Host: "localhost",
		Port: "8080",
	}

	flag.Usage = func() {
		fmt.Printf("Usage: %s [host address] [port] [flags]\n", os.Args[0])
		fmt.Printf("\thost\n\t\tThe host address to bind to (default localhost)\n")
		fmt.Printf("\tport\n\t\tThe port to bind to (default 8080)\n")
		fmt.Printf("flags:\n")
		flag.PrintDefaults()
	}

	flag.IntVar(&defaults.Autos, "nAutos", 5, "number of autos to run during the simulation")
	flag.IntVar(&defaults.Delta, "delta", 10, "the amount of passengers to change during a stop [rand(-delta,delta)]")
	flag.IntVar(&defaults.InitialCount, "passengers", 50, "the amount of passengers autos start with")
	flag.IntVar(&defaults.StopProbability, "probability", 75, "the probability that a bus stops")
	flag.Int64Var(&defaults.Seed, "seed", time.Now().UnixNano(), "the seed to pass to the RNG -- by default the seed is the current time")
	// Why 11 and 5?
	// Because there're prime!
	// So there's no chance that a bus is perpetually stuck in a "Loading" state because the stop and loading time are multiples of each other.
	flag.DurationVar(&defaults.StopPeriod, "period", 11*time.Second, "The periodicity of auto stops")
	flag.DurationVar(&defaults.StopDuration, "duration", 5*time.Second, "The length of time a bus is stopped for")
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
