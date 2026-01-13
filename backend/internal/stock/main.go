package main

import (
	"fmt"

	_ "github.com/leebrouse/emergency-material-system/common/config"
	"github.com/spf13/viper"
)

func main() {
	viper.GetString("")
	fmt.Println("hello world")
}
