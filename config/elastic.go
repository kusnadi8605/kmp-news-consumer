package config

import (
	"gopkg.in/olivere/elastic.v7"
)

//ElasticConn ..
func ElasticConn(url string) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		// Handle error
		panic(err)
	}
	return client
}
