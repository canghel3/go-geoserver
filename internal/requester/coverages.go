package requester

import (
	"bytes"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/models/coverages"
	"github.com/canghel3/go-geoserver/pkg/models/customerrors"
	"io"
	"net/http"
)

type CoverageRequester struct {
	store string
	data  *internal.GeoserverData
}

func (cr *CoverageRequester) Create(content []byte) error {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages", cr.data.Connection.URL, cr.data.Workspace, cr.store), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := cr.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusCreated, http.StatusOK:
		return nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

// TODO: implement
//func (cr *CoverageRequester) GetAll() (*coveragestores.AllCoverageStoreRetrievalWrapper, error) {
//	return nil, errors.New("not implemented")
//}

func (cr *CoverageRequester) Get() (*coverages.CoverageWrapper, error) {
	return nil, nil
}

func (cr *CoverageRequester) Delete(content []byte) error {
	return nil
}
