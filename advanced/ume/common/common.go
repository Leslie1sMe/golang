package common

import (
	"fmt"
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		os.Exit(1)
	}
}
