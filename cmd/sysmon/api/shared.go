package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

type Response struct {
	Msg string `json:"msg"`
}

func MakeSudo() {
	user := os.Getenv("USER")
	password, _ := ioutil.ReadFile("/home/" + user + "/password")
	c1 := exec.Command("echo", string(password))
	c2 := exec.Command("sudo", "-S", "ls")

	var stderr bytes.Buffer
	read_end, write_end := io.Pipe()
	c1.Stdout = write_end
	c2.Stdin = read_end
	c2.Stderr = &stderr
	c1.Start()

	go func() {
		defer write_end.Close()
		c1.Wait()
	}()

	err := c2.Run()

	if err != nil {
		fmt.Println(stderr)
	}
}
