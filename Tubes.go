package main

import "fmt"

func main() {
	var tipekendaraan string
	var jam1Motor, jam1Mobil, jam2Motor, jam2Mobil int
	var jam, menit, tarif pilihan int
	var stop bool
	stop = false
	jam1Motor = 3000
	jam1Mobil = 2000
	jam2Motor = 5000
	jam2Mobil = 3000
	for !stop {
		fmt.Print("Masukkan tipe kendaraan : ")
		fmt.Scan(&tipekendaraan)
		for tipekendaraan != "Motor" && tipekendaraan != "Mobil" {
			fmt.Println(" Input invalid, kendaraan hanya Motor/Mobil")
			fmt.Print("Masukkan tipe kendaraan : ")
			fmt.Scan(&tipekendaraan)
		}
		fmt.Print("Masukkan jam parkir : ")
		fmt.Scan(&jam)
		for jam < 0 {
			fmt.Println(" Input invalid, jam hanya bisa bernilai 0 atau lebih")
			fmt.Print("Masukkan jam parkir : ")
			fmt.Scan(&jam)
		}
		fmt.Print("Masukkan menit parkir : ")
		fmt.Scan(&menit)
		for menit < 0 {
			fmt.Println(" Input invalid, menit hanya bisa bernilai 0 atau lebih")
			fmt.Print("Masukkan menit parkir : ")
			fmt.Scan(&menit)
			
		}

	}