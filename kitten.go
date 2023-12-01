package kitten

import (
	"github.com/shravanibhojane/cat"
)

func Bark() string {
	return "Meow!"
}

func Barks() string {
	return "Meow!Meow!Meow!"
}

func BigBark() string {
	return cat.WhenGrownUp(Bark())
}
func BigBarks() string {
	return cat.WhenGrownUp(Barks())
}
