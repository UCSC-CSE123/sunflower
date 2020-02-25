package cmd

import "os"

type args struct {
	Host string
	Port string
}

func getFlags() args {
	defaults := args{
		Host: "localhost",
		Port: "8080",
	}

	arguments := os.Args[1:]

	if len(arguments) >= 1 {
		defaults.Host = arguments[0]
	}

	if len(arguments) >= 2 {
		defaults.Port = arguments[1]
	}

	return defaults
}
