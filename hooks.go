package golog

import (
	"fmt"
	"time"
)

func timeHook(c *Core) interface{} {
	h, m, s := time.Now().Local().Clock()

	return fmt.Sprintf("%v:%v:%v", h, m, s)
}
