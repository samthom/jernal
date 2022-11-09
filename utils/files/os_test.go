package files

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewOsUtil(t *testing.T) {
    want := osUtils{}
    got := newOsUtil()

    assert.Equal(t, want, got, "unexpected value fron newOsUtil")
}

func TestGetHomeDir(t *testing.T) {
    o := newOsUtil()

    want, _ := os.UserHomeDir()
    got, err := o.getHomeDir()

    assert.Nil(t, err, "err is not nil")
    assert.Equal(t, want, got, "unexpected home directory")
}

func TestGetDateStr(t *testing.T) {
    o := newOsUtil()

    want := time.Now().Format("01-02-2006")
    got := o.getDateStr()
    assert.Equal(t, want, got, "unexpected date string")
}
