package main

import (
	"os"
	"reflect"
	"sort"
	"testing"
)

type Test struct {
	FilePath       string
	ExpectedOutput []string
}

var tests = []Test{
	Test{
		FilePath: "test/test1.txt",
		ExpectedOutput: []string{
			"Lisa: -93",
			"Quincy: error",
			"Tom: 500",
		},
	},
	Test{
		FilePath: "test/test2.txt",
		ExpectedOutput: []string{
			"Chloe: 1694",
			"Dan: error",
			"Neil: 500",
			"Rainna: error",
			"Sarah: 99",
		},
	},
	Test{
		FilePath: "test/test3.txt",
		ExpectedOutput: []string{
			"Bacon: -40",
			"Cheese: -240",
			"Hollandaise: error",
			"Macaroni: 0",
		},
	},
}

func TestBasic(t *testing.T) {
	for _, test := range tests {
		f, err := os.Open(test.FilePath)
		if err != nil {
			t.Errorf("Failed to open input file", err)
		}
		defer f.Close()

		output := processInput(f)
		if output == nil {
			t.Fatalf("Error processing input")
		}
		sort.Strings(output)
		if !reflect.DeepEqual(output, test.ExpectedOutput) {
			t.Errorf("Expected %s but got %s", test.ExpectedOutput, output)
		}
	}
}
