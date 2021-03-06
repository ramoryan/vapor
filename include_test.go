package vapor

import (
	"os"
	"testing"
)

func TestInclude(t *testing.T) {
	// invalid file
	c, err := include("./examples/myplainhtml.html")
	if err == nil {
		t.Error("Including not exists file must returns error!")
	}

	// valid html file
	c, err = include("./examples/html/myplainhtml.html")
	if c == nil || err != nil || len(c) <= 0 {
		t.Error(err)
	}

	dir, _ := os.Open("./examples/vapr/")

	filenames := make([]string, 0)
	fi, _ := dir.Stat()

	if fi.IsDir() {
		fis, _ := dir.Readdir(-1) // -1 means return all the FileInfos

		for _, fileinfo := range fis {
			if !fileinfo.IsDir() {
				filenames = append(filenames, fileinfo.Name())
			}
		}
	}

	defer dir.Close()

	if len(filenames) <= 0 {
		t.Error("Vapor examples are not found!")
	}

	clearVariables()
	str := "abcdefgh"
	AddStrVar("vaporMainStr", str)

	strSlice := []string{"aaa", "bbb", "ccc"}
	AddStrSliceVar("vaporStrSlice", strSlice)

	intSlice := []int{956, 1848, 1956}
	AddIntSliceVar("vaporIntSlice", intSlice)

	vaporMap := map[string]interface{}{"Béla": 1, "Géza": 2, "Kálmán": 3}
	AddMapVar("vaporMap", vaporMap)

	for _, name := range filenames {
		_, err = include("./examples/vapr/" + name)
		if err != nil {
			t.Error(err)
		}
	}
}
