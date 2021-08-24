package main

import (
	"flag"
	"log"

	B "github.com/yettera765/mmdb-builder/builder"
	"github.com/yettera765/mmdb-builder/config"
)

var (
	confPath   string
	outputPath string
)

func init() {
	flag.StringVar(&confPath, "c", "mmdb.yaml", "path to config file, in yaml format")
	flag.StringVar(&outputPath, "o", "mmdb/Country.mmdb", "path to output file")
	flag.Parse()
}

func main() {
	log.Println("mmdb-builder start")
	cfg, err := config.Init(confPath)
	if err != nil {
		log.Println(err)
		log.Fatalln("initial config error")
	}

	builder := B.New()
	builder.Process(cfg)
	builder.Build(outputPath)
	log.Println("build done")
}
