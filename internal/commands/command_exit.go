package commands

import "os"

func commandExit(cfg *Config, param string) error {
	os.Exit(0)
	return nil
}
