package main

import (
	"context"
	"fmt"
	"log"
	"redis" // Import package Redis tadi, supaya bisa menggunakan fungsi CRUD serta koneksi yang telah dibuat
	"time"
)

func main() {
	var ctx context.Context = context.TODO()

	// Membuat koneksi Redis
	rdsClient, err := redis.ConnectRedis(ctx)
	if err != nil {
		log.Panicf("error koneksi: %v", err)
		return
	}

	// Mengisi struct Redis
	rds := redis.ConstructorRedis(rdsClient)

	// Example redis-set
	var expire = 4*time.Minute
	if err := rds.RedisSet(ctx, "buku:bainn:belajarredis", "ceritanya ini buku", expire); err != nil {
		fmt.Printf("error redis_set: %v", err)
	} else {
		fmt.Println("berhasil tambah data di redis")
	}

	// Example redis-get
	value, err := rds.RedisGet(ctx, "buku:bainn:belajarredis")
	if err != nil {
		fmt.Printf("error redis_get: %v", err)
	} else {
		fmt.Printf("hasil redis_get: %v\n", value)
	}

	// Example redis-del
	if err := rds.RedisDel(ctx, "buku:bainn:belajarredis"); err != nil {
		fmt.Printf("error redis_del: %v", err)
	} else {
		fmt.Println("berhasil hapus data di redis")
	}

}