package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Martumpol/models"
	"github.com/josepnapitupulu/API_Martumpol/utils"
)

var NewMartumpol models.Martumpol
// var tmpl * template.Template
// func init(){
// 	tmpl = template.Must(template.ParseFiles("jemaat.html"))
// }

// func Index(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/jemaat.html")
// 	temp.Execute(w, nil)
// }

// func Create(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/createJemaat.html")
// 	temp.Execute(w, nil)
// }

func GetMartumpol(w http.ResponseWriter, r *http.Request) {
	newMartumpols := models.GetAllMartumpols()
	res, _ := json.Marshal(newMartumpols)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMartumpolById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	martumpolId := vars["martumpolId"]
	IdDaftar, err := strconv.ParseInt(martumpolId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	martumpolDetails, _ := models.GetMartumpolbyId(IdDaftar)
	res, _ := json.Marshal(martumpolDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMartumpol(w http.ResponseWriter, r *http.Request) {
	CreateMartumpol := &models.Martumpol{}
	utils.ParseBody(r, CreateMartumpol)
	b := CreateMartumpol.CreateMartumpol()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMartumpol(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	martumpolId := vars["martumpolId"]
	ID, err := strconv.ParseInt(martumpolId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	martumpol := models.DeleteMartumpol(ID)
	res, _ := json.Marshal(martumpol)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMartumpol(w http.ResponseWriter, r *http.Request) {
    var updateMartumpol = &models.Martumpol{}
    utils.ParseBody(r, updateMartumpol)
    vars := mux.Vars(r)
    martumpolId := vars["martumpolId"]
    ID, err := strconv.ParseInt(martumpolId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    martumpolDetails, db := models.GetMartumpolbyId(ID)
    if updateMartumpol.NamaMempelaiLaki != "" {
        martumpolDetails.NamaMempelaiLaki = updateMartumpol.NamaMempelaiLaki
    }
    if updateMartumpol.AlamatMempelaiLaki != "" {
        martumpolDetails.AlamatMempelaiLaki = updateMartumpol.AlamatMempelaiLaki
    }
    if updateMartumpol.TanggalMartumpol != "" {
        martumpolDetails.TanggalMartumpol = updateMartumpol.TanggalMartumpol
    }
    if updateMartumpol.NamaAyahMempelaiLaki != "" {
        martumpolDetails.NamaAyahMempelaiLaki = updateMartumpol.NamaAyahMempelaiLaki
    }
    if updateMartumpol.NamaIbuMempelaiLaki != "" {
        martumpolDetails.NamaIbuMempelaiLaki = updateMartumpol.NamaIbuMempelaiLaki
    }
	if updateMartumpol.NamaLengkapMempelaiPerempuan != "" {
        martumpolDetails.NamaLengkapMempelaiPerempuan = updateMartumpol.NamaLengkapMempelaiPerempuan
    }
	if updateMartumpol.AlamatMempelaiPerempuan != "" {
        martumpolDetails.AlamatMempelaiPerempuan = updateMartumpol.AlamatMempelaiPerempuan
    }
	if updateMartumpol.NamaAyahMempelaiPerempuan != "" {
        martumpolDetails.NamaAyahMempelaiPerempuan = updateMartumpol.NamaAyahMempelaiPerempuan
    }
	if updateMartumpol.NamaIbuMempelaiPerempuan != "" {
        martumpolDetails.NamaIbuMempelaiPerempuan = updateMartumpol.NamaIbuMempelaiPerempuan
    }
    db.Save(&martumpolDetails)
    res, _ := json.Marshal(martumpolDetails)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}