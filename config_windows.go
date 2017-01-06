// +build windows

package main

import ()

func set_yara_bin_name() error {
	// TODO
	Config.YaraBinName = "yara64.exe"
	return nil
}

func set_line_break() error {
	Config.LineBreak = "\r\n"
	return nil
}
