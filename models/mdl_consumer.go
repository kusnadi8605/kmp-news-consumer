package models

import (
	"context"
	"encoding/json"
	conf "kmp-news-consumer/config"
	dts "kmp-news-consumer/datastruct"
	log "kmp-news-consumer/logging"

	"github.com/segmentio/kafka-go"
)

//Consumer ..
func Consumer(reader *kafka.Reader, dbConn *conf.Connection, elasticURL string) {
	defer reader.Close()
	for {
		m, err := reader.FetchMessage(context.Background())

		if err != nil {
			log.Errorf("read kafka error : %s", err.Error())
		}

		//log.Logf("Got incoming message: %v / %v / %v : %s", m.Topic, m.Partition, m.Offset, string(m.Value))

		kafkaJSON := dts.NewsJSON{}

		log.Logf("got incoming message, subject %s", m.Value)
		json.Unmarshal([]byte(m.Value), &kafkaJSON)

		// save news  to database
		id, err := SaveNews(dbConn, kafkaJSON.Author, kafkaJSON.Body)

		if err != nil {
			log.Logf("Error when to save database  : %v", err.Error())
		}

		log.Logf("save to database ID : %v", id)
		reader.CommitMessages(context.Background(), m)

		// save news to elastic
		jsonNews := id[0]
		elasticConn := conf.ElasticConn(elasticURL)
		CreateDoc(elasticConn, jsonNews.ID, conf.Param.ElasticIndex, jsonNews)
	}
}
