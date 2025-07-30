package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/formats"
	"github.com/canghel3/go-geoserver/pkg/layers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLayerGroupIntegration_Create(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, formats.GeoTIFF)
	addTestCoverageStore(t, formats.EHdr)
	addTestCoverage(t, formats.GeoTIFF)
	addTestCoverage(t, formats.EHdr)

	t.Run("201 Created", func(t *testing.T) {
		t.Run("Mode Container", func(t *testing.T) {
			layerInputs := []layers.LayerInput{
				{
					Type: layers.TypeLayer,
					Name: testdata.CoverageGeoTiffName,
				},
				{
					Type: layers.TypeLayer,
					Name: testdata.CoverageEHdrName,
				},
			}
			err := geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs))
			assert.NoError(t, err)
		})

		t.Run("In Workspace", func(t *testing.T) {
			layerInputs := []layers.LayerInput{
				{
					Type: layers.TypeLayer,
					Name: testdata.CoverageGeoTiffName,
				},
				{
					Type: layers.TypeLayer,
					Name: testdata.CoverageEHdrName,
				},
			}

			err := geoclient.LayerGroups().Delete(testdata.LayerGroupName)
			assert.NoError(t, err)

			err = geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs))
			assert.NoError(t, err)
		})

		t.Run("Not In Workspace", func(t *testing.T) {
			//requires layer input name to contain workspace in the name
			layerInputs := []layers.LayerInput{
				{
					Type: layers.TypeLayer,
					Name: fmt.Sprintf("%s:%s", testdata.Workspace, testdata.CoverageGeoTiffName),
				},
				{
					Type: layers.TypeLayer,
					Name: fmt.Sprintf("%s:%s", testdata.Workspace, testdata.CoverageEHdrName),
				},
			}
			err := geoclient.LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs))
			assert.NoError(t, err)
		})

		t.Run("Layer Group In Layer Group", func(t *testing.T) {
			t.Skip("not implemented")
		})
	})

	t.Run("Invalid Name", func(t *testing.T) {
		err := geoclient.LayerGroups().Publish(layers.NewGroup(testdata.InvalidName, layers.ModeContainer, nil))
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})
}

func TestLayerGroupIntegration_Get(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, formats.GeoTIFF)
	addTestCoverageStore(t, formats.EHdr)
	addTestCoverage(t, formats.GeoTIFF)
	addTestCoverage(t, formats.EHdr)

	layerInputs := []layers.LayerInput{
		{
			Type: layers.TypeLayer,
			Name: fmt.Sprintf("%s:%s", testdata.Workspace, testdata.CoverageGeoTiffName),
		},
		{
			Type: layers.TypeLayer,
			Name: fmt.Sprintf("%s:%s", testdata.Workspace, testdata.CoverageEHdrName),
		},
	}
	err := geoclient.LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs))
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		group, err := geoclient.LayerGroups().Get(testdata.LayerGroupName)
		assert.NoError(t, err)
		assert.NotNil(t, group)
		assert.Equal(t, 2, len(group.Publishables.Entries))
		assert.Equal(t, layers.ModeContainer, group.Mode)
		assert.Nil(t, group.Workspace)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		group, err := geoclient.LayerGroups().Get("none")
		assert.Error(t, err)
		assert.Nil(t, group)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})
}

func TestLayerGroupIntegration_Delete(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, formats.GeoTIFF)
	addTestCoverageStore(t, formats.EHdr)
	addTestCoverage(t, formats.GeoTIFF)
	addTestCoverage(t, formats.EHdr)

	//requires layer input name to contain workspace in the name
	layerInputs := []layers.LayerInput{
		{
			Type: layers.TypeLayer,
			Name: fmt.Sprintf("%s:%s", testdata.Workspace, testdata.CoverageGeoTiffName),
		},
		{
			Type: layers.TypeLayer,
			Name: fmt.Sprintf("%s:%s", testdata.Workspace, testdata.CoverageEHdrName),
		},
	}
	err := geoclient.LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs))
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		err = geoclient.LayerGroups().Delete(testdata.LayerGroupName)
		assert.NoError(t, err)
	})
}

func TestLayerGroupIntegration_Update(t *testing.T) {
	t.Run("Invalid Previous Name", func(t *testing.T) {
		err := geoclient.LayerGroups().Update(testdata.InvalidName, layers.NewGroup("some", layers.ModeContainer, nil))
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})

	t.Run("Invalid New Name", func(t *testing.T) {
		err := geoclient.LayerGroups().Update("some", layers.NewGroup(testdata.InvalidName, layers.ModeContainer, nil))
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})
}
