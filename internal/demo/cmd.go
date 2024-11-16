package demo

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/axgle/mahonia"
)

func Cmd() {
	cmd := exec.Command("cmd")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in //绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出

	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr

	go func() {
		in.WriteString("/C svn log")
	}()
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	enc := mahonia.NewEncoder("gbk")

	log.Println(enc.ConvertString(out.String()))
	log.Println(stdErr.String())
}
