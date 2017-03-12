package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/as27/setenv"
)

func main() {
	file := flag.String("f", "", "Load values from file")
	printEnv := flag.Bool("print", false, "Prints out all Enviroment variables")
	flag.Parse()
	if *file != "" {
		filename := *file
		envs, err := setenv.ParseFile(filename)
		if err != nil {
			log.Println(err)
		}
		err = setenv.SetEnv(envs)
		if err != nil {
			log.Println(err)
		}
	}
	if *printEnv {
		envs := os.Environ()
		for _, e := range envs {
			fmt.Println(e)
		}

	}

}
