package main

import "fmt"
import "strings"
import "jvmgo/ch02/classpath"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classjre:%v classpath:%v class:%v args:%v\n",
		cmd.XjreOption, cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	fmt.Printf("className:%v\n", className)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
