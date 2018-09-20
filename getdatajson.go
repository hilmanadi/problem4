package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
	"log"
	"encoding/csv"
	"os"
)
type Data struct{
	Data []Dataa `json:"data"`
}
type Dataa struct{
	MuseumId string `json:"museum_id"` 
	KodePengelolaan string `json:"kode_pengelolaan"`
	Nama string `json:"nama"`
	Sdm string `json:"sdm"`
	AlamatJalan string `json:"alamat_jalan"`
	DesaKelurahan string `json:"desa_kelurahan"`
	Kecamatan string `json:"kecamatan"`
	KabupatenKota string `json:"kabupaten_kota"`
	Propinsi string `json:"propinsi"`
	Lintang string `json:"lintang"`
	Bujur string `json:"bujur"`
	Koleksi string `json:"koleksi"`
	SumberDana string `json:"sumber_dana"`
	Pengelola string `json:"pengelola"`
	Tipe string `json:"tipe"`
	Standar string `json:"standar"`
	TahunBerdiri string `json:"tahun_berdiri"`
	Bangunan string `json:"bangunan"`
	LuasTanah string `json:"luas_tanah"`
	StatusKepemilikan string `json:"status_kepemilikan"`
}

func main() {
    //fmt.Println("Starting the application...")
    response, err := http.Get("http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?nama=museum")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    }
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err!=nil{
		log.Fatal(err)
	}
	data = bytes.TrimPrefix(data,[]byte("\xef\xbb\xbf"))
	var respon Data
	json.Unmarshal(data,&respon)
	
	for _,data := range respon.Data{
		if (data.KabupatenKota == "Kota Malang"){
			var d[] 
			d = append(d,data.MuseumId,data.KodePengelolaan,data.Nama,data.Sdm,data.AlamatJalan,
			data.DesaKelurahan,data.Kecamatan,data.KabupatenKota,data.Propinsi,data.Lintang,data.Bujur,
			data.Koleksi,data.SumberDana,data.Pengelola,data.Tipe,data.Standar,data.TahunBerdiri,data.Bangunan,
			data.LuasTanah,data.StatusKepemilikan)
			file,err := os.Create("result.csv")
			fmt.Println(err)
			defer file.Close()
			
			writer := csv.NewWriter(file)
			defer writer.Flush()
			
			for _,value := range d {
				err  := writer.Write(value)
			}
			//fmt.Println("======================= MALANG =============================")
			//fmt.Println(data.MuseumId)
			//fmt.Println(data.KodePengelolaan)
			//fmt.Println(data.Nama)
			//fmt.Println(data.Sdm)
			//fmt.Println(data.AlamatJalan)
			//fmt.Println(data.DesaKelurahan)
			//fmt.Println(data.Kecamatan)
			//fmt.Println(data.KabupatenKota)
			//fmt.Println(data.Propinsi)
			//fmt.Println(data.Lintang)
			//fmt.Println(data.Bujur)
			//fmt.Println(data.Koleksi)
			//fmt.Println(data.SumberDana)
			//fmt.Println(data.Pengelola)
			//fmt.Println(data.Tipe)
			//fmt.Println(data.Standar)
			//fmt.Println(data.TahunBerdiri)
			//fmt.Println(data.Bangunan)
			//fmt.Println(data.LuasTanah)
			//fmt.Println(data.StatusKepemilikan)
			//fmt.Println("====================================================")
			
		}else if (data.KabupatenKota == "Kota Jakarta Pusat"){
			fmt.Println("======================== JAKARTA PUSAT ============================")
			fmt.Println(data.MuseumId)
			fmt.Println(data.KodePengelolaan)
			fmt.Println(data.Nama)
			fmt.Println(data.Sdm)
			fmt.Println(data.AlamatJalan)
			fmt.Println(data.DesaKelurahan)
			fmt.Println(data.Kecamatan)
			fmt.Println(data.KabupatenKota)
			fmt.Println(data.Propinsi)
			fmt.Println(data.Lintang)
			fmt.Println(data.Bujur)
			fmt.Println(data.Koleksi)
			fmt.Println(data.SumberDana)
			fmt.Println(data.Pengelola)
			fmt.Println(data.Tipe)
			fmt.Println(data.Standar)
			fmt.Println(data.TahunBerdiri)
			fmt.Println(data.Bangunan)
			fmt.Println(data.LuasTanah)
			fmt.Println(data.StatusKepemilikan)
			fmt.Println("====================================================")
		}
	}
	
}