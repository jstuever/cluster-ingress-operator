// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/defaults/cluster-ingress.yaml (565B)
// assets/router/cluster-role-binding.yaml (329B)
// assets/router/cluster-role.yaml (654B)
// assets/router/deployment.yaml (2.28kB)
// assets/router/metrics/cluster-role-binding.yaml (285B)
// assets/router/metrics/cluster-role.yaml (259B)
// assets/router/metrics/role-binding.yaml (297B)
// assets/router/metrics/role.yaml (291B)
// assets/router/namespace.yaml (257B)
// assets/router/service-account.yaml (213B)
// assets/router/service-cloud.yaml (628B)
// assets/router/service-internal.yaml (512B)

package manifests

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

var _assetsDefaultsClusterIngressYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xb1\x6e\xdc\x30\x10\x44\x7b\x7e\xc5\xe0\xae\x3e\x05\x4e\xa9\x36\x49\x61\x20\x45\x80\x00\xe9\xf7\xa8\xd1\x69\x73\xd4\x92\x20\x57\x36\x92\xaf\x0f\x64\x49\x29\xec\x2b\x39\x58\x0e\xdf\x3e\x9e\xf1\x25\x2d\xcd\x59\x9f\xed\x56\xd9\x1a\x5e\xd5\x27\x0c\x1c\x65\x49\x8e\x17\x49\x0b\x5b\x38\xe3\xd9\x9a\x4b\x4a\xac\x88\xd9\x46\xbd\xa1\x15\x46\x1d\x35\xee\x23\x90\x4a\x48\x29\x49\x39\x40\x1c\x75\x31\xd7\x99\x5d\xb8\xab\x0d\xfd\xbb\x37\x82\x14\xfd\xc5\xda\x34\x5b\x0f\xdd\xb2\x2e\x17\x5a\x9b\x74\xf4\x4e\xf3\xa7\x97\x27\x49\x65\x92\xa7\x30\xd3\x65\x10\x97\x3e\x00\x26\x33\xfb\x03\x6d\x3f\xb7\x22\x91\x3d\xfe\x5f\xbe\xec\x75\x97\x5c\x58\xc5\x73\x0d\xc0\xa8\x26\x49\xff\xb2\xb6\xb5\xe5\x8c\x6f\xd6\x96\x4a\xf8\x24\x8e\x6c\xe9\x0f\x7c\x22\x8e\x79\x44\x31\x0c\x4c\x74\xbe\xe5\x87\x89\xb8\x6d\x70\xe0\x22\x5f\x7f\x33\x7a\xf7\xa1\xfe\xf2\x78\xa1\x03\x2b\x66\xf3\x9a\x57\x91\x61\x35\xb8\x01\xed\x5a\xbe\xe6\x59\xd4\xa0\xed\x81\x48\x2c\x4d\xed\xf6\x46\xa4\xef\xbe\x62\x65\xb0\x3c\xf0\x47\x92\xc8\x99\xe6\x6b\xe9\x16\xfd\x64\x62\xf4\x5c\xb7\x04\x98\xc5\xe3\xf4\x5d\xae\x4c\xed\x88\x80\xd3\x3a\x79\xa9\x39\xb1\xbb\x2f\x57\x56\xa3\xb3\xad\xcc\xaf\xb9\xde\x59\x4f\x3d\x4e\xa7\x00\x54\x96\xa4\x51\x5a\x8f\xcf\xe1\x5f\x00\x00\x00\xff\xff\xd7\x3f\xbd\x29\x35\x02\x00\x00")

func assetsDefaultsClusterIngressYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDefaultsClusterIngressYaml,
		"assets/defaults/cluster-ingress.yaml",
	)
}

func assetsDefaultsClusterIngressYaml() (*asset, error) {
	bytes, err := assetsDefaultsClusterIngressYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/defaults/cluster-ingress.yaml", size: 565, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x81, 0x95, 0x12, 0x92, 0x14, 0x74, 0x7b, 0x3, 0x61, 0x27, 0x74, 0xc8, 0x72, 0xaa, 0x82, 0x62, 0x1e, 0x43, 0x89, 0x16, 0x9e, 0x8b, 0xdc, 0xa5, 0xe0, 0xcc, 0xf0, 0x44, 0x7c, 0xe4, 0x79, 0xf6}}
	return a, nil
}

var _assetsRouterClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x31\x4e\xc4\x40\x0c\x45\xfb\x39\x85\x25\xea\x0c\xa2\x43\xd3\x01\x37\x58\x24\x7a\xef\xc4\xbb\x31\x49\xec\xc8\xf6\xa4\xe0\xf4\x28\x4a\x44\xc3\x4a\x29\x2d\xf9\xbf\xff\xfe\x13\xbc\xb3\xf4\x0e\x31\x10\x98\xb6\x20\x03\xd3\x89\x20\x14\x38\x1c\x3e\xc9\x56\xae\x04\x6f\xb5\x6a\x93\xc8\x69\x64\xe9\x0b\x7c\x4c\xcd\x83\xec\xa2\x13\x6d\x71\x96\x7b\xc2\x85\xbf\xc8\x9c\x55\x0a\xd8\x15\x6b\xc6\x16\x83\x1a\xff\x60\xb0\x4a\x1e\x5f\x3d\xb3\x3e\xaf\x2f\x69\xa6\xc0\x1e\x03\x4b\x02\x10\x9c\xa9\x80\x2e\x24\x3e\xf0\x2d\x3a\x96\xbb\x91\x7b\xb7\x9b\x24\x6f\xd7\x6f\xaa\xe1\x25\x75\xb0\x17\x1f\x3e\x87\xce\x1f\xe1\xf8\xdf\x4f\x5f\xb0\x3e\xa2\xa6\x6d\xd8\x85\x6e\x5b\xf1\xbf\x19\xe7\x32\x27\xf0\xdf\x00\x00\x00\xff\xff\x83\x13\xa9\xa6\x49\x01\x00\x00")

func assetsRouterClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterClusterRoleBindingYaml,
		"assets/router/cluster-role-binding.yaml",
	)
}

func assetsRouterClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsRouterClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/cluster-role-binding.yaml", size: 329, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x12, 0x9a, 0xeb, 0xd3, 0x79, 0x1a, 0xa6, 0x75, 0xb5, 0xda, 0x10, 0x70, 0xf9, 0x80, 0xd1, 0x51, 0x55, 0x98, 0x4d, 0x11, 0x98, 0x79, 0x5c, 0x46, 0x99, 0xa3, 0x39, 0x68, 0xe9, 0xa2, 0x72, 0x56}}
	return a, nil
}

var _assetsRouterClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\x31\x4f\x03\x31\x0c\x85\xf7\xfc\x0a\xab\xcc\x77\x15\x1b\xba\x95\x81\x1d\x21\x76\x5f\xe2\x72\xa6\x69\x1c\xd9\xce\x55\xe2\xd7\xa3\xbb\x16\x84\x5a\x90\xe8\x96\x17\xd9\xdf\x7b\x7e\x77\xf0\x98\x9b\x39\x29\x58\x94\x4a\x09\x54\x32\xc1\x4e\x14\x54\x9a\x93\x5a\x0f\x2f\x13\x1b\xd8\x24\x2d\x27\x18\x09\xd0\x40\xc9\x5c\x39\x3a\xcf\xab\xac\x62\xc6\x63\xa6\x3e\xec\xb9\xa4\xe1\x8b\xf8\x2c\x99\x02\x56\x7e\x25\x35\x96\x32\x80\x8e\x18\x7b\x6c\x3e\x89\xf2\x07\x3a\x4b\xe9\xf7\x0f\xd6\xb3\x6c\xe7\xfb\x70\x20\xc7\x84\x8e\x43\x00\x28\x78\xa0\x01\xa4\x52\xb1\x89\x77\xde\x71\x79\x53\x32\xeb\x4e\x91\x82\xb6\x4c\x36\x84\x0e\xb0\xf2\x93\x4a\xab\xb6\x2c\x75\xb0\xd9\x04\x58\xb2\x49\xd3\x48\xe7\x3f\x2a\xa9\x0a\x17\xb7\x55\x2d\x60\xab\x18\xe9\x24\x8d\x74\xe6\x93\x98\x49\xc7\xf3\x4a\x66\xf3\xf5\x71\x44\x8f\x53\xb8\xf6\x59\x4e\xa0\xe2\x1c\x7f\xde\x70\x6d\xed\xb2\xa7\xa2\x34\x33\x1d\x2f\x1c\xa2\x12\x3a\xfd\x41\xbe\x2c\xe7\x1a\x6c\x6d\x7c\xa7\xe8\x18\x23\x99\xdd\x66\xb0\x36\xd8\x7f\x37\xfb\x2b\x7e\x9d\xb9\xb5\x93\xff\x83\xb7\xe6\xe8\xed\x82\xdf\x6a\x5a\x02\x7f\x06\x00\x00\xff\xff\xef\xa4\xbd\x5b\x8e\x02\x00\x00")

func assetsRouterClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterClusterRoleYaml,
		"assets/router/cluster-role.yaml",
	)
}

func assetsRouterClusterRoleYaml() (*asset, error) {
	bytes, err := assetsRouterClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/cluster-role.yaml", size: 654, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8e, 0xeb, 0x9d, 0xb3, 0x5c, 0xa1, 0x1f, 0xfb, 0x73, 0x10, 0xc0, 0x9e, 0x1c, 0x20, 0xc8, 0xf7, 0x31, 0x30, 0x5, 0x6a, 0xbd, 0xba, 0xa9, 0x91, 0xf8, 0xce, 0xb6, 0x64, 0x47, 0x80, 0x47, 0xc8}}
	return a, nil
}

var _assetsRouterDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x4d\x73\x1a\x39\x10\xbd\xf3\x2b\xba\xec\x43\x4e\x63\xe2\xc4\x9b\xda\xd5\x8d\x82\xc9\x16\xb5\xc1\xa6\x80\xe4\x4a\xc9\x9a\x06\x54\x68\xd4\xda\x56\x0f\xf6\xec\xaf\xdf\x12\x03\x84\x19\xb0\x2b\x7b\x5b\x9d\xa8\xe9\x7e\xaf\x9f\xfa\x0b\xdd\xc2\x08\x83\xa3\xba\x44\x2f\xf0\x62\x65\x03\x05\xae\x74\xe5\x04\x76\xda\x55\x18\x7b\xb7\x30\x74\x55\x14\x64\x18\xfb\x35\x63\x8c\x10\x03\x1a\xbb\xb2\xe6\xe0\x01\x9a\x11\x74\x08\xce\x62\x01\x5a\x80\x2b\x2f\xb6\xc4\xbb\xde\xd6\xfa\x42\x9d\xd1\xf7\x74\xb0\x3f\x90\xa3\x25\xaf\x12\x20\xf6\x77\xf7\xbd\x12\x45\x17\x5a\xb4\xea\x01\xdc\xc2\xa3\x2e\x11\x6c\x84\x88\xd2\xa2\x02\xf0\xba\xc4\x18\xb4\x41\x05\x14\xd0\xc7\x8d\x5d\x49\x66\x1b\x45\x3d\x00\xa7\x9f\xd1\xc5\x44\x02\x89\x5a\x01\x53\x25\xc8\xbd\xa4\x35\x7d\x8d\xe8\xd0\x08\x71\xe3\x51\x6a\x31\x9b\x6f\x67\x90\x36\x08\x40\xb0\x0c\x4e\x0b\x1e\xdc\xcf\x34\xa6\xe3\x5a\xc8\x2e\xb6\x39\x07\x69\x77\x27\xb1\x77\x96\xfa\x86\xca\x40\x1e\xbd\xa8\xa3\x3d\x33\xe4\x85\xc9\xb9\x03\xf4\xa8\x77\xff\x1b\x79\x67\x0d\x0e\x8c\xa1\xca\x4b\x4a\x4d\x27\x8a\xa7\x02\xe7\xad\x8b\xa5\xf3\x8c\xa2\xef\xb6\xd5\x33\xb2\x47\xc1\x98\xe2\x52\x54\xe0\xac\xaf\x5e\x4f\x4e\x09\x9a\x31\x39\xec\x78\xbe\x10\x6f\x91\x15\x7c\xf8\x70\xcc\xcb\x6a\x65\xbd\x95\xfa\x27\x7f\xa0\x62\xe0\xc5\x0e\x2e\x0c\x00\x8c\x7f\x57\x96\xb1\x18\x55\x6c\xfd\x7a\x6e\x36\x58\x54\xce\xfa\xf5\x78\xed\xe9\xf4\x39\x7f\x45\x53\x49\x6a\x82\x33\x64\x06\x42\x81\x1c\xad\xeb\xbf\xb0\x56\x70\xd3\x56\xb5\xa1\x28\xa9\x01\x6e\xce\x10\x87\x32\x5c\xde\xbf\x39\xfb\x12\xe7\xaf\x21\x25\xd9\x92\x8f\x5d\x7b\x06\xdb\x14\xe8\xfd\x32\x75\x30\x90\x5a\x8f\x75\x8a\x06\x63\x7f\x61\x6c\xc6\xa1\x1b\x28\x85\x7a\xa3\xd8\x00\x81\x2d\xb1\x95\x7a\xe8\x74\x8c\x4d\x85\x63\x1d\x05\xcb\xcc\x34\x43\x97\x19\xb6\x62\x8d\x76\x07\x40\x62\xd0\xd6\x23\x9f\xc5\xc9\xf6\xc3\x71\xd1\x81\x69\x9e\x6c\xa9\xd7\x6f\x0c\xd4\xf1\xec\x5d\xa6\x95\x73\x53\x72\xd6\xd4\x0a\xc6\xab\x47\x92\x29\x63\x6c\x27\x20\x10\x4b\x6c\x97\xac\x09\xbb\x11\x09\xad\x3b\x9f\x34\x4e\x89\x45\xc1\xef\x1f\x5b\xd6\xc0\x24\x64\xc8\x29\x58\x0c\xa7\x6f\xd0\xc5\xf7\xf8\x1e\x1e\x3e\xff\x27\xc2\x12\x85\xad\x79\x97\xf2\xfe\x8f\xcf\x5f\x7e\x89\xf3\x16\x26\xc8\xeb\xce\x9e\xfb\x69\x46\xbf\x53\x2d\xef\x28\x5a\x22\x54\x11\x39\x69\x01\xed\x0b\x08\x3a\xc6\x17\xe2\x62\xbf\x36\xd7\xe8\x53\x43\xb5\x08\xaf\x5c\x61\xbe\x18\x2c\xe6\xcb\xe9\xd3\x6c\xd1\x52\xb9\x6f\x38\x05\x37\x49\xfe\xcd\x15\xd8\xec\xe9\xfb\x22\x9f\x2d\xe7\xf9\xec\xc7\x78\x98\x2f\x1f\x07\x93\x7c\x3e\x1d\x0c\xf3\x6b\x24\xd7\x96\x6a\x97\x6f\x94\x7f\x1d\x7c\xff\xb6\x58\x0e\xf3\xd9\x62\xfc\x75\x3c\x1c\x2c\xf2\xe5\x68\x3c\xbb\x46\xd7\x47\x31\xfd\xb0\xb5\x7d\x71\xb1\x1f\xd8\xee\xb4\x9c\x5f\xcc\xd9\x1d\x7a\x8c\x71\xca\xf4\x8c\xed\x89\x49\x4b\xc5\x6a\x37\x42\xa7\xeb\x39\x1a\xf2\x45\x54\x70\xdf\xee\xa1\xd4\x23\x7f\xa2\x74\x47\x2d\x68\xd9\x28\xe8\x6f\x50\x3b\xd9\xfc\xd3\x35\x5e\xab\x34\xa3\x2e\xec\xff\x43\x48\xa4\x8a\x4d\x77\x7d\xa4\x85\x8a\x51\x2e\x96\x8a\x09\x55\xd2\xf2\xb1\xec\x6e\x3d\x2c\x89\x6b\x05\x9f\x7e\xfb\x32\xb1\x67\xb6\x1d\xb9\xaa\xc4\x49\xfa\x1f\xe9\xcc\x70\x99\xbe\x4d\x1b\xbd\xef\xd7\x0c\x0e\x5d\x70\x78\x1f\x64\x06\x59\xd2\x33\xa0\xeb\x95\x72\xfa\xe4\x5d\xad\x40\xb8\x3a\x9a\x1a\x01\xa7\xd8\xd9\x2f\x70\x45\x34\xdc\x4e\xed\xc1\x7b\x42\x05\x2a\x78\xf8\xf4\xb1\x35\x6a\xf3\xbd\xfb\xe5\x03\x22\x6b\x86\xf4\xdf\x00\x00\x00\xff\xff\x78\xfb\x8a\x72\xe8\x08\x00\x00")

func assetsRouterDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterDeploymentYaml,
		"assets/router/deployment.yaml",
	)
}

func assetsRouterDeploymentYaml() (*asset, error) {
	bytes, err := assetsRouterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/deployment.yaml", size: 2280, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd8, 0xac, 0x8f, 0x9f, 0xaf, 0x84, 0xcd, 0xc2, 0x20, 0x84, 0x77, 0xa2, 0x9c, 0x90, 0x27, 0x1f, 0x29, 0x87, 0x36, 0x17, 0x83, 0x9b, 0x1, 0xab, 0x3e, 0xc8, 0xcd, 0x33, 0x59, 0x15, 0xab, 0x73}}
	return a, nil
}

var _assetsRouterMetricsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xc1\x4a\xc4\x40\x0c\x86\xef\xf3\x14\x79\x81\x56\xbc\x2d\x73\x53\x0f\xde\x57\xf0\x9e\x9d\xa6\x36\xb6\x93\x0c\x49\xa6\x07\x9f\x5e\x8a\x22\xc2\x42\xaf\x81\x7c\xdf\xff\xad\x2c\x53\x86\x97\xad\x7b\x90\x5d\x75\xa3\x67\x96\x89\xe5\x23\x61\xe3\x77\x32\x67\x95\x0c\x76\xc3\x32\x62\x8f\x45\x8d\xbf\x30\x58\x65\x5c\x2f\x3e\xb2\x3e\xec\x8f\xa9\x52\xe0\x84\x81\x39\x01\x08\x56\xca\x60\xda\x83\x6c\xa8\x2a\x1c\x6a\x07\xcc\xfb\xed\x93\x4a\x78\x4e\x03\xfc\x18\xdf\xc8\x76\x2e\xf4\x54\x8a\x76\x89\xbf\xd7\x66\x5a\x29\x16\xea\x3e\xac\x17\xff\x3d\x7b\xc3\x42\x19\xb4\x91\xf8\xc2\x73\xfc\x27\x9b\x6e\x74\xa5\xf9\x90\xdf\xa5\x9c\x0c\x02\xc0\xc6\xaf\xa6\xbd\x9d\xd4\xa5\xef\x00\x00\x00\xff\xff\x7f\xc0\x4a\x40\x1d\x01\x00\x00")

func assetsRouterMetricsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterMetricsClusterRoleBindingYaml,
		"assets/router/metrics/cluster-role-binding.yaml",
	)
}

func assetsRouterMetricsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsRouterMetricsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/metrics/cluster-role-binding.yaml", size: 285, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa2, 0xce, 0x4e, 0xd1, 0x37, 0xde, 0x79, 0x91, 0x5c, 0x71, 0xd1, 0x88, 0x1b, 0xdb, 0xaf, 0x1, 0xe5, 0x8c, 0x81, 0xb3, 0xfd, 0x30, 0xe3, 0x5d, 0xb0, 0x59, 0x8b, 0x2a, 0x47, 0xf9, 0xa0, 0xdf}}
	return a, nil
}

var _assetsRouterMetricsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xce\xb1\x4e\xc4\x40\x0c\x84\xe1\x7e\x9f\xc2\x12\x75\x72\xa2\x43\x69\x29\xe8\x29\xe8\x9d\xec\x70\xb1\x2e\x6b\xaf\x6c\xef\x49\xf0\xf4\xe8\x44\x90\xa8\xe7\x93\xfe\x79\xa2\xd7\x63\x44\xc2\xc9\xed\x40\x90\x02\x15\x95\xd6\x2f\xea\x6e\x0d\xb9\x63\x04\xa5\x51\x6c\xce\x1d\xe4\x36\x1e\xb6\x21\x5d\xb6\x20\x68\xed\x26\x9a\x85\xbb\x7c\xc0\x43\x4c\x17\xf2\x95\xb7\x99\x47\xee\xe6\xf2\xcd\x29\xa6\xf3\xed\x25\x66\xb1\xcb\xfd\xb9\xdc\x44\xeb\xf2\xd7\x7c\xb7\x03\xa5\x21\xb9\x72\xf2\x52\x88\x94\x1b\x96\x33\x32\x35\x53\x49\x73\xd1\x6b\xf1\x71\x20\x96\x32\x11\x77\x79\x73\x1b\x3d\x1e\x7a\xfa\x95\xb3\x75\x68\xec\xf2\x99\xb3\x58\x21\x72\x84\x0d\xdf\xf0\xdf\x78\x5c\xce\xcf\x85\xe8\x0e\x5f\xcf\xf1\x8a\x2c\x3f\x01\x00\x00\xff\xff\x4f\xd5\xdf\xe0\x03\x01\x00\x00")

func assetsRouterMetricsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterMetricsClusterRoleYaml,
		"assets/router/metrics/cluster-role.yaml",
	)
}

func assetsRouterMetricsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsRouterMetricsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/metrics/cluster-role.yaml", size: 259, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x98, 0x77, 0x73, 0x9c, 0x6, 0x33, 0xf2, 0x91, 0x6f, 0x3b, 0x35, 0x49, 0xf3, 0xa5, 0xfc, 0x1d, 0x2e, 0x2e, 0xa6, 0x5b, 0x95, 0xa6, 0x7e, 0x8d, 0xfe, 0x7e, 0xf4, 0x62, 0x30, 0xa5, 0x37, 0x61}}
	return a, nil
}

var _assetsRouterMetricsRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\x31\x4e\xc5\x40\x0c\x04\xd0\x7e\x4f\xe1\x0b\x24\x88\xee\x6b\x3b\x68\xe8\x3f\x12\xbd\xb3\x71\x12\x93\xac\xbd\xb2\xbd\x29\x38\x3d\x42\x8a\x44\x05\xd2\x6f\x47\x33\x9a\x87\x8d\x3f\xc8\x9c\x55\x32\xd8\x84\x65\xc4\x1e\x9b\x1a\x7f\x61\xb0\xca\xb8\xdf\x7c\x64\x7d\x3a\x9f\xd3\xce\x32\x67\xb8\xeb\x41\xaf\x2c\x33\xcb\x9a\x2a\x05\xce\x18\x98\x13\x80\x60\xa5\x0c\xcd\xb4\x52\x6c\xd4\x7d\xd8\x6f\x7e\xc5\xde\xb0\x50\x06\x6d\x24\xbe\xf1\x12\x03\xcb\x6a\xe4\x9e\x4c\x0f\xba\xd3\xf2\x33\xc7\xc6\x6f\xa6\xbd\xfd\x63\x48\x00\xbf\x84\xbf\x1e\xbd\x4f\x9f\x54\xc2\x73\x1a\xae\xf6\x3b\xd9\xc9\x85\x5e\x4a\xd1\x2e\xf1\xa0\xb4\xaa\x70\xa8\xb1\xac\x90\xbe\x03\x00\x00\xff\xff\x15\x9f\x30\x56\x29\x01\x00\x00")

func assetsRouterMetricsRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterMetricsRoleBindingYaml,
		"assets/router/metrics/role-binding.yaml",
	)
}

func assetsRouterMetricsRoleBindingYaml() (*asset, error) {
	bytes, err := assetsRouterMetricsRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/metrics/role-binding.yaml", size: 297, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xff, 0xef, 0x8, 0xfb, 0x1f, 0xa3, 0xc7, 0xfb, 0xbc, 0x6, 0x78, 0xad, 0x0, 0x28, 0x90, 0xc8, 0xe8, 0xf5, 0x7d, 0xf8, 0xd0, 0xeb, 0x52, 0xf, 0xd4, 0x81, 0xce, 0x69, 0xb8, 0x8c, 0x26, 0x8c}}
	return a, nil
}

var _assetsRouterMetricsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x6e\xeb\x30\x0c\x45\x77\x7d\x05\x91\x37\x3b\x0f\xdd\x02\xfd\x40\xf7\x0e\xdd\x19\xe9\x36\x26\x62\x8b\x02\x49\xb9\x68\xbf\xbe\x88\xe3\x02\x9d\x78\xef\x01\xc1\xc3\x7f\xf4\xa6\x0b\x9c\x1a\x50\x51\xe9\xfa\x45\xdd\x74\x45\xcc\x18\x4e\xa1\xe4\xc5\xb8\x83\x4c\x47\xc0\x68\x45\x98\x14\x27\xb4\xda\x55\x5a\x24\xee\xf2\x0e\x73\xd1\x96\xc9\xae\x5c\xce\x3c\x62\x56\x93\x6f\x0e\xd1\x76\xbe\x5f\xfc\x2c\xfa\x7f\x7b\x49\x77\x69\x35\xef\xae\xb4\x22\xb8\x72\x70\x4e\x44\x8d\x57\xe4\x3f\xca\xe9\x7e\xf1\x03\x7b\xe7\x82\x4c\xda\xd1\x7c\x96\x8f\x98\xa4\xdd\x0c\xee\xc9\xc6\x02\xcf\x69\x22\xee\xf2\x6a\x3a\xba\x3f\x2e\x4d\x74\x3a\x25\x22\x83\xeb\xb0\x82\x83\x39\x6c\x93\x02\xdf\xcb\xef\xd7\xcf\xd6\xb5\x3e\xc2\x06\xbb\x1e\xcb\x37\xc4\x3e\x17\xf1\x67\xf8\xe4\x28\x73\xfa\x09\x00\x00\xff\xff\x67\x78\x6f\x08\x23\x01\x00\x00")

func assetsRouterMetricsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterMetricsRoleYaml,
		"assets/router/metrics/role.yaml",
	)
}

func assetsRouterMetricsRoleYaml() (*asset, error) {
	bytes, err := assetsRouterMetricsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/metrics/role.yaml", size: 291, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0xe3, 0x68, 0x2c, 0xfd, 0xf1, 0x27, 0x9d, 0xa0, 0x3b, 0x10, 0x3e, 0xca, 0x3b, 0x76, 0x39, 0xf4, 0xb1, 0x37, 0x7b, 0xa3, 0xa7, 0x11, 0xc0, 0x6, 0x4b, 0x47, 0xbb, 0x93, 0x4b, 0xb7, 0xc2}}
	return a, nil
}

var _assetsRouterNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x41\x4a\x43\x41\x10\x44\xf7\x73\x8a\xe2\xbb\x8e\xe2\x76\xee\xa0\x1b\xc1\x7d\x67\x7e\x25\x69\x33\xd3\xfd\x99\xee\xc4\xeb\x8b\x46\x30\xe0\xba\x1e\x8f\x57\x67\xb5\xb5\xe2\x55\x06\x63\x93\xc6\x22\x9b\xbe\x73\x86\xba\x55\x5c\x9f\xcb\x60\xca\x2a\x29\xb5\x00\x26\x83\x15\xbe\xd1\xe2\xa4\x87\xdc\xa9\x1d\x27\x23\x0a\x20\x66\x9e\x92\xea\x16\xdf\x20\xfe\xa0\x47\xf5\x27\xf3\x95\xbb\x60\x67\x4b\x9f\x15\xcb\x52\x80\x2e\x7b\xf6\x5f\xf8\x01\xd2\xbb\x7f\xde\x99\x87\x9b\xa6\x4f\xb5\x23\xd2\xd1\xdd\xcf\x38\xf8\xc4\x1b\xe7\x55\x1b\x5f\x6e\x2b\x7c\xff\xc1\x96\x01\x35\xe4\x49\xe3\xa7\xef\x76\xe2\x5f\x42\xeb\x97\x48\xce\x3b\x71\xc5\x92\xf3\xc2\xa5\x7c\x05\x00\x00\xff\xff\xfd\xbd\x46\x74\x01\x01\x00\x00")

func assetsRouterNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterNamespaceYaml,
		"assets/router/namespace.yaml",
	)
}

func assetsRouterNamespaceYaml() (*asset, error) {
	bytes, err := assetsRouterNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/namespace.yaml", size: 257, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x93, 0x80, 0x53, 0x86, 0xac, 0xef, 0x4d, 0xd1, 0x80, 0xe0, 0x94, 0x53, 0xb, 0x1e, 0x3f, 0xd3, 0x6c, 0x99, 0x8c, 0x3c, 0x6d, 0x61, 0x35, 0xa0, 0x7a, 0x7c, 0x51, 0x10, 0x53, 0xc7, 0xc5, 0x86}}
	return a, nil
}

var _assetsRouterServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\xb1\x4e\xc4\x30\x10\x84\xe1\xde\x4f\x31\xd2\xd5\x9c\x44\xeb\x8e\x92\x16\x24\x7a\xb3\x99\xbb\x5b\x91\x78\xcd\xee\x3a\x88\xb7\x47\x41\x29\xa7\x98\x5f\xdf\x05\x2f\x22\x36\x7b\xe2\x66\x0e\xb7\x99\xf4\x80\x38\x5b\x72\xc1\xe7\x2f\xf2\x41\xd8\xa0\xb7\x34\xbf\xe2\x35\xf1\xa3\xeb\x0a\xe7\xf7\x54\x27\x64\x9d\x91\x74\x84\xd8\xe0\x52\x2e\x18\xf4\x4d\x23\xd4\x7a\xc0\xb9\xfe\x57\xd2\xf0\x76\x84\x31\xdc\x84\x11\xda\xef\xd7\xf2\xa5\x7d\xa9\x78\xa7\xef\x2a\x3c\x0d\xa5\x0d\xfd\xa0\x1f\xef\x8a\xfd\xb9\x6c\xcc\xb6\xb4\x6c\xb5\x00\xbd\x6d\xac\x27\xf0\x9c\x31\x9a\xb0\x1e\xba\x1e\x0f\xbd\xe5\x93\xf6\xbb\x33\xa2\xfc\x05\x00\x00\xff\xff\x33\xdc\xda\x8c\xd5\x00\x00\x00")

func assetsRouterServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterServiceAccountYaml,
		"assets/router/service-account.yaml",
	)
}

func assetsRouterServiceAccountYaml() (*asset, error) {
	bytes, err := assetsRouterServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/service-account.yaml", size: 213, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd0, 0xe3, 0x6, 0x3a, 0x88, 0x2e, 0x33, 0xe3, 0x24, 0xf0, 0xf0, 0xe9, 0x43, 0xc8, 0x46, 0x6c, 0x60, 0x9, 0x69, 0x84, 0x3, 0xd8, 0xc3, 0x80, 0xb, 0xab, 0x37, 0x13, 0xce, 0xf2, 0xeb, 0x60}}
	return a, nil
}

var _assetsRouterServiceCloudYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\x6b\x14\x41\x10\x85\xef\xfd\x2b\x1e\xe4\x9c\xa0\x98\x83\xcc\x31\x39\x09\x41\x16\x5c\xbc\x57\x7a\x6a\x76\x9a\xf4\x54\x35\x55\x35\xab\xfb\xef\xa5\x7b\x37\xa0\x28\x1e\xe7\x31\xfd\xd5\x7b\xdf\x1d\x5e\x94\x66\x3c\x51\x25\xc9\x6c\xf8\xc6\x76\x2e\x99\x11\x8a\x56\x29\x33\x8a\x60\x31\x95\x80\x2e\x88\x95\x61\xba\x07\x5b\x8f\x73\xd5\x7d\x06\xcb\xb9\x98\xca\xc6\x12\xfe\x90\xee\xf0\x5c\x77\xef\x3f\x7c\x91\x93\xb1\x3b\xbc\x71\x2e\x4b\xc9\x38\x53\xdd\xd9\x41\xc6\xa0\xd6\x6a\xe1\x19\x14\xb0\x5d\xa2\x6c\xfc\x90\xde\x8a\xcc\xd3\xfb\xf9\x44\xad\x7c\x67\xf3\xa2\x32\xe1\xfc\x31\x6d\x1c\x34\x53\xd0\x94\x80\x3b\x7c\xa5\x8d\x51\x1c\xce\xf1\x07\x02\x10\xda\xd8\x1b\x65\x9e\xa0\x8d\xc5\xd7\xb2\xc4\x7d\xb9\x36\x49\x40\xa5\x57\xae\xde\x21\xe8\x1d\xa6\xdb\x98\xd4\x3b\xf6\x34\x2e\x8d\xa7\x21\xe4\xdd\x47\x02\x9c\x2b\xe7\x50\xfb\xfb\x59\xef\x72\x5c\x8b\x83\xaa\x2b\x56\xf2\x21\x88\x97\x85\xf3\xd0\xb5\x91\xbd\x15\x39\xe1\xe5\x09\x4d\xb5\x22\xc8\x4e\x1c\x0e\x72\xec\xb2\x32\xd5\x58\x2f\xf8\xb1\xb2\x40\x74\xc0\x6e\x6e\x9b\xce\x57\x4f\xcd\xd8\xb9\xab\x17\x10\x44\x67\xc6\x2b\xaf\x45\xe6\x71\xc7\xaf\xaa\xfa\x6c\xfe\x19\x6c\x42\xf5\x68\xb4\x2c\x25\x1f\xb4\x96\x7c\xe9\x43\x32\xd5\x04\x34\xb5\x18\xab\xef\x87\xa0\x09\x6b\x44\x1b\x6b\x9a\x69\x68\xd6\x3a\xe1\xf8\x7c\xb8\x26\x6a\x31\xe1\xf3\x87\xf1\x71\x2d\x7c\x18\xd1\xed\xcd\xef\x08\xff\x2f\xe3\xf1\xf1\xd3\x3f\x21\x9e\x7e\x05\x00\x00\xff\xff\x14\xac\xd6\xf5\x74\x02\x00\x00")

func assetsRouterServiceCloudYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterServiceCloudYaml,
		"assets/router/service-cloud.yaml",
	)
}

func assetsRouterServiceCloudYaml() (*asset, error) {
	bytes, err := assetsRouterServiceCloudYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/service-cloud.yaml", size: 628, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0x35, 0x3d, 0xd7, 0x8, 0xf9, 0xba, 0x0, 0x54, 0xbd, 0x2a, 0xeb, 0x98, 0x83, 0x6f, 0x28, 0x5e, 0xda, 0xd8, 0xa9, 0x45, 0x65, 0xe2, 0x35, 0x98, 0xd0, 0x6, 0x64, 0xc4, 0x82, 0x36, 0x14}}
	return a, nil
}

var _assetsRouterServiceInternalYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\x31\x4f\x23\x41\x0c\x05\xe0\x7e\x7e\xc5\x93\x52\xe7\x74\x51\xa2\xd3\x31\x1d\x4a\x95\x06\x45\x02\xd1\x9b\x59\x27\xb1\x98\x9d\x19\xd9\xde\x20\xfe\x3d\xda\x0d\x11\x0b\x34\x29\xd7\xfb\xfc\xf9\xcd\x02\xdb\x3c\x98\xb3\xe2\x91\xf5\x2c\x89\xf1\x26\x7e\x42\xc7\x07\x1a\xb2\xe3\x4c\x79\x60\x0b\x5f\xa9\x5d\x39\x2a\x9b\xc1\x1a\x27\x39\x48\x02\x95\x52\x9d\x5c\x6a\x31\x90\x32\xa8\xb5\x2c\xdc\x81\x1c\x3a\x14\x97\x9e\xff\x84\x57\x29\x5d\xbc\x1e\x08\xd4\xe4\x99\xd5\xa4\x96\x88\xf3\x2a\xf4\xec\xd4\x91\x53\x0c\xc0\x02\x0f\xd4\x33\xa8\x74\xb8\xff\xe1\x1a\xfb\x37\x13\x28\xd4\xb3\x35\x4a\x1c\x51\x1b\x17\x3b\xc9\xc1\x97\x72\xe9\x17\x80\x4c\x2f\x9c\x6d\x54\x31\x96\x8a\xd0\x3a\x38\x6b\x18\x9b\x8f\x53\x7f\x6f\x1c\xaf\xef\xda\xed\x03\x60\x9c\x39\x79\xd5\xdf\x3b\x40\xab\xea\x13\xb6\x9c\xee\x46\x9c\xdc\xdb\x94\x1b\xff\x44\xfc\xff\x7b\xf9\xd0\xea\x35\xd5\x1c\xf1\xb4\xdd\x4f\x13\x27\x3d\xb2\xef\xa7\xd0\xe7\xce\x9c\xb0\x99\xb1\xd9\xac\x6f\x44\x6c\xa6\xf4\xec\x2a\x69\xee\xac\xee\xd6\xff\x6e\x80\xa6\xd8\x47\x00\x00\x00\xff\xff\x30\x30\x02\xf6\x00\x02\x00\x00")

func assetsRouterServiceInternalYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterServiceInternalYaml,
		"assets/router/service-internal.yaml",
	)
}

func assetsRouterServiceInternalYaml() (*asset, error) {
	bytes, err := assetsRouterServiceInternalYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/service-internal.yaml", size: 512, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x88, 0x3, 0x6d, 0x72, 0xb, 0x1e, 0xd8, 0x4, 0x82, 0xd2, 0xb1, 0x70, 0xa9, 0x3f, 0xf, 0x83, 0x3a, 0x2b, 0xeb, 0x18, 0x2b, 0x1e, 0xd2, 0xd6, 0xc3, 0xbe, 0x58, 0x72, 0xaa, 0xee, 0x2f, 0xe3}}
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
	"assets/defaults/cluster-ingress.yaml": assetsDefaultsClusterIngressYaml,

	"assets/router/cluster-role-binding.yaml": assetsRouterClusterRoleBindingYaml,

	"assets/router/cluster-role.yaml": assetsRouterClusterRoleYaml,

	"assets/router/deployment.yaml": assetsRouterDeploymentYaml,

	"assets/router/metrics/cluster-role-binding.yaml": assetsRouterMetricsClusterRoleBindingYaml,

	"assets/router/metrics/cluster-role.yaml": assetsRouterMetricsClusterRoleYaml,

	"assets/router/metrics/role-binding.yaml": assetsRouterMetricsRoleBindingYaml,

	"assets/router/metrics/role.yaml": assetsRouterMetricsRoleYaml,

	"assets/router/namespace.yaml": assetsRouterNamespaceYaml,

	"assets/router/service-account.yaml": assetsRouterServiceAccountYaml,

	"assets/router/service-cloud.yaml": assetsRouterServiceCloudYaml,

	"assets/router/service-internal.yaml": assetsRouterServiceInternalYaml,
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
	"assets": {nil, map[string]*bintree{
		"defaults": {nil, map[string]*bintree{
			"cluster-ingress.yaml": {assetsDefaultsClusterIngressYaml, map[string]*bintree{}},
		}},
		"router": {nil, map[string]*bintree{
			"cluster-role-binding.yaml": {assetsRouterClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml":         {assetsRouterClusterRoleYaml, map[string]*bintree{}},
			"deployment.yaml":           {assetsRouterDeploymentYaml, map[string]*bintree{}},
			"metrics": {nil, map[string]*bintree{
				"cluster-role-binding.yaml": {assetsRouterMetricsClusterRoleBindingYaml, map[string]*bintree{}},
				"cluster-role.yaml":         {assetsRouterMetricsClusterRoleYaml, map[string]*bintree{}},
				"role-binding.yaml":         {assetsRouterMetricsRoleBindingYaml, map[string]*bintree{}},
				"role.yaml":                 {assetsRouterMetricsRoleYaml, map[string]*bintree{}},
			}},
			"namespace.yaml":        {assetsRouterNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml":  {assetsRouterServiceAccountYaml, map[string]*bintree{}},
			"service-cloud.yaml":    {assetsRouterServiceCloudYaml, map[string]*bintree{}},
			"service-internal.yaml": {assetsRouterServiceInternalYaml, map[string]*bintree{}},
		}},
	}},
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
