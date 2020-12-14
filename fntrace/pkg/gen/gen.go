package gen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

func hasFuncDecl(f *ast.File) bool {
	for _, decl := range f.Decls {
		if _, ok := decl.(*ast.FuncDecl); ok {
			return true
		}
	}

	return false
}

func Rewrite(filename string) ([]byte, error) {
	fset := token.NewFileSet()
	oldAST, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s: %w", filename, err)
	}
	//fmt.Printf("%#v\n", *oldAST)

	if !hasFuncDecl(oldAST) {
		return nil, nil
	}

	// add import declaration
	astutil.AddImport(fset, oldAST, "github.com/bingoohuang/fntrace")
	//fmt.Printf("added=%#v\n", added)

	// inject code into each function declaration
	addDeferTraceIntoFuncDecls(oldAST)

	buf := &bytes.Buffer{}
	err = format.Node(buf, fset, oldAST)
	if err != nil {
		return nil, fmt.Errorf("error formatting new code: %w", err)
	}
	return buf.Bytes(), nil
}

func addDeferTraceIntoFuncDecls(f *ast.File) {
	for _, decl := range f.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			// inject code to fd
			addDeferStmt(fd)
		}
	}
}

func addDeferStmt(fd *ast.FuncDecl) (added bool) {
	stmts := fd.Body.List

	// check whether "defer fntrace.Trace()()" has already exists
	for _, stmt := range stmts {
		ds, ok := stmt.(*ast.DeferStmt)
		if !ok {
			continue
		}
		// it is a defer stmt
		ce, ok := ds.Call.Fun.(*ast.CallExpr)
		if !ok {
			continue
		}

		se, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		x, ok := se.X.(*ast.Ident)
		if !ok {
			continue
		}
		if (x.Name == "fntrace") && (se.Sel.Name == "Trace") {
			// already exist , return
			return false
		}
	}

	// not found "defer fntrace.Trace()()"
	// add one
	ds := &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "fntrace",
					},
					Sel: &ast.Ident{
						Name: "Trace",
					},
				},
			},
		},
	}

	newList := make([]ast.Stmt, len(stmts)+1)
	newList[0] = ds
	copy(newList[1:], stmts)
	fd.Body.List = newList
	return true
}
