package main

import "fmt"

func main() {
	var tipekendaraan string
	var jam1Motor, jam1Mobil, jam2Motor, jam2Mobil, jam1Sepeda, jam2Sepeda, jam1Bus, jam2Bus int
	var jam, menit, tarif, pilihan, pembayaran int
	var stop bool
	stop = false
	jam1Motor = 3000
	jam2Motor = 2000
	jam1Mobil = 5000
	jam2Mobil = 3000
	jam1Sepeda = 1000
	jam2Sepeda = 500
	jam1Bus = 10000
	jam2Bus = 7000
	for !stop {
		fmt.Print("Masukkan tipe kendaraan (Motor/Mobil/Sepeda/Bus): ")
		fmt.Scan(&tipekendaraan)
		for tipekendaraan != "Motor" && tipekendaraan != "Mobil" && tipekendaraan != "Sepeda" && tipekendaraan != "Bus" {
			fmt.Println("Input invalid, kendaraan hanya Motor/Mobil/Sepeda/Bus")
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
			} else if tipekendaraan == "Sepeda" {
				if menit != 0 && jam == 0 {
					tarif = jam1Sepeda
				} else {
					tarif = ((jam - 1) * jam2Sepeda) + jam1Sepeda
					if menit != 0 {
						tarif += jam2Sepeda
					}
				}
			} else if tipekendaraan == "Bus" {
				if menit != 0 && jam == 0 {
					tarif = jam1Bus
				} else {
					tarif = ((jam - 1) * jam2Bus) + jam1Bus
					if menit != 0 {
						tarif += jam2Bus
					}
				}
			}
			fmt.Printf("Tarif parkir yang harus dibayar adalah Rp. %d dengan tipe kendaraan %s \n", tarif, tipekendaraan)
			fmt.Println("\nMetode Pembayaran:")
			fmt.Println("1. QRIS")
			fmt.Println("2. Cash")
			fmt.Print("Pilihan (1/2): ")
			fmt.Scan(&pembayaran)
			for pembayaran != 1 && pembayaran != 2 {
				fmt.Println("Pembayaran tidak valid, masukkan 1 atau 2")
				fmt.Print("Pilihan (1/2): ")
				fmt.Scan(&pembayaran)
			}
			if pembayaran == 1 {
				fmt.Println("Silahkan Scan QRIS.")
			} else {
				fmt.Println("Silahkan Bayar Tunai ke Petugas.")
			}
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
