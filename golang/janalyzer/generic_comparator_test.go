package janalyzer

import "testing"

func TestCompareEqualTo(t *testing.T) {
	var oper Comparator
	var bVal bool
	oper = &CompareEqualTo{}

	bVal = oper.CompareValue(true, false)
	if bVal {
		t.Errorf("CompareEqualTo(bool) got %v, want %v", bVal, false)
	}

	bVal = oper.CompareValue("Gowdas", "Gowdas")
	if !bVal {
		t.Errorf("CompareEqualTo(string) got %v, want %v", bVal, true)
	}

	bVal = oper.CompareValue(10.456, 10.456)
	if !bVal {
		t.Errorf("CompareEqualTo(float64) got %v, want %v", bVal, true)
	}
}

func TestCompareGreaterThanEqual(t *testing.T) {
	var oper Comparator
	var bVal bool
	oper = &CompareGreaterThanEqual{}

	bVal = oper.CompareValue("gowdas", "Gowdas")
	if !bVal {
		t.Errorf("CompareGreaterThanEqual(string) got %v, want %v", bVal, true)
	}

	bVal = oper.CompareValue(10.5, 10.59)
	if bVal {
		t.Errorf("CompareGreaterThanEqual(float64) got %v, want %v", bVal, false)
	}

	bVal = oper.CompareValue(true, false)
	if bVal {
		t.Errorf("CompareGreaterThanEqual(bool) got %v, want %v", bVal, false)
	}
}

func TestCompareGreaterThan(t *testing.T) {
	var oper Comparator
	var bVal bool
	oper = &CompareGreaterThan{}

	bVal = oper.CompareValue("gowdas", "GOWDAS")
	if !bVal {
		t.Errorf("CompareGreaterThan(string) got %v, want %v", bVal, true)
	}

	bVal = oper.CompareValue(12345.12345, 12345.123)
	if !bVal {
		t.Errorf("CompareGreaterThan(float64) got %v, want %v", bVal, true)
	}

	bVal = oper.CompareValue(true, false)
	if bVal {
		t.Errorf("CompareGreaterThan(bool) got %v, want %v", bVal, false)
	}
}

func TestCompareLessThanEqual(t *testing.T) {
	var oper Comparator
	var bVal bool
	oper = &CompareLessThanEqual{}

	bVal = oper.CompareValue("GOWDAS", "gowdas")
	if !bVal {
		t.Errorf("CompareLessThanEqual(string) got %v, want %v", bVal, true)
	}

	bVal = oper.CompareValue(12345.0, 12345.01)
	if !bVal {
		t.Errorf("CompareLessThanEqual(float64) got %v, want %v", bVal, true)
	}

	bVal = oper.CompareValue(false, false)
	if bVal {
		t.Errorf("CompareLessThanEqual(bool) got %v, want %v", bVal, false)
	}
}

func TestCompareLessThan(t *testing.T) {
	var oper Comparator
	var bVal bool
	oper = &CompareLessThan{}

	bVal = oper.CompareValue("GOWDA", "HYMAN")
	if !bVal {
		t.Errorf("CompareLessThan(string) got %v, want %v", bVal, true)
	}

	bVal = oper.CompareValue(1.09, 1.001)
	if bVal {
		t.Errorf("CompareLessThanEqual(float64) got %v, want %v", bVal, false)
	}

	bVal = oper.CompareValue(true, true)
	if bVal {
		t.Errorf("CompareLessThan(bool) got %v, want %v", bVal, false)
	}
}

func TestCompareToString(t *testing.T) {
	var oper Comparator
	oper = &CompareEqualTo{}

	if oper.ToString() != "==" {
		t.Errorf("CompareEqualsToString : wants : %v, got %v", "==", oper.ToString())
	}

	oper = &CompareGreaterThanEqual{}
	if oper.ToString() != ">=" {
		t.Errorf("CompareGreaterThanEqualToString : wants : %v, got %v", ">=", oper.ToString())
	}

	oper = &CompareGreaterThan{}
	if oper.ToString() != ">" {
		t.Errorf("CompareGreaterThanToString : wants : %v, got : %v", ">", oper.ToString())
	}

	oper = &CompareLessThanEqual{}
	if oper.ToString() != "<=" {
		t.Errorf("CompareLessThanEqualToString : wants : %v, got : %v", "<=", oper.ToString())
	}

	oper = &CompareLessThan{}
	if oper.ToString() != "<" {
		t.Errorf("CompareLessThanEqual : wants : %v, got : %v", "<", oper.ToString())
	}
}
