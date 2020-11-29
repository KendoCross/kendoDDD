package helper

import (
	"io"
	"os"
	"path"
	"path/filepath"
)

func MakesureFileExist(filepath string) {
	dir := path.Dir(filepath)

	var err error

	_, err = os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err.Error())
		}

		if err = os.MkdirAll(dir, os.ModeDir|0755); err != nil {
			panic(err.Error())
		}
	}

	_, err = os.Stat(filepath)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err.Error())
		}

		file, err := os.Create(filepath)
		if err != nil {
			panic(err.Error())
		}

		file.Close()
	}
}

func IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

func GetFileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if nil != err {
		return -1, err
	}
	return fi.Size(), err
}

// CopyFile copies the source file to the dest file.
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		if sourceinfo, e := os.Stat(source); nil != e {
			err = os.Chmod(dest, sourceinfo.Mode())
			return
		}
	}

	return nil
}

// CopyDir copies the source directory to the dest directory.
func CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, err := os.Open(source)
	if err != nil {
		return err
	}

	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			err = CopyDir(srcFilePath, destFilePath)
			if err != nil {
				continue
			}
		} else {
			err = CopyFile(srcFilePath, destFilePath)
			if err != nil {
				continue
			}
		}
	}

	return nil
}
