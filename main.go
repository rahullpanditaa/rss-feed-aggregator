package main

import (
	"fmt"
	"os"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/config"
)

func main() {
	contents, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(contents)

	// contents holds json file contents
	err = contents.SetUser("rahul")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	c, _ := config.Read()
	fmt.Println(c)
}
