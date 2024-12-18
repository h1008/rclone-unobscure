package main

import (
	"fmt"

	"github.com/rclone/rclone/fs/config"
	"github.com/rclone/rclone/fs/config/configfile"
	"github.com/rclone/rclone/fs/config/obscure"
)

func main() {
	obscured_fields := []string{"pass", "password", "password2"}
	s := &configfile.Storage{}
	config.SetData(s)
	fmt.Println("Loading configuration:", config.GetConfigPath())
	for _, section := range config.FileSections() {
		fmt.Println("[" + section + "]")
		for _, f := range obscured_fields {
			pass, ok := config.FileGetValue(section, f)
			if ok {
				fmt.Println(f, "=", obscure.MustReveal(pass))
			}
		}
		fmt.Println()
	}
}
