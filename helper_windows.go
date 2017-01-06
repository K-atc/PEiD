// +build windows

package main

import (
	"errors"
	"os"
	"os/exec"
	// "strings"
	// "github.com/Sirupsen/logrus"
	"fmt"
	"path/filepath"
	"strings"
)

func Find(file string) (found_path string, err error) {
	var out []byte
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	defer os.Chdir(cwd)

	_, file_name := filepath.Split(file)
	for _, drive := range []string{"C:", "E:"} {
		err = os.Chdir(drive)
		if err != nil {
			return "", err
		}

		out, _ = exec.Command("cmd", "/C", "dir", file_name, "/b/s").Output()

		for _, v := range strings.Split(string(out), "\n") {
			if strings.Contains(v, Config.YaraRuleIndex) {
				found_path = v
			}
		}
	}
	found_path = strings.Trim(found_path, " \t\r\n") // NOTE: widnows CRLF "\r\n"
	if len(found_path) > 0 {
		return found_path, nil
	} else {
		msg := fmt.Sprintf("file '%s' not found", file)
		return "", errors.New(msg)
	}
}

// func Where(bin_name string) (path string, err error) {
// 	out, err := exec.Command("cmd", "/C", "where", bin_name).Output()
// 	if err != nil {
// 		return "", err
// 	}
// 	found_path := strings.Split(string(out), "\n")[0]
// 	strings.Trim(found_path, " \t\r\n")
// 	return found_path, nil
// }
