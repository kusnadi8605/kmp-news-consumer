# kmp-news-consumer
Golang, Kafka, ElasticSearch &amp; MySql

# Create table news
CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author` text,
  `body` text,
  `created` timestamp 
  PRIMARY KEY (`id`)
);

# Post News
curl -X POST http://localhost:8181/api/save_news -H 'Content-Type: application/json' -d '{"author":"kusnadi","body":"ini adalah body"}'
