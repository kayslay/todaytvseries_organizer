package main

import (
	"flag"
	"log"
	"os"

	"github.com/kayslay/todaytvseries_organizer/config"
	"github.com/kayslay/todaytvseries_organizer/helpers"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	err := helpers.Start(*config.InitConfig())
	if err != nil {
		log.Println(err)
	}
}
