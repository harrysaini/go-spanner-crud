package libs

import (
	"context"
	"fmt"
	"go-spanner-crud/src/models"
	"log"

	"cloud.google.com/go/spanner"
)

var client *spanner.Client = nil

func getSpannerClient(conf *models.Configuration) *spanner.Client {
	ctx := context.Background()
	databaseURL := fmt.Sprintf("projects/%s/instances/%s/databases/%s", conf.Gcp.Project, conf.Gcp.Instance, conf.Gcp.Database)
	log.Println(databaseURL)
	clint, err := spanner.NewClient(ctx, databaseURL)
	log.Println("Spanner", err)
	if err != nil {
		log.Fatalln(err)
	}

	return clint
}

// GetSpannerClientInstance returns spanner instance
func GetSpannerClientInstance() *spanner.Client {
	if client != nil {
		return client
	}
	client = getSpannerClient(Conf)
	return client
}
