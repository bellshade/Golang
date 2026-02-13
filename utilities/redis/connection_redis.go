package redis

import (
	"fmt"
	"context"

	"github.com/redis/go-redis/v9" // Library Redis untuk Go versi terbaru
)

// Fungsi untuk membuat koneksi ke Redis
func ConnectRedis(ctx context.Context) (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:     "127.21.57.1:6379", // Alamat IP serta Port Redis berjalan
		Password: "Password",         // Password Redis kalian. *jika menggunakan
		DB:       0,
	})

	// Ping koneksi Redis untuk memastikan apakah sudah terkoneksi
	if err := rds.Ping(ctx).Err(); err != nil {
		// Return pesan error untuk tracing jika terjadi error
		return nil, fmt.Errorf("terjadi error: %w", &err)
	}
	return rds, nil
}