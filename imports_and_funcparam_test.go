package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddImport(t *testing.T) {
	testCode := `package main

func main() {
	fmt.Printf("Hello, Golang\n")
}
`
	rewrited, err := addImportsAndRegister("ttt", []byte(testCode))
	assert.Equal(t, nil, err)

	fmt.Println(string(rewrited))
}
