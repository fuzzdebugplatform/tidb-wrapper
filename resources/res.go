// Code generated by go-bindata.
// sources:
// resource/context.go
// resource/route.go
// resource/trace.go
// DO NOT EDIT!

package resource

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _resourceContextGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xdd\x6f\xdb\x36\x10\x7f\x96\xfe\x8a\xab\x80\x05\x52\xe6\xd0\x9b\xb3\x27\x03\x1e\x90\x06\x2d\x3a\x20\x4d\x03\x3b\x45\x1f\x3c\x23\x51\xe4\x93\xcc\x85\x22\x65\x92\xca\x07\x0c\xff\xef\xc3\x91\x94\xa5\xb4\xd9\xfa\xd4\x61\x7e\xb1\x74\x1f\xbf\x3b\xde\xc7\x8f\x6a\xf2\xe2\x3e\xaf\x10\xac\xce\x0b\xbc\x69\x2d\x17\x37\xbf\xc4\x31\xaf\x1b\xa5\x2d\xa4\x71\x94\xa0\x2c\xd4\x9a\xcb\x6a\xbc\xc1\xa7\x64\xf8\xfe\x97\x51\x92\x04\x15\xb7\x9b\xf6\x8e\x15\xaa\x1e\x9b\x26\x57\x22\x2f\x0a\x3e\xae\x5b\x5d\xb7\xfa\x94\xf4\x42\x55\xf4\x27\xd1\x8e\x37\xd6\x36\xf4\xac\xb1\xc2\x27\xf7\x64\xac\xe6\xb2\x32\xee\xf1\x59\x16\x49\x9c\xc5\xf1\x43\xae\x29\x72\xdd\x42\xf7\x23\x15\x9b\x7f\xf9\xd8\x5a\x7c\x8a\xa3\x42\xb5\xd2\xa2\x5e\xa0\x85\x19\xd4\xf9\x3d\xa6\x75\xde\x2c\x3d\xd2\xea\xf8\x5c\x3d\xa0\x3e\xf7\x26\xd9\x00\x4e\x35\x28\x01\x60\x06\x65\x2e\x0c\xfa\xf7\x8f\xed\x4b\xe8\x2c\x8e\xcb\x56\x16\xb0\xe6\x15\x1a\x9b\x9a\xad\x00\x0f\x9b\x85\x7f\xd8\x51\x5e\xfa\x74\x02\xd3\x19\x84\x33\xb2\x4b\x7c\x3c\x9d\xa4\x59\xd0\xb0\x2f\x9a\x5b\x4c\x97\xab\xbb\x67\x8b\x04\x91\x65\x71\xe4\x8a\x86\xc1\xe9\x74\xc2\x16\x6d\x9d\x4a\x2e\xb2\x38\xd2\x68\x5b\x2d\x61\x83\x4f\xec\x9d\x33\xba\x56\x0b\x17\x2a\xf5\x3e\x59\xbc\x0f\x49\xb9\x33\xa5\x9b\x5c\xae\x05\x86\x74\x46\x70\x27\x54\x71\x7f\xd9\xd6\xc0\xa5\xcd\x7c\x76\x6c\x7e\xa1\x8a\x7b\xca\x67\x8d\x25\x6a\x20\xc9\x67\x29\xbc\x2c\x8e\x4a\xa5\xe1\x66\x04\x05\x25\xa3\x73\x59\x21\x0c\xea\xb9\x8b\xa3\xa8\xa0\x02\x92\xb6\x60\xf5\xd2\x87\x5b\xc5\x9d\x9c\xd5\x2d\xeb\xe0\x83\xe4\x9a\x26\x67\xb9\x9c\xac\xb8\xb4\x3b\x2f\xba\x52\x66\xd9\x65\x76\x3c\x59\x8d\xe0\x35\xf1\xcf\xbf\xae\xf6\x2b\x98\xd1\x51\xda\xc2\xee\xf6\xbb\xfd\x30\xc8\x21\xe3\x68\x4f\x15\xb0\xcf\x0d\xc2\x62\x2b\x5c\x77\x83\x0b\x65\xbb\xd8\x0a\x9a\x0f\xdf\x1c\xff\xbb\xa5\xc1\x9c\x26\x66\x2b\x92\xdb\x38\x72\xd9\xc1\x72\x75\xfc\x9e\x0b\xf4\xde\xc1\xc0\x4d\x7c\x72\x7b\x40\xef\x0d\x7a\x78\x92\xf5\xe8\xc1\xb1\xe4\x82\xfc\xa2\x0b\x2e\x09\xd9\x9f\xbc\x53\x0a\x2e\x87\xa0\x9f\x1a\x94\x0b\x9b\xdb\xd6\x0c\x50\xdf\x62\xa9\x34\xc2\x9d\x52\xa2\x73\xbb\x73\x22\x42\x3d\x2b\x2d\x6a\x78\xa1\xcc\x49\xe4\x41\xdd\x24\x70\xc9\x6d\xea\xba\x5d\x29\x20\x89\x7f\x89\x7c\xb3\x5c\xef\x8e\xe6\x6e\xc5\x3e\x78\x89\x2b\xad\xc6\x6a\x04\x37\xae\xed\x4e\xc7\xce\x55\xdd\x70\x81\xe9\xed\xd8\x95\x62\x7c\x7c\x9b\xf5\x20\xcc\xbb\xbe\x27\x74\xe7\xe9\xe2\x3c\xd2\x6c\x6b\xa0\x3d\x66\x73\x34\x8d\x92\x06\xdd\xbc\xeb\x11\x68\xdc\xb6\x68\x2c\x1c\x07\xad\x7b\xf3\x89\x45\xde\x8f\x7d\xc0\x7c\x8d\x3a\xcd\xd8\x02\x6d\x9a\x9c\x15\x05\x1a\x73\x72\xae\xa4\xd5\x4a\x9c\x9c\x09\xa1\x1e\x4f\x3e\x69\x5e\x71\x99\x8c\x20\x39\x4e\xb2\xd7\x5c\xcf\xd6\xeb\x7f\x70\xf5\x26\x86\x7c\x49\x81\xd2\x9e\x5c\x3f\x37\xf8\x3a\x8c\xcb\xa0\x08\x66\xd4\x2a\x72\xcb\x9b\x46\xf0\x22\xb7\x5c\x49\x4f\x6f\xb4\x30\x51\xe4\xe9\x80\x2a\x17\xe8\x8a\x5d\x6b\x5e\x5f\x69\x2c\xf9\x53\x1a\x8e\xcd\x3e\xcf\x2f\xd8\x55\x6e\x37\x23\x48\x42\x3d\x83\xf7\x70\x63\xfc\x7c\xff\x21\x4b\xe5\x16\xec\xb0\x78\x4b\x1f\x62\xe5\xa2\xa1\x40\x8b\x69\xaf\x1c\x05\x3e\xca\x02\x5a\xbf\x1a\x51\xc4\x4b\xe8\x21\x67\x33\x90\x5c\xf8\x8a\x77\x07\x76\xdd\x09\xa7\x76\x8d\xf1\xd3\x78\xa9\xec\x7b\xd5\xca\x75\xf6\x8d\x6d\xc7\x5c\x89\x54\x16\x4a\x4a\x22\xc9\xbc\x95\xa7\x2a\x7a\xdc\xbb\x83\x19\xc7\x21\x47\xdd\x56\xee\x16\x5b\x31\xed\xb3\x61\x66\x2b\x68\xec\x1c\xe1\xd0\xca\x5c\xe6\x35\x06\x26\x18\x72\x4f\x67\x5e\x87\xbc\xcb\xf3\xce\xe2\xe8\xb0\x91\x3b\x7a\x9a\x1e\x50\x1c\x6c\x4f\x15\xf3\xbe\xbc\x2e\x98\xa3\x98\xaf\x42\x78\x96\x0a\x21\x42\x0c\xe6\x16\x78\x06\x79\xd3\xa0\x5c\xa7\x03\x61\xe0\x55\x0f\xf9\x75\xb4\x61\xf9\x23\x53\x04\xe4\x03\x4c\x27\x19\x81\x07\xcc\xfa\x82\xd1\x48\xbd\x7d\xb6\x68\x46\x80\xda\x1d\x91\x24\xec\x63\xae\xcd\x26\x17\xa9\x29\xba\x96\x92\xf6\xcd\xb0\x99\x42\x55\xec\x4a\x73\x69\xcb\x34\x79\xa7\xb5\xd2\x53\xe7\x0a\xc1\x15\x7e\x7a\xf8\x93\x76\x06\xf5\x30\xdc\x8b\xae\x1e\x62\x93\xc1\xde\x4d\x66\x20\x84\x57\xf8\xc0\xb8\x19\xf9\x71\x74\x30\x1e\x87\x91\x86\x93\xdf\xc1\x6c\x85\x9b\x26\xcf\x92\xd3\x6f\xaf\xf4\xee\xae\x13\x28\x07\x5b\x91\x75\xeb\x30\x68\x3f\x75\xdf\x03\xff\xdb\x05\xd7\xc5\xea\x96\x0e\xe8\xaa\x0b\x59\xec\x3b\xd0\xc1\x8d\xf9\xfd\xc6\x39\xb8\xff\x4d\xf3\x1e\xb9\x2d\x36\x3f\xae\x79\xfe\xc3\x69\xc0\x69\x6a\x71\x68\xdd\x51\x7f\xdf\xed\xfc\x2d\x37\x25\xf3\x7d\xe7\x07\x33\x78\x43\xff\x03\x37\xe6\x2f\xbc\x19\x1c\xe4\x1e\x7f\xb8\x66\xdf\xa9\x7f\x40\xfa\x6f\x1b\xe0\x2a\x73\xc1\x8d\x45\x79\x26\xd7\x0b\xd4\x0f\x98\x26\xd3\xdf\x4e\x27\x93\x49\x32\x82\x50\x79\xfa\x7c\x49\xe9\x13\xee\xef\x00\x00\x00\xff\xff\xa6\x55\xfa\xcb\x67\x0b\x00\x00")

func resourceContextGoBytes() ([]byte, error) {
	return bindataRead(
		_resourceContextGo,
		"resource/context.go",
	)
}

func resourceContextGo() (*asset, error) {
	bytes, err := resourceContextGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/context.go", size: 2919, mode: os.FileMode(436), modTime: time.Unix(1572016663, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceRouteGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x4f\x6b\xb3\x40\x10\x87\xcf\xbb\x9f\x62\xf0\xf0\xa2\x22\x9b\x1c\x72\x7a\x4b\xae\x21\x87\xb6\x84\x24\xa5\x87\x52\xc2\xa2\x13\x95\x26\xeb\x76\x9c\xed\x1f\x82\xdf\xbd\xb8\x6e\x42\x5a\x42\x43\x41\x70\x1c\xc7\xe7\xf7\x8c\x6a\x75\xfe\xa2\x4b\x04\x26\x9d\xe3\xc6\x71\xbd\xdb\x8c\xa5\xac\xf7\xb6\x21\x86\x58\x8a\xc8\x20\x8f\x2a\x66\x1b\x49\x11\x11\x96\xf8\x61\x23\x99\x48\xc9\x9f\x16\x81\x1a\xc7\x08\x2d\x93\xcb\x19\x0e\x52\x58\xcd\x8c\x64\x20\x1d\x06\xd5\xd2\x9f\xa4\xa8\xb4\x29\x76\x48\xd0\x73\xd4\x7c\xb8\x90\x5d\x80\x0c\x43\xa1\x7b\x06\xf3\xf0\x16\x9e\x9e\x53\x5f\xf5\xf3\x5b\x67\x72\x88\x2b\x48\xbf\x3d\x93\x40\x28\xe2\xcb\xf9\x19\x5c\xca\x4f\xfa\x8c\x4a\x85\x94\x29\x68\x6b\xd1\x14\xf1\xb1\x93\xc1\x3f\x5f\x1c\x02\xf3\x04\xe9\x92\xeb\x26\x33\x67\xf2\xab\x32\x3d\x22\xf6\x46\x4b\x6c\x6d\x63\x5a\x7c\xa4\x9a\x91\x32\x48\x43\xf7\xd5\x61\xcb\xc9\xdf\x45\xcf\xb6\xf4\x26\x21\x31\xf9\x5d\x7d\x85\xf4\x86\xf3\xf5\x7a\x11\xbf\xc3\x45\x2d\xfa\x21\xd6\x7b\x6d\x1b\x82\x4d\x16\x7e\x84\xff\x53\x20\x6d\x4a\x84\x93\xed\x41\x0a\x51\x6f\x87\xdb\x2a\x08\xaa\x3b\xcd\x79\xb5\x62\xaa\x4d\x19\x93\x7a\x58\xde\xaa\x85\xe6\xca\xe3\xc4\xf0\xd5\x55\x30\x56\x67\x4e\x19\x50\xe2\x07\x90\x1d\x19\x29\x44\x27\xfb\x63\x34\x02\xd3\xc0\xf1\x5d\xef\x7b\x34\x16\x37\xd0\xa2\x29\x60\x32\x9e\x00\x85\x25\xa4\xf0\xee\xf7\x0d\xcf\x1a\x67\x8a\xc0\xeb\xbe\x02\x00\x00\xff\xff\x0b\xf4\x42\x19\xff\x02\x00\x00")

func resourceRouteGoBytes() ([]byte, error) {
	return bindataRead(
		_resourceRouteGo,
		"resource/route.go",
	)
}

func resourceRouteGo() (*asset, error) {
	bytes, err := resourceRouteGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/route.go", size: 767, mode: os.FileMode(436), modTime: time.Unix(1571915942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTraceGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\xbd\x0a\xc2\x30\x14\xc5\xf1\x39\xf7\x29\x0e\x1d\x1d\x44\x7c\x85\xce\x82\x88\xe0\x50\x42\x09\x21\x94\x60\xf3\x61\x72\x23\x96\x92\x77\x97\xb4\x8b\x8e\xe7\xf2\xbf\xbf\xa8\xf4\x53\x4d\x06\x9c\x94\x36\x63\x61\x3b\x8f\x27\x22\xeb\x62\x48\x8c\x2e\x2f\x5e\x77\x44\xbc\x44\x83\x3e\xbc\x4d\x42\xe6\x54\x34\x63\x25\xe1\x0a\x00\xb4\xe2\x78\x7b\x5c\x0a\x9b\x0f\x89\x7b\x53\xe0\x54\x1c\x86\xb3\xb4\x9e\xe5\x9e\xaf\x95\xc4\x35\x64\x00\x43\xbb\x52\xfd\x25\xfb\x50\x3c\xff\xc9\xf9\x35\xb7\x65\xfd\x44\xc2\x6d\xda\xbe\xe4\x61\x7b\xa0\x4a\xdf\x00\x00\x00\xff\xff\x97\x3c\x43\x33\xb6\x00\x00\x00")

func resourceTraceGoBytes() ([]byte, error) {
	return bindataRead(
		_resourceTraceGo,
		"resource/trace.go",
	)
}

func resourceTraceGo() (*asset, error) {
	bytes, err := resourceTraceGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/trace.go", size: 182, mode: os.FileMode(436), modTime: time.Unix(1572009054, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"resource/context.go": resourceContextGo,
	"resource/route.go": resourceRouteGo,
	"resource/trace.go": resourceTraceGo,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"resource": &bintree{nil, map[string]*bintree{
		"context.go": &bintree{resourceContextGo, map[string]*bintree{}},
		"route.go": &bintree{resourceRouteGo, map[string]*bintree{}},
		"trace.go": &bintree{resourceTraceGo, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

