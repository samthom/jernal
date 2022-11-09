package files

import (
	"os"

	"github.com/spf13/afero"
)

const dirName = ".jernal"

type system struct {
    fs afero.Afero
    os OsUtils
}

type FS interface {
	Call(string) bool
    create(string) error
}

func NewSystem(fs afero.Fs, os OsUtils) FS {
    return system{
        fs: afero.Afero{
            Fs: fs,
        },
        os: os,
    }
}

func (s system) Call(date string) bool {
	return false
}

func (s system) Init() error {
    dir, err := s.os.getHomeDir()
    if err != nil {
        return err
    }
    appDir := dir + "/" + dirName
    exists, err := s.fs.DirExists(appDir)
    if err != nil {
        return err
    }
    if !exists {
        s.create(appDir)
    }
    return nil
}

func (s system) create(name string) error {
    return s.fs.Mkdir(name, os.ModeDir)
}
