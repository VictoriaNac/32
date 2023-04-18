package repository

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type DataBase struct {
	rdb *redis.Client
}


func NewDataBase() (*DataBase, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.pass"),
		DB:       0,
	})

	return &DataBase{rdb: rdb}, nil
}


func (db *DataBase) Close() error {
	err := db.rdb.Close()
	if err != nil {
		return err
	}

	return nil
}


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (db *DataBase) GetLastId() (int, error) {
	ctx := context.Background()
	lastId := 0

	iter := db.rdb.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		id, err := strconv.Atoi(iter.Val())
		lastId = max(lastId, id)

		if err != nil {
			return 0, err
		}
	}

	if err := iter.Err(); err != nil {
		return 0, err
	}

	lastId++
	return lastId, nil
}