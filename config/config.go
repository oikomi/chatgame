

package config

import (
	"os"
	"encoding/json"
	"flag"
	//"fmt"
	"log"
)

type Config struct {
	Listen string
	Logfile    string
}

func LoadConfig(configpath string) (cfg Config, err error) {
	//log.Println(configpath)
	var configfile string
	flag.StringVar(&configfile, "config", configpath, "config file")
	flag.Parse()

	file, err := os.Open(configfile)
	if err != nil {
		log.Fatalln("Open configfile failed")
		return
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&cfg)
	if err != nil {
		return
	}
	return
}

func DumpConfig(cfg *Config) {
	//fmt.Printf("Mode: %s\nListen: %s\nServer: %s\nLogfile: %s\n", 
	//cfg.Mode, cfg.Listen, cfg.Server, cfg.Logfile)
}