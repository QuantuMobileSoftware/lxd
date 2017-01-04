package oci

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"encoding/json"
)

const (
	ccOCIBinaryPath = "cc-oci-runtime"

	logPath         = "/tmp/log.json"
	pidFile = "/tmp/pid"

	defaultConsoleName = "/dev/console"

	containerInfoTimeFormat = "2006-01-02T15:04:05.999999999Z"
)

var (
	defaultArgs = []string{
		"--log-format", "json",
		"--log", logPath,
		"--debug",
	}
)

type ContainerInfo struct {
	Version string `json:"ociVersion"`
	ID string `json:"id"`
	PID int `json:"pid"`
	BundlePath string `json:"bundlePath"`
	CommsPath string `json:"commsPath"`
	ProcessPath string `json:"processPath"`
	Status string `json:"status"`
	Created string `json:"created"`
	Mounts []struct{
		Destination string `json:"destination"`
	} `json:"Mounts"`
	Console struct{
		Socket bool `json:"socket"`
		Path string `json:"path"`
	} `json:"console"`
	VM struct {
		HypervisorPath string `json:"hypervisor_path"`
		ImagePath string `json:"image_path"`
		KernelPath string `json:"kernel_path"`
		WorkloadPath string `json:"workload_path"`
		KernelParams string `json:"kernel_params"`
	} `json:"vm"`
}

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
		return errors.Wrapf(err, "Can't create container %v with bundle %v", name, bundlePath)
	}

	return nil
}

func Start(name string) error {
	cmd := ccOCICmd("start", name)
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "Can't start container %v", name)
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



func State(cname string) (ContainerInfo, error) {
	cmd := ccOCICmd("list", cname)
	data, err := cmd.Output()
	if err != nil {
		return ContainerInfo{}, errors.Wrapf(err, "Error while getting state info for %v", cname)
	}
	result := &ContainerInfo{}

	if err := json.Unmarshal(data, result); err != nil {
		return ContainerInfo{}, errors.Wrapf(err, "Error while getting state info for %v", cname)
	}

	return *result, nil
}
