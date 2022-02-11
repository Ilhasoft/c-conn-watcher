package main

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewDB() *sqlx.DB {
	db, err = sqlx.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

const selectActiveConnectionCountSQL = `
SELECT 
	count(*)
FROM 
	channels_channelconnection
WHERE
	channel_id = $1 AND
	(status = 'W' OR status = 'R' OR status = 'I')
`

const selectChannelByIDSQL = `
SELECT
	id,
	name,
	uuid
FROM
	channels_channel
WHERE
	id = $1
`

func selectActiveConnectionsCount(channelId int64) (int64, error) {
	var count int64 = 0
	cbg := context.Background()
	ctx, cancel := context.WithTimeout(cbg, 30*time.Second)
	defer cancel()
	err := db.GetContext(ctx, &count, selectActiveConnectionCountSQL, channelId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func selectChannelByID(ID int64) (Channel, error) {
	ch := Channel{}
	cbg := context.Background()
	ctx, cancel := context.WithTimeout(cbg, 30*time.Second)
	defer cancel()
	err := db.GetContext(ctx, &ch, selectChannelByIDSQL, ID)
	if err != nil {
		return ch, err
	}
	return ch, nil
}
