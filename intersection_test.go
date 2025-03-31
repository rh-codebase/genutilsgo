package genutilsgo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestIntersectionInt(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 3, 7, 5, 6, 5, 8, 4}
	fmt.Println(a)
	fmt.Println(b)

	ao, bo, ab := IntersectionI(a, b)
	fmt.Println(ao)
	fmt.Println(bo)
	fmt.Println(ab)
}

func TestIntersectionIntHash(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 3, 7, 5, 6, 5, 8, 4}
	fmt.Println(a)
	fmt.Println(b)

	ab := IntersectionHash(a, b)
	fmt.Println(ab)
}

func TestIntersectionIntSorted(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 3, 7, 5, 6, 5, 8, 4}
	fmt.Println(a)
	fmt.Println(b)

	ab := IntersectionSorted(a, b)
	fmt.Println(ab)
}

func TestIntersectionIntSorted2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 3, 7, 5, 6, 20, 8, 4}
	fmt.Println(a)
	fmt.Println(b)

	ab := IntersectionSorted(a, b)
	fmt.Println(ab)
}

func TestIntersectionIntHash2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 3, 7, 5, 6, 20, 8, 4}
	fmt.Println(a)
	fmt.Println(b)

	ab := IntersectionHash(a, b)
	fmt.Println(ab)
}

func TestIntersectionStringHash(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e"}
	b := []string{"z", "a", "x", "y", "q", "b", "r", "c"}
	fmt.Println(a)
	fmt.Println(b)

	ab := IntersectionHash(a, b)
	fmt.Println(ab)
}

func TestIntersectionStringHash2(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e"}
	b := []string{"z", "a", "x", "y", "q", "b", "r", "c", "z"}
	fmt.Println(a)
	fmt.Println(b)

	ab := IntersectionHash(a, b)
	fmt.Println(ab)
}

func TestIntersectionStringHash3(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e"}
	b := []string{"z", "a", "x", "y", "q", "b", "r", "c", "a"}
	fmt.Println(a)
	fmt.Println(b)

	ab := IntersectionHash(a, b)
	fmt.Println(ab)
}

func TestFilePaths(t *testing.T) {

	var fp FilePaths
	fp.Fn = make(map[string][]MyFileInfo, 1)
	fp.FilenameDuplicates("a", "/b", 23)
	fp.FilenameDuplicates("a", "/c", 44)
	fp.FilenameDuplicates("a", "/d", 33)
	fp.FilenameDuplicates("z", "/d/e/f", 23)
	fp.FilenameDuplicates("a", "/c/z/q", 33)
	for k, v := range fp.Fn {
		fmt.Println(k, ":")
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}

func TestFilePaths2(t *testing.T) {
	var fp FilePaths
	fp.Fn = make(map[string][]MyFileInfo, 1)

	err := filepath.Walk("/home/rh/proj/goutils",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(info.Name(), path)
			fp.FilenameDuplicates(info.Name(), path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range fp.Fn {
			if len(v) > 1 {
				fmt.Println(k, ":")
				for _, vv := range v {
					fmt.Println(vv.Path, vv.Size)
				}
			}
		}
	}
}
