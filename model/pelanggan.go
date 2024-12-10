package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataPelanggan struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`    // MongoDB ObjectID
	NamaPelanggan  string             `bson:"namapelanggan" json:"namapelanggan"`   // Nama pelanggan
	NomorPelanggan string             `bson:"nomorpelanggan" json:"nomorpelanggan"` // Nomor pelanggan
	NamaPasangan   string             `bson:"namapasangan" json:"namapasangan"`     // Nama pasangan
	Alamat         string             `bson:"alamat" json:"alamat"`                 // Alamat
	Lokasi         string             `bson:"lokasi" json:"lokasi"`                 // Lokasi
	Tanggal        string             `bson:"tanggal" json:"tanggal"`               // Tanggal (format string, bisa diubah ke time.Time jika perlu)
	Paket          string             `bson:"paket" json:"paket"`                   // Paket
	Harga          float64            `bson:"harga" json:"harga"`                   // Harga (tipe float untuk menyimpan angka desimal)
}
