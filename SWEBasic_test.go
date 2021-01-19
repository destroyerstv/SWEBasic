package main

import (
	"testing"
)

func TestInterface(t *testing.T) {
	inc := &Incrementor{
		curNum:   5,
		maxValue: 10,
	}

	if _, ok := interface{}(inc).(Incrementator); !ok {
		t.Errorf("Incrementor struct not implementation Incrementator interface")
	}
}

func TestGetNumber(t *testing.T) {
	inc := &Incrementor{
		curNum:   5,
		maxValue: 10,
	}

	result := inc.GetNumber()

	if result != inc.curNum {
		t.Errorf("wrong result, expected %#v, got %#v", inc.curNum, result)
	}
}

//IncrementNumber increment by 1 current number from Incrementator struct
func TestIncrementNumber(t *testing.T) {
	inc := &Incrementor{
		curNum:   9,
		maxValue: 10,
	}

	inc.IncrementNumber()
	result := inc.GetNumber()

	if result != 10 {
		t.Errorf("wrong result, expected %#v, got %#v", 10, result)
	}

	inc.IncrementNumber()
	result = inc.GetNumber()

	if result != 0 {
		t.Errorf("wrong result, expected %#v, got %#v", 0, result)
	}

}

//SetMaximumValue sets the maximum value specified in the variable maximumValue
func TestSetMaximumValue(t *testing.T) {

	var mVal int = 11

	inc := &Incrementor{
		curNum:   10,
		maxValue: 10,
	}

	inc.SetMaximumValue(mVal)

	inc.IncrementNumber()
	result := inc.GetNumber()

	if result != mVal {
		t.Errorf("wrong result, expected %#v, got %#v", mVal, result)
	}

	inc.IncrementNumber()
	result = inc.GetNumber()

	if result != 0 {
		t.Errorf("wrong result, expected %#v, got %#v", 0, result)
	}
}
