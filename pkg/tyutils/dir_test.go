package tyutils

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestFile() string {
	f, _ := os.Create("test_file")
	return f.Name()
}

func cleanTestFile() {
	_ = os.Remove("test_file")
}

func TestExists(t *testing.T) {
	pwd, err := os.Getwd()
	assert.Nil(t, err)
	tf := createTestFile()
	defer cleanTestFile()

	p := path.Join(pwd, tf)
	assert.True(t, Exists(p))

	assert.False(t, Exists(p+"_no_exist"))
}

// 判断所给路径是否为文件夹
func TestIsDir(t *testing.T) {
	pwd, err := os.Getwd()
	assert.Nil(t, err)
	tf := createTestFile()
	defer cleanTestFile()

	p := path.Join(pwd, tf)
	assert.True(t, IsDir(pwd))
	assert.False(t, IsDir(p))
}

func TestIsFile(t *testing.T) {
	pwd, err := os.Getwd()
	assert.Nil(t, err)
	tf := createTestFile()
	defer cleanTestFile()

	p := path.Join(pwd, tf)
	assert.True(t, IsFile(p))
	assert.False(t, IsFile(pwd))
}
