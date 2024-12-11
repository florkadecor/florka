package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataPengeluaran struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB ObjectID
	IDPelanggan primitive.ObjectID `json:"idpelanggan" bson:"idpelanggan"`    // Pelanggan ID as ObjectID
	Jenis       string             `json:"jenis" bson:"jenis"`
	Objek       string             `json:"objek" bson:"objek"`
	Harga       string             `json:"harga" bson:"harga"`
	Keterangan  string             `json:"keterangan" bson:"keterangan"`
	Tanggal     string             `json:"tanggal" bson:"tanggal"`
}
