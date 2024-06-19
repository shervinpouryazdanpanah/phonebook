package config

import "phonebook/config"

func Get() config.Config {
	return configuration
}
