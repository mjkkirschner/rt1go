package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Println("setup vs")
	//setup visual studio vars for cuda compiler
	varsPath := `"C:\\Program Files (x86)\\Microsoft Visual Studio\\2019\\Community\\VC\\Auxiliary\\Build\\vcvars64.bat"`
	nvccCommand := "nvcc render.cu -o render.dll --shared"
	cmd := exec.Command("cmd")
	cmd.Dir = "./cuda"
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/c " + varsPath + " && " + nvccCommand}
	for _, v := range cmd.Args {
		fmt.Println(v)
	}
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))

	//copy the generated dll to root.
	CopyFile("./cuda/render.dll", "./render.dll")

	//build rtCLI
	compileCMD := exec.Command("go", "run", "./rtCLI/main.go")
	output, _ = compileCMD.CombinedOutput()
	fmt.Println(string(output))
}

func CopyFile(srcPath string, destPath string) {

	srcFile, err := os.Open(srcPath)
	check(err)
	defer srcFile.Close()

	destFile, err := os.Create(destPath) // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	check(err)

	err = destFile.Sync()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}
