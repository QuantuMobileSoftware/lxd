package main

import (
	"time"
	"io"
	"os"
	"github.com/lxc/lxd/shared"
	"github.com/lxc/lxd/lxd/oci"
	"github.com/pkg/errors"
	"gopkg.in/lxc/go-lxc.v2"
)

type OCIContainer struct {
	// Properties
	architecture int
	cType        containerType
	creationDate time.Time
	lastUsedDate time.Time
	ephemeral    bool
	id           string
	name         string
	stateful     bool

	// Config
	expandedConfig  map[string]string
	expandedDevices shared.Devices
	fromHook        bool
	localConfig     map[string]string
	localDevices    shared.Devices
	profiles        []string

	// Cache
	c        *lxc.Container
	daemon   *Daemon
	idmapset *shared.IdmapSet
	storage  storage
}

func containerOCILoad(d *Daemon, cname string) (container, error) {
	state, err := oci.State(cname)
	if err != nil {
		return nil, errors.Wrapf(err, "Error while attempting to get container %v info", cname)
	}
	return OCIContainer{
		name: state.ID,
	}, nil
}

func (OCIContainer) Freeze() error {
	panic("implement me")
}

func (OCIContainer) Shutdown(timeout time.Duration) error {
	panic("implement me")
}

func (OCIContainer) Start(stateful bool) error {
	panic("implement me")
}

func (OCIContainer) Stop(stateful bool) error {
	panic("implement me")
}

func (OCIContainer) Unfreeze() error {
	panic("implement me")
}

func (OCIContainer) Restore(sourceContainer container) error {
	panic("implement me")
}

func (OCIContainer) Migrate(cmd uint, stateDir string, function string, stop bool, actionScript bool) error {
	panic("implement me")
}

func (OCIContainer) Snapshots() ([]container, error) {
	panic("implement me")
}

func (OCIContainer) Rename(newName string) error {
	panic("implement me")
}

func (OCIContainer) Update(newConfig containerArgs, userRequested bool) error {
	panic("implement me")
}

func (OCIContainer) Delete() error {
	panic("implement me")
}

func (OCIContainer) Export(w io.Writer, properties map[string]string) error {
	panic("implement me")
}

func (OCIContainer) CGroupGet(key string) (string, error) {
	panic("implement me")
}

func (OCIContainer) CGroupSet(key string, value string) error {
	panic("implement me")
}

func (OCIContainer) ConfigKeySet(key string, value string) error {
	panic("implement me")
}

func (OCIContainer) FileExists(path string) error {
	panic("implement me")
}

func (OCIContainer) FilePull(srcpath string, dstpath string) (int, int, os.FileMode, string, []string, error) {
	panic("implement me")
}

func (OCIContainer) FilePush(srcpath string, dstpath string, uid int, gid int, mode int) error {
	panic("implement me")
}

func (OCIContainer) FileRemove(path string) error {
	panic("implement me")
}

func (OCIContainer) Exec(command []string, env map[string]string, stdin *os.File, stdout *os.File, stderr *os.File, wait bool) (int, int, error) {
	panic("implement me")
}

func (OCIContainer) Render() (interface{}, interface{}, error) {
	panic("implement me")
}

func (OCIContainer) RenderState() (*shared.ContainerState, error) {
	panic("implement me")
}

func (OCIContainer) IsPrivileged() bool {
	panic("implement me")
}

func (OCIContainer) IsRunning() bool {
	panic("implement me")
}

func (OCIContainer) IsFrozen() bool {
	panic("implement me")
}

func (OCIContainer) IsEphemeral() bool {
	panic("implement me")
}

func (OCIContainer) IsSnapshot() bool {
	panic("implement me")
}

func (OCIContainer) IsStateful() bool {
	panic("implement me")
}

func (OCIContainer) IsNesting() bool {
	panic("implement me")
}

func (OCIContainer) OnStart() error {
	panic("implement me")
}

func (OCIContainer) OnStop(target string) error {
	panic("implement me")
}

func (OCIContainer) Id() int {
	panic("implement me")
}

func (OCIContainer) Name() string {
	panic("implement me")
}

func (OCIContainer) Architecture() int {
	panic("implement me")
}

func (OCIContainer) CreationDate() time.Time {
	panic("implement me")
}

func (OCIContainer) LastUsedDate() time.Time {
	panic("implement me")
}

func (OCIContainer) ExpandedConfig() map[string]string {
	panic("implement me")
}

func (OCIContainer) ExpandedDevices() shared.Devices {
	panic("implement me")
}

func (OCIContainer) LocalConfig() map[string]string {
	panic("implement me")
}

func (OCIContainer) LocalDevices() shared.Devices {
	panic("implement me")
}

func (OCIContainer) Profiles() []string {
	panic("implement me")
}

func (OCIContainer) InitPID() int {
	panic("implement me")
}

func (OCIContainer) State() string {
	panic("implement me")
}

func (OCIContainer) Path() string {
	panic("implement me")
}

func (OCIContainer) RootfsPath() string {
	panic("implement me")
}

func (OCIContainer) TemplatesPath() string {
	panic("implement me")
}

func (OCIContainer) StatePath() string {
	panic("implement me")
}

func (OCIContainer) LogFilePath() string {
	panic("implement me")
}

func (OCIContainer) LogPath() string {
	panic("implement me")
}

func (OCIContainer) StorageStart() error {
	panic("implement me")
}

func (OCIContainer) StorageStop() error {
	panic("implement me")
}

func (OCIContainer) Storage() storage {
	panic("implement me")
}

func (OCIContainer) IdmapSet() *shared.IdmapSet {
	panic("implement me")
}

func (OCIContainer) LastIdmapSet() (*shared.IdmapSet, error) {
	panic("implement me")
}

func (OCIContainer) TemplateApply(trigger string) error {
	panic("implement me")
}

func (OCIContainer) Daemon() *Daemon {
	panic("implement me")
}
