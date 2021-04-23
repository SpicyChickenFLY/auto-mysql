package progress

import (
	"fmt"

	"github.com/SpicyChickenFLY/auto-mysql/utils/colorful"
)

// Check is a built-in func for checking error and output its result
func Check(info string, err error) error {
	if err != nil {
		fmt.Printf("[ %s ] %s\n",
			colorful.RenderStr("FAIL", "highlight", "black", "red"), info)
		fmt.Println(err)
	} else {
		fmt.Printf("[  %s  ] %s\n",
			colorful.RenderStr("OK", "highlight", "black", "green"), info)
	}
	return err
}
