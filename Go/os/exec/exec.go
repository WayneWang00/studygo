package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	//goVersion()
	//pwd()
	//ls()
	//testStdoutPipe()
	testCommand()
}

func goVersion() {
	cmd := exec.Command("go", "version")
	fmt.Println(cmd.Path, cmd.Args)
}

func pwd() {
	cmd := exec.Command("pwd")
	fmt.Println(cmd.Path, cmd.Args)
}

func ls() {
	cmd := exec.Command("ls", "-lah")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}

	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("cmd.Run() failed:", err)
	//}

	//bts, err := cmd.CombinedOutput()
	//if err != nil {
	//	fmt.Println("cmd.Run failed:", err)
	//	return
	//}
	//fmt.Println("out:", string(bts))

	//var stdout, stderr bytes.Buffer
	//cmd.Stdout = &stdout
	//cmd.Stderr = &stderr
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("cmd.Run failed:", err)
	//	return
	//}
	//fmt.Printf("stdout:%s\nstderr:%s\n", stdout.String(), string(stderr.Bytes()))

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("stdoutPipe failed:", err)
		return
	}
	defer stdout.Close()

	if err = cmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		return
	}
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("readAll failed:", err)
		return
	}
	fmt.Println(string(content))
}

func testStdoutPipe() {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("stdoutPipe failed:", err)
		return
	}
	if err = cmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		return
	}

	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		fmt.Println("decode failed:", err)
		return
	}
	if err = cmd.Wait(); err != nil {
		fmt.Println("wait failed:", err)
		return
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

func testCommand() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println("run failed:", err)
		return
	}
	fmt.Println("out:", out.String())
}
