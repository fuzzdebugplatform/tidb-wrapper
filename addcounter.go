package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type BlockInfo struct {
	Pos []int
}

func (b *BlockInfo) Add(start, end int)  {
	b.Pos = append(b.Pos, start, end)
}

var GoTestFile = errors.New("go test file")

func addCounter(relativePath string, fileName string, src []byte) ([]byte, *BlockInfo, error) {
	name := fileName[:strings.IndexRune(fileName, '.')]
	if strings.HasSuffix(name, "_test") {
		return nil, nil, GoTestFile
	}

	return addCounterGo(relativePath, name, src)
}

func addCounterGo(relativePath string, gofileName string, src []byte) ([]byte, *BlockInfo, error) {
	fileSet, aFile, err := parse(src)
	if err != nil {
		return nil, nil, err
	}
	coverVar := "_" + gofileName + "_00000"

	counterHandle := fmt.Sprintf(`var %s = "%s"`, coverVar, relativePath)
	bInfo := &BlockInfo{}
	scanner := NewBlockScanner(fileSet, src, func(startLine, endLine, blockNum int) (insert string) {
		bInfo.Add(startLine, endLine)
		return fmt.Sprintf("trace_util_0.Count(%s, %d);", coverVar, blockNum)
	})

	ast.Walk(scanner, aFile)

	rewriteBuf := bytes.NewBuffer(scanner.Res())
	rewriteBuf.WriteRune('\n')
	rewriteBuf.WriteString(counterHandle)

	return rewriteBuf.Bytes(), bInfo, nil
}

func parse(content []byte) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	aFile, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	return fset, aFile, nil
}
