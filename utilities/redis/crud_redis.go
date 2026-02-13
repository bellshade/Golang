package redis

import (
	"fmt"
	"time"
	"context"

	"github.com/redis/go-redis/v9"
)

// redisCrud adalah interface yang mendefinisikan kontrak untuk fungsi CRUD pada Redis.
// Menggunakan interface membuat kode menjadi lebih fleksibel dan mudah di test
type redisCrud interface {
	RedisGet(ctx context.Context, key string) (any, error)
	RedisSet(ctx context.Context, key string, value any, ttl time.Duration) error
	RedisDel(ctx context.Context, key string) error
}

// Untuk mengimplementasikan interface redisCrud
type rds struct {
	rdsConnection *redis.Client
}

// Untuk menginisialisasi struct Redis dan mengembalikan interface, sehingga fungsi CRUD bisa digunakan
func ConstructorRedis(redisConnection *redis.Client) redisCrud {
	return &rds{
		rdsConnection: redisConnection,
	}
}


// Function untuk mendapatkan data di Redis berdasarkan Key
// *note fungsi get akan mengembalikan nil apabila key tidak ada di Redis
// Parameter - 1.ctx context.Context 2.key string
func (r *rds) RedisGet(ctx context.Context, key string) (any, error) { 
	// Keyword Syntax = Get
	result, err := r.rdsConnection.Get(ctx, key).Result() // Menggunakan method Result supaya mengembalikan string dan err
	if err != nil {
		return nil, fmt.Errorf("terjadi error saat redis_get: %w", err) // Kembalikan error jika ada
	} else if err == redis.Nil {
		return nil, fmt.Errorf("key tersebut tidak ada didalam redis") // Kembalikan pesan jika key tidak ada di redis
	}
	return result, nil
}

// Function untuk menambahkan data ke Redis serta untuk mengupdate data di redis
// *note fungsi set bisa digunakan untuk mengupdate karena fungsi set akan menimpa value sebelumnya apabila key yang digunakan sama
// Parameter - 1.ctx context.Context 2.key string 3.value any 4.ttl time.Duration
func (r *rds) RedisSet(ctx context.Context, key string, value any, ttl time.Duration) error {
	// Keyword Syntax = Set
	// key   -> adalah tanda pengenal untuk mengambil data yang disimpan
	// value -> nilai yang akan disimpan
	// ttl   -> berfungsi sebagai kapan data ini expire atau hangus. misal 13 menit, maka lebih dari itu data hilang
	err := r.rdsConnection.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("terjadi error saat redis_set: %w", err) // Kembalikan error jika ada
	}
	return nil
}

// Function untuk menghapus data di Redis
// *note menghapus data berdasarkan key yang diberikan
// Parameter - 1.ctx context.Context 2.key string
func (r *rds) RedisDel(ctx context.Context, key string) error {
	// Keyword Syntax = Del
	err := r.rdsConnection.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("terjadi error saat redis_delete: %w", err) // Kembalikan error jika ada
	}
	return nil
}

