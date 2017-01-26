package react

const testVersion = 4

type MyCell struct {
    value     int
    callbacks []CallbackHandle
}


// Define a function New() Reactor and the stuff that follows from
// implementing Reactor.
func New() Reactor {
	return MyCell{0, nil}
}


// implemention of Reactor
func (c *MyCell) CreateInput(value int) MyCell {
	c.value = value
	c.callbacks = nil
	return c
}

func (c *MyCell) CreateCompute1(d *MyCell, func(int) int) MyCell {
	d = New()
	
}

// implemention of Cell
func (c *MyCell) Value() int {
	return c.value
}

// implementation of InputCell
func (c *MyCell) SetValue(value int) {
	c.value = value
}




// A Reactor manages linked cells.
type Reactor interface {
	// CreateInput creates an input cell linked into the reactor
	// with the given initial value.
	CreateInput(int) InputCell

	// CreateCompute1 creates a compute cell which computes its value
	// based on one other cell. The compute function will only be called
	// if the value of the passed cell changes.
	CreateCompute1(Cell, func(int) int) ComputeCell

	// CreateCompute2 is like CreateCompute1, but depending on two cells.
	// The compute function will only be called if the value of any of the
	// passed cells changes.
	CreateCompute2(Cell, Cell, func(int, int) int) ComputeCell
}


// A ComputeCell always computes its value based on other cells and can
// call callbacks upon changes.
type ComputeCell interface {
	Cell

	// AddCallback adds a callback which will be called when the value changes.
	// It returns a callback handle which can be used to remove the callback.
	AddCallback(func(int)) CallbackHandle

	// RemoveCallback removes a previously added callback, if it exists.
	RemoveCallback(CallbackHandle)
}
