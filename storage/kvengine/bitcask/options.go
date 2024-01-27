package bitcask

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

const subDirData = "data"

var (
	ErrStorageHomeDirEmpty = errors.New("storage home dir must not be empty")
	ErrMkStorageDirFail    = errors.New("failed to mk storage dir")
)

var Options = &BitcaskOption{
	LockFilename:   "kve.lock",
	DataFileSuffix: ".data",
	HintFileSuffix: ".hint",

	// 1 << 27 --> 128 MB
	// 1 << 28 --> 256 MB
	// 1 << 29 --> 512 MB
	// 1 << 30 --> 1024 MB
	// 1 << 31 --> 2048 MB
	// 1 << 26 --> 64 MB
	MaxFileSize: 1 << 26,
}

// BitcaskOption stores the config info.
type BitcaskOption struct {
	// The home dir must specify.
	StorageHomeDir string
	DataDir        string
	DataFileSuffix string
	HintFileSuffix string
	LockFilename   string
	MaxFileSize    uint64
}

func (o *BitcaskOption) SetStorageHomeDir(storageDir string) {
	if storageDir == "" {
		panic(ErrStorageHomeDirEmpty)
	}

	o.StorageHomeDir = storageDir
	o.createDir()
}

func (o *BitcaskOption) createDir() {
	o.DataDir = filepath.Join(o.StorageHomeDir, subDirData)

	if err := os.MkdirAll(o.DataDir, os.ModePerm); err != nil {
		panic(ErrMkStorageDirFail)
	}
}

// GetDataFile gets the data file path by given fileId, e.g. fileId=0, return path/000.data
func (o *BitcaskOption) GetDataFile(fileId uint32) string {
	filename := fmt.Sprintf("%03d%s", fileId, o.DataFileSuffix)
	return filepath.Join(o.DataDir, filename)
}

// GetHintFile gets the hint file path by given fileId, e.g. fileId=0, return path/000.hint
func (o *BitcaskOption) GetHintFile(fileId uint32) string {
	filename := fmt.Sprintf("%03d%s", fileId, o.HintFileSuffix)
	return filepath.Join(o.DataDir, filename)
}

// GetLockFile gets the lock file for the storage engine.
func (o *BitcaskOption) GetLockFile() string {
	return filepath.Join(o.StorageHomeDir, o.LockFilename)
}

// Print all of kvs options
func (o *BitcaskOption) Print() {
	t := reflect.TypeOf(*o)
	v := reflect.ValueOf(*o)
	optionsFieldNum := v.NumField()

	for k := 0; k < optionsFieldNum; k++ {
		fmt.Printf("\t%s: %v\n", t.Field(k).Name, v.Field(k).Interface())
	}
}
