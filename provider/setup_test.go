package provider

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/avast/retry-go/v3"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations"
	"github.com/go-openapi/loads"
)

var serverPort int

func TestMain(m *testing.M) {
	swaggerSpec, err := loads.Embedded(server.SwaggerJSON, server.FlatSwaggerJSON)
	must(err)
	api := operations.NewCloudViewAPIsAPI(swaggerSpec)
	srv := server.NewServer(api)
	srv.ConfigureAPI()

	go func() {
		if err := srv.Serve(); err != nil {
			log.Fatalln(err)
		}
	}()

	must(retry.Do(func() error {
		serverPort = srv.Port
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", serverPort))
		return err
	}, retry.Delay(10*time.Millisecond), retry.Attempts(10)))

	status := m.Run()

	must(srv.Shutdown())

	os.Exit(status)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
