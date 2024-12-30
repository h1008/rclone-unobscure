package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/rclone/rclone/fs/config"
	"github.com/rclone/rclone/fs/config/configfile"
	"github.com/rclone/rclone/fs/config/obscure"
)

func main() {
	obscured_fields := []string{"pass", "password", "password2"}
	s := &configfile.Storage{}
	config.SetData(s)

	if len(os.Args) == 2 {
		config.SetConfigPath(os.Args[1])
	}

	fmt.Println("# Configuration:", config.GetConfigPath())
	for _, section := range config.LoadedData().GetSectionList() {
		fmt.Println("[" + section + "]")
		for _, f := range config.LoadedData().GetKeyList(section) {
			v, ok := config.LoadedData().GetValue(section, f)
			if !ok {
				log.Fatalf("Failed to find key %v in section %v", f, section)

			}
			if slices.Contains(obscured_fields, f) {
				v = obscure.MustReveal(v)
			}
			fmt.Println(f, "=", v)
		}
		fmt.Println()
	}
}
