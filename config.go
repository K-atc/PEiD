package main

import (
    "os/exec"
    // [ホームディレクトリを取得するのにos/userを使うとエラーになる場合がある - Qiita](http://qiita.com/hironobu_s/items/da2f97c2154075d3fbbe)
    "github.com/mitchellh/go-homedir"
    "github.com/Sirupsen/logrus"
    // "fmt"    
    "strings"
    "errors"
)

var Config struct {
    YaraRulesPath string
}

func set_yara_rules_path() error {
    hdir, err := homedir.Dir()
    if err != nil {
        return errors.New("cannot get HOME value")
    }
    // out, err := exec.Command("find", hdir, "-name", "index.yar", "|", "grep", "rules/index.yar").Output()
    out, err := exec.Command("find", hdir, "-name", "index.yar").Output()
    if err != nil {
        logrus.Warn("exec find error")
    }
    var found_path string
    for _, v := range strings.Split(string(out), "\n") {
        if strings.Contains(v, "rules/index.yar") {
            found_path = v
        }
    }
    if len(found_path) > 0 {
        Config.YaraRulesPath = found_path
    } else {
        return errors.New("Cannot find YaraRules")
    }
    return nil
}

func Configure() {
    var err error
    err = set_yara_rules_path()
    if err != nil {
        logrus.Fatal(err)
    }
}