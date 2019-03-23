package main

import (
	"github.com/hikaru7719/jsontagchecker"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(jsontagchecker.Analyzer)
}
