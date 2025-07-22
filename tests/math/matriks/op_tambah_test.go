package matriks

import (
	"testing"

	"github.com/bellshade/Golang/math/matriks"
	"github.com/bellshade/Golang/tests"
)

func TestTambah(t *testing.T) {
	testingData := []struct {
		name    string
		a       matriks.Matriks
		b       matriks.Matriks
		want    matriks.Matriks
		wantErr bool
	}{
		{
			name: "dua matriks 2x2 valid",
			a: matriks.Matriks{
				{1, 2},
				{3, 4},
			},
			b: matriks.Matriks{
				{5, 6},
				{7, 8},
			},
			want: matriks.Matriks{
				{6, 8},
				{10, 12},
			},
			wantErr: false,
		},
		{
			name: "matriks 3x1",
			a: matriks.Matriks{
				{1},
				{2},
				{3},
			},
			b: matriks.Matriks{
				{4},
				{5},
				{6},
			},
			want: matriks.Matriks{
				{5},
				{7},
				{9},
			},
			wantErr: false,
		},
		{
			name: "jumlah baris tidak sama",
			a: matriks.Matriks{
				{1, 2},
			},
			b: matriks.Matriks{
				{3, 4},
				{5, 6},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "jumlah kolom tidak sama di baris pertama",
			a: matriks.Matriks{
				{1, 2, 3},
				{4, 5},
			},
			b: matriks.Matriks{
				{1, 2},
				{3, 4},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "kedua matriks kosong",
			a:       matriks.Matriks{},
			b:       matriks.Matriks{},
			want:    matriks.Matriks{},
			wantErr: false,
		},
		{
			name: "matriks dengan baris kosong (valid)",
			a: matriks.Matriks{
				{},
				{},
			},
			b: matriks.Matriks{
				{},
				{},
			},
			want: matriks.Matriks{
				{},
				{},
			},
			wantErr: false,
		},
	}

	for _, tt := range testingData {
		t.Run(tt.name, func(t *testing.T) {
			got, err := matriks.Tambah(tt.a, tt.b)

			if tt.wantErr {
				tests.Assert(t, err != nil, "error malah none reks")
				return
			}

			tests.OK(t, err)
			tests.Equals(t, tt.want, got)
		})
	}
}
