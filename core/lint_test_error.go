package core

import "fmt"

// intentional lint error to verify CI catches failures — remove after confirming
func lintError() string {
	return fmt.Sprintf("%s", 42)
}
