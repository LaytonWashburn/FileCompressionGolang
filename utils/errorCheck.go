package utils
import (		
	"log"
	// "fmt"
	// "os"
)

// checks for error
func ErrorChecker(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}