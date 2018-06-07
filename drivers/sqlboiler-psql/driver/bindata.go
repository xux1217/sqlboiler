// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.tpl
// override/templates/singleton/psql_upsert.tpl
// override/templates_test/singleton/psql_main_test.tpl
// override/templates_test/singleton/psql_suites_test.tpl
// override/templates_test/upsert.tpl
// DO NOT EDIT!

package driver

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

var _templates17_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x6f\x6f\xdb\xb6\x13\x7e\x2d\x7d\x8a\xab\xf1\x03\x22\xfd\xe0\x28\x7b\x9d\xc1\x2f\x9a\xf4\xcf\x8a\xae\xa9\xd7\xb4\x2b\xb0\xa2\x08\x68\xe9\x64\x13\xa1\x48\x95\xa2\x9c\x7a\xaa\xbe\xfb\x70\x14\x65\x49\xb6\x52\xbb\xed\xba\x75\xaf\x12\xf1\xcf\xdd\xc3\xe7\x1e\xde\x1d\x5d\x55\xa7\xf0\x3f\xc3\x16\x02\xaf\x58\x86\xd7\x5c\x2e\x4b\xc1\x34\x9c\xcf\x20\x7a\x4d\xa3\x11\x0d\xc3\x27\x28\xda\x99\x4f\x60\xb8\x11\x78\xc9\x0a\x84\xd3\xba\xf6\xad\x81\x35\xd3\x47\x6f\x8f\x59\x86\x62\xb8\xbd\x88\x57\x98\x31\xbb\x61\x7f\x6b\x74\xdd\xcd\xda\x0d\x3c\x85\xe8\x61\x92\x3c\x15\x6a\xc1\x84\x35\x72\x76\x06\x6f\xf2\x02\xb5\x79\x0a\xcc\x18\xcc\x72\x53\x00\x93\xc0\x25\x8d\x4d\x81\xc9\x04\x12\x85\x76\xac\xcc\x13\x66\x10\x94\x06\xbe\x94\x4a\x23\x28\x09\xb1\x92\xa9\xe0\xb1\x89\xfc\xb4\x94\x31\x04\x0a\xfe\x5f\x55\xfb\xa4\xd4\x75\xd8\xba\x09\x2c\x0a\xa9\x0c\x44\x57\xea\x52\x49\x83\x1f\x4d\x5d\xc7\xe6\x23\xd9\xa2\x8f\xc8\x0d\x4e\xa1\xaa\x50\x26\x84\xd2\xb9\x7e\x29\x2f\x9d\x3b\x58\x28\x25\xa6\x5b\xef\x97\x4a\x94\x99\x2c\xe0\xdd\xfb\xc2\x68\x2e\x97\x53\xb7\xc1\x8d\x4f\xdd\x71\xda\x65\x0b\xc5\x45\xe4\x3e\x42\x40\xad\x95\x86\xca\xf7\x34\x9a\x52\x4b\x50\x51\x83\xb4\x01\xda\x07\x69\xf7\x3d\x45\xf3\xe8\x22\x08\xab\x0a\x45\x81\x16\xf8\x14\xda\x09\xb7\xd2\xcd\xcb\xa4\xae\xa7\x7b\xd0\xf7\x50\x7f\x1e\x6c\xe8\xd7\xbe\xbf\x25\xc2\x6f\x62\x48\x51\xe9\xc5\x91\xfe\x9d\x33\xc9\xe3\x9d\x88\xce\xbf\x2d\xa4\x60\x6d\x16\x34\x66\x39\x3a\x3e\xc6\xf3\x1f\x2e\xc8\x95\xef\xf1\x94\x8e\x41\x97\xe4\x07\x8b\xf0\xcf\x16\xd7\x83\x19\x48\x2e\x08\xa8\x97\x13\xef\x81\x75\xf9\x56\xb3\xfc\xb1\xd6\x01\x6a\x1d\x86\xbe\x57\x8f\xa9\xe1\x9e\xf0\x8f\x45\x1f\x4a\xca\x26\xf4\x8d\x1f\x31\x2e\x8d\xd2\x5f\x72\xc5\x7b\xa6\xf3\xaf\x94\xc6\x7c\x9f\x73\x42\xd2\xf0\xfb\xd8\x61\xea\x31\xbf\xaf\x97\x6e\xb9\x1b\xea\xed\x1a\x8f\xc7\x3f\xa4\xa3\x11\xb5\xf7\xd5\x4d\xb8\xff\x55\xad\x6c\xa3\xf7\x3d\x74\x71\x8d\x38\x60\x0a\x12\x15\x97\x19\x4a\xc3\x0c\x57\x12\x52\xa5\x61\xa5\xee\xc0\x28\xc8\xb5\xca\x51\x8b\x0d\x94\x05\x0e\xcf\x6a\x3d\x0e\x8e\x7b\xb4\xac\xfe\xe3\xaa\xda\x96\x20\x9e\x82\x82\x59\x17\x5d\x57\x92\xec\x7c\x11\x5d\xe1\x5d\x30\xa9\xaa\x68\x7e\xbb\x24\x1a\xea\xfa\x1c\xa4\x82\xaa\xea\x15\xfd\xba\x26\x82\xd7\x3c\xc1\xc4\x92\x5e\x5a\x7a\x26\x56\x0e\xbe\x47\x1d\x03\x85\x5e\x50\x30\x27\x86\x67\x58\x18\x96\xe5\x37\xcd\xaa\x9b\x15\x8a\x1c\xf5\x04\x22\xa8\x9b\xd5\x9d\xaa\x7f\x51\xea\xb6\xb0\x3a\x1a\xe8\x3f\x51\x17\x98\x2a\x8d\x4d\x14\xec\xa2\xa3\x2f\xc3\xbe\x96\xbb\xd3\x12\x5c\x8b\xd6\x92\xef\xfb\x9e\xfc\xf3\x11\xa6\xac\x14\xa6\x20\xc7\x1f\x4a\xd4\x1c\x8b\xe8\x4a\xc9\x3f\x50\x2b\x37\x75\x8d\xa4\x83\xdd\x86\xaa\xae\x1d\xcd\x6f\xb9\x59\xb9\x95\x53\x50\xa1\xef\x7b\x67\x67\x70\x51\x72\x91\x40\xcc\xe2\x15\xc2\x2d\x6e\x80\xcb\x53\xc1\x25\x42\xb9\x14\x5c\x6c\xe0\x14\xb2\x4d\xf1\x41\xc0\xba\x80\x9c\xfe\xe6\x5a\x2d\x04\x66\x85\xef\x2d\xca\x94\x90\x14\x46\x67\x4c\x2e\x05\x52\x71\xb8\x28\xd3\x14\x75\x10\x5a\x8e\xf6\xf4\x42\x27\x5c\x94\x69\xf4\x56\x73\x83\x17\x1b\x83\xc1\x89\x39\xa1\xc0\x00\xe9\x72\x6c\x3a\xb5\xd3\xfe\xee\x70\x44\xc3\x14\xdc\x9b\x29\xc4\x04\x42\x33\xb9\xc4\x3d\x25\x0e\x0c\x5e\x5b\x51\x06\xf1\xfd\x06\x77\x97\x16\x46\xc7\x4a\xae\xa3\x67\x46\xb1\x60\xa0\xe5\xe8\x39\x97\x49\x38\x8a\x61\xb8\xee\x52\x89\xbf\x17\xc6\x30\x39\xdc\x0f\x63\xb8\xee\x6b\x60\xec\xdb\xec\x29\xf0\x33\xb6\x48\x43\xe7\x33\xa0\x59\x37\x11\xfa\x5e\x27\x92\x79\xd9\x8a\x64\x51\xa6\xa1\xbd\x63\xfb\x7a\x6d\x2e\xd3\x25\x69\xf2\x45\x69\xa2\x57\xbf\xaa\xf8\x96\xcc\x58\x95\x4e\x1b\xb1\x26\xe4\xe5\xc0\xe6\x77\xb7\xb8\x79\x7f\x9c\x8b\x37\x52\x34\x4e\x7c\x6f\xcd\xb4\xbd\x97\x36\xe7\xf8\x56\xca\x0f\x9c\x4b\x3a\x77\xdb\x45\x6a\x34\x04\x61\xc8\xf4\xb3\xde\x17\xdd\x46\xdf\xf3\x46\xdd\xb7\x49\xf1\xc0\x7c\xff\xc6\x1e\xb1\x54\x95\xa6\xbf\xba\x0b\x18\x7d\x86\xbe\xe7\xb9\x2a\x76\x3e\xdb\xd1\xe9\x9b\xde\xd7\xb7\xc0\x9e\x6b\x9e\x31\xbd\x79\x8e\x9b\xde\x4a\xe2\x94\x48\x14\x28\xdd\x35\x0a\x29\xc3\xff\x64\xd9\x3c\x9c\xe0\x4b\x69\x9f\x76\x46\xb9\x54\xbe\x9b\xee\xa9\x02\x95\x22\xb1\x09\x77\x61\x93\x99\x3b\x65\x6c\x21\x80\xe0\x85\x4d\xff\x36\xff\x7b\x6d\x8e\x20\x0e\x76\xf2\x45\x87\xb2\x9d\xe8\xe3\xdc\x6e\x9c\x41\xc6\x6e\x31\xe8\xca\x1c\xed\x38\x8a\x0c\xba\xaa\x64\x28\xdf\x6c\x3d\x4c\x47\x25\xbc\xbf\xd3\xc2\xf7\x9a\x0b\x10\x51\xf2\xdf\xc0\xac\x39\x6d\x23\xe4\xdf\x68\x68\xae\x0a\xb3\xd4\x58\x04\x09\x67\x02\xc9\xf8\xa4\xaa\xfa\xef\xe3\xba\x9e\x8c\x35\x60\x1a\x4d\x3b\xdc\x95\xf3\xb6\x5e\xdb\xe8\x35\x7e\xd7\x4c\x94\xf8\x82\xe5\xb9\x3d\x36\x5d\x91\xae\x10\x5d\x70\x99\xb8\xa9\x51\x32\x5e\x6f\x72\x1c\x3f\xec\xd6\x60\xeb\xcf\x6b\x0b\x6c\xaf\x30\x0e\x2a\xa3\xa5\xc2\x85\x4a\xa3\x09\x69\x61\x1b\x25\x0b\x54\xa3\xf9\x7e\x30\xc9\x23\xb9\x1a\x01\x39\x44\x69\x61\xd6\x4d\xdf\x61\xa9\xb3\xa9\x14\x53\x0a\x4d\xf4\x4c\x26\x5c\x63\x6c\x82\x76\xe0\x77\x5a\xf1\x32\x0d\x14\xa9\x64\xcd\xc4\xa0\xcc\xdb\xc9\xe2\x89\x56\x59\x0b\xde\x1a\x74\xa9\x70\x10\x98\xb0\x49\x60\x0d\x12\xea\xc6\xb8\x34\xa8\x53\x16\x63\xd5\xb4\x2e\x56\xe0\x3b\x34\xf5\x28\x6c\x37\x76\xce\xe7\x46\xdf\xef\xba\x67\xa3\x39\x29\x4f\x9b\xd6\xee\x11\x2e\xca\xe5\x0b\x95\x34\x75\x3d\xcd\x4c\xf4\x24\xd7\x5c\x1a\x21\x83\x6e\xde\xd6\x0f\xdd\xda\xb2\xba\x0e\x0f\xaf\x26\x76\x3a\x6f\x07\xce\xb3\xd3\x17\x37\x1d\x9c\xd7\xa8\x82\x9a\xb0\xc8\x5e\x9d\x57\xea\x2e\xe8\x81\x68\x7c\x44\x51\x14\x46\xd7\x31\xb3\x2a\x23\x52\x68\xc0\x9a\xb4\xfd\xca\xbd\x96\x9c\xab\xc0\xb6\x7c\x5f\x62\xd5\x3d\x54\xb6\xda\x9a\xcd\xa0\xf8\x20\xa2\xc7\x5a\x5f\xa9\x57\xea\xae\xa9\xbb\xce\x23\x89\xee\xec\x0c\xda\x3b\x6f\x1f\x2a\xf2\xc4\xb8\xc0\x03\x93\x1b\xb3\xa2\x17\xcd\xdd\x0a\x25\x98\x15\x6a\x3c\x29\xa8\x59\x6e\xee\xb9\x53\x66\xd7\x78\x8d\xd3\x74\xd3\xde\x1f\x7b\x3e\x7a\x11\x8c\xb3\xb4\x4b\xca\xfe\xbe\xc3\x9c\x0c\x29\xe8\xda\xec\xd1\xf6\x98\x6a\x05\xbd\xf6\xe8\xa9\x67\xd3\xdc\x97\x54\x8c\x49\x27\x9e\x7e\x61\x3f\xa2\x4d\x68\x1b\x91\x43\x6b\x6d\xe3\x01\xb3\xe6\xa0\xc7\x99\xde\x36\x20\xde\x67\x5e\x1e\xdb\x1f\xe8\x12\xf5\x30\x35\xa8\xbf\xea\xd5\xe1\xde\x15\xdb\x68\x39\xa3\x92\x8b\xfe\x8b\xa3\xf6\xff\x0a\x00\x00\xff\xff\xfa\xc7\xe8\xee\xde\x15\x00\x00")

func templates17_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertTpl,
		"templates/17_upsert.tpl",
	)
}

func templates17_upsertTpl() (*asset, error) {
	bytes, err := templates17_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.tpl", size: 5598, mode: os.FileMode(420), modTime: time.Unix(1528396430, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSingletonPsql_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5f\x6b\xdb\x3e\x14\x7d\x96\x3e\xc5\xad\xa0\xd4\x06\xe1\xfe\xfa\xfa\x83\x3c\xb4\xb1\xdb\x65\x04\xbb\x89\xed\x6d\x30\xf6\xe0\xd8\xd7\xa9\xc0\x91\x33\xfd\xc9\x56\xd6\x7c\xf7\x21\xc7\xae\xd3\xa6\xa3\x14\x82\x12\x74\xef\x39\x3a\xf7\xe8\x28\x97\x97\xb0\xb2\xa2\xa9\xf2\xad\x46\x65\x16\x16\xd5\xe3\x7d\xab\xcd\x5a\xa1\x3e\x14\x34\x14\x90\x2e\xe6\xa0\x4d\x61\x70\x83\xd2\x80\x36\x4a\xc8\x35\x58\xed\x56\xf3\x80\x60\x3b\x6c\x58\x98\x02\xb6\xaa\xdd\x89\x0a\xab\x80\xd6\x56\x96\xff\xa4\xf6\x2a\x51\x40\xa5\xc4\x0e\x95\x0e\x42\x51\x34\x58\x1a\x0e\xa6\x58\x35\x18\x17\x1b\xec\x8f\xe0\x60\xb7\x55\x61\x30\x91\xd3\x56\xd6\x8d\x28\x0d\xac\xda\xb6\xe1\xa0\xd0\x0c\x35\x0e\x65\x5f\xe3\xf0\xeb\x41\x18\x6c\x84\x36\xf0\xfd\xc7\x81\xc1\x1f\xc4\xfe\xa1\x64\xe8\x83\x89\xdb\xdc\x14\x72\xdd\x60\x30\xab\x50\x9a\x85\x6d\x0d\xa6\x8d\x28\xd1\xe9\x0a\xe6\x0b\x0e\xee\x7b\xb9\x18\xc9\x7d\x4a\x46\xf6\x8f\x10\x3c\xa3\x7c\x4a\x14\x7e\x0c\xab\xd0\xf8\x94\x92\x95\xad\xe1\xff\x63\xdc\x1d\x9a\x1b\x5b\xd7\xa8\x3c\x9f\x92\x0a\x6b\x54\x47\xc5\x7b\x3b\x14\x57\xb6\x76\xf0\xb2\x6d\xec\x46\x6a\x47\xc1\xc2\xe8\xf6\x3a\x9f\x67\xf0\xe5\x7a\x9e\x47\x29\xa3\x44\xd4\xd0\xa0\xf4\x46\x95\x70\x36\x81\xff\x9c\x5d\xcf\xb8\x09\xd4\x1b\x13\xa4\x5b\x25\xa4\xa9\x3d\xe6\x9d\x6b\xbf\xc7\x83\xfb\xcd\x38\x25\x84\x1c\x6c\xd6\xc1\xe7\x56\x1c\xb1\x71\x60\x1c\x98\x3f\x74\x0c\x0a\x9b\xa2\xc4\x87\xb6\xa9\x50\x75\x41\x08\x72\x8d\x33\x59\xe1\xef\xe3\x02\x7f\xa5\x8b\xc3\x15\x87\x2b\xdf\xa7\x64\x4f\x29\x71\x8a\x6e\x7b\x45\x94\x38\x87\xdc\x19\x6c\x16\xa7\xd1\x32\x83\x59\x9c\x25\x70\xae\xdd\x27\x89\x61\x9a\xc4\xb7\xf3\xd9\x34\x83\x4e\xe9\x73\xc6\xf8\x38\x22\xa7\xc4\x19\x25\x6a\x38\x3b\x09\xdc\xd3\x53\x27\xe4\xb0\xef\xc3\x64\x70\x67\x65\xeb\xe0\xab\x12\x06\xd3\x6e\x72\x8f\x85\x09\xc4\x49\xf6\x69\x16\xdf\x31\x27\x12\xb0\xd1\xf8\xb2\xf3\xe6\xd1\xa0\x77\xe1\x5d\xf8\x6f\xc0\x5f\xf8\x37\x26\xba\xb3\xef\xad\x7e\xe6\x43\x98\x40\x7e\x1f\x5e\x67\x11\xa4\x51\x06\xcc\x4d\x40\xea\x56\x81\xe0\xb0\x73\x97\xad\x0a\xb9\xc6\xfe\x95\x74\x42\xdc\x80\x62\xbc\xdf\x13\x65\xbc\x53\x46\xf6\x6e\xf9\xe9\x52\x59\xbd\x8c\xdd\x18\xd7\x93\xa4\xee\x3a\xe4\x6b\x91\x07\x92\x37\x4b\x0c\x26\x10\x7d\x9b\xce\xf3\x30\x0a\x03\xf6\x0e\x7a\x7f\xb8\xf4\x3e\xab\xee\x55\x8c\x53\x9c\x12\x2f\xa3\x2c\x5f\xc6\xb3\xf8\x0e\xd8\xbb\x4e\x77\x7f\x24\x83\xc9\xee\x0c\x85\xc6\x2a\x09\x0e\xd4\xf7\xfb\x74\x4f\xff\x06\x00\x00\xff\xff\x76\xcb\x6a\x7a\x25\x05\x00\x00")

func templatesSingletonPsql_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonPsql_upsertTpl,
		"templates/singleton/psql_upsert.tpl",
	)
}

func templatesSingletonPsql_upsertTpl() (*asset, error) {
	bytes, err := templatesSingletonPsql_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/psql_upsert.tpl", size: 1317, mode: os.FileMode(420), modTime: time.Unix(1528415503, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_main_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x6d\x6f\xe3\xb8\x11\xfe\x2c\xfd\x8a\xa9\x81\x1c\xa4\xad\xc2\x1c\xfa\xf2\x25\x07\x63\x91\x38\x4e\xba\xb8\xac\xe3\xb3\xdd\x1e\x8a\x6e\xdb\xa3\xad\x91\x42\x44\x22\x19\x92\x8a\xd7\x3d\xe4\xbf\x17\x43\x4a\xb6\x9c\xb5\x93\x6d\x8b\x05\xfa\xcd\x26\x9f\x19\xce\x33\xef\x7a\xe2\x06\x4c\xf9\x79\x7a\x73\xfd\x80\x1b\x18\x82\xc1\x12\x3f\x6b\xf6\xb1\xb1\x6e\xa4\x6a\x2d\x2a\x4c\x7e\x49\xde\xd7\xe9\x3f\x2e\x6e\x17\xe3\x19\x2c\x2e\x2e\x6f\xc7\x70\x37\xb9\xfd\x2b\xb0\x77\x9f\xe4\x27\xfb\xdb\x8b\xab\x2b\x18\xdd\x4d\xe6\x8b\xd9\xc5\x87\xc9\x02\xd8\xbb\xf7\x70\x7d\x37\x1b\x7f\xb8\x99\xc0\x8f\x63\x42\xbd\xff\xe1\x93\xfc\x25\x8d\x63\xb7\xd1\x08\xba\x5c\xa0\x75\x68\xc0\x3a\xd3\xac\x1c\xfc\x1a\x47\xf9\x72\xa4\xa4\x84\x77\xf6\xb1\x62\x57\x97\x31\x1d\x4c\x78\x8d\x40\x10\x21\xcb\x38\xba\x57\xd6\x01\xec\xfe\x37\x16\x4d\xff\xbf\xe6\xd6\xf6\xff\x5b\x5b\xd5\x2a\xc7\xdd\xbd\x32\x5e\x5e\x48\x17\xc7\x91\x2e\xa7\xdc\xda\x6b\x51\x6d\x01\x71\xe4\xd0\xba\xab\x4b\xff\x6a\x7b\xf6\x1c\xc7\x45\x23\x57\x20\xa4\x70\x49\x1a\xcc\xfc\xc8\x85\x84\x21\x7c\xd7\x71\xf8\xf5\x99\x60\x67\x67\x60\xd1\x35\x1a\xf2\xa6\xd6\x16\xdc\x3d\x42\xce\x1d\x5f\x72\x8b\x60\x57\xf7\x58\x73\xe0\x32\x07\x51\x93\x19\x16\x84\x23\x3b\x14\x70\x70\x48\x47\xdc\x6c\xc0\x70\x99\xab\xba\xda\x90\xae\x12\x25\x1a\xee\x30\x07\x32\xaa\xa7\x4a\x81\xbb\xe7\xce\x9f\x5a\x58\x71\x09\x4b\x04\xd3\x48\xe0\x25\x17\xd2\x3a\x52\xdc\x58\x21\x4b\xb2\x60\x5f\x91\x7d\xac\x96\x4a\x54\x68\xe0\x6e\xf6\x11\x34\x5f\x3d\xf0\x12\x59\xe0\x97\x68\x78\xd7\xf1\x49\x03\x91\x24\x05\x34\x46\x19\x22\x4d\xd9\x81\xc6\x84\x83\x38\x8e\x9e\x84\x46\xc3\xe6\xe8\xae\xb0\xe0\x4d\xe5\x92\x81\xa6\xb0\x05\x9e\x83\x0c\x06\xba\x59\x56\x62\x35\x48\x8f\x42\xc9\x0b\x83\x0c\xfe\xf8\x87\xdf\xff\xee\x38\xa8\x8d\x20\x29\x34\xf8\xd8\x08\x83\x83\x94\x42\xc7\xda\xd4\x18\x42\x10\xbc\x41\x37\xf7\xf1\x6a\xe5\xf2\xa5\xe4\x35\x61\x23\xcd\x7c\xd6\x1c\x03\xd2\x65\x80\xf9\x64\x3a\x06\xa3\xcb\x00\xf3\x39\x76\x0c\x46\x97\x2d\x8c\x52\xad\x07\xfb\x20\xf7\x78\x7b\x4c\x97\x9e\xc7\xb4\x75\xe4\x89\x31\xf9\x7e\x08\x4f\xbc\xe2\xec\x12\x4b\x21\xff\xc2\x2b\x91\x73\x27\x94\x4c\x52\xd6\xfe\xc1\x24\x8e\x22\x0f\x09\x6a\x26\xca\x8d\x6b\xed\x36\x49\x20\x47\x41\xd9\x71\xc9\x8e\x62\xc9\x25\x1d\x36\xb8\x67\x8b\x9d\x28\x97\xf8\x1f\xe3\xc7\x86\x57\x36\x09\x3c\x33\xf8\xbe\xc3\x07\x72\xaf\x28\x0f\x71\xeb\xe0\x5d\x98\x8e\xe3\x5b\x1f\x74\x02\x5b\x97\x64\x71\x94\xb2\xd1\x3d\xae\x1e\x12\x72\x8f\x28\x7c\x76\xfe\x66\x08\x52\x54\x94\xaf\x91\x41\xd7\x18\x49\xa7\x71\xf4\x1c\xc7\xd1\xd9\x19\x8c\x0c\x72\x87\xc0\xdb\x32\x13\xff\xc2\x1c\xf2\x25\x90\x09\x8c\xe2\xd1\x2b\xfe\xe1\x0e\xc3\xe6\x8e\x2f\x2b\x0c\x17\x5b\x06\xbd\x47\x87\xa0\x59\xcd\x1f\x70\x7a\xd3\xf5\x93\x24\xfd\xe1\x2d\x73\x7a\xb2\xb9\x51\x7a\xe1\x9f\x7e\x53\xae\x2f\xb6\xf2\x6c\xbe\x52\x30\x8e\xa8\x29\x8d\xea\x1c\xce\x87\x80\x9f\x71\xc5\x46\xaa\xae\xb9\xcc\x93\x81\x2e\xff\x49\x77\x54\x62\xa7\xa7\xa1\x7e\x4f\x95\xac\x36\x83\x0c\x76\x64\x3b\x71\x36\x96\x4f\x30\x04\xae\x35\xca\x3c\x51\x96\xfe\x0b\x43\x49\x48\x68\x5d\x8e\xe5\x53\x92\x32\xc6\xd2\x38\x0a\xf6\x1d\x7e\xd2\x3e\x56\x5e\xfd\xce\xe3\x7d\x81\xaf\x7f\x24\x8e\x4c\x06\x6b\x7a\x40\x28\x36\x15\x1a\x93\x9e\xa9\x73\x97\xab\x86\x8a\x70\xdd\xd7\x3d\x77\xb9\x6f\xde\x12\xd7\xd7\x3f\xe2\xe6\x0a\xad\x33\x6a\x83\x26\xd9\xce\xbe\x0c\xcc\x5e\x74\x77\xfa\xb8\x71\xaf\x7a\x5a\x19\xcb\x7e\x36\x5c\x27\x68\xa8\xda\x0a\x2e\x2a\x6a\xdf\x0a\x2c\x89\x42\xeb\x69\x58\x05\x3f\x50\x13\xe8\x87\xb4\x6f\xe3\xff\xfa\x92\x7d\xac\xf6\x9f\x39\xc0\xe7\x67\x2e\x0e\x3d\x52\xd4\x8e\x4d\x8d\x90\xae\x92\xa4\x3d\xfd\xba\x77\xd7\x5c\x38\x28\x94\x39\x4c\x32\x8e\xd6\x6c\x54\x29\x8b\x49\x0a\x67\x67\x70\x51\xd0\xe0\xef\x32\x52\x58\xc8\x95\xc4\x0c\x56\x84\xf0\x73\x73\x6d\x84\x43\x40\x99\x83\x2a\xfc\x81\x16\x1a\xe3\x83\xbe\xfa\x46\x2c\x0e\x38\xb0\x95\x97\xa2\xda\x2e\x05\xfb\x43\xd3\x34\x72\x54\xe7\x89\xa5\x0c\xcb\x3a\xe9\x76\x8f\xc8\x80\x9b\xd2\x02\x63\x2c\xfc\xef\x8d\xd6\xd5\x81\x12\x69\x85\x83\x54\x5b\x4f\xff\x59\x61\x88\x02\x2a\x94\xc1\x98\x94\x3c\xf3\xbd\xf7\xcb\xaa\x57\x02\xc1\x12\xcb\x26\xb8\x9e\x21\xcf\xd1\xb4\xe8\x40\xd7\x86\xf2\x39\x1f\xc2\x77\xcb\x8d\x43\xcb\x2e\x9b\xa2\xf0\xbb\x0e\x5d\x91\xbb\x0f\x5d\xad\xfa\x85\x17\x54\x6c\x0f\x43\xe8\x82\xf0\x36\x96\xe7\x43\xa0\xeb\x59\x23\xdf\x88\x62\x17\x26\xd3\x48\x29\x64\x79\x3e\xd8\xba\x38\x78\x29\x7d\x81\x0f\x8f\xb7\x13\x25\x49\x0f\x5c\xa3\x31\x7b\xd7\x2f\x5b\xe6\x9b\x01\x6f\x3d\x0e\x7f\xfb\x7b\x70\x25\xd9\xdc\x0a\x75\x47\x1d\x8b\xb9\xa6\x77\x8b\x64\x30\xbd\xf9\xd3\xdd\x7c\x31\x3c\xb1\xbe\x01\xd2\x7c\xf5\xd3\xef\x05\x66\x7a\x37\x5b\x0c\x4f\x72\x8f\xa1\x99\x7a\x08\xf3\xe7\xf9\x78\xd6\xe9\xa1\x99\x7e\x50\xcf\xc5\x7c\x7e\xfd\xe1\x76\xdc\xe1\x76\x3b\x2f\xa1\x9f\x8f\xf0\x7a\x39\xcd\x76\xb9\xea\x6a\x9d\x75\x61\x13\xaa\x71\xa2\x62\x0b\xac\xb5\x87\x0d\xfc\xda\x57\x76\x3b\xd0\x6b\x23\xf9\x68\x01\x86\xba\x06\xa5\x69\xb3\x81\x42\x54\xd8\x55\x1f\x11\xbb\x6e\x89\x79\x2b\x06\x27\xf6\xfc\x24\x3f\xd7\xca\xba\xd2\xa0\x3d\xef\x79\xb4\xf3\xda\xd6\x33\xdb\x72\x08\xfb\x5b\xaf\x1e\xbe\x54\xdb\x29\xf2\x40\xdf\xa1\x77\x98\x4a\x12\x28\x7d\xc5\x9c\x93\xa3\x86\x74\x9b\xcf\xff\x91\x49\xbb\xf1\xfb\x0d\xcd\xea\x27\x1d\x0c\xc1\xd5\x9a\xf9\x4d\x2a\xdd\xd6\x0a\x1d\xb5\xd3\xe1\x48\x42\xee\xef\x3a\xbb\x74\x6c\x15\x68\xd6\xb6\x5e\x9f\x82\x01\x9c\x2f\xbf\xd8\x30\x0e\xeb\xee\xaf\x5f\x6f\x68\x26\xa8\xd7\x3b\x38\x3d\x15\xc5\x29\x7e\x16\xd6\xd9\x43\xcf\x9c\x9d\x81\x43\x6e\x72\xb5\x96\xbe\xaf\x37\x0e\x2d\xac\x2a\xe4\xb2\xd1\xe0\xb8\x7d\xb0\xb0\xbe\x47\xe9\x47\x5b\xf8\x8e\x2b\x84\x14\xf6\xbe\x6b\x6e\x87\xec\xec\x14\x1e\xff\x2a\xdb\x5b\x2a\xfd\xb7\x74\xe7\xd6\xb7\xd6\xca\x0e\x0f\x1e\xf1\x5f\xaf\xa7\x5b\xb7\x29\xcb\x66\x58\xab\x27\xda\x97\x7b\x2d\xe7\x58\x74\x95\x24\x56\x49\xfb\xe1\x9f\x05\x3a\xfe\x5b\x5b\x14\x5b\x2e\x07\x9e\xed\xae\x32\x6f\xb5\x37\xe0\x85\x47\x76\x88\x76\xf8\x3c\x56\xec\x4e\xa3\x4c\x06\x5d\xdf\x18\x64\x90\x1b\xf1\x84\x86\x4d\xe7\x3f\xdd\x5e\x36\xa2\xca\x7f\x6a\xd0\x6c\xda\xc1\xd0\x7d\x3a\x85\x2c\xff\xb2\x68\x5e\x96\x54\xfb\x81\x92\xbe\xd6\x00\xa5\xa8\xb2\x2f\x5c\xb6\xcf\xe5\x39\xfe\x77\x00\x00\x00\xff\xff\x00\xef\xd0\xbf\x8f\x11\x00\x00")

func templates_testSingletonPsql_main_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_main_testTpl,
		"templates_test/singleton/psql_main_test.tpl",
	)
}

func templates_testSingletonPsql_main_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_main_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_main_test.tpl", size: 4495, mode: os.FileMode(420), modTime: time.Unix(1528388322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_suites_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\xb1\x8e\x83\x30\x10\x44\x7b\xbe\x62\x84\x28\xe0\x04\xfe\x80\x93\xae\xba\xea\xae\x48\x11\x91\x0f\x70\xc2\x82\x2c\x39\x1b\x84\x17\x29\x92\xf1\xbf\x47\x18\x8b\x90\xce\xe3\x99\xb7\x3b\xdb\xcf\x7c\x43\x4b\x4e\x2e\xa3\xa3\x49\x4a\xc1\x97\x90\x13\xc3\x83\x6a\x2b\xf8\x0c\xf0\xbe\xc1\xa4\x79\x20\x14\x86\x3b\x7a\xd6\x28\x44\x5f\x2d\xe1\xfb\x07\xaa\x5d\x5f\x2e\x84\x94\x33\x7d\x32\xd5\x9f\xfb\x7f\x18\x8e\x36\x9a\xdd\x27\xeb\x8e\x72\xcb\x9e\xf4\x3d\x0e\x4b\x64\x94\x0b\x46\x3b\x4f\xda\x62\x81\x18\xb1\xf4\xab\x77\x50\xd4\x79\xe6\x32\xf7\xfe\x4d\x87\x90\xd7\x58\x6b\x7f\x7e\x6e\x27\x55\x71\x19\x71\x77\xec\x91\x54\xc8\x5e\x01\x00\x00\xff\xff\x2f\xea\xf2\xb5\x00\x01\x00\x00")

func templates_testSingletonPsql_suites_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_suites_testTpl,
		"templates_test/singleton/psql_suites_test.tpl",
	)
}

func templates_testSingletonPsql_suites_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_suites_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_suites_test.tpl", size: 256, mode: os.FileMode(420), modTime: time.Unix(1528388322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4d\x6f\x1a\x31\x10\x3d\xaf\x7f\xc5\x14\xb5\xd5\x6e\xb5\x71\xd4\x6b\x2a\x0e\x09\xc9\x21\xaa\x8a\x50\x58\x7e\x80\xb3\x3b\x4b\x2c\x8c\xbd\xf2\xce\x26\x50\xc7\xff\xbd\xb2\x81\x84\x4f\x09\xb5\xaa\xaa\x1e\x40\xf2\xf0\x66\xde\x9b\x37\x33\x38\x77\x01\x1f\x49\x3c\x2a\x1c\x8a\x39\x8e\xa5\x9e\x76\x4a\x58\xb8\xea\x03\x2f\x42\x94\x87\x30\xbc\x42\xbb\xf9\xe5\x15\x48\x92\xc2\x81\x68\x11\x2e\xbc\x67\xbb\x05\x46\xaa\xb3\x42\x1d\xa6\x37\xab\xf8\xd1\xe4\x67\x61\xcf\x4a\x2d\xc5\x1c\xd5\xd1\xd4\xb3\x64\xef\xa6\xd7\x9d\x2e\x81\xb0\x25\xe7\xf6\xd5\x7b\x3f\x69\x5a\xb4\x94\x12\x7c\x09\x08\xa9\xa7\xbc\xc8\xc0\xb1\x84\xf8\x48\x58\xa1\x14\xaa\x34\x63\x2c\x91\x35\x28\xd4\xa9\x73\xfb\x3a\xbc\x1f\x18\xd5\xcd\x75\x9b\x41\xbf\x7f\x12\x33\xb2\x72\x2e\xec\xf2\x3b\x2e\xdf\xd0\x8e\x25\x09\xf1\xf1\x4c\x36\x69\x2f\x7c\x37\x52\x4f\x21\xca\x83\x17\x49\x4f\x60\xb4\x5a\x42\xb3\xca\x83\x19\x2e\xa1\x5c\x65\xf6\x32\x96\x78\xc6\x92\x16\xb1\x0a\x26\x58\xa1\x2b\x33\x97\x3f\x91\x0f\xf1\x65\x8c\x58\xa5\x19\x4b\x9e\x85\x05\xb4\xf1\x63\x2c\x4b\x2e\x2f\xe1\x9a\x08\xe7\x0d\x01\x3d\x21\xdc\x0f\xc7\x77\x0f\x05\xb4\xb2\x42\x30\x35\x08\x0d\x93\x51\x88\xb0\xc4\x84\x8a\xdb\x36\xbd\xb7\xe0\x7c\x74\x21\x54\xdd\x26\x1d\x93\xed\x4a\x4a\x83\x9a\x1c\x3e\x9b\x1c\x8e\xb5\x7f\x7b\x53\x2c\x1b\x6c\x73\x20\xdb\x61\xf6\x2d\x16\xf9\xd0\x07\x2d\xd5\xda\x86\xbb\xa0\xb3\x4e\x7b\x13\x1d\x0d\x20\xf3\xce\x70\x42\x0e\xb4\x91\xf8\x0a\x3e\xb5\xbd\x3c\x14\x5c\xdb\xe2\x9c\xac\x41\x1b\x02\x3e\x34\x03\xa3\x09\x17\xe4\x7d\x49\x8b\xd0\x58\xb9\x7a\xf3\x1b\x51\xce\xa6\xd6\x74\xba\x4a\x33\xe7\x50\x57\xde\xb3\x64\x05\xf9\xd1\xb5\x54\x2c\xd2\x58\x65\xbb\xc2\x41\xe0\xd1\x48\xc5\x6f\x70\x2a\x75\xac\xa1\x5a\xdc\x8e\x15\x8b\xb4\xa4\x45\x1e\x3a\xdc\x30\x9c\x05\xca\x58\x52\x61\x8d\x16\x68\xc1\x1f\x8c\x52\x8f\xa2\x9c\x85\x79\xbe\x19\x6f\xf8\x7a\x61\x4f\xf5\x19\x06\x80\xba\x0a\x8b\x0f\xe1\x55\x0b\xd5\x62\xe4\xc8\x21\x52\xdf\xeb\x1a\x6d\x9a\xed\xbe\xce\x9b\x49\x17\xa9\x4f\x0c\xe4\x60\x12\xa5\xe9\x34\xc5\xc0\xfe\x52\x6d\x6e\x2f\xcd\xf8\x20\x80\xce\x6c\xe6\xdd\x87\x43\x9d\xe9\x86\x37\x40\x22\x73\x00\x7d\xdd\x81\xf4\x5e\x84\x26\x30\x1a\xc1\x62\x69\x6c\x95\xc3\xd4\xd0\x55\x2f\x5f\xe1\xd7\xaa\xf7\x4e\x65\x32\xba\xbd\x2e\xee\x8e\x9d\xca\x1f\xdf\xc2\x7a\x32\x67\xfd\x5d\x70\xce\xff\xee\xd9\xfc\xfe\x82\x85\x93\xfe\xd7\xfb\xf5\xbf\xac\x97\x67\xbf\x02\x00\x00\xff\xff\xb9\xcc\xee\xf9\x84\x07\x00\x00")

func templates_testUpsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertTpl,
		"templates_test/upsert.tpl",
	)
}

func templates_testUpsertTpl() (*asset, error) {
	bytes, err := templates_testUpsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.tpl", size: 1924, mode: os.FileMode(420), modTime: time.Unix(1528388322, 0)}
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
	"templates/17_upsert.tpl": templates17_upsertTpl,
	"templates/singleton/psql_upsert.tpl": templatesSingletonPsql_upsertTpl,
	"templates_test/singleton/psql_main_test.tpl": templates_testSingletonPsql_main_testTpl,
	"templates_test/singleton/psql_suites_test.tpl": templates_testSingletonPsql_suites_testTpl,
	"templates_test/upsert.tpl": templates_testUpsertTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.tpl": &bintree{templates17_upsertTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_upsert.tpl": &bintree{templatesSingletonPsql_upsertTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_main_test.tpl": &bintree{templates_testSingletonPsql_main_testTpl, map[string]*bintree{}},
			"psql_suites_test.tpl": &bintree{templates_testSingletonPsql_suites_testTpl, map[string]*bintree{}},
		}},
		"upsert.tpl": &bintree{templates_testUpsertTpl, map[string]*bintree{}},
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

