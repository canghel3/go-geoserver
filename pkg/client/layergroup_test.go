package client

import (
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

		t.Run("With Style", func(t *testing.T) {
			t.Skip("cannot be tested until style management is implemented")
			layerInputs := []layers.LayerInput{
				{
					Type:  layers.TypeLayer,
					Name:  testdata.CoverageGeoTiffName,
					Style: testdata.StyleGenericName,
				},
				{
					Type:  layers.TypeLayer,
					Name:  testdata.CoverageEHdrName,
					Style: testdata.StyleGenericName,
				},
			}

			err := geoclient.LayerGroups().Delete(testdata.LayerGroupName + "_in_group")
			assert.NoError(t, err)

			err = geoclient.LayerGroups().Delete(testdata.LayerGroupName)
			assert.NoError(t, err)

			err = geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs))
			assert.NoError(t, err)

			group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
			assert.NoError(t, err)
			assert.NotNil(t, group)
			assert.Equal(t, testdata.StyleGenericName, group.Styles.Style[0].Name)
		})

		t.Run("With Keywords", func(t *testing.T) {
			t.Run("Single Keyword", func(t *testing.T) {
				layerInputs := []layers.LayerInput{
					{
						Type:  layers.TypeLayer,
						Name:  testdata.CoverageGeoTiffName,
						Style: testdata.StyleGenericName,
					},
					{
						Type:  layers.TypeLayer,
						Name:  testdata.CoverageEHdrName,
						Style: testdata.StyleGenericName,
					},
				}

				err := geoclient.LayerGroups().Delete(testdata.LayerGroupName + "_in_group")
				assert.NoError(t, err)

				err = geoclient.LayerGroups().Delete(testdata.LayerGroupName)
				assert.NoError(t, err)

				err = geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs, options.LayerGroup.Keywords("single")))
				assert.NoError(t, err)

				group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
				assert.NoError(t, err)
				assert.NotNil(t, group)
				assert.Len(t, group.Keywords.Keywords, 1)
			})

			t.Run("Multi Keyword", func(t *testing.T) {
				layerInputs := []layers.LayerInput{
					{
						Type:  layers.TypeLayer,
						Name:  testdata.CoverageGeoTiffName,
						Style: testdata.StyleGenericName,
					},
					{
						Type:  layers.TypeLayer,
						Name:  testdata.CoverageEHdrName,
						Style: testdata.StyleGenericName,
					},
				}

				err := geoclient.LayerGroups().Delete(testdata.LayerGroupName)
				assert.NoError(t, err)

				err = geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs, options.LayerGroup.Keywords("first", "second")))
				assert.NoError(t, err)

				group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
				assert.NoError(t, err)
				assert.NotNil(t, group)
				assert.Len(t, group.Keywords.Keywords, 2)
			})

		})

		t.Run("With Title", func(t *testing.T) {
			layerInputs := []layers.LayerInput{
				{
					Type:  layers.TypeLayer,
					Name:  testdata.CoverageGeoTiffName,
					Style: testdata.StyleGenericName,
				},
				{
					Type:  layers.TypeLayer,
					Name:  testdata.CoverageEHdrName,
					Style: testdata.StyleGenericName,
				},
			}

			err := geoclient.LayerGroups().Delete(testdata.LayerGroupName)
			assert.NoError(t, err)

			err = geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs, options.LayerGroup.Title("title")))
			assert.NoError(t, err)

			group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
			assert.NoError(t, err)
			assert.NotNil(t, group)
			assert.Equal(t, group.Title, "title")
		})

		t.Run("Single Layer With Default Style", func(t *testing.T) {
			layerInputs := []layers.LayerInput{
				{
					Type: layers.TypeLayer,
					Name: testdata.CoverageGeoTiffName,
				},
			}

			err := geoclient.LayerGroups().Delete(testdata.LayerGroupName)
			assert.NoError(t, err)

			err = geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName, layers.ModeContainer, layerInputs))
			assert.NoError(t, err)

			group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
			assert.NoError(t, err)
			assert.NotNil(t, group)
			assert.Len(t, group.Styles.Style, 1)
			assert.Len(t, group.Publishables.Entries, 1)
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
			Name: testdata.CoverageGeoTiffName,
		},
		{
			Type: layers.TypeLayer,
			Name: testdata.CoverageEHdrName,
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

	group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
	assert.NoError(t, err)
	assert.NotNil(t, group)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("Update Mode", func(t *testing.T) {
			g := group
			g.Mode = layers.ModeSingle
			err = geoclient.LayerGroups().Update(testdata.LayerGroupName, *g)
			assert.NoError(t, err)

			group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
			assert.NoError(t, err)
			assert.NotNil(t, group)
			assert.Equal(t, layers.ModeSingle, group.Mode)
		})

		t.Run("Add New Layer", func(t *testing.T) {
			t.Run("Styles Is Nil", func(t *testing.T) {
				g := group
				g.Publishables = nil
				g.Styles = nil
				g.AddPublishables(layers.LayerInput{
					Type:  layers.TypeLayer,
					Name:  testdata.CoverageGeoTiffName,
					Style: "",
				})

				err = geoclient.LayerGroups().Update(testdata.LayerGroupName, *g)
				assert.NoError(t, err)

				group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
				assert.NoError(t, err)
				assert.NotNil(t, group)
				assert.Len(t, group.Publishables.Entries, 1)
			})

			t.Run("Styles Exist", func(t *testing.T) {
				g := group
				g.AddPublishables(layers.LayerInput{
					Type:  layers.TypeLayer,
					Name:  testdata.CoverageGeoTiffName,
					Style: "",
				})

				err = geoclient.LayerGroups().Update(testdata.LayerGroupName, *g)
				assert.NoError(t, err)

				group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
				assert.NoError(t, err)
				assert.NotNil(t, group)
				assert.Len(t, group.Publishables.Entries, 2)
			})
		})

		t.Run("Add New Layer Group", func(t *testing.T) {
			err := geoclient.Workspace(testdata.Workspace).LayerGroups().Publish(layers.NewGroup(testdata.LayerGroupName+"interm", layers.ModeContainer, layerInputs))
			assert.NoError(t, err)

			g := group
			g.AddPublishables(layers.LayerInput{
				Type:  layers.TypeLayerGroup,
				Name:  testdata.LayerGroupName + "interm",
				Style: "",
			})

			err = geoclient.LayerGroups().Update(testdata.LayerGroupName, *g)
			assert.NoError(t, err)

			group, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
			assert.NoError(t, err)
			assert.NotNil(t, group)
			assert.Len(t, group.Publishables.Entries, 3)
		})

		t.Run("With Style", func(t *testing.T) {
			t.Skip("test wont work until styles management is implemented")
			g := group
			g.AddPublishables(layers.LayerInput{
				Type:  layers.TypeLayer,
				Name:  testdata.CoverageGeoTiffName,
				Style: testdata.StyleGenericName,
			})

			err = geoclient.LayerGroups().Update(testdata.LayerGroupName, *g)
			assert.NoError(t, err)

			gg, err := geoclient.Workspace(testdata.Workspace).LayerGroups().Get(testdata.LayerGroupName)
			assert.NoError(t, err)
			assert.NotNil(t, gg)
			assert.Len(t, gg.Publishables.Entries, 4)
			assert.Equal(t, gg.Styles.Style[4].Name, testdata.StyleGenericName)
		})
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
