package env

import "os"

// Remember to convert to SetKey
// SetKey takes a string and combines with the KEY environment variable as an added layer of security
func setKey() string {
	os.Setenv("KEY", "some-key")
	return os.Getenv("KEY")

}
