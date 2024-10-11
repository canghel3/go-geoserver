package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	"github.com/rs/zerolog/log"
	"io"
	"path/filepath"

	//postgres library
	_ "github.com/lib/pq"
	"net"
	"net/http"
	"os"
	"testing"
)

var (
	host = "" //automatically filled in

	target   string
	username = "admin"
	password = "geoserver"

	databaseUser     = "admin"
	databasePassword = "geoserver"
	databaseName     = "geoserver"
	databasePort     = "" //automatically filled in
	databaseIP       = "" //automatically filled in
)

func TestMain(m *testing.M) {
	host = getOutboundIP()
	databaseIP = host

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Error().Msgf("could not connect to docker: %s", err)
		return
	}

	database, err := startDatabase(pool)
	if err != nil {
		log.Error().Msg(err.Error())
		purge(pool, database)
		return
	}

	geoserver, err := startGeoserver(pool)
	if err != nil {
		log.Error().Msg(err.Error())
		purge(pool, database, geoserver)
		return
	}

	target = fmt.Sprintf("http://%s:%s", host, geoserver.GetPort("8080/tcp"))

	code := m.Run()
	purge(pool, database, geoserver)
	os.Exit(code)
}

func startGeoserver(pool *dockertest.Pool) (*dockertest.Resource, error) {
	//wording directory will be path to where this file is located ./service
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	localMountLocation := fmt.Sprintf("%s/testdata/geoserver/data", wd)
	mount := fmt.Sprintf("%s:/opt/geoserver/data", localMountLocation)

	options := &dockertest.RunOptions{
		Repository: "kartoza/geoserver",
		Tag:        "2.22.2",
		Env: []string{
			//"DB_BACKEND=POSTGRES",
			"RECREATE_DATADIR=TRUE",
			//"POSTGRES_DB=" + vectorLayersDBName,
			//"POSTGRES_USER=" + vectorLayersDBUser,
			//"POSTGRES_PASS=" + vectorLayersDBPassword,
			//"POSTGRES_PORT=" + dbPort,
			"GEOSERVER_ADMIN_PASSWORD=geoserver",
			"GEOSERVER_ADMIN_USER=admin",
			"STABLE_EXTENSIONS=css-plugin",
			"HOST=" + host,
			"GEOSERVER_DATA_DIR=/opt/geoserver/data"},
		Mounts:     []string{mount},
		Privileged: false,
	}

	resource, err := pool.RunWithOptions(options)
	if err != nil {
		return nil, err
	}

	port := resource.GetPort("8080/tcp")
	url := fmt.Sprintf("http://%s:%s/geoserver/web", host, port)

	err = pool.Retry(func() error {
		response, err := http.DefaultClient.Get(url)
		if err != nil {
			return err
		}

		if response.StatusCode != http.StatusOK {
			return errors.New("waiting for status OK")
		}
		return nil
	})
	if err != nil {
		return resource, err
	}

	f := filepath.Join(wd, "testdata", "geotiff", "rasters", "shipments_2_geocoded.tif")
	err = copyFileToDir(f, localMountLocation)

	return resource, err
}

func startDatabase(pool *dockertest.Pool) (*dockertest.Resource, error) {
	resource, err := pool.Run("postgis/postgis", "latest",
		[]string{
			"POSTGRES_USER=" + databaseUser,
			"POSTGRES_PASSWORD=" + databasePassword,
			"POSTGRES_DB=" + databaseName},
	)
	if err != nil {
		return resource, err
	}

	databasePort = resource.GetPort("5432/tcp")

	err = pool.Retry(func() error {
		db, err := sqlx.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable", databaseUser, databasePassword, host, databaseName, databasePort))
		if err != nil {
			return err
		}

		return db.Ping()
	})
	if err != nil {
		return resource, err
	}

	files := []string{"./testdata/migrations/table_for_feature.sql"}
	return resource, migrate(files...)
}

func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func purge(pool *dockertest.Pool, resources ...*dockertest.Resource) {
	for _, r := range resources {
		err := pool.Purge(r)
		if err != nil {
			log.Printf("error purging resource %v: %s", r, err.Error())
		}
	}
}

func migrate(files ...string) error {
	connection := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable", databaseUser, databasePassword, host, databaseName, databasePort)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return err
	}
	defer db.Close()

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return err
		}
	}

	return nil
}

// copyFileToDir copies a source file to a target directory.
// The new file will have the same name as the source file.
func copyFileToDir(srcFilePath, targetDir string) error {
	// Open the source file for reading
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Extract the filename from the source file path
	srcFileName := filepath.Base(srcFilePath)

	// Create the destination file path
	destFilePath := filepath.Join(targetDir, srcFileName)

	// Create the destination file for writing
	destFile, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy the contents from the source file to the destination file
	_, err = io.Copy(destFile, srcFile)
	return err
}
