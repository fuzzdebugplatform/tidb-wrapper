package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"strings"
)

func addImportsAndRegister(mod string, relPath string, content []byte) ([]byte, error) {
	fset, file, err := parse(content)
	if err != nil {
		return nil, err
	}

	addImport(mod, file)
	addRegister(file, relPath)

	buf := &bytes.Buffer{}
	err = format.Node(buf, fset, file)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type specialFunc struct {
	path   string
	fun    string
}

var specialVars = map[specialFunc]string{
	{"server/server_11111.go", "ProcessKey"}: "key",
	{"server/conn.go", "handleQuery"}: "sql",
}

func addRegister(file *ast.File, relPath string) {
	for _, decl := range file.Decls {
		if fDecl, ok := decl.(*ast.FuncDecl); ok {
			pathFun := specialFunc{relPath, fDecl.Name.Name}
			if regVar, ok := specialVars[pathFun]; ok {
				// add `trace_util_0.Register(regVar)`
				newBodyList := make([]ast.Stmt, 0, len(fDecl.Body.List))
				newBodyList = append(newBodyList, &ast.ExprStmt{X:
					&ast.CallExpr{
						Fun:&ast.SelectorExpr{
							X:&ast.Ident{
								Name:"trace_util_0",
							},
							Sel:&ast.Ident{
								Name:"Register",
							},
						},
						Args:[]ast.Expr{
							&ast.Ident{
								Name:regVar,
							},
						},
					}})
				newBodyList = append(newBodyList, fDecl.Body.List...)
				fDecl.Body.List = newBodyList
			}
		}
	}
}

func addImport(mod string, file *ast.File) {
	hasImports := false
	for _, decl := range file.Decls {
		if gDecl, ok := decl.(*ast.GenDecl); ok {
			if gDecl.Tok == token.IMPORT {
				hasImports = true
				gDecl.Specs = append(gDecl.Specs, &ast.ImportSpec{
					Path: &ast.BasicLit{Kind: token.STRING,
						Value: fmt.Sprintf(`"%s/%s"`, mod, extraPackage)},
				})
				break
			}
		}
	}

	if !hasImports {
		newDecl := make([]ast.Decl, 0)
		newDecl = append(newDecl, &ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{&ast.ImportSpec{
				Path: &ast.BasicLit{Kind: token.STRING,
					Value: fmt.Sprintf(`"%s/%s"`, mod, extraPackage)},
			},},
		})
		newDecl = append(newDecl, file.Decls...)
		file.Decls = newDecl
	}

}

type CallWalker struct {
	internalMods map[string]bool
}

var builtInFunc = map[string]bool{
	"append": true,
	"delete": true,
}

func (c *CallWalker) Visit(node ast.Node) (w ast.Visitor) {
	if cExpr, ok := node.(*ast.CallExpr); ok {
		if nameIdent, ok := cExpr.Fun.(*ast.Ident); ok {
			if !builtInFunc[nameIdent.Name] {
				cExpr.Args = append(cExpr.Args, &ast.Ident{Name: "digest_0"})
			}
		}

		if sExpr, ok := cExpr.Fun.(*ast.SelectorExpr); ok {
			root, err := findSelectorRoot(sExpr)
			if err != nil {
				return c
			}

			// internal mod call
			if c.internalMods[root] {
				cExpr.Args = append(cExpr.Args, &ast.Ident{Name: "digest_0"})
			}
		}
	}

	return c
}

var WithOutRoot = errors.New("selector without root")

func findSelectorRoot(parenSelector *ast.SelectorExpr) (string, error) {

	switch node := parenSelector.X.(type) {
	case *ast.Ident:
		return node.Name, nil
	case *ast.SelectorExpr:
		return findSelectorRoot(node)
	default:
		return "", WithOutRoot
	}
}

func addFuncParam(mod string, file *ast.File) {
	interalMods := make(map[string]bool)
	for i := 0; i < len(file.Imports); i++ {
		if strings.HasPrefix(file.Imports[i].Path.Value, mod) {
			imName := file.Imports[i].Path.Value[
				strings.LastIndex(file.Imports[i].Path.Value, "/")+1:]
			interalMods[imName] = true
		}
	}

	for _, decl := range file.Decls {
		if fDecl, ok := decl.(*ast.FuncDecl); ok {

			// add func param
			fDecl.Type.Params.List = append(fDecl.Type.Params.List, &ast.Field{
				Type:  &ast.Ident{Name: "string"},
				Names: []*ast.Ident{{Name: "digest_0"}},
			})

			// add call func param
			cWalker := &CallWalker{internalMods: interalMods}
			ast.Walk(cWalker, fDecl)

		}
	}
}
