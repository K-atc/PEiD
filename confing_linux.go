// +build: linux

package main

import (
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"os/exec"
	"strings"
)

func set_yara_bin_name() error {
	Config.YaraBinName = "yara"
	return nil
}

func set_yara_rules_path() error {
	var out []byte
	var err error

	need_find := false
	if check_if_command_exists("locate", "-v") {
		out, err = exec.Command("locate", "rules/index.yar").Output()
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
		out, err = exec.Command("find", Config.HomeDir, "-name", "index.yar").Output()
		if err != nil {
			fmt.Println(err)
			logrus.Warn("exec find error")
		}
	}

	var found_path string
	for _, v := range strings.Split(string(out), "\n") {
		if strings.Contains(v, Config.YaraRuleIndex) {
			found_path = v
		}
	}
	found_path = strings.Trim(found_path, " \t\r\n")
	if len(found_path) > 0 {
		Config.YaraRulesPath = found_path
	} else {
		return errors.New("Cannot find YaraRules")
	}
	return nil
}

func set_line_break() error {
	Config.LineBreak = "\n"
	return nil
}
