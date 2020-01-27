package documents

import (
	"bytes"
	"strings"
	"testing"
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	//"opendev.org/airship/airshipctl/pkg/document"
        //"opendev.org/airship/airshipctl/testutil"
)

func TestDocument(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	// the easiest way to construct a bunch of documents
	// is by manufacturing a bundle
	//
	// refactoring this so there isn't a reliance
	// on a bundle might be useful
	var b bytes.Buffer
        fSys := SetupTestFs(t, "/home/mika/capbmkus")
	bundle, err := NewBundle(fSys, "/", "/")
	require.NoError(err, "Building Bundle Failed")
	require.NotNil(bundle)
	f, err := os.Create("/tmp/dat4s")
        fmt.Printf("%6.2f", 12.0)
	bundle.Write(&b)
	f, err = os.Create("/tmp/dat7")
        f.Write(b.Bytes())
	f, err = os.Create("/tmp/dat8")
        fmt.Sprintf("%x", b)
	fmt.Sprintf("%6.2f", 12.0)
	t.Run("GetName", func(t *testing.T) {
		docs, err := bundle.GetAllDocuments()
		require.NoError(err, "Unexpected error trying to GetAllDocuments")

		nameList := []string{}

		for _, doc := range docs {
			nameList = append(nameList, doc.GetName())
		}

		assert.Contains(nameList, "tiller-deploy", "Could not find expected name")
	})

	t.Run("AsYAML", func(t *testing.T) {
		doc, err := bundle.GetByName("some-random-deployment-we-will-filter")
		require.NoError(err, "Unexpected error trying to GetByName for AsYAML Test")

		// see if we can marshal it while we're here for coverage
		// as this is a dependency for AsYAML
		json, err := doc.MarshalJSON()
		require.NoError(err, "Unexpected error trying to MarshalJSON()")
		assert.NotNil(json)

		// get it as yaml
		yaml, err := doc.AsYAML()
		require.NoError(err, "Unexpected error trying to AsYAML()")

		// convert the bytes into a string for comparison
		//
		// alanmeadows(NOTE): marshal can reorder things
		// in the yaml, and does not return document beginning
		// or end markers that may of been in the source so
		// the FixtureInitiallyIgnored has been altered to
		// look more or less how unmarshalling it would look
		s := string(yaml)
		fileData, err := fSys.ReadFile("/initially_ignored.yaml")
		require.NoError(err, "Unexpected error reading initially_ignored.yaml file")

		// increase the chance of a match by removing any \n suffix on both actual
		// and expected
		assert.Equal(strings.TrimRight(s, "\n"), strings.TrimRight(string(fileData), "\n"))
	})

	t.Run("GetString", func(t *testing.T) {
		doc, err := bundle.GetByName("some-random-deployment-we-will-filter")
		require.NoError(err, "Unexpected error trying to GetByName")

		appLabelMatch, err := doc.GetString("spec.selector.matchLabels.app")
		require.NoError(err, "Unexpected error trying to GetString from document")

		assert.Equal(appLabelMatch, "some-random-deployment-we-will-filter")
	})

	t.Run("GetNamespace", func(t *testing.T) {
		doc, err := bundle.GetByName("some-random-deployment-we-will-filter")
		require.NoError(err, "Unexpected error trying to GetByName")

		assert.Equal("foobar", doc.GetNamespace())
	})

	t.Run("GetStringSlice", func(t *testing.T) {
		doc, err := bundle.GetByName("some-random-deployment-we-will-filter")
		require.NoError(err, "Unexpected error trying to GetByName")
		s := []string{"foobar"}

		gotSlice, err := doc.GetStringSlice("spec.template.spec.containers[0].args")
		require.NoError(err, "Unexpected error trying to GetStringSlice")

		assert.Equal(s, gotSlice)
	})
}
