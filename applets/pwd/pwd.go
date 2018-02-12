package pwd

import (
	"os"
)

func Pwd(call []string) error {

	pwd, _ := os.Getwd()
	println(pwd)
	return nil
}
