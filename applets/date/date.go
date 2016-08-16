package date

import "time"
import "fmt"

func Date(call []string) error {
	println(fmt.Sprintf("%s", time.Now()))
	return nil
}
