package main

import "fmt"

func CheckInputString(str string, canBeEmpty bool, minLength int, maxLength int) error {
	if !canBeEmpty {
		if str == "" {
			return fmt.Errorf("String cannot be empty")
		}
	}
	if len(str) < minLength && minLength != -1 {
		return fmt.Errorf("String is shorter than minLength: Expected: %d Has: %d", minLength, len(str))
	}
	if len(str) > maxLength && maxLength != -1 {
		return fmt.Errorf("String is longer than maxLength: Expected: %d Has: %d", maxLength, len(str))
	}
	return nil
}

func CheckInputInt(integer int, canBeZero bool, minValue int, maxValue int) error {
	if !canBeZero {
		if integer == 0 {
			return fmt.Errorf("Integer cannot be zero")
		}
	}
	if integer < minValue {
		return fmt.Errorf("Integer is lower than minValue: Expected: %d Has: %d", minValue, integer)
	}
	if integer > maxValue {
		return fmt.Errorf("Integer is longer than maxValue: Expected: %d Has: %d", maxValue, integer)
	}
	return nil
}
