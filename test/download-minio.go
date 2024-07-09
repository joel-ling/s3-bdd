package test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func downloadMinio() (binPath string, e error) {
	const (
		errFmt = "File ‘%s’ already there; not retrieving.\n"
	)

	var (
		stderr bytes.Buffer
		wget   *exec.Cmd
	)

	binPath = filepath.Join(os.TempDir(), "minio")

	wget = exec.Command("wget",
		"-O", binPath,
		"-nc", // skips download if file exists, but exits with status 1!
		fmt.Sprintf("https://dl.min.io/server/minio/release/%s-%s/minio",
			runtime.GOOS,
			runtime.GOARCH,
		),
	)

	wget.Stderr = &stderr

	e = wget.Run()
	if e != nil {
		if stderr.String() != fmt.Sprintf(errFmt, binPath) {
			return
		}

		e = nil
	}

	e = os.Chmod(binPath, 0755)
	if e != nil {
		return
	}

	return
}
