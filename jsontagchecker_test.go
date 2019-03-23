package jsontagchecker_test

import (
	"github.com/hikaru7719/jsontagchecker"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, jsontagchecker.Analyzer, "a")
}
