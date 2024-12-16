package main

import (
	"flag"
	"fmt"
	"os"
)

func TestingFlags() {
	flags := flag.NewFlagSet("MyFlags", flag.ExitOnError)
	// Лучше писать через NewFlagSet, чтобы собирать сразу все флаги в одном месте
	var output bool
	var configPath string
	flags.BoolVar(&output, "output", false, "Нужен ли вывод программы?")
	flags.StringVar(&configPath, "config", "config.cfg", "Конфиг для программы")
	// Теперь, если не указать -output или --output, по дефолту будет false.
	// Если же написать -output=true или --output=true, то будет true
	// Если передать флаг -h, --h, -help, --help, то выведится пояснение для всех созданных флагов. В данном случае будет
	// -output Нужен ли вывод программы?

	err := flags.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	fmt.Println(output, configPath)
}

func main() {
	for _, value := range os.Args {
		fmt.Println(value)
	}
	fmt.Println()
	fmt.Println(os.Getenv("PWD"))
	fmt.Println()

	TestingFlags()
}
