package documents

import (
        "io/ioutil"
        "path/filepath"
        "testing"

        "github.com/stretchr/testify/require"
        "sigs.k8s.io/kustomize/v3/pkg/fs"

        //"opendev.org/airship/airshipctl/pkg/document"
)

// SetupTestFs help manufacture a fake file system for testing purposes. It
// will iterate over the files in fixtureDir, which is a directory relative
// to the tests themselves, and will write each of those files (preserving
// names) to an in-memory file system and return that fs
func SetupTestFs(t *testing.T, fixtureDir string) fs.FileSystem {
        t.Helper()

        x := fs.MakeFakeFS()

        files, err := ioutil.ReadDir(fixtureDir)
        require.NoErrorf(t, err, "Failed to read fixture directory %s", fixtureDir)
        for _, file := range files {
                fileName := file.Name()
                filePath := filepath.Join(fixtureDir, fileName)

                fileBytes, err := ioutil.ReadFile(filePath)
                require.NoErrorf(t, err, "Failed to read file %s, setting up testfs failed", filePath)
                err = x.WriteFile(filepath.Join("/", file.Name()), fileBytes)
                require.NoErrorf(t, err, "Failed to write file %s, setting up testfs failed", filePath)
        }
        return x
}

// NewTestBundle helps to create a new bundle with FakeFs containing documents from fixtureDir
func NewTestBundle(t *testing.T, fixtureDir string) Bundle {
        t.Helper()
        b, err := NewBundle(SetupTestFs(t, fixtureDir), "/", "/")
        require.NoError(t, err, "Failed to build a bundle, setting up TestBundle failed")
        return b
}

