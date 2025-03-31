package genutilsgo

import (
	goi "github.com/juliangruber/go-intersect"
)

type MyFileInfo struct {
	Path string
	Size int64
}

// FilePaths is a structure hold a map of filenames and a slice
// of paths indicating where that file is located.
type FilePaths struct {
	Fn map[string][]MyFileInfo
}

func IntersectionI(a, b []int) (aonly, bonly, inter []int) {
	m := make(map[int]uint8)
	for _, k := range a {
		m[k] |= (1 << 0)
	}
	for _, k := range b {
		m[k] |= (1 << 1)
	}

	var inAAndB, inAButNotB, inBButNotA []int
	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		switch {
		case a && b:
			inAAndB = append(inAAndB, k)
		case a && !b:
			inAButNotB = append(inAButNotB, k)
		case !a && b:
			inBButNotA = append(inBButNotA, k)
		}
	}

	return inAButNotB, inBButNotA, inAAndB
}

func IntersectionHash(a, b interface{}) (inter interface{}) {
	inter = goi.Hash(a, b)
	return inter
}

func IntersectionSorted(a, b interface{}) (inter interface{}) {
	inter = goi.Sorted(a, b)
	return inter
}

func (fp *FilePaths) FilenameDuplicates(fn, path string, size int64) {
	if val, ok := (fp.Fn)[fn]; ok {
		fi := MyFileInfo{}
		fi.Path = path
		fi.Size = size
		val = append(val, fi)
		(fp.Fn)[fn] = val
	} else {
		ns := make([]MyFileInfo, 1)
		ns[0].Path = path
		ns[0].Size = size
		(fp.Fn)[fn] = ns
	}
}
