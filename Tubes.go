package main

import "fmt"

func main() {
	var tipekendaraan string
	var jam1Motor, jam1Mobil, jam2Motor, jam2Mobil int
	var jam, menit, tarif, pilihan int
	var stop bool
	stop = false
	jam1Motor = 3000
	jam2Motor = 2000
	jam1Mobil = 5000
	jam2Mobil = 3000
	for !stop {
		fmt.Print("Masukkan tipe kendaraan (Motor/Mobil): ")
		fmt.Scan(&tipekendaraan)
		for tipekendaraan != "Motor" && tipekendaraan != "Mobil" {
			fmt.Println("Input invalid, kendaraan hanya Motor/Mobil")
			fmt.Print("Masukkan tipe kendaraan: ")
			fmt.Scan(&tipekendaraan)
		}
		fmt.Print("Masukkan jam parkir: ")
		fmt.Scan(&jam)
		for jam < 0 {
			fmt.Println("Input invalid, jam hanya bisa bernilai 0 atau lebih")
			fmt.Print("Masukkan jam parkir : ")
			fmt.Scan(&jam)
		}
		fmt.Print("Masukkan menit parkir: ")
		fmt.Scan(&menit)
		for menit < 0 || menit > 59 {
			fmt.Println("Input invalid, menit hanya bisa bernilai 0-59")
			fmt.Print("Masukkan menit parkir: ")
			fmt.Scan(&menit)
		}
		if menit == 0 && jam == 0 {
			fmt.Println("Kendaraan tidak parkir")
		} else {
			if tipekendaraan == "Motor" {
				if menit != 0 && jam == 0 {
					tarif = jam1Motor
				} else {
					tarif = ((jam - 1) * jam2Motor) + jam1Motor
					if menit != 0 {
						tarif += jam2Motor
					}
				}
			} else if tipekendaraan == "Mobil" {
				if menit != 0 && jam == 0 {
					tarif = jam1Mobil
				} else {
					tarif = ((jam - 1) * jam2Mobil) + jam1Mobil
					if menit != 0 {
						tarif += jam2Mobil
					}
				}
			}
			fmt.Printf("Tarif parkir yang harus dibayar adalah Rp. %d dengan tipe kendaraan %s \n", tarif, tipekendaraan)
		}
		fmt.Printf("\nApakah ingin memasukkan tarif parkir lagi? \n")
		fmt.Println("1. Iya")
		fmt.Println("2. Tidak")
		fmt.Print("Pilihan (1/2): ")
		fmt.Scan(&pilihan)
		for pilihan != 1 && pilihan != 2 {
			fmt.Println("Pilihan tidak valid, masukkan 1 atau 2")
			fmt.Print("Pilihan (1/2): ")
			fmt.Scan(&pilihan)
		}
		stop = pilihan == 2
	}
	fmt.Println("Program Selesai")
}
