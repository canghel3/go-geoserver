package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/formats"
	"github.com/canghel3/go-geoserver/pkg/layers"
	"github.com/canghel3/go-geoserver/pkg/options"
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
			err := geoclient.LayerGroups().Delete(testdata.LayerGroupName)
			assert.NoError(t, err)

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

			err = geoclient.LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs, options.LayerGroup.Workspace(testdata.Workspace)))
			assert.NoError(t, err)
		})

		t.Run("Layer Group In Layer Group", func(t *testing.T) {
			layerInputs := []layers.LayerInput{
				{
					Type: layers.TypeLayer,
					Name: testdata.CoverageGeoTiffName,
				},
				{
					Type: layers.TypeLayer,
					Name: testdata.CoverageEHdrName,
				},
				{
					Type: layers.TypeLayerGroup,
					Name: testdata.LayerGroupName,
				},
			}
			err := geoclient.LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName+"_in_group", layers.ModeContainer, layerInputs, options.LayerGroup.Workspace(testdata.Workspace)))
			assert.NoError(t, err)
		})
	})

	t.Run("Invalid Group Name", func(t *testing.T) {
		err := geoclient.LayerGroups().Publish(layers.NewGroup(testdata.InvalidName, layers.ModeContainer, nil))
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})

	t.Run("Invalid Workspace Name", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.InvalidName, layers.ModeContainer, nil))
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})

	t.Run("Workspace Not Specified", func(t *testing.T) {
		t.Run("And It Is Unkown", func(t *testing.T) {
			err := geoclient.LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeSingle, nil))
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, "workspace is required. use Workspace from options.LayerGroup")
		})

		t.Run("And It Is Known", func(t *testing.T) {
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

			err := geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName+"new", layers.ModeSingle, layerInputs))
			assert.NoError(t, err)
		})
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
			Name: testdata.CoverageGeoTiffName,
		},
		{
			Type: layers.TypeLayer,
			Name: testdata.CoverageEHdrName,
		},
	}
	err := geoclient.LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs, options.LayerGroup.Workspace(testdata.Workspace)))
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
		assert.NoError(t, err)
		assert.NotNil(t, group)
		assert.Equal(t, 2, len(group.Publishables.Entries))
		assert.Equal(t, layers.ModeContainer, group.Mode)
		assert.NotNil(t, group.Workspace)
		assert.Equal(t, testdata.Workspace, group.Workspace.Name)
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
	err := geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs))
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		err = geoclient.Workspace(testdata.Workspace).LayerGroups().Delete(testdata.LayerGroupName)
		assert.NoError(t, err)
	})
}

func TestLayerGroupIntegration_Update(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, formats.GeoTIFF)
	addTestCoverageStore(t, formats.EHdr)
	addTestCoverage(t, formats.GeoTIFF)
	addTestCoverage(t, formats.EHdr)

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

	group, err := geoclient.LayerGroups().Get(fmt.Sprintf("%s:%s", testdata.Workspace, testdata.LayerGroupName))
	assert.NoError(t, err)
	assert.NotNil(t, group)

	t.Run("200 Ok", func(t *testing.T) {
		group.Mode = layers.ModeSingle
		err = geoclient.LayerGroups().Update(testdata.LayerGroupName, *group)
		assert.NoError(t, err)

		group, err := geoclient.LayerGroups().Get(fmt.Sprintf("%s:%s", testdata.Workspace, testdata.LayerGroupName))
		assert.NoError(t, err)
		assert.NotNil(t, group)
		assert.Equal(t, layers.ModeSingle, group.Mode)
	})

	//t.Run("404 Not Found", func(t *testing.T) {
	//	group.Name = "none"
	//	err = geoclient.LayerGroups().Update("none", *group)
	//	assert.Error(t, err)
	//	assert.Nil(t, group)
	//	assert.IsType(t, &customerrors.NotFoundError{}, err)
	//})

	t.Run("Invalid Previous Name", func(t *testing.T) {
		err = geoclient.LayerGroups().Update(testdata.InvalidName, *group)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})

	t.Run("Invalid New Name", func(t *testing.T) {
		group.Name = testdata.InvalidName
		err = geoclient.LayerGroups().Update("some", *group)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})

	t.Run("Workspace Not Specified", func(t *testing.T) {
		t.Run("And It Is Unknown", func(t *testing.T) {
			group.Workspace = nil
			group.Name = testdata.LayerGroupName
			err = geoclient.LayerGroups().Update("existing", *group)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, "group.Workspace is required")
		})

		t.Run("And It Is Known", func(t *testing.T) {
			group.Name = testdata.LayerGroupName
			err = geoclient.Workspace(testdata.Workspace).LayerGroups().Update(testdata.LayerGroupName, *group)
			assert.NoError(t, err)
		})
	})
}
