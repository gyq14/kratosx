package base

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

// ModulePath returns go module path.
func ModulePath(filename string) (string, error) {
	modBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return modfile.ModulePath(modBytes), nil
}

// ModuleVersion returns module version.
func ModuleVersion(path string) (string, error) {
	stdout := &bytes.Buffer{}
	fd := exec.Command("go", "mod", "graph")
	fd.Stdout = stdout
	fd.Stderr = stdout
	if err := fd.Run(); err != nil {
		return "", err
	}
	rd := bufio.NewReader(stdout)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			return "", err
		}
		str := string(line)
		i := strings.Index(str, "@")
		if strings.Contains(str, path+"@") && i != -1 {
			return path + str[i:], nil
		}
	}
}

// KratosMod returns kratos mod.
func KratosMod() string {
	// go 1.15+ read from env GOMODCACHE
	cacheOut, _ := exec.Command("go", "env", "GOMODCACHE").Output()
	cachePath := strings.Trim(string(cacheOut), "\n")
	pathOut, _ := exec.Command("go", "env", "GOPATH").Output()
	gopath := strings.Trim(string(pathOut), "\n")
	if cachePath == "" {
		cachePath = filepath.Join(gopath, "pkg", "mod")
	}
	if path, err := ModuleVersion("github.com/go-kratos/kratos/v2"); err == nil {
		return filepath.Join(cachePath, path)
	}
	return filepath.Join(gopath, "src", "github.com", "go-kratos", "kratos")
}

// KratosxCliMod returns kratos mod.
func KratosxCliMod() string {
	path := os.Getenv("AUTOCODE_TEMP_PATH")
	if path != "" {
		return path
	}
	// go 1.15+ read from env GOMODCACHE
	cacheOut, _ := exec.Command("go", "env", "GOMODCACHE").Output()
	cachePath := strings.Trim(string(cacheOut), "\n")
	pathOut, _ := exec.Command("go", "env", "GOPATH").Output()
	gopath := strings.Trim(string(pathOut), "\n")
	if cachePath == "" {
		cachePath = filepath.Join(gopath, "pkg", "mod")
	}
	cliPath := cachePath + "/kratosx/cmd"

	files, err := os.ReadDir(cliPath)
	if err != nil {
		return filepath.Join(gopath, "src", "github.com", "gyq14", "kratosx", "cmd", "kratosx")
	}

	var lastKratosxDir string
	for _, file := range files {
		if file.IsDir() {
			dirName := file.Name()
			if strings.HasPrefix(dirName, "kratosx@") {
				lastKratosxDir = dirName
			}
		}
	}

	return cliPath + "/" + lastKratosxDir
}
