package main

import (
	"fmt"
	"testing"
)

func Example_capitalize() {
	fmt.Println(capitalize("hello"))
	fmt.Println(capitalize(""))
	fmt.Println(capitalize("t"))
	// output:
	// Hello
	//
	// T
}

func Example_between() {
	fmt.Println(between("AB", "A", "B"))
	fmt.Println(between("co  ow", "c", "w"))
	fmt.Println(between("", "", ""))
	fmt.Println(between(" ,\\__/,", ",", ","))
	fmt.Println(between("AB", "B", "A"))
	fmt.Println(between("hhii", "h", "i"))
	// output:
	//
	// o  o
	//
	// \__/
	//
	// hi
}

func Example_betweenQuotes() {
	fmt.Println(betweenQuotes(`"hi"`))
	fmt.Println(betweenQuotes(`asdf "hi" asdf`))
	fmt.Println(betweenQuotes(`asdf ""hi"" asdf`))
	fmt.Println(betweenQuotes(`asdf "'hi'" asdf`))
	fmt.Println(betweenQuotes(`asdf '"hi asdf`))
	fmt.Println(betweenQuotes(`asdf '"hi'" asdf`))
	// output:
	// hi
	// hi
	// "hi"
	// 'hi'
	//
	// hi'
}

func Example_betweenQuotesOrAfterEquals() {
	fmt.Println(betweenQuotesOrAfterEquals("a = 123"))
	fmt.Println(betweenQuotesOrAfterEquals(`"asdf"`))
	fmt.Println(betweenQuotesOrAfterEquals(`x = "z = fnufnu"`))
	// output:
	// 123
	// asdf
	// z = fnufnu
}

func TestHas(t *testing.T) {
	if !has("a b c", "b") {
		t.Fail()
	}
	if !has("a b b c", "b") {
		t.Fail()
	}
	if has("", "a") {
		t.Fail()
	}
	if has(" ", "a") {
		t.Fail()
	}
	if has(" ", " ") {
		t.Fail()
	}
	if has("   ", " ") {
		t.Fail()
	}
	if !has(`"a"`, `"a"`) {
		t.Fail()
	}
}
