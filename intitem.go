package gopareto

import "fmt"

type WithIsEmpty struct {
	isEmpty bool
}

func (i *WithIsEmpty) IsEmpty() bool {
	return i.isEmpty
}

type aInt struct {
	WithIsEmpty
	value int
}

func Int(v int) Item {
	return &aInt{value: v}
}

func (a *aInt) IsGreaterThan(b Item) (bool, error) {
	if b.(*aInt) == nil {
		return false, fmt.Errorf(
			"tried to compare aInt %v to non aInt b %v",
			a, b)
	}
	if a.IsEmpty() || b.IsEmpty() {
		return false, nil
	}
	return a.value > b.(*aInt).value, nil
}

var (
	emptyInt = aInt{WithIsEmpty: WithIsEmpty{isEmpty: true}}
)

func (a *aInt) Empty() Item {
	v := emptyInt // should be copy by value
	return &v
}

func EmptyInt() Item {
	return Int(0).Empty()
}

func (a *aInt) String() string {
	if a.IsEmpty() {
		return "X"
	}

	return fmt.Sprintf("%d", a.value)
}
