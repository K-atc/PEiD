package main

import (
	"os/exec"
	// "strings"
)

func check_if_command_exists(cmd_name string, opt_version string) bool {
	cmd := exec.Command(cmd_name, opt_version)
	err := cmd.Start()
	if err != nil {
		return false
	}
	return true
}

// func go_find(base_dir string) (found_path string) {

//     var files = []string{}

// 	// quit = make(chan bool)
// 	go func() {
// 		filepath.Walk(base_dir, func(path string, info os.FileInfo, err error) error {
// 			if terminating {
// 				return errors.New("terminate")
// 			}
// 			if info == nil {
// 				return err
// 			}
// 			if !info.IsDir() {
// 				if p, err := filepath.Rel(cwd, path); err == nil {
// 					path = p
// 				}
// 				path = filepath.ToSlash(path)
// 				mutex.Lock()
// 				files = append(files, path)
// 				mutex.Unlock()
// 				dirty = true
// 				timer.Reset(duration)
// 			} else if strings.HasPrefix(info.Name(), ".") {
// 				return filepath.SkipDir
// 			}
// 			return nil
// 		})
// 		scanning = -1
// 		// quit <- true
// 	}()
//     for _, v := range files {
//         if strings.Contains(v, "")
//     }
//     return files[0]
// }
