package echo

import "fmt"

// Service returns a simple message
func Service() string {
	return fmt.Sprintf("Howdy: %d", 10)
}
