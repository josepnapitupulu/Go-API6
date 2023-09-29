package models

import (
	"github.com/jinzhu/gorm"
	"github.com/josepnapitupulu/API_Martumpol/config"
	)

var db *gorm.DB

type Martumpol struct {
	gorm.Model
	NamaMempelaiLaki string `json:"nama_mempelai_laki"`
	AlamatMempelaiLaki string `json:"alamat_mempelai_laki"`
	TanggalMartumpol string `json:"tanggal_martumpol"`
	NamaAyahMempelaiLaki string `json:"nama_ayah_mempelai_laki"`
	NamaIbuMempelaiLaki string `json:"nama_ibu_mempelai_laki"`
	NamaLengkapMempelaiPerempuan string `json:"nama_lengkap_mempelai_perempuan"`
	AlamatMempelaiPerempuan string `json:"alamat_mempelai_perempuan"`
	NamaAyahMempelaiPerempuan string `json:"nama_ayah_mempelai_perempuan"`
	NamaIbuMempelaiPerempuan string `json:"nama_ibu_mempelai_perempuan"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Martumpol{})
}

func (b *Martumpol) CreateMartumpol() *Martumpol {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllMartumpols() []Martumpol {
	var Martumpols []Martumpol
	db.Find(&Martumpols)
	return Martumpols
}

func GetMartumpolbyId(id int64) (*Martumpol, *gorm.DB) {
	var getMartumpol Martumpol
	db := db.Where("ID=?", id).Find(&getMartumpol)
	return &getMartumpol, db
}

func DeleteMartumpol(id int64) Martumpol {
	var martumpol Martumpol
	db.Where("ID=?", id).Delete(martumpol)
	return martumpol 
}