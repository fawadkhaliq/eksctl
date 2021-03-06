// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/schema.json (20.826kB)

package v1alpha5

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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x5f\x6f\xe3\xb8\x11\x7f\xcf\xa7\x10\x74\x7d\x2a\xf6\x36\x2d\x70\xbd\x87\x7d\x73\x9c\xdc\xae\xb1\x9b\xc4\x88\x0f\x5b\xa0\x87\x7b\x18\x51\x63\x99\x17\x8a\x54\x49\xca\xb1\xaf\xd8\xef\x5e\x48\xb2\x62\xc9\x12\x65\x99\xfa\x63\xa7\xdd\x87\x03\x6e\x65\x0d\xf9\x9b\xe1\xfc\xe7\x28\xff\xb9\x72\x1c\xf7\x2f\x8a\xac\x30\x04\xf7\x83\xe3\xae\xb4\x8e\x3e\x5c\x5f\xff\xa1\x04\xff\x31\x7b\xfa\x5e\xc8\xe0\xda\x97\xb0\xd4\x3f\xfe\xed\xa7\xeb\xec\xd9\x0f\xee\xbb\x94\x4e\xe2\x32\x21\xfa\xe1\xda\xc7\x25\xe5\x54\x53\xc1\xd5\xf5\x94\xc5\x4a\xa3\x9c\x0a\xbe\xa4\x41\xf6\x62\xe1\x67\xf7\x83\x93\xec\xe9\x38\x6e\xfe\x1e\x13\xb1\xff\x4f\xd0\x64\xf5\xfa\x93\xe3\xb8\x91\x14\x11\x4a\x4d\x51\x15\x9e\x3a\x8e\x4b\x32\xa2\x2f\x22\x08\x28\x0f\x4a\xbf\xd9\x33\xb2\xa7\x6f\x64\xe8\x15\x68\xbe\xfb\x2b\xe9\xb7\xdd\xff\x7d\xcb\x57\x73\xc1\xf7\x53\x6a\x60\xf3\x22\x27\x4b\x60\x0a\x5f\x5f\xd2\xdb\x08\x93\xed\x84\xf7\x07\x12\x9d\x2d\xb7\x5b\xa2\x2a\x9d\x2a\xcb\x46\x21\x21\x07\x8f\xe1\xaf\xdb\xe8\xe0\x07\xc7\x71\xa9\xc6\xf0\xf0\x61\x01\x89\xd2\xb2\xc4\x58\x81\xa5\xd2\x7b\x20\x25\x6c\x47\xe1\x3f\xd3\xa2\x02\xd3\x12\xff\x1d\x53\x89\xbe\xfb\xc1\xf9\x6d\xcf\x72\x88\x1a\x7c\xd0\x90\x63\xfa\xfd\xdd\x31\x21\x3d\x53\xee\x1f\x4a\xe7\x40\x0c\x25\xd6\x71\x03\x61\xc4\xd2\x55\x7e\x2b\x0b\xaf\x0c\xb5\xf0\xdb\xef\x57\x35\x52\x74\x21\xa2\x5f\x51\x2a\x2a\x78\x3f\xdb\xe3\xb3\x22\x9a\xbd\xa7\xe2\x7a\xfd\x77\x60\xd1\x0a\xfe\x71\x1c\xc3\xab\xb8\xc6\x34\xa0\x7b\xdc\x9f\x4f\x19\x0e\x85\x70\x54\x24\xb3\xc9\x7d\x3d\x90\x75\x44\x46\x05\xf2\x75\x3e\xad\x07\xc2\x85\x8f\x1f\xa5\x88\xa3\xb6\x16\xdc\x11\xa6\x11\xe8\x43\x0e\xe4\x74\xa7\x50\xd4\x37\xe0\x10\xa0\xff\x70\x29\x5c\xdd\x1f\xe0\xe9\xc4\xdc\x12\x64\x00\x1a\xe7\x52\x2c\x29\x6b\xed\x72\x87\x62\xed\x97\x12\x9a\x4e\x8c\xc1\x1a\x28\x03\x8f\x32\xaa\xb7\xff\x12\x7c\xb4\x68\xf2\xae\x18\xef\x6b\xb2\x83\x3e\xe4\xd7\x36\xd6\xd7\xc3\x52\x48\x24\x6a\x75\xc7\x89\xdc\x46\xba\xc6\x95\x0f\x83\x6e\x51\xd9\xb6\x1e\x9d\x06\x1d\x57\x0e\x6b\x50\x81\x2d\xb2\x2d\x6b\xe1\x04\x54\x8f\x83\xe5\x23\xd5\x63\xe4\x24\x77\xdc\x8f\x04\xe5\x5a\xb5\xc9\xc5\x22\x49\xd7\xa0\x71\x42\x08\xaa\xca\x91\xe4\x9b\x79\x42\x30\x04\xc3\x69\x46\xb1\xc7\x28\x39\x75\x81\xc1\xb8\x4f\x62\x67\x0b\xbe\x15\xca\x35\x25\xf8\x24\x18\x4e\x9e\x1e\x8e\x64\x3a\x06\x1b\x7b\x5d\x61\x8e\x32\xa4\x2a\x49\x99\xd4\x8d\x88\xb9\x0f\x72\x6b\xb3\x62\xee\xa8\x85\x7f\xb7\x41\x12\x27\xd2\xe8\x80\xcf\xb0\x5a\x4f\x58\x5f\xa8\x5e\x3d\xce\x6e\xa7\x56\x2a\xb3\x13\xdd\x84\x10\x11\x97\xf5\xd4\x39\x43\x44\xda\x2b\xce\xa2\x84\xeb\x82\xca\x8c\xd9\xe4\x3e\xcd\x4e\x5b\x28\x36\x87\x10\x6d\x0e\x34\xa1\x53\x11\x10\x2b\x62\x06\x1e\xb2\xca\x39\x46\xa0\x35\x4a\x3e\xaf\x47\x9a\xbe\xf2\xfe\xaf\x95\x67\x8d\x01\x7a\x2f\x64\xf3\xa9\x14\xe5\x78\x08\x14\x38\x17\x1a\xca\x95\xfe\x85\xa1\x1d\x52\x89\x0e\x14\xbc\x85\x3a\x9d\xa5\x12\xcb\xd5\xbd\xfe\x08\xb5\x06\xb2\x9a\x0b\x46\xc9\x76\xf2\xf4\x70\x86\xa4\xaf\x88\xa0\x37\x2d\x32\x9c\xb2\x96\x31\xf6\xa8\xff\x51\x3f\xde\xff\x0c\x19\x5c\x45\x7b\x9b\x52\x3a\x0d\xc1\x77\xfb\xce\x25\xd4\xc2\xca\xe5\x89\x69\xc6\x60\x3c\x1c\x46\xb9\xfa\x6e\x5a\x1a\xe3\x0a\xc7\x2d\x31\x28\x14\x1b\xc7\x7b\x6b\xb6\x31\x72\xb7\x8f\x05\xe5\xba\x55\x47\xed\xbb\x2e\xe7\x7a\xf0\x30\xf9\xb5\x8d\xde\x26\x09\xee\x0b\xb4\x77\x61\x83\xe1\x6d\x6f\x6a\xb8\xab\xce\x6c\x54\x81\x24\x0b\x2e\x29\x49\x0a\xb6\x58\xaf\x84\xa4\x7a\x7b\x5b\x13\x9c\x9b\x1a\xb5\x21\xfa\xf4\x90\xc0\x71\x5c\x8f\x72\x90\xdb\x3b\x4e\x84\x9f\x75\xf1\x5d\x0f\x14\xfe\xfc\x53\x29\x50\xd6\x47\x43\x69\xa5\xd6\x4a\x03\x79\x7e\x38\xc5\x10\x87\x3b\xbc\xd8\xe3\x78\x52\xb9\xdc\x9b\x3d\x76\xae\x69\xcc\x8d\x51\xd4\x2f\x42\x3e\xf7\x99\x3c\x67\x95\x7e\x7f\xbc\x0f\x8a\x7b\x30\x6d\xf9\x3a\x9f\xb6\xd1\x14\x7a\xec\xf6\xa6\xde\xc2\xa9\x2f\xed\x9a\x11\x24\x4e\xbc\x41\xd6\x2d\xb6\x59\xa0\x62\x03\xce\xe0\x29\x5d\x6e\x77\xb5\x80\x70\xa3\x25\x4c\x67\xb7\x4f\x67\x48\xf0\xd5\x0a\x64\xd6\x7d\x5f\x74\x95\x2b\xc4\x5a\x4c\x18\x13\x89\xcf\x9e\xcd\xd7\x3f\x5b\xf5\x4b\x38\x8c\xd4\xa1\x2c\x44\xdf\x7a\xf5\x34\x77\x18\x87\x47\xb5\xdf\xb5\xc1\x3b\x65\x7d\xc8\x31\xf5\xa6\x7f\x47\x73\x70\x45\x72\x72\x3a\xac\x90\x21\xd1\x42\xaa\xe1\x33\xe2\xa8\x9f\x5e\xe5\x1e\xf1\x25\x5d\x4e\x2d\x76\xa8\xba\xf9\x92\x7a\xc7\x3a\xbc\x13\xfb\xff\xaa\x1b\x0c\x27\xd7\xca\x76\xb2\xb6\xe7\x49\xb6\xf2\x3f\xda\x29\xed\xff\x5c\x3e\xd2\x56\x6d\x46\x89\x91\x18\x27\x98\x3c\x25\x3b\xd5\x1e\x4d\x82\x0c\xca\x4a\x33\x20\x8e\xc7\x7c\xb7\x5a\x2c\x9e\x10\x5a\x69\x09\x51\x35\x08\x0c\x88\xa9\x72\x29\x3f\x80\x42\x54\x46\x1b\x4e\x0e\x6f\xbb\x1a\x6c\x57\x28\x14\xd4\x7e\xb8\x30\x07\x21\xfd\x05\x42\xca\xac\x9a\xa5\x94\x2b\x0d\x9c\xa4\x03\x6f\x36\xf4\x3e\xaa\x44\x28\x53\x88\x80\x50\x6d\x84\x40\xb9\xc6\x00\x0d\xfa\x14\x52\xbe\xa0\x7f\x1a\xb7\x6f\xa6\x85\x8d\x35\xed\x5a\xb0\x38\x44\x6b\xf2\x0b\x98\xf0\x50\xaa\x3a\xda\xd1\x3c\x89\xb4\x58\x7c\x7a\x8b\xbe\xbf\xd1\xc2\x6c\xaa\x96\x4b\x4e\x40\x8a\xe6\x59\x33\xdc\xd7\x7c\xc0\xe5\xf1\xbc\x01\x5c\x24\x6a\x49\x89\x9a\x0a\x96\x64\x31\xe5\xb6\xb1\xc1\x47\x06\x12\x78\xcc\x20\x29\x55\xdb\x7b\xc3\x22\x91\x85\x5f\x0a\x33\x98\x6f\xb6\xcc\xca\x3b\x4d\x97\xd1\xcc\x19\x80\xbf\x3e\x43\x6c\xbd\x4f\xaf\x3c\x4d\x63\x5c\xe5\xe9\xec\x71\xbe\x70\x0f\x4c\x6e\x8c\x98\xfd\xe6\x42\x7d\x4e\xaf\x6e\x69\xf2\x9a\x17\x8f\x37\xbb\xb7\xf7\x6e\xb5\x18\x2e\x35\x3c\xbf\x91\x20\xd3\x4f\x44\x2d\x75\x79\x47\x6a\xbe\xed\xd3\x9a\x8f\x86\xc6\xdb\x5b\xce\x4f\x41\x05\x4d\xd1\xd6\x69\x50\xe2\xc1\x66\xbf\x2b\x78\xba\xd8\x07\x7a\xea\x31\xd2\x34\xa4\x7f\xa2\x31\x86\x35\xea\x5c\xc7\x0c\xbe\x10\x16\x6c\xae\xcd\x53\xea\x93\x6e\x09\x2b\xd4\xbb\x69\xe4\x4e\xec\x7f\x0e\xd5\x67\xdc\xce\x6e\xed\x51\xa4\x21\xd0\x52\xbb\xe7\xc2\x57\x73\x94\x89\x25\x5a\x2d\xf1\x66\xea\x0e\x0d\x75\x97\x0a\x17\x08\x94\x30\x50\x8a\x92\x2f\x02\xfc\x1b\x60\x49\xb4\x94\x89\x92\x9e\x25\xfe\xc9\x00\x75\xea\xa0\xcf\x33\x0c\x57\x57\x1f\x0f\x1c\x88\x4c\xf5\xf5\x68\x9f\x69\x19\x2a\xc1\x52\xff\x4e\x6b\x86\x52\x90\x67\x1c\xe9\xea\xee\x15\xd3\x4d\x71\x6b\x43\x32\x82\x37\x79\x7f\x71\x2a\xc2\x10\xb8\x7f\x06\xc5\x11\x6b\x94\x92\xfa\x15\x28\x56\x95\x56\x76\x43\x78\xfb\x60\xf4\xb1\x4d\xd4\xcf\xb1\x87\x0c\xf5\x5d\x7a\xe3\x7c\xf8\x41\xa7\xd3\xc9\x05\x0d\x38\xd5\x79\x75\xf0\x76\x8f\x05\xe3\x4d\xbd\xf6\x1e\xf9\xa0\x77\xe2\x87\x94\x4f\x05\x4f\xbc\x38\x1a\x0b\xdf\x23\x09\xae\xd6\x94\xf7\x98\xd1\xbf\x4d\xf9\xb7\xfc\x86\xe5\xfc\x93\xd0\x79\xbd\xba\xbb\x38\xb0\xbc\xfd\xcd\x57\xe9\x70\x81\x5c\x5c\xc2\x36\x5d\x2c\xae\xd1\xe3\x57\x32\x13\xdf\x17\x3c\x3d\xa4\xaa\xde\x8e\x10\x9f\xca\xdb\x8f\xa5\xbe\x26\xa6\x0d\xfd\x27\x1a\x42\x80\x37\x31\x65\x3e\xca\x62\xa3\x08\x62\x2d\x16\x04\x58\xf9\x29\x6e\x12\x67\x00\x2c\x71\xf7\xc5\x20\x80\x52\x67\xb7\x4b\xe5\x35\xa2\xe8\x1e\xd5\xca\x2d\x17\x45\xc5\x7f\x2e\xd5\xa6\xf4\xeb\xb2\xf4\x2b\x30\x6f\xc6\x03\x89\xaa\xf4\x74\xf3\x04\xdb\xd2\xe6\x95\x2f\x32\x8f\x77\xb6\x4a\x5c\xdb\xb8\xcb\x82\x78\x6c\xc8\x8b\x72\xb4\xa1\x2f\x0a\xdc\x0a\xfe\xee\x64\xac\xb0\x7b\x76\x5f\x2b\x26\x67\x6d\xb5\xdf\xd2\x6e\xbf\x82\xf6\xd8\x90\xa7\x6a\x66\x75\x36\xc6\x0f\x93\xcd\xe4\x43\xba\x84\x23\x8d\x4d\xb3\x8d\x14\x3a\xaa\x67\x08\x70\x49\xf5\x2d\xa9\x79\xd4\x83\xc7\xa1\x67\x2a\xbb\x05\xbf\xc5\x24\xa1\xbd\x01\x85\x9d\x9a\x63\xf9\x42\x73\x94\x04\xb9\x86\x00\x27\x9e\x58\x63\xe7\x75\x55\x24\x74\x7e\x30\x73\x21\xaa\x2d\x82\xd6\xab\xec\x26\x1c\xa9\xe0\x0b\x2d\x41\x63\x70\xce\x91\xfc\x52\xc7\xf2\x68\xe8\x49\x42\xf4\x22\x9d\xf7\x74\x0f\x02\xf7\x17\x41\x80\xb5\xf7\xe6\x59\x36\x36\xbb\x3d\x83\x96\x16\x78\xb0\xf1\x15\x7b\x66\x2f\xc1\x55\x24\xe5\xfd\xd1\x53\x03\xc6\xc4\xcb\x09\x67\x93\xbe\x6e\x23\x9b\x6c\xaa\xf4\x33\x6e\xe7\xa0\x8d\xae\xb4\x71\x42\x32\x5f\xa0\x13\xb1\x6d\x4e\xab\x44\x2c\x49\x79\x8a\x79\x36\x5a\xb1\xdf\xbf\x92\x3c\x56\x47\xb4\x8c\x67\x9e\x76\x3d\x6d\x64\xd6\x69\xbe\x2f\xb1\xa5\x4f\xc8\x2a\x8d\xa8\x31\x4d\xa9\x66\x70\xd8\xf8\x67\x10\x52\xfd\xb0\x61\x54\xe2\x9a\xda\x7e\xd2\x26\x62\x1d\xc5\xfa\x24\x8b\xea\x5f\x4a\x4f\xe5\x91\x43\xa3\x88\x62\x69\xa5\x45\x9e\x04\x6e\x4e\xbe\x1a\x8d\x1e\xf4\xea\x0c\x41\x64\xc9\xe2\x8d\xad\x93\x8b\x95\xb9\x04\x68\xa2\xc3\x10\xa8\x95\x74\x77\xf7\xa9\x8b\xc5\xa7\x53\x5d\x73\xff\x8a\x54\xfd\xab\x37\x2d\xb4\xea\x19\xb7\x67\xfa\xe0\xf6\x2a\xf9\xef\xdb\xd5\x7f\x03\x00\x00\xff\xff\x81\x7a\xc9\x26\x5a\x51\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x91, 0x79, 0x69, 0x4, 0xb9, 0x68, 0xf8, 0x6d, 0x7c, 0xa9, 0x96, 0x39, 0xeb, 0x7a, 0x3a, 0x99, 0xca, 0xb3, 0xd8, 0x69, 0xb, 0x18, 0xfc, 0xbe, 0x85, 0x27, 0x50, 0x95, 0xa7, 0x44, 0x95}}
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
	"schema.json": schemaJson,
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
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
