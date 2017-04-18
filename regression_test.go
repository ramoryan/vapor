package vapor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestRegression(t *testing.T) {
	files, err := filepath.Glob("../../test/regression/*.vapr")
	if err != nil {
		t.Error(err)
	} else {
		for _, file := range files {
			dst := ParseFile(file)
			html := strings.TrimSuffix(file, filepath.Ext(file)) + ".html"
			chk, _ := ioutil.ReadFile(html)

			fmt.Println("MEEEH: ", string(chk))

			if strings.TrimSpace(dst) != strings.TrimSpace(string(chk)) {
				out := strings.TrimSuffix(file, filepath.Ext(file)) + ".err"
				ioutil.WriteFile(out, []byte(dst), 0)
				t.Error(errors.New("result/expected mismatch:" + out))
			} else {
				fmt.Println("Test ", file, ": OK")
			}
		}
	}
}
