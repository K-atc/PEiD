// +build: linux

package main

import ()

func set_yara_bin_name() error {
	Config.YaraBinName = "yara"
	return nil
}

func set_line_break() error {
	Config.LineBreak = "\n"
	return nil
}
