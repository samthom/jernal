package files

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewSystem(t *testing.T) {
    osUtil := newOsUtil()
    fs := afero.NewMemMapFs()

    want := system{
        fs: afero.Afero{
            Fs: fs,
        },
        os: osUtil,
    }
    got := NewSystem(fs, osUtil)

    assert.Equal(t, want, got, "invalid system instance returned")
}

func TestCreate(t *testing.T) {
    osUtil := newOsUtil()
    fs := afero.NewMemMapFs()
    dirPath := "/.jernal"

    system := NewSystem(fs, osUtil)

    t.Run("create dir without error", func(t *testing.T) {
        err := system.create(dirPath)

        assert.NoError(t, err, "unexpected error from create new dir")
        // check if the dir is created with correct permissions
        info, err := fs.Stat(dirPath)
        assert.NoError(t, err, "unexpected error from create new dir")
        assert.True(t, info.IsDir(), "not a directory")
        assert.Equal(t, os.ModeDir, info.Mode(), "unexpected permissions for dir")
    })

    t.Run("create dir returns error", func(t *testing.T) {
        err := system.create("")

        assert.Error(t, err, "create is not returning error as expected")
    })
}
