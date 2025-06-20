package testdata

import (
	"github.com/canghel3/go-geoserver/internal"
	"io"
	"os"
	"path/filepath"
)

func GeoserverInfo(client internal.HTTPClient) internal.GeoserverData {
	return internal.GeoserverData{
		Client: client,
		Connection: internal.GeoserverConnection{
			URL: GeoserverUrl,
			Credentials: internal.GeoserverCredentials{
				Username: GeoserverUsername,
				Password: GeoserverPassword,
			},
		},
		DataDir:   GeoserverDataDir,
		Workspace: Workspace,
	}
}

func Read(file string) ([]byte, error) {
	return os.ReadFile(file)
}

func Copy(src, dst string) error {
	sr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sr.Close()

	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}

	ds, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer ds.Close()

	_, err = io.Copy(ds, sr)
	if err != nil {
		return err
	}

	return nil
}
