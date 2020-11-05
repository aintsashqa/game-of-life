package resources

// Generated by github.com/jteeuwen/go-bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _resources_alive_png = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x62\x60\x60\x70\x00\x61\x0e\x26\x06\x06\x06\x55\xee\x67\x9d\x0c\x0c\x0c\x8c\xc5\x41\xee\x4e\x0c\xeb\xce\xc9\xbc\x64\x60\x60\x60\x49\x77\xf4\x75\x64\x60\xd8\xd8\xcf\xfd\x27\x91\x95\x81\x81\x81\xb3\xc0\x23\xb2\x98\x81\x81\xef\x30\x08\x33\x1e\xcf\x5f\x91\xc2\xc0\xc0\x10\xe1\xe9\xe2\x18\x92\xe1\xfc\xf6\x3c\x23\x2f\x83\x81\x00\xf3\xc2\x9f\xe6\x77\x4f\x2a\xcc\x97\xe1\xab\xb0\x5e\xd7\x10\xe6\xa0\xbd\x40\xf4\x02\x57\x06\xeb\x16\xa6\x28\xc6\xb5\xc4\x73\x6f\x7d\x3c\xf0\x8d\xf1\x54\xf0\xa5\x43\x47\xac\xf6\xde\x60\x60\x60\x60\xf0\x74\xf5\x73\x59\xe7\x94\xd0\x04\x08\x00\x00\xff\xff\xf0\x15\xf4\xe7\xc3\x00\x00\x00")

func resources_alive_png() ([]byte, error) {
	return bindata_read(
		_resources_alive_png,
		"resources/alive.png",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"resources/alive.png": resources_alive_png,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"resources/alive.png": &_bintree_t{resources_alive_png, map[string]*_bintree_t{}},
}}