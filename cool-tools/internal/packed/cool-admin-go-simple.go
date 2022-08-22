package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}