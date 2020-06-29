// Code generated by go-bindata. DO NOT EDIT.
// sources:
// res/Firefox.lnk (1.224kB)

package assets

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _firefoxLnk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x5f\x48\x34\x55\x18\xc6\x7f\x9b\xab\x69\xdd\x28\x49\xa4\x18\x4e\x10\xa2\xd5\xae\x3b\xeb\xfe\xef\x0f\xda\xe8\xba\xe5\xae\x2e\x6a\x61\x32\x17\xae\x39\xeb\x2e\xee\xb2\xb6\x5b\x30\x4a\x48\x5d\x84\x24\x65\x54\x17\x2a\x11\x2e\x52\x64\x50\x0b\xfd\xb9\x09\x45\xbc\xe9\xa2\xa0\x30\x50\x88\x02\xd3\xa0\x8b\x82\x12\x0c\x82\x4a\xf0\x63\xce\x8c\xe2\xfa\x7d\x1f\x7e\x17\xdf\x0b\xf3\x9e\x99\xf3\x9e\xf3\x3e\xcf\x79\xce\x33\x51\xc0\x56\x7f\x07\x46\x6c\x8b\x4c\xf8\x3d\x40\x02\x96\x16\x57\x3a\x9e\xd8\xdb\xb5\xad\x7d\x69\x8e\xc5\xea\x1f\xbe\x69\xfb\x6e\xd7\xb6\x73\x5a\x25\x16\xda\x28\x8f\xb7\x6d\xf5\x34\xc7\x0f\x06\x76\xa4\x3f\x42\xe9\xda\xb5\x1f\xab\x79\xd8\xe5\x7a\xbf\x81\x76\x25\xa4\x72\x7d\x2c\x23\x8b\x71\xb2\xdf\xbf\x59\x47\x7c\x70\xa0\x77\xb0\x6b\xce\x0d\xaf\x50\x83\x9d\xa3\xad\xf6\x58\xe8\x31\xa3\xe6\x04\x9a\xab\xca\x01\x9f\x15\x59\x4f\x3e\x60\x8b\x93\x27\xc7\x24\x79\x12\x64\x91\x08\x93\x26\x83\x46\x01\x89\x56\x74\x02\xf8\x68\x03\x3a\x29\x90\x42\x23\x43\x86\x0e\xdc\x38\x99\x10\xef\x8f\xe0\xc0\x8d\x4c\x00\x19\x3f\x70\x1f\xa9\x73\x5e\xbe\xcd\x5a\x62\x03\xa3\x4f\x45\xa3\x73\x32\xc4\x2d\x5e\xc6\xbc\xf1\x18\xbc\xb6\x4a\x06\x27\xfb\xa5\x93\xd9\x87\xfd\xc4\xc8\x31\x2b\xb8\x64\x48\x58\xbc\xf2\x68\x24\xc9\xa1\x0b\x9c\x71\xdc\x18\x52\x3e\xd7\xff\xd7\x6b\x12\xc9\x74\x5e\x4b\xe6\x74\xa7\xa6\x6b\x44\x6e\x80\x34\xf5\x85\x81\x54\x73\x09\x69\xac\xf7\x1f\x92\x65\x9d\x9d\x68\xe8\x68\x40\x23\x90\x06\x9a\x2c\xe5\x8c\xd1\x6b\xed\x4b\x01\x0d\x40\x05\xa0\x6c\x34\xa9\xb5\x40\x77\xfa\xc1\x17\x5e\xd2\x24\x05\x25\xa4\xc6\xf3\xb9\xc9\x7c\x22\x2b\x85\xd3\x19\xad\x20\xb5\xea\x01\x5f\x9b\x1a\xcb\xcd\xa6\x33\x99\x84\x14\x36\x99\xaa\x17\x19\x13\xc0\x89\x13\xb5\x2c\xdf\xda\xd5\xa8\x57\x48\xa5\xde\xe4\x80\x2d\x28\x84\x6e\x1b\x8a\x21\x40\x25\x14\x1f\x02\xbe\xb6\xc4\xba\x1b\x8a\x47\x9d\xa3\x2f\xff\x5f\x3c\xe9\x9b\x7f\x7a\xfb\xf8\xe0\xab\xff\x16\x8d\xda\x98\x29\x5c\x71\xc4\x12\x73\xfa\x79\x47\x2a\x31\xa1\x15\xce\x2e\x65\xa2\x54\x7a\x71\xe6\xf8\x93\x9e\x77\x5e\xfd\xed\xe3\x9f\xd6\xe7\xed\x1f\xbe\xf5\xe8\x7e\x53\xd7\xef\x75\x2b\x53\xc3\x9f\x57\xdc\xb5\x10\xbd\xaa\xee\xb2\x89\x8b\x2e\xbe\x0e\xc8\x43\xf1\xa1\x5f\x17\x46\xc2\x9b\xd1\x80\xb2\x71\x72\xcf\xbb\x2d\xcb\xd9\xef\xb3\x67\x96\x6b\x06\x0c\x6f\x0c\xe1\x40\xc6\x81\xd7\xf2\xb2\x83\x0e\xfc\x04\xf1\xe0\x21\x48\x10\x9f\x10\xc1\xa8\x05\xf1\xe2\xc2\x23\xbe\xdd\x04\xc4\x4a\x2f\x1e\x5c\xb8\x84\xff\x8d\x1d\x3e\xd1\xcd\x9c\x31\x63\xdc\x22\xf2\x4c\x6b\x74\x75\x66\x35\xd8\xb7\xbe\x73\xb8\xeb\x38\xdc\xfb\xbb\xce\x14\xc9\x4c\xa7\xa7\xa6\xbd\x2a\xcf\x98\x19\xd5\x1e\xfc\x28\x84\x45\x6f\x1f\x3d\xc8\xa2\xb3\x22\x58\xca\x3c\x79\xee\xe3\xa0\x85\xf0\xd9\xbd\xd9\xee\x4f\xdf\x9c\x8e\x7c\x14\xe9\x74\x7e\xf0\xb8\xfe\xc6\xfd\x96\x57\x89\x00\xbf\xec\x37\xfe\x39\xff\xef\x51\x64\xa9\xf4\x73\xdf\x9d\xdf\xce\xb9\x2e\xfe\x08\xd7\x02\x00\x00\xff\xff\xb2\x29\x10\x72\xc8\x04\x00\x00")

func firefoxLnkBytes() ([]byte, error) {
	return bindataRead(
		_firefoxLnk,
		"Firefox.lnk",
	)
}

func firefoxLnk() (*asset, error) {
	bytes, err := firefoxLnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Firefox.lnk", size: 1224, mode: os.FileMode(0666), modTime: time.Unix(1572378490, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x24, 0xd0, 0x59, 0xcd, 0x33, 0x8f, 0xe0, 0x86, 0xf7, 0xdd, 0x55, 0x4b, 0xa2, 0x82, 0xe8, 0xfe, 0x37, 0x18, 0xa1, 0x8, 0x93, 0x93, 0xef, 0xd2, 0x87, 0x41, 0x56, 0xe1, 0xcb, 0xc2, 0xb0, 0x31}}
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
	"Firefox.lnk": firefoxLnk,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"Firefox.lnk": &bintree{firefoxLnk, map[string]*bintree{}},
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
