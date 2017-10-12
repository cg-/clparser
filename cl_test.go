package clparser

import (
	"fmt"
	"testing"
)

var clp *CraigsListParser

func init() {
	clp = &CraigsListParser{
		city:     "Testville",
		category: "Testcat",
		terms:    []string{"term1", "term 2"},
	}
}

func TestNewCraigsListParser(t *testing.T) {
	clp := NewCraigsListParser()
	if clp.category != "" || clp.city != "" || clp.terms != nil {
		t.Error("Variables weren't nil!")
	}
}

func TestGet(t *testing.T) {
	clp := NewCraigsListParser()
	clp.SetCategory("laf")
	clp.SetCity("seattle")
	clp.AddTerms("car", "mine")
	out, err := clp.Get()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	fmt.Println(out)
	t.Skip("Get test not implemented.")
}

func TestSetCity(t *testing.T) {
	clp := NewCraigsListParser()
	before := clp.city
	testVar := "TEST TEST TEST"
	if before == testVar {
		testVar = "TEST TEST"
	}
	clp.SetCity(testVar)
	if clp.city != testVar {
		t.Log("City wasn't set properly.")
		t.Fail()
	}
}

func TestSetCategory(t *testing.T) {
	clp := NewCraigsListParser()
	before := clp.category
	testVar := "TEST TEST TEST"
	if before == testVar {
		testVar = "TEST TEST"
	}
	clp.SetCategory(testVar)
	if clp.category != testVar {
		t.Log("Category wasn't set properly.")
		t.Fail()
	}
}

func TestSetTerms(t *testing.T) {
	tests := make([][]string, 0)
	tests = append(tests, []string{"test"})
	tests = append(tests, []string{"test1", "test2"})
	for i := range tests {
		testSet := tests[i]
		clp := NewCraigsListParser()
		if clp.terms != nil {
			clp.terms = nil
		}

		clp.SetTerms(testSet)
		checks := make([]bool, len(testSet))
		for i := range testSet {
			checks[i] = false
		}

		for i := range clp.terms {
			for j := range testSet {
				if clp.terms[i] == testSet[j] {
					checks[j] = true
				}
			}
		}

		for i := range checks {
			if checks[i] == false {
				t.Fail()
			}
		}
	}
}

func TestAddTerms(t *testing.T) {

	// test add 1
	clp := NewCraigsListParser()
	if clp.terms != nil {
		t.Log("Didn't initialize nil terms")
		t.Fail()
	}

	clp.AddTerms("asd")

	found := false
	for i := range clp.terms {
		if clp.terms[i] == "asd" {
			found = true
		}
	}
	if !found {
		t.Fail()
	}

	// test add 2
	clp = NewCraigsListParser()
	if clp.terms != nil {
		t.Log("Didn't initialize nil terms")
		t.Fail()
	}

	clp.AddTerms("asdasdasd", "asd 11")

	found1 := false
	found2 := false
	for i := range clp.terms {
		if clp.terms[i] == "asdasdasd" {
			found1 = true
		}
		if clp.terms[i] == "asd 11" {
			found2 = true
		}
	}
	if !found1 || !found2 {
		t.Fail()
	}
}

func TestDeleteTerms(t *testing.T) {
	// delete first
	clp := NewCraigsListParser()
	if clp.terms != nil {
		t.Log("Failed delete first.")
		t.Fail()
	}

	clp.AddTerms("test1", "test2", "test3")
	clp.DeleteTerms("test1")

	found1 := false
	found2 := false
	found3 := false
	for i := range clp.terms {
		if clp.terms[i] == "test1" {
			found1 = true
		}
		if clp.terms[i] == "test2" {
			found2 = true
		}
		if clp.terms[i] == "test3" {
			found3 = true
		}
	}

	if found1 == true || found2 == false || found3 == false {
		t.Log("Failed delete first.")
		t.Fail()
	}

	// delete second
	clp = NewCraigsListParser()
	if clp.terms != nil {
		t.Log("Failed delete second.")
		t.Fail()
	}

	clp.AddTerms("test1", "test2", "test3")
	clp.DeleteTerms("test2")

	found1 = false
	found2 = false
	found3 = false
	for i := range clp.terms {
		if clp.terms[i] == "test1" {
			found1 = true
		}
		if clp.terms[i] == "test2" {
			found2 = true
		}
		if clp.terms[i] == "test3" {
			found3 = true
		}
	}

	if found1 == false || found2 == true || found3 == false {
		t.Log("Failed delete second.")
		t.Fail()
	}

	// delete last
	clp = NewCraigsListParser()
	if clp.terms != nil {
		t.Log("Failed delete third.")
		t.Fail()
	}

	clp.AddTerms("test1", "test2", "test3")
	clp.DeleteTerms("test3")

	found1 = false
	found2 = false
	found3 = false
	for i := range clp.terms {
		if clp.terms[i] == "test1" {
			found1 = true
		}
		if clp.terms[i] == "test2" {
			found2 = true
		}
		if clp.terms[i] == "test3" {
			found3 = true
		}
	}

	if found1 == false || found2 == false || found3 == true {
		t.Log("Failed delete third.")
		t.Fail()
	}

	// delete two
	clp = NewCraigsListParser()
	if clp.terms != nil {
		t.Log("Failed delete two.")
		t.Fail()
	}

	clp.AddTerms("test1", "test2", "test3")
	clp.DeleteTerms("test2", "test3")

	found1 = false
	found2 = false
	found3 = false
	for i := range clp.terms {
		if clp.terms[i] == "test1" {
			found1 = true
		}
		if clp.terms[i] == "test2" {
			found2 = true
		}
		if clp.terms[i] == "test3" {
			found3 = true
		}
	}

	if found1 == false || found2 == true || found3 == true {
		t.Log("Failed delete two.")
		t.Fail()
	}

	// delete all
	clp = NewCraigsListParser()
	if clp.terms != nil {
		t.Log("Failed delete all.")
		t.Fail()
	}

	clp.AddTerms("test1", "test2", "test3")
	clp.DeleteTerms("test1", "test2", "test3")

	found1 = false
	found2 = false
	found3 = false
	for i := range clp.terms {
		if clp.terms[i] == "test1" {
			found1 = true
		}
		if clp.terms[i] == "test2" {
			found2 = true
		}
		if clp.terms[i] == "test3" {
			found3 = true
		}
	}

	if found1 == true || found2 == true || found3 == true {
		t.Log("Failed delete all.")
		t.Fail()
	}
}
