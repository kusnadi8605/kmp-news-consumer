# kmp-news-consumer
Golang, Kafka, ElasticSearch &amp; MySql

## Add Library
go get -v github.com/go-sql-driver/mysql  
go get -v gopkg.in/olivere/elastic.v7   
go get -v github.com/segmentio/kafka-go  
go get -v kmp-news-consumer/parser  
go get -v github.com/go-yaml/yaml  


## Create Table news
CREATE TABLE `news` (  
  `id` int(11) NOT NULL AUTO_INCREMENT,  
  `author` text,  
  `body` text,  
  `created` timestamp    
  PRIMARY KEY (`id`)  
);  

# Running Kafka
## Running Zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties
## Running Rafka Server
bin/kafka-server-start.sh config/server.properties

# Running App
go run main.go

## Post News
curl -X POST http://localhost:8181/api/save_news -H 'Content-Type: application/json' -d '{"author":"kusnadi","body":"ini adalah body"}'
