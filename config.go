package main

import (
	"github.com/Sirupsen/logrus"
	// [ホームディレクトリを取得するのにos/userを使うとエラーになる場合がある - Qiita](http://qiita.com/hironobu_s/items/da2f97c2154075d3fbbe)
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

var Config struct {
	HomeDir       string
	YaraRulesPath string
	YaraRuleIndex string
	YaraBinName   string
	YaraBinPath   string
	LineBreak     string
}

func set_home_dir() error {
	hdir, err := homedir.Dir()
	if err != nil {
		return err
	}
	Config.HomeDir = hdir
	return nil
}

func set_yara_bin_path() error {
	found_path, err := Where(Config.YaraBinName)
	if err != nil {
		return err
	}
	if len(found_path) > 0 {
		Config.YaraBinPath = found_path
	} else {
		msg := fmt.Sprintf("Cannot find '%s'", Config.YaraBinName)
		return errors.New(msg)
	}
	return nil
}

func set_yara_rules_path() error {
	found_path, err := Find(Config.YaraRuleIndex)
	if err != nil {
		return err
	}
	if len(found_path) > 0 {
		Config.YaraRulesPath = found_path
	} else {
		return errors.New("Cannot find YaraRules")
	}
	return nil
}

func set_yara_rule_index() error {
	Config.YaraRuleIndex = filepath.Join("rules", "index.yar")
	return nil
}

func Configure() {
	var err error
	err = set_home_dir()
	if err != nil {
		logrus.Fatal(err)
	}
	err = set_yara_bin_name()
	if err != nil {
		logrus.Fatal(err)
	}
	err = set_yara_bin_path()
	if err != nil {
		logrus.Info("try run with --prepare option")
		logrus.Fatal(err)
	}
	msg := fmt.Sprintf("yara = '%s'", Config.YaraBinPath)
	logrus.Info(msg)
	err = set_yara_rule_index()
	if err != nil {
		logrus.Fatal(err)
	}
	err = set_yara_rules_path()
	if err != nil {
		logrus.Fatal(err)
	}
	err = set_line_break()
	if err != nil {
		logrus.Fatal(err)
	}
}
