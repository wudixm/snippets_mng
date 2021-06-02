package models

import (
	"github.com/cenkalti/backoff/v4"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"time"
)

var ES *elasticsearch.Client

func ConnectES() {

	//https://github.com/elastic/go-elasticsearch

	//es, _ := elasticsearch.NewDefaultClient()
	//log.Println(elasticsearch.Version)
	//log.Println(es.Info())
	retryBackoff := backoff.NewExponentialBackOff()

	es, err := elasticsearch.NewClient(elasticsearch.Config{

		//Addresses: []string{"http://127.0.0.1:9200"},
		//Addresses: []string{"http://localhost:9200"},
		Addresses: []string{"http://192.168.200.217:9200"},
		// Retry on 429 TooManyRequests statuses
		//
		RetryOnStatus: []int{502, 503, 504, 429},

		// Configure the backoff function
		//
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},

		// Retry up to 5 attempts
		//
		MaxRetries: 3,
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	//NOTE: It is critical to both close the response body and to consume it,
	//in order to re-use persistent TCP connections in the default HTTP transport.
	//If you're not interested in the response body, call io.Copy(ioutil.Discard, res.Body).
	log.Println(res)

	ES = es

}
