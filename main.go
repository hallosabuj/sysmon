package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

const (
	ProtocLinux   = "https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-linux-x86_64.zip"
	ProtocWindows = "https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-win64.zip"
	ProtocDarwin  = "https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-osx-x86_64.zip"
	GoogleAPIs    = "https://github.com/googleapis/googleapis/archive/master.zip"
)

func Unzip(source, destination, pattern string) error {
	data, err := ioutil.ReadFile(source)
	if nil != err {
		return err
	}
	_ = os.Remove(destination)
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if nil != err {
		return err
	}
	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		_, name := filepath.Split(file.Name)
		if matched, err := filepath.Match(pattern, name); nil == err {
			if !matched {
				continue
			}
		} else {
			return err
		}
		err := func() error {
			reader, err := file.Open()
			if err != nil {
				return err
			}
			defer reader.Close()
			path := filepath.Join(destination, file.Name)
			_ = os.Remove(path)
			err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
			if err != nil {
				return err
			}
			writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer writer.Close()
			_, err = io.Copy(writer, reader)
			if err != nil {
				return err
			}
			return nil
		}()
		if nil != err {
			return err
		}
	}
	return nil
}

func DownloadFile(url, path string) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if nil != err {
		return err
	}
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func FilenameFromURL(value string) (string, error) {
	uri, err := url.Parse(value)
	if nil != err {
		return "", err
	}
	_, name := filepath.Split(uri.Path)
	return name, nil
}

func DownloadProtoc() error {
	dir, err := os.Getwd()
	if nil != err {
		return err
	}
	if err := os.MkdirAll(filepath.Join(dir, "tools"), os.ModePerm); nil != err {
		return err
	}
	if runtime.GOOS == "windows" {
		if filename, err := FilenameFromURL(ProtocWindows); nil == err {
			if _, err := os.Stat(filepath.Join(dir, "tools", filename)); os.IsNotExist(err) {
				err = DownloadFile(ProtocWindows, filepath.Join(dir, "tools", filename))
				if nil != err {
					return err
				}
			}
			err = Unzip(
				filepath.Join(dir, "tools", filename),
				filepath.Join(dir, "tools", "protoc"),
				"*",
			)
			if nil != err {
				return err
			}
		} else {
			return err
		}
	}
	if runtime.GOOS == "linux" {
		if filename, err := FilenameFromURL(ProtocLinux); nil == err {
			if _, err := os.Stat(filepath.Join(dir, "tools", filename)); os.IsNotExist(err) {
				err = DownloadFile(ProtocLinux, filepath.Join(dir, "tools", filename))
				if nil != err {
					return err
				}
			}
			err = Unzip(
				filepath.Join(dir, "tools", filename),
				filepath.Join(dir, "tools", "protoc"),
				"*",
			)
			if nil != err {
				return err
			}
		} else {
			return err
		}
	}
	if runtime.GOOS == "darwin" {
		if filename, err := FilenameFromURL(ProtocDarwin); nil == err {
			if _, err := os.Stat(filepath.Join(dir, "tools", filename)); os.IsNotExist(err) {
				err = DownloadFile(ProtocDarwin, filepath.Join(dir, "tools", filename))
				if nil != err {
					return err
				}
			}
			err = Unzip(
				filepath.Join(dir, "tools", filename),
				filepath.Join(dir, "tools", "protoc"),
				"*",
			)
			if nil != err {
				return err
			}
		} else {
			return err
		}
	}
	{
		if _, err := os.Stat(filepath.Join(dir, "tools", "googleapis.zip")); os.IsNotExist(err) {
			err = DownloadFile(GoogleAPIs, filepath.Join(dir, "tools", "googleapis.zip"))
			if nil != err {
				return err
			}
		}
		err = Unzip(
			filepath.Join(dir, "tools", "googleapis.zip"),
			filepath.Join(dir, "tools", "googleapis"),
			"*.proto",
		)
		if nil != err {
			return err
		}
	}
	return nil
}

func GoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

func ProtoGenerate() {
	fmt.Println("Generating code for protocol buffers ")
	{
		err := DownloadProtoc()
		if nil != err {
			fmt.Println("Failed downloading protoc binaries")
			fmt.Println("Error: ", err)
			return
		}
	}
	{
		files := make([]string, 0)
		if entries, err := ioutil.ReadDir(filepath.Join("proto")); nil == err {
			for _, entry := range entries {
				if match, err := filepath.Match("*.proto", entry.Name()); nil == err {
					if match {
						files = append(files, entry.Name())
					}
				}
			}
		}
		for _, file := range files {
			fmt.Println("Compiling: ", file)
			name := strings.TrimSuffix(file, filepath.Ext(file))
			_ = os.MkdirAll(filepath.Join("proto", fmt.Sprintf("%spb", name)), os.ModePerm)
			arguments := strings.Join(
				[]string{
					"--proto_path=proto",
					fmt.Sprintf("-I %s", filepath.Join("tools", "protoc", "include")),
					fmt.Sprintf("-I %s", filepath.Join("tools", "googleapis", "googleapis-master")),
					fmt.Sprintf("--go_out=proto/%spb --go_opt=paths=source_relative", name),
					fmt.Sprintf("--go-grpc_out=proto/%spb --go-grpc_opt=paths=source_relative", name),
					fmt.Sprintf("--grpc-gateway_out=proto/%spb --grpc-gateway_opt=paths=source_relative", name),
					filepath.Join("proto", file),
				},
				" ",
			)
			cmd := exec.Command(
				filepath.Join("tools", "protoc", "bin", "protoc"),
				strings.Split(arguments, " ")...,
			)
			cmd.Env = os.Environ()
			cmd.Env = append(cmd.Env, fmt.Sprintf("PATH=%s", strings.Join(
				[]string{
					os.Getenv("PATH"),
					GoPath(),
					filepath.Join(GoPath(), "bin"),
				},
				string(os.PathListSeparator),
			)))
			output, err := cmd.CombinedOutput()
			if nil != err {
				fmt.Println("Error: ", err)
				fmt.Println(string(output))
			}
		}
	}
}

func BuildWeb() error {
	fmt.Println("Building ui ...")
	_ = os.RemoveAll("./web/build")
	cmd := exec.Command("npm", strings.Split("run build", " ")...)
	cmd.Dir = filepath.Join("web")
	output, err := cmd.CombinedOutput()
	if nil != err {
		fmt.Println(string(output))
	}
	return err
}

func Build() {
	if err := BuildWeb(); nil != err {
		fmt.Println("Error: ", err)
		return
	}

	cmd := exec.Command(
		filepath.Join("go"),
		strings.Split("build -o bin/sysmon.bin cmd/sysmon/main.go", " ")...,
	)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, strings.Join([]string{"GOARCH", "amd64"}, "="))
	cmd.Env = append(cmd.Env, strings.Join([]string{"CGO_ENABLED", "0"}, "="))
	output, err := cmd.CombinedOutput()
	if nil != err {
		fmt.Println("Error: ", err)
		fmt.Println(string(output))
	}
}

func main() {
	ProtoGenerate()
	Build()
}
