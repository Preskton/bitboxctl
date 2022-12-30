/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/preskton/bitboxctl/cmd"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
	host "periph.io/x/host/v3"
)

var banner = `
█████      ███   █████    █████                                    █████    ████ 
░░███      ░░░   ░░███    ░░███                                    ░░███    ░░███ 
 ░███████  ████  ███████   ░███████   ██████  █████ █████  ██████  ███████   ░███ 
 ░███░░███░░███ ░░░███░    ░███░░███ ███░░███░░███ ░░███  ███░░███░░░███░    ░███ 
 ░███ ░███ ░███   ░███     ░███ ░███░███ ░███ ░░░█████░  ░███ ░░░   ░███     ░███ 
 ░███ ░███ ░███   ░███ ███ ░███ ░███░███ ░███  ███░░░███ ░███  ███  ░███ ███ ░███ 
 ████████  █████  ░░█████  ████████ ░░██████  █████ █████░░██████   ░░█████  █████
░░░░░░░░  ░░░░░    ░░░░░  ░░░░░░░░   ░░░░░░  ░░░░░ ░░░░░  ░░░░░░     ░░░░░  ░░░░░ 
`

func main() {
	log.Debug("Setting viper global defaults")
	// TODO figure out how to hook viper to logrus
	viper.SetConfigName("bitboxctl")
	viper.AddConfigPath("config")
	viper.AddConfigPath("$HOME/.bitboxctl")

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.WithField("filename", e.Name).Infof("Config file reloaded from disk")
	})
	viper.WatchConfig()

	hideBanner := viper.GetBool("hideBanner")

	if !hideBanner {
		fmt.Print(banner)
		fmt.Print("\n")
	}

	log.Debug("Initializing host...")
	host.Init()

	cmd.Execute()
}
