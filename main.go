package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const extraPackage = "trace_util_0"

var dir string
var target string

var rootCmd = &cobra.Command{
	Use:   "tidb-wrapper",
	Short: "wrap tidb to statistic code path of one sql",
	Run: func(cmd *cobra.Command, args []string) {

		if isDirExist(target) {
			log.Fatalln("Fatal Error: dump directory already exist")
		}
		err := os.MkdirAll(target, os.ModePerm)
		if err != nil {
			log.Fatalf("Fatal Error: target dir %s create fail %v\n", target, err)
		}

		// init filter map
		filteredMap[filepath.Join(target, ".idea")] = true
		filteredMap[filepath.Join(target, ".git")] = true
		filteredMap[filepath.Join(target, extraPackage)] = true

		// step 1: copy all project file to another directory
		if dir == "." {
			filtered[target] = true
		}
		Copy(dir, target)

		// step 2: add counter in every block start
		bWalker := NewBlockWalker()
		walk(target, bWalker.writeCounter)

		// step 3: write extra package trace_util_0
		extraDir := filepath.Join(target, extraPackage)
		err = os.Mkdir(extraDir, os.ModePerm)
		if err != nil {
			log.Fatalf("%s dir create fail %v\n", extraPackage, err)
		}
		err = writeExtraPackage(target, bWalker.BlockMap)
		if err != nil {
			log.Fatalf("write extra package error %v\n", err)
		}

		// step 4: add import trace_util_0
		modName := getModName(dir)
		mWalker := &ModWalker{modName:modName, blockMap:bWalker.BlockMap}
		walk(target, mWalker.addImports)

		// step 5: go get -u github.com/spaolacci/murmur3 in target dir
		log.Println("go get -u github.com/spaolacci/murmur3")
		shellCmd := exec.Command("go", "get", "-u", "github.com/spaolacci/murmur3")
		shellCmd.Dir = target
		buf := &bytes.Buffer{}
		shellCmd.Stdout = buf
		err = shellCmd.Run()
		if err != nil {
			log.Fatalf("go get error %v\n%s\n", err, buf.String())
		}

		log.Printf("Info: code generation ok in %s\n", target)
	},
}

func getModName(dir string) string {
	modPath := filepath.Join(dir, "go.mod")
	f, err := os.Open(modPath)
	if err != nil {
		log.Fatalf("%s open error %v\n", modPath, err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("read line error %v\n", err)
	}

	return strings.TrimSpace(strings.Split(line, " ")[1])
}

var filteredMap = make(map[string]bool)

func init() {
	rootCmd.Flags().StringVarP(&dir, "dir", "D", ".", "directory to transfer")
	rootCmd.Flags().StringVarP(&target, "target", "T", "mywrapper", "target dir to rewrite source")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(rootCmd.UsageString())
		os.Exit(1)
	}
}

func walk(root string, walkFn filepath.WalkFunc)  {
	err := filepath.Walk(root, walkFn)
	if err != nil {
		log.Fatalf("Walk directory error %v\n", err)
	}
}

type BlockWalker struct {
	// relative path -> *BlockInfo
	BlockMap map[string]*BlockInfo
}

func NewBlockWalker() *BlockWalker {
	return &BlockWalker{BlockMap:make(map[string]*BlockInfo)}
}

func (b *BlockWalker) writeCounter(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if filteredMap[path] {
		return filepath.SkipDir
	}

	if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
		relativePath, err := filepath.Rel(target, path)
		if err != nil {
			log.Fatalln("should not run here")
		}

		src, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("%s read error\n", path)
		}

		rewrited, bInfo, err := addCounter(relativePath, info.Name(), src)
		if err == GoTestFile {
			return nil
		}
		if err != nil {
			log.Fatalf("rewrite error %v\n", err)
		}
		b.BlockMap[relativePath] = bInfo
		err = ioutil.WriteFile(path, rewrited, os.ModePerm)
		if err != nil {
			log.Fatalf("%s write error %v\n", path, err)
		}
	}

	return nil
}

func isDirExist(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

type ModWalker struct {
	modName string
	blockMap map[string]*BlockInfo
}


func (m *ModWalker) addImports(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if filteredMap[path] {
		return filepath.SkipDir
	}

	relativePath, err := filepath.Rel(target, path)
	if err != nil {
		log.Fatalf("filepaht rel should not err %v\n", err)
	}

	bInfo := m.blockMap[relativePath]

	if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") && bInfo != nil &&
		len(bInfo.Pos)>0 {
		src, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("%s read error\n", path)
		}

		rewrited, err := addImportsAndRegister(m.modName, relativePath, src)
		if err != nil {
			log.Fatalf("%s %v\n", path, err)
		}

		err = ioutil.WriteFile(path, rewrited, os.ModePerm)
		if err != nil {
			log.Fatalf("%s write error %v\n", path, err)
		}
	}

	return nil
}