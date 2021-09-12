package main

import (
	"flag"
	"fmt"
	"g-sig/pkg/config"
	"io/ioutil"
	"os"
)

var version = "0.1.0"

func init() {

	file, err := os.Open("g-sig.conf")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	config := config.NewConfig(buffer)
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Printf("g-sig version is %s", version)
		return
	}

}
