package requester

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/fonts"
	"io"
	"net/http"
)

type FontsRequester struct {
	data internal.GeoserverData
}

func (fr *FontsRequester) Get() (*fonts.Fonts, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/fonts", fr.data.Connection.URL), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(fr.data.Connection.Credentials.Username, fr.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := fr.data.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var font fonts.Fonts
		err = json.NewDecoder(response.Body).Decode(&font)
		if err != nil {
			return nil, err
		}

		return &font, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
