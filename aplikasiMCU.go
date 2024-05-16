package main

import "fmt"

const NMAX int = 100

// type Patient struct{
// 	ID int
// 	Age int
// 	Gender string
// 	Name string
// }

// type tabPatient struct{
// 	Daftar [NMAX]Patient
// 	N int
// }

type Package struct {
	ID       int
	Price    int
	Name     string
	Category string
}

type tabPackage struct {
	Daftar [NMAX]Package
	N      int
}

// type Record struct{
// 	ID int
// 	Year int
// 	Month int
// 	Day int
// 	Package Package
// 	Patient Patient
// 	Result string
// }

// type tabRecord struct{
// 	Daftar [NMAX]Record
// 	N int
// }

// var patients tabPatient
var packages tabPackage

// var records tabRecord

func mainHeader() {
	fmt.Println("==================================")
	fmt.Println("|    Medical Check-Up Program    |")
	fmt.Println("|      By Athaya & M.Regian      |")
	fmt.Println("|     S-1 Informatika IF47-03    |")
	fmt.Println("|              2024              |")
	fmt.Println("============ MAIN MENU ===========")
	fmt.Println("1. Pakcage Data                   ")
	fmt.Println("2. Patient Data                   ")
	fmt.Println("3. Record Data                    ")
	fmt.Println("4. Reports                        ")
	fmt.Println("0. Exit                           ")
	fmt.Println("Choice :                          ")
}

func mainMenu() {
	mainHeader()
	var input int
	fmt.Scan(&input)
	if input == 1 {
		packageManagement()
		// } else if input == 2{

		// } else if input == 3{

		// } else if input == 0{

	}
}

func packageHeader() {
	fmt.Println("===========================================")
	fmt.Println("Main menu >> Package Data                  ")
	fmt.Println("===========================================")
	fmt.Println("Package List                               ")
	fmt.Println("===========================================")
	fmt.Println("ID     ||Name      ||Category    ||Price   ")
	for i := 0; i < packages.N; i++ {
		fmt.Println(packages.Daftar[i].ID, packages.Daftar[i].Name, packages.Daftar[i].Category, packages.Daftar[i].Price)
	}
	fmt.Println("===========================================")
	fmt.Println("Total package registered :        ", packages.N)
	fmt.Println("1. Add package                    ")
	fmt.Println("2. Edit package                   ")
	fmt.Println("3. Delete package                 ")
	fmt.Println("0. Return                         ")
}

func packageManagement() {
	packageHeader()
	var input int
	fmt.Scan(&input)
	if input == 1 {
		addDataPackage()
		packageManagement()
		// } else if input == 2{

		// } else if input == 3{

		// } else if input == 0{

	}
}

func addDataPackage() {
	fmt.Print("Package name :")
	fmt.Scan(&packages.Daftar[packages.N].Name)

	fmt.Print("Package Category (Basic/Standard/Advanced) :")
	fmt.Scan(&packages.Daftar[packages.N].Category)

	fmt.Print("Package price :")
	fmt.Scan(&packages.Daftar[packages.N].Price)
	
	packages.Daftar[packages.N].ID = 10000 + packages.N + 1
	packages.N += 1
}

func main() {
	mainMenu()
}

