package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/StevenZack/tools/strToolkit"
)

func main() {
	root, e := os.Getwd()
	if e != nil {
		fmt.Println(`getwd error :`, e)
		return
	}
	e = filepath.Walk(root+"/assets", func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if !strToolkit.EndsWith(f.Name(), ".html") && !strToolkit.EndsWith(f.Name(), ".js") && !strToolkit.EndsWith(f.Name(), ".css") {
			return nil
		}
		fo, e := os.OpenFile(root+"/lib/"+getFirstName(f.Name())+".dart", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if e != nil {
			fmt.Println(`open error :`, e)
			return nil
		}
		defer fo.Close()
		_, e = fo.WriteString(`String Str_` + getFirstName(f.Name()) + `='''
		`)
		if e != nil {
			fmt.Println(`write error :`, e)
			return nil
		}
		fi, e := os.OpenFile(path, os.O_RDONLY, 0644)
		if e != nil {
			fmt.Println(`fi error :`, e)
			return nil
		}
		defer fi.Close()
		_, e = io.Copy(fo, fi)
		if e != nil {
			fmt.Println(`copy error :`, e)
			return nil
		}
		fo.WriteString("\n''';")
		fmt.Println(root + "/" + getFirstName(f.Name()) + ".dart")
		return nil
	})
	if e != nil {
		fmt.Println(`walk error :`, e)
		return
	}
}

func getFirstName(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "." {
			return s[:i]
		}
	}
	return s
}
