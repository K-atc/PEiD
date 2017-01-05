// +build windows

package main

import (
	"errors"
	// "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
)

func set_yara_bin_name() error {
	// TODO
	Config.YaraBinName = "yara64"
	return nil
}

func set_yara_rules_path() error {
	var cwd string
	var out []byte
	var err error

	cwd, err = os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(cwd)

	var found_path string
	for _, drive := range []string{"C:", "E:"} {
		err = os.Chdir(drive)
		if err != nil {
			return err
		}

		out, _ = exec.Command("cmd", "/C", "dir", "index.yar", "/b/s").Output()

		for _, v := range strings.Split(string(out), "\n") {
			if strings.Contains(v, Config.YaraRuleIndex) {
				found_path = v
			}
		}
	}
	found_path = strings.Trim(found_path, " \t\r\n") // NOTE: widnows CRLF "\r\n"

	if len(found_path) > 0 {
		Config.YaraRulesPath = found_path
	} else {
		return errors.New("Cannot find YaraRules")
	}
	return nil
}

func set_line_break() error {
	Config.LineBreak = "\r\n"
	return nil
}
