package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCoverGoContent(t *testing.T) {

	blockMap := map[string]*BlockInfo{
		"cmd/cmd.go": {
			Pos: []int{
				5, 8,
				14, 17,
				8, 10,
				10, 12,
			},
		},
		"data/res.go":{
			Pos: []int{
				1, 2,
				3, 4,
			},
		},
	}

	content := getCoverGoContent(blockMap)

	expected := `package trace_util_0

func NewDigestCounter(sql string) *CoverCounter {
	// handle -> Cover
	pathCounter := &CoverCounter{
		sql:sql,
		m: map[string]*Cover{

			"cmd/cmd.go": {
				Pos: []int{ 

     5,       8, //[0]
    14,      17, //[1]
     8,      10, //[2]
    10,      12, //[3]
				},
			},

			"data/res.go": {
				Pos: []int{ 

     1,       2, //[0]
     3,       4, //[1]
				},
			},

		},
	}

	return pathCounter
}

func Register(sql string) string {
	myCounter := NewDigestCounter(sql)
	mu.Lock()
	defer mu.Unlock()
	dgest := digest(sql)
	counterSet[dgest] = myCounter
	return dgest
}`
	//fmt.Println(string(content))
	assert.Equal(t, expected, string(content))
}
