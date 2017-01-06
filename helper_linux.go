// +build linux

package main

import (
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"os/exec"
	"path/filepath"
	"strings"
)

func Find(file string) (found_path string, err error) {
	var out []byte

	_, file_name := filepath.Split(file)
	need_find := false
	if check_if_command_exists("locate", "-v") {
		out, err = exec.Command("locate", file).Output()
		if err != nil {
			need_find = true
			logrus.Warn("exec locate error")
		}
		if strings.Contains(string(out), "locate: stat ()") {
			need_find = true
			logrus.Warn("locate db error: run `sudo updatedb` first")
		}
	} else {
		need_find = true
	}
	if need_find {
		logrus.Info("install 'mlocate' package to speed up automatic configuration")
		out, err = exec.Command("find", Config.HomeDir, "-name", file_name).Output()
		if err != nil {
			fmt.Println(err)
			logrus.Warn("exec find error")
		}
	}

	for _, v := range strings.Split(string(out), "\n") {
		if strings.Contains(v, file) {
			found_path = v
		}
	}
	found_path = strings.Trim(found_path, " \t\r\n")
	if len(found_path) > 0 {
		return found_path, nil
	} else {
		msg := fmt.Sprintf("file '%s' not found", file)
		return "", errors.New(msg)
	}
}

// func Where(bin_name string) (path string, err error) {
// 	if len(bin_name) == 0 {
// 		return "", errors.New("in Where(): arg is not given")
// 	}
// 	out, err := exec.Command("sh", "-c", "where", bin_name).CombinedOutput()
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(out), nil
// 	found_path := strings.Split(string(out), "\n")[0]
// 	strings.Trim(found_path, " \t\r\n")
// 	return found_path, nil
// }
