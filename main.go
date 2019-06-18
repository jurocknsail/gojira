package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/gemalto/gojira/cmd"
	"github.com/spf13/viper"
)

func main() {
	initConfigFolder()
	cmd.Execute()
}

func initConfigFolder() {

	usr, usrErr := user.Current()
	if usrErr != nil {
		log.Fatal(usrErr)
	}

	configPath := usr.HomeDir + "/.gojira"
	configFile := "gojira"
	configFilePath := configPath + "/gojira.yaml"
	dodFilePath := configPath + "/dod.yaml"
	fileMode := os.ModePerm

	//Create config folder if not exists
	if _, statErr := os.Stat(configPath); os.IsNotExist(statErr) {
		direrr := os.Mkdir(configPath, fileMode)
		if direrr != nil {
			panic(direrr)
		}
		//Create the config file if not exists
		if _, stateFileErr := os.Stat(configFilePath); os.IsNotExist(stateFileErr) {
			os.OpenFile(configFilePath, os.O_RDONLY|os.O_CREATE, 0666)
		}
		//Create the config file if not exists
		if _, stateDoDFileErr := os.Stat(dodFilePath); os.IsNotExist(stateDoDFileErr) {
			os.OpenFile(dodFilePath, os.O_RDONLY|os.O_CREATE, 0666)
		}
	}

	//Create the config file if not exists, in case the folder was there but empty
	if _, stateFileErr := os.Stat(configFilePath); os.IsNotExist(stateFileErr) {
		os.OpenFile(configFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	}
	//Create the config file if not exists, in case the folder was there but empty
	if _, stateDoDFileErr := os.Stat(dodFilePath); os.IsNotExist(stateDoDFileErr) {
		os.OpenFile(dodFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	}

	// Init the actual config using viper
	viper.SetConfigName(configFile) // name of config file (without extension)
	viper.AddConfigPath(configPath) // call multiple times to add many search paths
	viper.AddConfigPath(dodFilePath)

	viper.SetConfigType("yaml")
	readerr := viper.ReadInConfig() // Find and read the config file
	if readerr != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file %s", readerr))
	}
}
