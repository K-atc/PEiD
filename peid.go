package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
	"os"
	"os/exec"
	// [ホームディレクトリを取得するのにos/userを使うとエラーになる場合がある - Qiita](http://qiita.com/hironobu_s/items/da2f97c2154075d3fbbe)
	// "github.com/mitchellh/go-homedir"
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

func check_requirements() bool {
	var all_met = true
	need := []string{"git", "yara"} // required commands
	for _, v := range need {
		res := check_if_command_exists(v, "-v")
		if res == false {
			msg := fmt.Sprintf("command '%s' not exists. check your PATH", v)
			logrus.Warn(msg)
			all_met = false
		}
	}
	return all_met
}

func usage() {
	fmt.Println("usage: %s FILE", os.Args[0])
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

type YaraRecord struct {
	file_name     string
	matched_rules []string
}

func parse_line(line string) (string, string, bool) {
	res := strings.SplitN(line, " ", 2)
	if len(res) != 2 {
		return "", "", false
	}
	return res[0], res[1], true
}

func do_exam(out string) []YaraRecord {
	var result []YaraRecord
	var matched_rules []string
	var file_name string
	prev_file_name := ""
	for _, line := range strings.Split(out, "\n") {
		if rule_name, file_name, ok := parse_line(line); ok {
			if prev_file_name == file_name {
				matched_rules = append(matched_rules, rule_name)
			} else {
				if prev_file_name != "" {
					result = append(result, YaraRecord{file_name: file_name, matched_rules: matched_rules})
				}
				matched_rules = append(make([]string, 0), rule_name)
				prev_file_name = file_name
			}
		}
	}
	file_name = prev_file_name // restore
	result = append(result, YaraRecord{file_name: file_name, matched_rules: matched_rules})
	for _, v := range result {
		var msg string
		if comment, ok := add_comment(v); ok {
			msg = fmt.Sprintf("%s =>\n%q\n%s", v.file_name, v.matched_rules, comment)
		} else {
			msg = fmt.Sprintf("%s =>\n%q\n", v.file_name, v.matched_rules)
		}
		fmt.Println(msg)
	}
	return result
}

func Examine(file string) {
	var err error
	var msg string

	if x_exists, _ := exists(file); !x_exists {
		msg = fmt.Sprintf("file '%s' not exists", file)
		logrus.Warn(msg)
		return
	}

	// other ways: http://tkuchiki.hatenablog.com/entry/2014/11/10/123447
	var cmd *exec.Cmd
    RULES_FILE := Config.YaraRulesPath
    fmt.Println("RULES_FILE = " + RULES_FILE)
	if fl, _ := os.Stat(file); fl.IsDir() {
		cmd = exec.Command("yara", "-w", "-r", RULES_FILE, file)
	} else {
		cmd = exec.Command("yara", "-w", RULES_FILE, file)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warn("yara execution error")
		fmt.Println(string(out))
		return
	}
	// fmt.Println(string(out))
	do_exam(string(out))
	// show_result(result)
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	if !check_requirements() {
		return
	} else {
		logrus.Info("all rewuirements met")
	}
    Configure()

	if len(os.Args) != 2 {
		usage()
	}
	file := os.Args[1]
	Examine(file)
}
