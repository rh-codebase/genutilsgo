package genutilsgo

import (
	"fmt"
	"testing"
)

type Mode []map[string]string

type Db struct {
	Path       string `yaml:"path"`
	ConfigMode Mode   `yaml:"solar.mode"`
}

func TestReadYaml(t *testing.T) {
	testfn := "test.yml"
	var db = &Db{}
	ReadYaml(testfn, db)
	if db.Path != "/a/b/c/testDb" {
		t.Fail()
	}
	for _, m := range db.ConfigMode {
		fmt.Println("New config map...")
		for k, v := range m {
			fmt.Println("k= ", k, " v= ", v)
		}
	}
}

func TestWriteYaml(t *testing.T) {
	// First read in test data
	testfn := "test.yml"
	var db = &Db{}
	ReadYaml(testfn, db)

	testofn := "/tmp/test_out.yml"
	WriteYaml(testofn, db)

	var db2 = &Db{}
	ReadYaml(testofn, db2)

	if db.Path != db2.Path {
		t.Fail()
	}
	for idx, m := range db.ConfigMode {
		for k, v := range m {
			if v != db2.ConfigMode[idx][k] {
				t.Fail()
			}
		}
	}
}
