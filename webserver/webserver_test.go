package webserver

import "testing"

func TestStartingHourIsInvalidIfLessThenNine(t *testing.T){
	result := IsBetween9and5(8)

	if result == true {
		t.Errorf("IsBetween9and5(8)= %t; want false", result)
	}
}


func TestStartingHourIsInvalidIfGreaterThen16(t *testing.T){
	result := IsBetween9and5(16)

	if result == true {
		t.Errorf("IsBetween9and5(16)= %t; want false", result)
	}
}
func TestStartingHourIsValidIfGreaterThen9(t *testing.T){
	result := IsBetween9and5(10)

	if result == false {
		t.Errorf("IsBetween9and5(10)= %t; want true", result)
	}
}

func TestStartingHourIsValidIfLessThen16(t *testing.T){
	result := IsBetween9and5(15)

	if result == false {
		t.Errorf("IsBetween9and5(15)= %t; want true", result)
	}
}
func TestStartingHourIsValidIfEqualTwoNine(t *testing.T){
	result := IsBetween9and5(9)

	if result == false {
		t.Errorf("IsBetween9and5(8)= %t; want true", result)
	}
}
