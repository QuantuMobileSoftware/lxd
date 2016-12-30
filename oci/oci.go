package oci

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/lxc/lxd/shared"
	"github.com/pkg/errors"
)

const (
	ccOCIBinaryPath = "./oci/cc-oci-runtime"

	logPath         = "/tmp/log.json"
	pidFile = "/tmp/pid"

	defaultConsoleName = "/dev/console"
)

var (
	defaultArgs = []string{
		"--log-format", "json",
		"--log", logPath,
		"--debug",
	}
)

func ccOCICmd(args ...string) *exec.Cmd {
	return exec.Command(
		ccOCIBinaryPath,
		append(defaultArgs, args...)...,
	)
}

func Create(name, bundlePath string) error {
	cmd := ccOCICmd(
		"create",
		"--bundle", bundlePath,
		"--pid-file", pidFile,
		"--console", defaultConsoleName,
		name)
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "%v create execution problem", ccOCIBinaryPath)
	}

	return nil
}

func Delete(name string) error {
	cmd := ccOCICmd(
		"delete",
		name)
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "%v delete execution problem", ccOCIBinaryPath)
	}

	return nil
}

func Start() {

}

func List() ([]shared.ContainerInfo, error) {
	cmd := ccOCICmd("list")
	reader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrapf(err, "%v list execution problem", ccOCIBinaryPath)
	}
	if err := cmd.Start(); err != nil {
		return nil, errors.Wrapf(err, "%v list execution problem", ccOCIBinaryPath)
	}
	scanner := bufio.NewScanner(reader)
	var results []shared.ContainerInfo
	isFirstLine := true
	for scanner.Scan() {
		if isFirstLine {
			isFirstLine = false
			continue
		}
		text := scanner.Text()
		container, err := shared.ParseContainerInfo(text)
		if err != nil {
			return nil, errors.Wrapf(err, "Error while parsing container string: %v", text)
		}
		results = append(results, container)
	}
	if err := cmd.Wait(); err != nil {
		return nil, errors.Wrapf(err, "%v list execution problem", ccOCIBinaryPath)
	}
	return results, nil
}
