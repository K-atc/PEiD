package main

import (
	"github.com/Sirupsen/logrus"
	// [ホームディレクトリを取得するのにos/userを使うとエラーになる場合がある - Qiita](http://qiita.com/hironobu_s/items/da2f97c2154075d3fbbe)
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

var Config struct {
	HomeDir       string
	YaraRulesPath string
	YaraRuleIndex string
	YaraBinName   string
	LineBreak string
}

func set_home_dir() error {
	hdir, err := homedir.Dir()
	if err != nil {
		return err
	}
	Config.HomeDir = hdir
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
