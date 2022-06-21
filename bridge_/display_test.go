package bridge_

import "testing"

func TestDisplay(t *testing.T) {
	// count display
	mathDisplay := NewDisplay(NewStringDisplayImpl("Hello Math"))
	engDisplay := NewDisplay(NewStringDisplayImpl("Hello English"))
	phyDisplay := NewMultipleDisplay(NewStringDisplayImpl("Hello Physics"))
	bioDisplay := NewRandomDisplay(NewStringDisplayImpl("Hello Biology"))

	mathDisplay.Display()
	engDisplay.Display()
	phyDisplay.CountDisplay(3)
	bioDisplay.RandomDisplay(10)

	// stair display
	starDisplay := NewStairDisplay(NewCharDisplayImpl('<', '*', '>'))
	hashDisplay := NewStairDisplay(NewCharDisplayImpl('|', '#', '-'))
	starDisplay.StairDisplay(5, 1)
	hashDisplay.StairDisplay(10, 2)
}
