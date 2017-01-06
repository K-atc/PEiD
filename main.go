package main

import (
	"flag"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var opt struct {
	verbose bool
	prepare bool
	version bool
}

var (
	version string
)

func check_requirements() bool {
	var all_met = true
	var need []string // required commands
	if runtime.GOOS == "linux" {
		need = []string{"find"}
	} else if runtime.GOOS == "windows" {
		need = []string{"cmd"}
	}
	for _, v := range need {
		res := check_if_command_exists(v, "-v")
		if res == false {
			msg := fmt.Sprintf("command '%s' not found", v)
			logrus.Warn(msg)
			all_met = false
		}
	}
	return all_met
}

func usage() {
	msg := fmt.Sprintf("usage: %s [OPTIONS] FILE", os.Args[0])
	fmt.Println(msg)
	flag.PrintDefaults()
	os.Exit(1)
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
	matched_tags  []string
	matched_rules []string
}

func parse_line(line string) (rule_name, tag_name, file_name string, parse_status bool) {
	res := strings.SplitN(line, " ", 3)
	if len(res) != 3 {
		return "", "", "", false
	}
	rule_name = res[0]
	tag_name = strings.Trim(res[1], "[]")
	file_name = strings.Trim(res[2], "\r")
	return rule_name, tag_name, file_name, true
}

func do_exam(out string) []YaraRecord {
	var result []YaraRecord
	var matched_rules []string
	var matched_tags []string
	var file_name string
	prev_file_name := ""
	for _, line := range strings.Split(out, "\n") {
		if rule_name, tag_name, file_name, ok := parse_line(line); ok {
			if prev_file_name == file_name {
				matched_rules = append(matched_rules, rule_name)
				matched_tags = append(matched_tags, tag_name)
			} else {
				if prev_file_name != "" {
					result = append(result, YaraRecord{file_name: file_name, matched_tags: matched_tags, matched_rules: matched_rules})
				}
				matched_rules = append(make([]string, 0), rule_name)
				matched_tags = append(make([]string, 0), tag_name)
				prev_file_name = file_name
			}
		}
	}
	file_name = prev_file_name // restore
	result = append(result, YaraRecord{file_name: file_name, matched_tags: matched_tags, matched_rules: matched_rules})
	for _, v := range result {
		msg := fmt.Sprintf("%s =>%s", v.file_name, Config.LineBreak)
		if opt.verbose {
			msg += fmt.Sprintf("%q%s%q%s", v.matched_rules, Config.LineBreak, v.matched_tags, Config.LineBreak)
		}
		if comment, ok := add_comment(v); ok {
			msg += fmt.Sprintf("%s", comment)
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
		cmd = exec.Command(Config.YaraBinPath, "-w", "-f", "-g", "-p", "2", "-r", RULES_FILE, file)
	} else {
		cmd = exec.Command(Config.YaraBinPath, "-w", "-f", "-g", RULES_FILE, file)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warn("yara execution error")
		fmt.Println(string(out))
		return
	}
	do_exam(string(out))
}

func show_version() {
	var msg string
	msg = fmt.Sprintf("PEiD - Yet another implementation of PEiD with yara")
	fmt.Println(msg)
	msg = fmt.Sprintf("  version: %s", version)
	fmt.Println(msg)
}

func main() {
	if opt.prepare {
		err := prepare()
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Info("prepare successfuly")
	} else if opt.version {
		show_version()
	} else {
		Configure()

		if !check_requirements() {
			return
		} else {
			logrus.Info("all requirements met")
		}
		if len(flag.Args()) != 1 {
			usage()
		} else {
			file := flag.Args()[0]
			Examine(file)
		}
	}
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	flag.BoolVar(&opt.version, "version", false, "version info")
	flag.BoolVar(&opt.prepare, "prepare", false, "prepare files to meet requirements")
	flag.BoolVar(&opt.verbose, "verbose", false, "verbose output")
	flag.BoolVar(&opt.verbose, "v", false, "verbose output")
	flag.Usage = func() {
		usage()
	}
	flag.Parse()
}
