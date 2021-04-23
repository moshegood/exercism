package react

type callback struct {
	function func(int)
	canceled bool
}

func (cb *callback) Cancel() {
	cb.canceled = true
}
func (cb *callback) Run(input int) {
	if !cb.canceled {
		cb.function(input)
	}
}

type cell struct {
	value     int
	callbacks []*callback
	computer  func() int
	// for optimization we track old values
	getInputs func() []int
	oldValues []int
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(input int) {
	if input == c.value {
		return
	}
	c.value = input
	for _, cb := range c.callbacks {
		cb.Run(input)
	}
}

func (c *cell) AddCallback(function func(int)) Canceler {
	cb := &callback{function: function}
	c.callbacks = append(c.callbacks, cb)
	return cb
}

type reactor struct {
	computeCells []*cell
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) addComputeCell(computeValue func() int, getInputs func() []int) *cell {
	newCell := &cell{
		computer:  computeValue,
		value:     computeValue(),
		getInputs: getInputs,
		oldValues: getInputs(),
	}
	r.computeCells = append(r.computeCells, newCell)
	return newCell
}

func (r *reactor) CreateInput(value int) InputCell {
	newCell := &cell{value: value}
	newCell.AddCallback(func(int) { r.wasUpdated() })
	return newCell
}

func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	computeValue := func() int { return f(c.Value()) }
	getInputs := func() []int { return []int{c.Value()} }
	return r.addComputeCell(computeValue, getInputs)
}

func (r *reactor) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	computeValue := func() int { return f(c1.Value(), c2.Value()) }
	getInputs := func() []int { return []int{c1.Value(), c2.Value()} }
	return r.addComputeCell(computeValue, getInputs)
}

func (r *reactor) wasUpdated() {
	for _, cell := range r.computeCells {
		// Optimization - don't run computer() call on unchanged inputs
		upstreamChanged := false
		getInputs := cell.getInputs()
		for i, value := range getInputs {
			if cell.oldValues[i] != value {
				upstreamChanged = true
			}
			cell.oldValues[i] = value
		}

		if upstreamChanged {
			cell.SetValue(cell.computer())
		}
	}
}

var _ Reactor = &reactor{}
