// +build ignore

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/markphelps/flipt/internal/fs"
	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/vfsgen"
)

func main() {
	source := filter.Skip(http.Dir("api"), func(path string, fi os.FileInfo) bool {
		// ignore the autogenerated flipt.swagger.json in favor of the hand edited version
		return strings.HasSuffix(path, "/flipt.swagger.json")
	})

	// Override all file mod times to be zero using ModTimeFS.
	err := vfsgen.Generate(fs.NewModTimeFS(source), vfsgen.Options{
		PackageName:  "swagger",
		VariableName: "Assets",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}