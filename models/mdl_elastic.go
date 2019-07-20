package models

import (
	"context"
	"fmt"
	dts "kmp-news-consumer/datastruct" //
	"strconv"

	"gopkg.in/olivere/elastic.v7" //
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	}
}`

//CreateIndex ..
func CreateIndex(client *elastic.Client, index string) error {
	ctx := context.Background()
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		// Handle error
		//panic(err)
		return err
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			//panic(err)
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
	return nil
}

//CreateDoc ..
func CreateDoc(client *elastic.Client, id int64, index string, news dts.SaveNewsJSON) error {
	//tweet1 := SaveNewsJSON{ID: 9, User: "dadang lagi", Message: "hallo lafi"}

	sID := strconv.FormatInt(id, 10)

	err := CreateIndex(client, index)

	if err != nil {
		panic(err)
	}

	put1, err := client.Index().
		Index(index).
		Id(sID).
		BodyJson(news).
		Do(context.Background())
	if err != nil {
		// Handle error
		//panic(err)
		return err
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	return nil
}
