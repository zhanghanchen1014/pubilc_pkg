package untils

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/zhanghanchen1014/viper/config"
	"log"
)

type ES struct {
	Id    string
	Index string
	Body  interface{}
}

func (e *ES) SyncES() error {
	// Build the request body.
	data, err := json.Marshal(e.Body)
	if err != nil {
		log.Fatalf("Error marshaling document: %s", err)
	}

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      e.Index,
		DocumentID: e.Id,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), config.Els)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	return err
}
