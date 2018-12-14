package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
)

func currPath() string {
	file, _ := exec.LookPath(os.Args[0])

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	rst := filepath.Dir(path)

	return rst
}

func clearTempFile(path string) {
	fileExtNames := []string{".ddp", ".dcu", ".tmp", ".~ddp", ".~dcu", ".elf", ".dproj.local", ".map", ".tvsconfig", ".identcache"}
	exp := []string{"ClearDelphi.exe"}
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return nil
		}
		if f.IsDir() {
			clearTempFile(path + f.Name())
			return nil
		}
		for _, str := range fileExtNames {
			if strings.HasSuffix(path, str) {
				var dd bool = false
				for _, ex := range exp {
					if ex == f.Name() {
						dd = true
						break
					}
				}
				if !dd {
					color.LightWhite.Printf(path + "\n")
					os.Remove(path)
				}
				//exec.Command("del " + fn)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	color.LightYellow.Printf("DELPHI开发环境临时文件清理工具 V1.1\n")
	color.LightYellow.Printf("Design by suikhan@126.com with Golang\n")
	color.Red.Printf(strings.Repeat("=", 50) + "\n")

	root := currPath()
	color.Red.Println("当前路径：" + root)
	clearTempFile(root)
	color.Red.Printf(strings.Repeat("=", 50) + "\n")
	color.LightYellow.Printf("清理完毕，按任意键返回...\n")
	fmt.Scanf("%s")
}
