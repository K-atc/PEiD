package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func check_if_command_exists(cmd_name string, opt_version string) bool {
	cmd := exec.Command(cmd_name, opt_version)
	err := cmd.Start()
	if err != nil {
		return false
	}
	return true
}

func Where(file_name string) (path string, err error) {
	env_path := os.Getenv("PATH")
	var path_spliter string
	if runtime.GOOS == "linux" {
		path_spliter = ":"
	} else if runtime.GOOS == "linux" {
		path_spliter = ";"
	}
	cwd, _ := os.Getwd()
	for _, v := range append(strings.Split(env_path, path_spliter), cwd) {
		file_infos, err := ioutil.ReadDir(v)
		if err == nil {
			for _, fi := range file_infos {
				if fi.Name() == file_name {
					return filepath.Join(v, fi.Name()), nil
				}
			}
		}
	}
	msg := fmt.Sprintf("'%s' not found", file_name)
	return "", errors.New(msg)
}

// thanks to: http://stackoverflow.com/a/24792688
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
