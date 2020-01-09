package main

import "testing"

func TestTrianglegetArea(t *testing.T) {
	tri := triangle{6, 3}
	if tri.getArea() != 9 {
		t.Errorf("Expected area of 9, but got %+v", tri.getArea())
	}

	tri = triangle{3, 6}
	if tri.getArea() != 9 {
		t.Errorf("Expected area of 9, but got %+v", tri.getArea())
	}
}

func TestSquareArea(t *testing.T) {
	s := square{6}
	if s.getArea() != 36 {
		t.Errorf("Expected area of 36, but got %+v", s.getArea())
	}
}
