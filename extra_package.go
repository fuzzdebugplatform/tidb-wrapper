package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"tidb-wrapper/resources"
)

var packageFiles = []string{"context.go", "route.go", "trace.go"}

func writeExtraPackage(dir string, blockMap map[string]*BlockInfo) error {
	packagePath := filepath.Join(dir, extraPackage)

	for _, file := range packageFiles {
		ioutil.WriteFile(filepath.Join(packagePath, file), getFileBytes(file), os.ModePerm)
	}

	// cover.go
	coverGoBytes := getCoverGoContent(blockMap)
	ioutil.WriteFile(filepath.Join(packagePath, "cover.go"), coverGoBytes, os.ModePerm)
	return nil
}

const resTmpl = "resource/%s"

func getFileBytes(name string) []byte {
	bs, err := resource.Asset(fmt.Sprintf(resTmpl, name))
	if err != nil {
		log.Printf("Asset Error: %v\n", err)
	}
	return bs
}

const coverGoTmpl = `package trace_util_0

func NewDigestCounter(sql string) *CoverCounter {
	// handle -> Cover
	pathCounter := &CoverCounter{
		sql:sql,
		m: map[string]*Cover{
{{range $file, $blocks := .}}
			"{{$file}}": {
				Pos: []int{ 
{{range $index, $line := $blocks.Pos}}{{ if iseven $index }}
{{printf "%6d" $line}}, {{ else }} {{printf "%6d" $line}}, //[{{div $index 2}}]{{ end }}{{end}}
				},
                Trace:make(map[[2]int]struct{}),
			},
{{end}}
		},
	}

	return pathCounter
}

func Register(sql string) string {
    openMu.RLock()
    if !open {
        openMu.RUnlock()
        return ""
    }
    openMu.RUnlock()
	myCounter := NewDigestCounter(sql)
	mu.Lock()
	defer mu.Unlock()
	dgest := digest(sql)
	counterSet[dgest] = myCounter
	return dgest
}`

var tmplFuncMap = template.FuncMap{
	"iseven": func(num int) string {
		if num % 2 == 0 {
			return "1"
		} else {
			return ""
		}
	},
	"div": func(a int, b int) int {
		return a / b
	},
}

var coverTmpl *template.Template

func init() {
	tmpl, err := template.New("test").Funcs(tmplFuncMap).Parse(coverGoTmpl)
	if err != nil {
		log.Fatalf("template shouldn't parse error %v\n", err)
	}
	coverTmpl = tmpl
}

func getCoverGoContent(blockMap map[string]*BlockInfo) []byte {
	buf := &bytes.Buffer{}
	err := coverTmpl.Execute(buf, blockMap)
	if err != nil {
		log.Fatalf("block info have error %v\n", err)
	}

	return buf.Bytes()
}



