package main

import (
	"io/ioutil"
	"os"
	"runtime"
)

//go:generate go-bindata -prefix data ./data

func prepare() error {
	goos := runtime.GOOS
	is_linux := (goos == "linux")
	is_windows := (goos == "windows")

	var yara_bin_name string
	if is_linux {
		yara_bin_name = "yara"
	}
	if is_windows {
		yara_bin_name = "yara64.exe"
	}
	yara_bin := MustAsset(yara_bin_name)
	ioutil.WriteFile(yara_bin_name, yara_bin, os.ModePerm)
	// _ = os.Chmod(yara_bin_name, 0755)

	yara_rules_zip_name := "rules.zip"
	yara_rules_zip := MustAsset(yara_rules_zip_name)
	ioutil.WriteFile(yara_rules_zip_name, yara_rules_zip, os.ModePerm)
	err := Unzip(yara_rules_zip_name, "./")
	if err != nil {
		return err
	}

	return nil
}
