package config

import (
	"fmt"

	"gopkg.in/olivere/elastic.v7"
)

//ElasticConn ..
func ElasticConn(url string) *elastic.Client {

	fmt.Println(url)
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		// Handle error
		panic(err)
	}
	return client
}
