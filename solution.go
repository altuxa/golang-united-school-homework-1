package solution

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	hello := emoji.Sprint("Hello :world_map:!")
	return hello
}
