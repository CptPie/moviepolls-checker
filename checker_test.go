package main

import "testing"

func Test_CheckInputString_canBeEmpty_isEmpty(t *testing.T) {
	str := ""
	canBeEmpty := true
	minLength := -1
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canBeEmpty_isNotEmpty(t *testing.T) {
	str := "foobar"
	canBeEmpty := true
	minLength := -1
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canBeEmpty_minLength_shorter(t *testing.T) {
	str := "foobar"
	canBeEmpty := true
	minLength := 10
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_CheckInputString_canBeEmpty_minLength_longer(t *testing.T) {
	str := "foobar"
	canBeEmpty := true
	minLength := 5
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canBeEmpty_minLength_equal(t *testing.T) {
	str := "foobar"
	canBeEmpty := true
	minLength := 6
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canBeEmpty_maxLength_shorter(t *testing.T) {
	str := "foobar"
	canBeEmpty := true
	minLength := -1
	maxLength := 10

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canBeEmpty_maxLength_longer(t *testing.T) {
	str := "foobar"
	canBeEmpty := true
	minLength := -1
	maxLength := 5

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_CheckInputString_canBeEmpty_maxLength_equal(t *testing.T) {
	str := "foobar"
	canBeEmpty := true
	minLength := -1
	maxLength := 6

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canNotBeEmpty_isEmpty(t *testing.T) {
	str := ""
	canBeEmpty := false
	minLength := -1
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_CheckInputString_canNotBeEmpty_isNotEmpty(t *testing.T) {
	str := "foobar"
	canBeEmpty := false
	minLength := -1
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canNotBeEmpty_minLength_shorter(t *testing.T) {
	str := "foobar"
	canBeEmpty := false
	minLength := 10
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_CheckInputString_canNotBeEmpty_minLength_longer(t *testing.T) {
	str := "foobar"
	canBeEmpty := false
	minLength := 5
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canNotBeEmpty_minLength_equal(t *testing.T) {
	str := "foobar"
	canBeEmpty := false
	minLength := 6
	maxLength := -1

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canNotBeEmpty_maxLength_shorter(t *testing.T) {
	str := "foobar"
	canBeEmpty := false
	minLength := -1
	maxLength := 10

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputString_canNotBeEmpty_maxLength_longer(t *testing.T) {
	str := "foobar"
	canBeEmpty := false
	minLength := -1
	maxLength := 5

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_CheckInputString_canNotBeEmpty_maxLength_equal(t *testing.T) {
	str := "foobar"
	canBeEmpty := false
	minLength := -1
	maxLength := 6

	err := CheckInputString(str, canBeEmpty, minLength, maxLength)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputInt_canBeZero_isZero(t *testing.T) {
	integer := 0
	canBeZero := true
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputInt_canBeZero_isNotZero(t *testing.T) {
	integer := 2
	canBeZero := true
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputInt_canBeZero_minValue_lower(t *testing.T) {
	integer := -20
	canBeZero := true
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}
func Test_CheckInputInt_canBeZero_minValue_higher(t *testing.T) {
	integer := 2
	canBeZero := true
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}
func Test_CheckInputInt_canBeZero_minValue_equal(t *testing.T) {
	integer := -10
	canBeZero := true
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputInt_canBeZero_maxValue_lower(t *testing.T) {
	integer := 5
	canBeZero := true
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}
func Test_CheckInputInt_canBeZero_maxValue_higher(t *testing.T) {
	integer := 15
	canBeZero := true
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}
func Test_CheckInputInt_canBeZero_maxValue_equal(t *testing.T) {
	integer := 10
	canBeZero := true
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputInt_canNotBeZero_isZero(t *testing.T) {
	integer := 0
	canBeZero := false
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_CheckInputInt_canNotBeZero_isNotZero(t *testing.T) {
	integer := 2
	canBeZero := false
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputInt_canNotBeZero_minValue_lower(t *testing.T) {
	integer := -20
	canBeZero := false
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_CheckInputInt_canNotBeZero_minValue_higher(t *testing.T) {
	integer := 2
	canBeZero := false
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputInt_canNotBeZero_minValue_equal(t *testing.T) {
	integer := -10
	canBeZero := false
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputInt_canNotBeZero_maxValue_lower(t *testing.T) {
	integer := 5
	canBeZero := false
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}

func Test_CheckInputInt_canNotBeZero_maxValue_higher(t *testing.T) {
	integer := 20
	canBeZero := false
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_CheckInputInt_canNotBeZero_maxValue_equal(t *testing.T) {
	integer := 10
	canBeZero := false
	minValue := -10
	maxValue := 10

	err := CheckInputInt(integer, canBeZero, minValue, maxValue)
	if err != nil {
		t.Errorf("Expected no error, got error %v", err.Error())
	}
}
