// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1619446565_postgres_make_anon_metrics_table.up.sql (326B)
// doc.go (382B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __1619446565_postgres_make_anon_metrics_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xdd\x4a\xc3\x40\x10\x85\xef\x7d\x8a\x73\xa9\xe0\x1b\x78\x35\x8d\x53\x8c\x6e\x92\xb2\xd9\x08\xbd\x5a\x96\x64\x28\x0b\xe6\x87\x9d\x35\xe0\xdb\x4b\x0d\x56\xb0\xf4\xf6\x7c\xdf\xcc\x39\x85\x65\x72\x0c\x47\x3b\xc3\x08\xcb\xe2\x47\xc9\x29\xf6\x8a\xfb\x3b\x00\x88\x03\x5a\xb6\x25\x19\x1c\x6c\x59\x91\x3d\xe2\x8d\x8f\x8f\x3f\x68\x14\xd5\x70\x12\x1f\x07\xbc\x93\x2d\x5e\xc8\x6e\xb9\xac\x32\xe5\xdf\x08\x75\xe3\x50\x77\xc6\x6c\x6c\x0d\x1f\x9f\x82\xd7\xb6\xa9\xff\x81\x73\xf3\x2a\x49\xe3\x3c\xdd\x38\x9d\x17\x49\x21\xc7\xe9\xe4\xf5\x4b\xb3\x8c\x37\x34\x15\x3d\x3f\xb9\x5a\xb5\xa4\xb9\x17\x55\x19\xb0\x6b\x1a\xc3\xf4\x37\x00\xcf\xbc\xa7\xce\x38\xec\xc9\xb4\xbc\xd9\x7d\x92\x90\x65\xf0\x21\xc3\x95\x15\xb7\x8e\xaa\xc3\xc5\x2b\x3a\x6b\xb9\x76\xfe\x42\x1e\x9e\xbe\x03\x00\x00\xff\xff\xa7\x96\x48\x08\x46\x01\x00\x00")

func _1619446565_postgres_make_anon_metrics_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1619446565_postgres_make_anon_metrics_tableUpSql,
		"1619446565_postgres_make_anon_metrics_table.up.sql",
	)
}

func _1619446565_postgres_make_anon_metrics_tableUpSql() (*asset, error) {
	bytes, err := _1619446565_postgres_make_anon_metrics_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1619446565_postgres_make_anon_metrics_table.up.sql", size: 326, mode: os.FileMode(0644), modTime: time.Unix(1619447043, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5a, 0xb0, 0x5d, 0xcb, 0xdc, 0x39, 0xb3, 0x38, 0x7e, 0x0, 0x42, 0xe6, 0x8c, 0x77, 0x74, 0x84, 0xf4, 0x4a, 0xfe, 0x5a, 0x22, 0xd8, 0x76, 0x23, 0xdd, 0xf8, 0x72, 0xdb, 0x4c, 0x7b, 0x8, 0xc8}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x3d\x6e\xec\x30\x0c\x84\x7b\x9d\x62\xb0\xcd\x36\xcf\x52\xf3\xaa\x74\x29\xd3\xe7\x02\x5c\x89\x96\x88\xb5\x24\x43\xa4\xf7\xe7\xf6\x81\x37\x01\xe2\x2e\xed\x87\xf9\x86\xc3\x10\xf0\x59\x44\x31\xcb\xc2\x10\x45\xe3\xc8\xaa\x34\x9e\xb8\x70\xa4\x4d\x19\xa7\x2c\x56\xb6\x8b\x8f\xbd\x06\x35\xb2\x4d\x27\xa9\xa1\x4a\x1e\x64\x1c\x6e\xff\x4f\x2e\x04\x44\x6a\x67\x43\xa1\x96\x16\x7e\x75\x29\xd4\x68\x98\xb4\x8c\xbb\x58\x01\x61\x1d\x3c\xcb\xc3\xe3\xdd\xb0\x30\xa9\xc1\x0a\xd9\x59\x61\x85\x11\x49\x79\xaf\x99\xfb\x40\xee\xd3\x45\x5a\x22\x23\xbf\xa3\x8f\xf9\x40\xf6\x85\x91\x96\x85\x13\xe6\xd1\xeb\xcb\x55\xaa\x8c\x24\x83\xa3\xf5\xf1\xfc\x07\x52\x65\x43\xa3\xca\xba\xfb\x85\x6e\x8c\xd6\x7f\xce\x83\x5a\xfa\xfb\x23\xdc\xfb\xb8\x2a\x48\xc1\x8f\x95\xa3\x71\xf2\xce\xad\x14\xaf\x94\x19\xdf\x39\xe9\x4d\x9d\x0b\x21\xf7\xb7\xcc\x8d\x77\xf3\xb8\x73\x5a\xaf\xf9\x90\xc4\xd4\xe1\x7d\xf8\x05\x3e\x77\xf8\xe0\xbe\x02\x00\x00\xff\xff\x4d\x1d\x5d\x50\x7e\x01\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 382, mode: os.FileMode(0644), modTime: time.Unix(1618237056, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0x2f, 0x1e, 0x64, 0x9, 0x93, 0xe4, 0x8b, 0xf2, 0x98, 0x5a, 0x45, 0xe2, 0x80, 0x88, 0x67, 0x7a, 0x2d, 0xd7, 0x4b, 0xd1, 0x73, 0xb6, 0x6d, 0x15, 0xc2, 0x0, 0x34, 0xcd, 0xa0, 0xdb, 0x20}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"1619446565_postgres_make_anon_metrics_table.up.sql": _1619446565_postgres_make_anon_metrics_tableUpSql,

	"doc.go": docGo,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"1619446565_postgres_make_anon_metrics_table.up.sql": &bintree{_1619446565_postgres_make_anon_metrics_tableUpSql, map[string]*bintree{}},
	"doc.go": &bintree{docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}