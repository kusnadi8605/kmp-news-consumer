package models

import (
	conf "kmp-news-consumer/config"
	dts "kmp-news-consumer/datastruct" //
	//
	"time"
)

//SaveNews .
func SaveNews(conn *conf.Connection, author string, body string) ([]dts.SaveNewsJSON, error) {
	arrNewsSave := []dts.SaveNewsJSON{}
	strNewsSave := dts.SaveNewsJSON{}

	tx, err := conn.Begin()
	if err != nil {
		return arrNewsSave, err
	}

	created := time.Now().Format("2006-01-02 15:04:05")

	stmt, err := tx.Prepare(`insert into news (author,body,created)
								   values (?,?,?)`)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(author, body, created)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	rows, err := res.LastInsertId()

	strNewsSave.ID = rows
	strNewsSave.Created = created
	arrNewsSave = append(arrNewsSave, strNewsSave)

	if err != nil {
		return nil, err
	}

	return arrNewsSave, nil
}
