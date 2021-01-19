package main

//Incrementator interface for increment
type Incrementator interface {
	GetNumber() int
	IncrementNumber()
	SetMaximumValue(maximumValue int)
}

//Incrementor struct implementing the Incrementator interface
type Incrementor struct {
	curNum   int //current number
	maxValue int //maximum value
}

//GetNumber return current number from Incrementator struct
func (inc *Incrementor) GetNumber() int {
	return inc.curNum
}

//IncrementNumber increment by 1 current number from Incrementator struct
func (inc *Incrementor) IncrementNumber() {
	if inc.curNum == inc.maxValue {
		inc.curNum = 0
		return
	}
	inc.curNum++
}

//SetMaximumValue sets the maximum value specified in the variable maximumValue
func (inc *Incrementor) SetMaximumValue(maximumValue int) {
	inc.maxValue = maximumValue
}
