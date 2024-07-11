package util

import (
	"log"
)

func LogError(err error) {
	red := "\033[31m"
	reset := "\033[0m"
	log.Printf(red+"error: %v"+reset, err)
}
