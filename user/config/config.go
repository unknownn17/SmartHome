package config

import (
	"log"
	"os"

	"github.com/spf13/cast"
	"github.com/subosito/gotenv"
)


type Config struct{
	MCPort string
}



func Connection()*Config{
	c:=Config{}
	err := gotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	c.MCPort=cast.ToString(os.Getenv("MC_PORT"))  

	return &Config{
		MCPort: c.MCPort,
	}
}