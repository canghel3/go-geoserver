package service

import (
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/utils"
	"gotest.tools/v3/assert"
	"math"
	"testing"
)

func TestCoverage(t *testing.T) {
	geoserverService := NewGeoserverService(geoserverURL, username, password)
	assert.NilError(t, geoserverService.CreateWorkspace("init"))

	t.Run("GeoTIFF", func(t *testing.T) {
		assert.NilError(t, geoserverService.CreateCoverageStore("init", "init", "file:/opt/geoserver/data/shipments_2_geocoded.tif", "GeoTIFF"))

		t.Run("CREATE", func(t *testing.T) {
			t.Run("BASIC", func(t *testing.T) {
				bbox := [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}
				assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", bbox, utils.ProjectionPolicyOption("FORCE_DECLARED")))
				c, err := geoserverService.GetCoverage("init", "init", "init")
				assert.NilError(t, err)

				assert.Equal(t, c.Coverage.ProjectionPolicy, "FORCE_DECLARED")
				assert.Equal(t, c.Coverage.Srs, "EPSG:3857")
				assert.Equal(t, c.Coverage.NativeBoundingBox.MinX, -13625746.1231970004737377)
				assert.Equal(t, c.Coverage.Title, "init")

				assert.NilError(t, geoserverService.DeleteCoverage("init", "init", "init", utils.RecurseOption(true)))
			})

			t.Run("WITH KEYWORDS", func(t *testing.T) {
				keywords := []string{"adu", "telefonul", "marian"}
				bbox := [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}
				assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", bbox, utils.ProjectionPolicyOption("FORCE_DECLARED"), utils.KeywordsOption(keywords)))

				c, err := geoserverService.GetCoverage("init", "init", "init")
				assert.NilError(t, err)

				assert.Equal(t, c.Coverage.ProjectionPolicy, "FORCE_DECLARED")
				assert.Equal(t, c.Coverage.Srs, "EPSG:3857")
				assert.Equal(t, c.Coverage.NativeBoundingBox.MinX, -13625746.1231970004737377)
				assert.Equal(t, c.Coverage.Title, "init")
				assert.DeepEqual(t, c.Coverage.Keywords.Keywords, keywords)

				assert.NilError(t, geoserverService.DeleteCoverage("init", "init", "init", utils.RecurseOption(true)))
			})

			t.Run("WITH TITLE", func(t *testing.T) {
				bbox := [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}
				assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", bbox, utils.ProjectionPolicyOption("FORCE_DECLARED"), utils.TitleOption("MARIAN")))

				c, err := geoserverService.GetCoverage("init", "init", "init")
				assert.NilError(t, err)

				assert.Equal(t, c.Coverage.ProjectionPolicy, "FORCE_DECLARED")
				assert.Equal(t, c.Coverage.Srs, "EPSG:3857")
				assert.Equal(t, c.Coverage.NativeBoundingBox.MinX, -13625746.1231970004737377)
				assert.Equal(t, c.Coverage.Title, "MARIAN")

				assert.NilError(t, geoserverService.DeleteCoverage("init", "init", "init", utils.RecurseOption(true)))
			})

			t.Run("WITHOUT PROJECTION POLICY OPTION", func(t *testing.T) {
				bbox := [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}
				assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", bbox, utils.TitleOption("MARIAN")))

				c, err := geoserverService.GetCoverage("init", "init", "init")
				assert.NilError(t, err)

				assert.Equal(t, c.Coverage.ProjectionPolicy, "NONE")
				assert.Equal(t, c.Coverage.Srs, "EPSG:3857")
				assert.Equal(t, c.Coverage.NativeBoundingBox.MinX, -13625746.1231970004737377)
				assert.Equal(t, c.Coverage.Title, "MARIAN")

				assert.NilError(t, geoserverService.DeleteCoverage("init", "init", "init", utils.RecurseOption(true)))
			})

			t.Run("WITH REPROJECTION POLICY", func(t *testing.T) {
				bbox := [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}
				assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", bbox, utils.ProjectionPolicyOption("REPROJECT_TO_DECLARED")))

				c, err := geoserverService.GetCoverage("init", "init", "init")
				assert.NilError(t, err)

				assert.Equal(t, c.Coverage.ProjectionPolicy, "REPROJECT_TO_DECLARED")
				assert.Equal(t, c.Coverage.Srs, "EPSG:3857")
				assert.Equal(t, c.Coverage.NativeBoundingBox.MinX, -13625746.1231970004737377)
				assert.Equal(t, c.Coverage.Title, "init")

				assert.NilError(t, geoserverService.DeleteCoverage("init", "init", "init", utils.RecurseOption(true)))
			})

			t.Run("WITH COVERAGE DIMENSION OPTION", func(t *testing.T) {
				t.Run("W/O DESCRIPTION", func(t *testing.T) {
					bbox := [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}

					coverage := utils.MakeCoverageDimension("BANDA", "REAL_64BITS", 0, 0, math.MaxFloat64)
					assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", bbox, utils.ProjectionPolicyOption("FORCE_DECLARED"), utils.TitleOption("MARIAN"), utils.CoverageDimensionsOption(coverage)))

					c, err := geoserverService.GetCoverage("init", "init", "init")
					assert.NilError(t, err)

					assert.Equal(t, c.Coverage.ProjectionPolicy, "FORCE_DECLARED")
					assert.Equal(t, c.Coverage.Srs, "EPSG:3857")
					assert.Equal(t, c.Coverage.NativeBoundingBox.MinX, -13625746.1231970004737377)
					assert.Equal(t, c.Coverage.Title, "MARIAN")
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Name, "BANDA")
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.DataType.Name, "REAL_64BITS")
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.NullValues.Double, float64(0))
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Range.Min, float64(0))
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Range.Max, math.MaxFloat64)

					assert.NilError(t, geoserverService.DeleteCoverage("init", "init", "init", utils.RecurseOption(true)))
				})

				t.Run("W/ DESCRIPTION", func(t *testing.T) {
					bbox := [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}

					coverage := utils.MakeCoverageDimension("BANDA", "REAL_32BITS", 0, 0, math.MaxFloat32, "a simple description goes a long way")
					assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", bbox, utils.ProjectionPolicyOption("FORCE_DECLARED"), utils.TitleOption("MARIAN"), utils.CoverageDimensionsOption(coverage)))

					c, err := geoserverService.GetCoverage("init", "init", "init")
					assert.NilError(t, err)

					assert.Equal(t, c.Coverage.ProjectionPolicy, "FORCE_DECLARED")
					assert.Equal(t, c.Coverage.Srs, "EPSG:3857")
					assert.Equal(t, c.Coverage.NativeBoundingBox.MinX, -13625746.1231970004737377)
					assert.Equal(t, c.Coverage.Title, "MARIAN")
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Name, "BANDA")
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Range.Min, float64(0))
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Range.Max, math.MaxFloat32)
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.DataType.Name, "REAL_32BITS")
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.NullValues.Double, float64(0))
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Description, "a simple description goes a long way")

					assert.NilError(t, geoserverService.DeleteCoverage("init", "init", "init", utils.RecurseOption(true)))
				})

				t.Run("W/ MULTILINE DESCRIPTION", func(t *testing.T) {
					bbox := [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}

					coverage := utils.MakeCoverageDimension("AVE CESAR", "REAL_64BITS", 0, 0, math.MaxFloat64, "a simple description goes a long way", "especially when written in hungarian")
					assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", bbox, utils.ProjectionPolicyOption("FORCE_DECLARED"), utils.TitleOption("MARIAN"), utils.CoverageDimensionsOption(coverage)))

					c, err := geoserverService.GetCoverage("init", "init", "init")
					assert.NilError(t, err)

					assert.Equal(t, c.Coverage.ProjectionPolicy, "FORCE_DECLARED")
					assert.Equal(t, c.Coverage.Srs, "EPSG:3857")
					assert.Equal(t, c.Coverage.NativeBoundingBox.MinX, -13625746.1231970004737377)
					assert.Equal(t, c.Coverage.Title, "MARIAN")
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Name, "AVE CESAR")
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Range.Min, float64(0))
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Range.Max, math.MaxFloat64)
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.DataType.Name, "REAL_64BITS")
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.NullValues.Double, float64(0))
					assert.Equal(t, c.Coverage.Dimensions.CoverageDimension.Description, "a simple description goes a long way\nespecially when written in hungarian")

					assert.NilError(t, geoserverService.DeleteCoverage("init", "init", "init", utils.RecurseOption(true)))
				})
			})

			t.Run("DUPLICATE", func(t *testing.T) {
				assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", [4]float64{0, 0, 0, 0}))
				assert.Error(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", [4]float64{0, 0, 0, 0}), "coverage init already exists")
			})
		})

		t.Run("GET", func(t *testing.T) {
			t.Run("NON-EXISTENT", func(t *testing.T) {
				_, err := geoserverService.GetCoverage("init", "init", "does_not_exist")
				assert.ErrorType(t, err, &customerrors.NotFoundError{})
				assert.Error(t, err, "coverage does_not_exist does not exist")
			})
		})

		t.Run("DELETE", func(t *testing.T) {
			t.Run("NON-EXISTENT", func(t *testing.T) {
				err := geoserverService.DeleteCoverage("init", "init", "does_not")
				assert.ErrorType(t, err, &customerrors.NotFoundError{})
				assert.Error(t, err, "coverage does_not does not exist")
			})
		})
	})

	assert.NilError(t, geoserverService.DeleteWorkspace("init", utils.RecurseOption(true)))
}
