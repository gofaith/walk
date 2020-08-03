package walk

import (
	"github.com/gen2brain/beeep"
)

func ShowToast(title, content string) {
	beeep.Notify(title, content, "")
}
