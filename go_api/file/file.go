package file

import (
	"go_api/err"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetAbsolutePath(filePath string) (string, error) {
	path, e := filepath.Abs(filePath)

	if nil != e {
		return "", err.New(err.UnknownErr, "file absolute path get failed")
	}

	return path, nil
}

func Open(filePath string) (*os.File, error) {
	fp, e := os.Open(filePath)
	if nil != e {
		if os.IsNotExist(e) {
			return nil, err.New(err.FileNotExists, "file not exists")
		} else if os.IsPermission(e) {
			return nil, err.New(err.FilePermissionDenied, "file unreadable")
		}

		return nil, err.New(err.UnknownErr, "file open failed")
	}

	return fp, nil
}

func GetExt(filePath string) string {
	ext := strings.ToLower(path.Ext(filePath))
	return strings.TrimPrefix(ext, ".")
}

func GetFileInfo(fp *os.File) (os.FileInfo, error) {
	stat, e := fp.Stat()
	if nil != e {
		return nil, err.New(err.UnknownErr, "file info get failed")
	}

	return stat, nil
}
