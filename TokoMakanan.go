package main

import (
	"fmt"
	"strings"
)

var menu []string = []string{"Tahu", "Tempe", "Sambal", "Nasi", "Lele", "Ayam"}

func main(){
	// Risky Kurniawan - ARS University
	var orders []string
	fmt.Println("===================================")
	fmt.Println("      TOKO MAKANAN INDONESIA")
	fmt.Println("===================================")
	for _, list := range menu {
		fmt.Println(list)
	}
	fmt.Println("===================================")

	PesanLagi:
	var pesanan, next string = "",""
	fmt.Print("Masukan pesanan anda dalam huruf ( eg: Tahu ) : ")
	fmt.Scanf("%s\n", &pesanan)
	if valid, err := validation(pesanan); valid == false {
		fmt.Println(err)
		goto PesanLagi
	}else{
		orders = append(orders, pesanan)
	}
	PilNext:
	fmt.Print("Lanjutkan Memesan ? (Y/T) ")
	fmt.Scanf("%s\n", &next)

	if strings.ToLower(next) == "y" {
		goto PesanLagi
	}else if strings.ToLower(next) != "y" && strings.ToLower(next) != "t"{
		fmt.Println("Keyword salah")
		goto PilNext
	}

	for _, order := range orders {
		fmt.Println("Pesanan anda : ",order)
	}	
	fmt.Println("===================================")
	fmt.Println("        ** Terimakasih **")
	fmt.Println("By Risky Kurniawan - ARS University")
	fmt.Println("===================================")
}

func validation(pesanan string) (bool, string) {
	if pesanan == "" {
		return false, "Pesanan anda belum ditulis, silahkan tuliskan kembali pesanan anda"
	}
	for _, list := range menu {
		if strings.ToLower(list) == strings.ToLower(pesanan){
			return true, ""
		}
	}

	return false, "Pesanan tidak tersedia pada menu, silahkan tuliskan kembali pesanan anda"
}