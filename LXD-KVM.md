# LXD-KVM project

## Usage

### Building
Build the whole thing with `make`.
Binaries are stored in `$GOPATH/bin`, both lxc and lxd.

### Running
Make sure that your global lxd installation is not running, then lxd binary.
Then run lxc:
```
sudo $GOPATH/bin/lxc launch <full path to image> --kvm
```

## Additional topics

### cc-oci-runtime overview
You can check relations between cc-oci-runtime and qemu (qemu-lite) here:
https://github.com/01org/cc-oci-runtime/blob/master/DESIGN.rst#112code-flow-with-docker-112

### Running cc-oci-runtime
Instructions for both running cc-oci-runtime standalone and with Docker are provided in cc-oci-runtime README:
https://github.com/01org/cc-oci-runtime#110running-stand-alone

### Running qemu (qemu-lite)
Use corresponding qemu binary according to your environment and architecture. 
```
<qemu-binary> -kernel "/boot/vmlinuz-$(uname -r)" \
  -initrd "/boot/initrd.img-$(uname -r)" \
  -fsdev local,id=r,path=<path to filesystem>,security_model=none \
  -device virtio-9p-pci,fsdev=r,mount_tag=r \
  -nographic \
  -append 'root=r rw rootfstype=9p rootflags=trans=virtio console=ttyS0 init=/bin/sh' \
  -m 1024
```

### Running in clearcontainers/ubuntu
Image with all environment set up for cc-oci-runtime and set as Docker runtime. 
Instructions for running images within this image are also provided there:
https://hub.docker.com/r/clearcontainers/ubuntu/
