package bundle

import (
	"os"
	"path/filepath"
	"github.com/pkg/errors"
	"text/template"
)

type Config struct {
	BundlePath string
}

type BundleFile struct {
	template *template.Template
	filename string
	config Config
}

func (b *BundleFile) Generate(c Config) error {
	file, err := os.Create(filepath.Join(c.BundlePath, b.filename))
	defer file.Close()
	if err != nil {
		return errors.Wrapf(err, "Can't create file %v for bundle %v", b.filename, c.BundlePath)
	}
	if err = b.template.Execute(file, c); err != nil {
		return errors.Wrapf(err, "Can't generate %v for bundle %v", b.filename, c.BundlePath)
	}
	return nil
}

var (
	bundleFiles = []BundleFile {
		BundleFile{
			template: template.Must(template.New("config_template").Parse(configTemplateString)),
			filename: "config.json",
		},
		BundleFile{
			template: template.Must(template.New("hostname_template").Parse(hostnameTemplateString)),
			filename: "rootfs/etc/hostname",
		},
		BundleFile{
			template: template.Must(template.New("hosts_template").Parse(hostsTemplateString)),
			filename: "rootfs/etc/hosts",
		},
		BundleFile{
			template: template.Must(template.New("resolvconf_template").Parse(resolvconfTemplateString)),
			filename: "rootfs/etc/resolv.conf",
		},
	}

)



func GenerateBundleMetadata(bundlePath string) error {
	c := Config{
		BundlePath: bundlePath,
	}
	var err error
	for _, b := range bundleFiles {
		if err = b.Generate(c); err != nil {
			return errors.Wrapf(err, "Can't create %v for %v bundle", b.filename, bundlePath)
		}
	}
	return nil
}
