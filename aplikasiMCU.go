package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func inputKalimat() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Menghapus newline di akhir input menggunakan TrimSpace untuk menangani semua kasus
	input = strings.TrimSpace(input)

	return input
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

const NMAX int = 100

type Patient struct {
	ID     int    `json:"patient_id"`
	Name   string `json:"patient_name"`
	Gender string `json:"patient_gender"`
	Age    int    `json:"patient_age"`
}

type tabPatient struct {
	Daftar [NMAX]Patient `json:"tab_patient_daftar"`
	N      int           `json:"tab_patient_n"`
}

type Package struct {
	ID       int    `json:"package_id"`
	Name     string `json:"package_name"`
	Category string `json:"package_category"`
	Price    int    `json:"package_price"`
}

type tabPackage struct {
	Daftar [NMAX]Package `json:"tab_package_daftar"`
	N      int           `json:"tab_package_n"`
}

type Record struct {
	ID      int     `json:"record_ID"`
	Patient Patient `json:"record_patient"`
	Package Package `json:"record_package"`
	Year    int     `json:"record_year"`
	Month   int     `json:"record_month"`
	Day     int     `json:"record_day"`
	Result  string  `json:"record_result"`
}

type tabRecord struct {
	Daftar [NMAX]Record `json:"tab_record_daftar"`
	N      int          `json:"tab_record_n"`
}

var patients tabPatient

var packages tabPackage

var records tabRecord

// MAIN MENU MAIN MENU MAIN MENU MAIN MENU MAIN MENU MAIN MENU  MAIN MENU MAIN MENU
// MAIN MENU MAIN MENU MAIN MENU MAIN MENU MAIN MENU MAIN MENU  MAIN MENU MAIN MENU
// MAIN MENU MAIN MENU MAIN MENU MAIN MENU MAIN MENU MAIN MENU  MAIN MENU MAIN MENU

func main_header() {
	fmt.Println("==================================")
	fmt.Println("|    Medical Check-Up Program    |")
	fmt.Println("|      By Athaya & M.Regian      |")
	fmt.Println("|     S-1 Informatika IF47-03    |")
	fmt.Println("|              2024              |")
	fmt.Println("============ MAIN MENU ===========")
	fmt.Println("1. Package Data                   ")
	fmt.Println("2. Patient Data                   ")
	fmt.Println("3. Record Data                    ")
	fmt.Println("4. Reports                        ")
	fmt.Println("0. Exit                           ")
	fmt.Print("Choice : ")
}

func main_menu() {
	ClearScreen()
	main_header()
	var choice int
	var validChoice bool = false
	for !validChoice {
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			package_main()
			validChoice = true
		case 2:
			patient_main()
			validChoice = true
		case 3:
			record_main()
			validChoice = true
		case 4:
			report_main()
			validChoice = true
		case 0:
			ClearScreen()
			validChoice = true
		default:
			fmt.Print("Invalid! Choice : ")
		}
	}

}

// PACKAGE MANAGEMENT PACKAGE MANAGEMENT PACKAGE MANAGEMENT PACKAGE MANAGEMENT
// PACKAGE MANAGEMENT PACKAGE MANAGEMENT PACKAGE MANAGEMENT PACKAGE MANAGEMENT
// PACKAGE MANAGEMENT PACKAGE MANAGEMENT PACKAGE MANAGEMENT PACKAGE MANAGEMENT

func package_header() {
	fmt.Println("===================================================================")
	fmt.Println("Main menu >> Package                       ")
	fmt.Println("===================================================================")
	fmt.Println("Package List                               ")
	fmt.Println("===================================================================")
	print_package_header()
	fmt.Println("===================================================================")
	fmt.Println("Total package registered :", packages.N)
	fmt.Println("1. Add package                    ")
	fmt.Println("2. Edit package                   ")
	fmt.Println("3. Delete package                   ")
	fmt.Println("0. Return                         ")
	fmt.Print("Choice : ")
}

func print_package_header() {
	fmt.Printf("%10v %30v %12v %12v\n", "ID", "Name", "Category", "Price")
	for i := 0; i < packages.N; i++ {
		fmt.Printf("%10v %30v %12v %12v\n", packages.Daftar[i].ID, packages.Daftar[i].Name, packages.Daftar[i].Category, packages.Daftar[i].Price)
	}
}

func package_main() {
	ClearScreen()
	package_header()
	var choice int
	var validChoice bool = false
	for !validChoice {
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			add_package()
			validChoice = true
		case 2:
			edit_package()
			validChoice = true
		case 3:
			delete_package()
			validChoice = true
		case 0:
			main_menu()
			validChoice = true
		default:
			fmt.Print("Invalid! Choice : ")
		}
	}

}

// PROCEDURE UNTUK MENAMBAH DATA PAKET YANG TERSEDIA
func add_package() {
	var paket Package
	var finish bool = false
	for !finish {
		fmt.Println("Type 'finish' if you have finished inputting")
		fmt.Print("Package name : ")
		paket.Name = inputKalimat()
		switch paket.Name {
		case "finish":
			finish = true
		default:

			fmt.Print("Package Category (Basic/Standard/Advanced) : ")
			fmt.Scanln(&paket.Category)

			for paket.Category != "Basic" && paket.Category != "Standard" && paket.Category != "Advanced" {
				fmt.Print("invalid!")
				fmt.Print("Package Category (Basic/Standard/Advanced) : ")
				fmt.Scanln(&paket.Category)
			}

			fmt.Print("Package price : ")
			fmt.Scanln(&paket.Price)

			packages.Daftar[packages.N].ID = 10000 + packages.N + 1
			packages.Daftar[packages.N].Name = paket.Name
			packages.Daftar[packages.N].Category = paket.Category
			packages.Daftar[packages.N].Price = paket.Price
			packages.N += 1

		}
	}
	package_main()
}

func edit_package() {
	var name string
	fmt.Print("Input package name : ")
	name = inputKalimat()
	idx := search_package(name)
	if idx != -1 {
		var newName, newCategory, accept string
		var newPrice int

		fmt.Print("Input new name : ")
		newName = inputKalimat()
		fmt.Print("Input new category (Basic/Standard/Advanced) : ")
		fmt.Scanln(&newCategory)
		for newCategory != "Basic" && newCategory != "Standard" && newCategory != "Advanced" {
			fmt.Println("invalid!")
			fmt.Print("input new category (Basic/Standard/Advanced) : ")
			fmt.Scanln(&newCategory)
		}
		fmt.Print("Input new price : ")
		fmt.Scanln(&newPrice)

		fmt.Print("Accept changes (Y/N)? ")
		fmt.Scanln(&accept)

		for accept != "Y" && accept != "N" {
			fmt.Println("Invalid!")
			fmt.Print("Accept changes (Y/N)? ")
			fmt.Scanln(&accept)
		}

		if accept == "Y" {
			packages.Daftar[idx].Name = newName
			packages.Daftar[idx].Category = newCategory
			packages.Daftar[idx].Price = newPrice
		}

		ClearScreen()
		package_main()
	} else {
		fmt.Print("Package not found! (input anything to return) ")
		var input string
		fmt.Scanln(&input)
		package_main()
	}
}

func search_package(name string) int {
	for i := 0; i < packages.N; i++ {
		if packages.Daftar[i].Name == name {
			return i
		}
	}
	return -1
}

func delete_package() {
	var name string
	fmt.Print("Input package name : ")
	name = inputKalimat()
	idx := search_package(name)

	if idx != -1 {
		var accept string
		fmt.Print("Delete package (Y/N)? ")
		fmt.Scanln(&accept)

		for accept != "Y" && accept != "N" {
			fmt.Println("Invalid!")
			fmt.Print("Delete package (Y/N)? ")
			fmt.Scanln(&accept)
		}

		if accept == "Y" {
			var i int
			for i = idx; i < packages.N-1; i++ {
				packages.Daftar[i].Name = packages.Daftar[i+1].Name
				packages.Daftar[i].Category = packages.Daftar[i+1].Category
				packages.Daftar[i].Price = packages.Daftar[i+1].Price
			}
			packages.Daftar[i].ID = 0
			packages.Daftar[i].Name = ""
			packages.Daftar[i].Category = ""
			packages.Daftar[i].Price = 0
			packages.N--
		}

		package_main()
	} else {
		fmt.Print("Package not found! (input anything to return) ")
		var input string
		fmt.Scanln(&input)
		package_main()
	}
}

// // PATIENT MANAGEMENT PATIENT MANAGEMENT PATIENT MANAGEMENT PATIENT MANAGEMENT
// // PATIENT MANAGEMENT PATIENT MANAGEMENT PATIENT MANAGEMENT PATIENT MANAGEMENT
// // PATIENT MANAGEMENT PATIENT MANAGEMENT PATIENT MANAGEMENT PATIENT MANAGEMENT

func patient_header() {
	fmt.Println("==============================================")
	fmt.Println("Main menu >> Patient                       ")
	fmt.Println("==============================================")
	fmt.Println("Patient List                               ")
	fmt.Println("==============================================")
	print_patient_header()
	fmt.Println("==============================================")
	fmt.Println("Total patient registered :", patients.N)
	fmt.Println("1. Add patient                    ")
	fmt.Println("2. Edit patient                   ")
	fmt.Println("3. Delete patient                 ")
	fmt.Println("4. Search patient                 ")
	fmt.Println("5. Patient detail                 ")
	fmt.Println("0. Return                         ")
	fmt.Print("Choice : ")
}

func print_patient_header() {
	fmt.Printf("%10v %17v %8v %8v\n", "ID", "Name", "Age", "Gender")
	for i := 0; i < patients.N; i++ {
		fmt.Printf("%10v %17v %8v %8v \n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Age, patients.Daftar[i].Gender)
	}
}

func patient_main() {
	ClearScreen()
	patient_header()
	var choice int
	var validChoice bool = false
	for !validChoice {
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			add_patient()
			validChoice = true
		case 2:
			edit_patient()
			validChoice = true
		case 3:
			delete_patient()
			validChoice = true
		case 4:
			search_patient_menu()
			validChoice = true
		case 5:
			main_patient_detail()
			validChoice = true
		case 0:
			main_menu()
			validChoice = true
		default:
			fmt.Print("Invalid! Choice : ")
		}
	}
}

// PROCEDURE UNTUK MENABAHKAN DATA PASIEN
func add_patient() {
	var pasien Patient
	var finish bool = false
	for !finish {
		fmt.Println("Type finish if you have finished inputting")
		fmt.Print("Patient name : ")
		pasien.Name = inputKalimat()
		switch pasien.Name {
		case "finish":
			finish = true
		default:
			fmt.Print("Patient age : ")
			fmt.Scanln(&pasien.Age)

			fmt.Print("Patient gender(M/F) : ")
			fmt.Scanln(&pasien.Gender)

			for pasien.Gender != "M" && pasien.Gender != "F" {
				fmt.Println("Invalid!")
				fmt.Print("Patient Gender : ")
				fmt.Scanln(&pasien.Gender)
			}

			patients.Daftar[patients.N].ID = 70000 + patients.N + 1
			patients.Daftar[patients.N].Name = pasien.Name
			patients.Daftar[patients.N].Age = pasien.Age
			patients.Daftar[patients.N].Gender = pasien.Gender

			patients.N += 1

		}
	}

	patient_main()
}

func edit_patient() {
	var name string
	fmt.Print("Input patient name : ")
	name = inputKalimat()
	idx := search_patient(name)
	if idx != -1 {
		var newName, newGender, accept string
		var newAge int

		fmt.Print("Input new name : ")
		newName = inputKalimat()
		fmt.Print("Input new age : ")
		fmt.Scanln(&newAge)
		fmt.Print("Input new gender (M/F) : ")
		fmt.Scanln(&newGender)
		for newGender != "M" && newGender != "F" {
			fmt.Println("Invalid!")
			fmt.Print("Input new gender (M/F) : ")
			fmt.Scanln(&newGender)
		}

		fmt.Print("Accept changes (Y/N)? ")
		fmt.Scanln(&accept)

		for accept != "Y" && accept != "N" {
			fmt.Println("Invalid!")
			fmt.Print("Accept changes (Y/N)? ")
			fmt.Scanln(&accept)
		}

		if accept == "Y" {
			patients.Daftar[idx].Name = newName
			patients.Daftar[idx].Age = newAge
			patients.Daftar[idx].Gender = newGender
		}

		patient_main()
	} else {
		fmt.Print("Patient not found! (input anything to return) ")
		var input string
		fmt.Scanln(&input)
		patient_main()
	}

}

func delete_patient() {
	var name string
	fmt.Print("Input patient name : ")
	name = inputKalimat()
	idx := search_patient(name)

	if idx != -1 {
		var accept string
		fmt.Print("Delete patient (Y/N)? ")
		fmt.Scanln(&accept)

		for accept != "Y" && accept != "N" {
			fmt.Println("Invalid!")
			fmt.Print("Delete patient (Y/N)? ")
			fmt.Scanln(&accept)
		}

		if accept == "Y" {
			var i int
			for i = idx; i < patients.N-1; i++ {
				patients.Daftar[i].Name = patients.Daftar[i+1].Name
				patients.Daftar[i].Age = patients.Daftar[i+1].Age
				patients.Daftar[i].Gender = patients.Daftar[i+1].Gender
			}
			patients.Daftar[i].ID = 0
			patients.Daftar[i].Name = ""
			patients.Daftar[i].Age = 0
			patients.Daftar[i].Gender = ""
			patients.N--
		}

		patient_main()
	} else {
		fmt.Print("Patient not found! (input anything to return) ")
		var input string
		fmt.Scanln(&input)
		patient_main()
	}
}

func patient_search() {
	fmt.Println("==============================================")
	fmt.Println("Main menu >> Patient >> Search             ")
	fmt.Println("==============================================")
	fmt.Println("Patient List                               ")
	fmt.Println("==============================================")
	print_patient_header()
	fmt.Println("==============================================")
	fmt.Println("Total patient registered :", patients.N)
	fmt.Println("1. Search by name                    ")
	fmt.Println("2. Search by package                   ")
	fmt.Println("3. Search by period                 ")
	fmt.Println("0. Return                         ")
	fmt.Print("Choice : ")
}

func search_patient(name string) int {
	for i := 0; i < patients.N; i++ {
		if patients.Daftar[i].Name == name {
			return i
		}
	}
	return -1
}

func search_binary_patient(name string) int {
	var left, right, mid int
	var temp tabPatient = patients
	sort_patient_by_name(&temp)
	idx := -1
	left = 0
	right = patients.N - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if temp.Daftar[mid].Name > name {
			right = mid - 1
		} else if temp.Daftar[mid].Name < name {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func sort_patient_by_name(A *tabPatient) {
	var temp Patient
	pass := 1
	for pass < A.N {
		idx := pass - 1
		i := pass
		for i < A.N {
			if A.Daftar[idx].Name > A.Daftar[i].Name {
				idx = i
			}
			i++
		}
		temp = A.Daftar[pass-1]
		A.Daftar[pass-1] = A.Daftar[idx]
		A.Daftar[idx] = temp
		pass++
	}
}

func search_patient_menu() {
	ClearScreen()
	patient_search()
	var choice int
	var validChoice bool = false
	for !validChoice {
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			search_by_name()
			validChoice = true
		case 2:
			search_by_package()
			validChoice = true
		case 3:
			search_by_period()
			validChoice = true
		case 0:
			patient_main()
			validChoice = true
		default:
			fmt.Println("Invalid! Choice:")
		}
	}
}

func patient_detail(idxPatient, idxRecord int) {
	ClearScreen()
	var temp tabPatient = patients
	sort_patient_by_name(&temp)
	fmt.Println("===============================================")
	fmt.Println("Main menu >> Patient >> Search >> Patient", temp.Daftar[idxPatient].ID)
	fmt.Println("===============================================")
	fmt.Println("ID patient        :", temp.Daftar[idxPatient].ID)
	fmt.Println("Patient name      :", temp.Daftar[idxPatient].Name)
	fmt.Println("Patient age       :", temp.Daftar[idxPatient].Age)
	fmt.Println("Patient gender    :", temp.Daftar[idxPatient].Gender)
	if idxRecord != -1 {
		fmt.Println("Package taken     :", records.Daftar[idxRecord].Package.Name)
		fmt.Printf("Date of MCu       : %v-%v-%v \n", records.Daftar[idxRecord].Day, records.Daftar[idxRecord].Month, records.Daftar[idxRecord].Year)
		fmt.Println("Result            :", records.Daftar[idxRecord].Result)
	} else {
		fmt.Println("Package           : <null>")
		fmt.Println("Date of MCU       : <null>")
		fmt.Println("Result            : <null>")
	}
	fmt.Println("===============================================")
}

func search_by_name() {
	var name string
	fmt.Print("Input name : ")
	name = inputKalimat()
	idxPatient := search_binary_patient(name)
	idxRecord := -1
	for i := 0; i < records.N; i++ {
		if records.Daftar[i].Patient.Name == name {
			idxRecord = i
		}
	}
	if idxPatient == -1 {
		fmt.Println("Patient not found!")
	} else {
		patient_detail(idxPatient, idxRecord)
	}
	fmt.Print("Input anything to return : ")
	var input string
	fmt.Scanln(&input)
	search_patient_menu()
}

func search_by_package() {
	var paket string
	fmt.Print("Input package name : ")
	paket = inputKalimat()
	ClearScreen()
	fmt.Println("==================================================================")
	fmt.Println("Main menu >> Patient >> Search >> Package", paket)
	fmt.Println("==================================================================")

	fmt.Printf("%10v %12v %10v %10v %20v\n", "ID", "Date", "Patient", "Package", "Result")
	for i := 0; i < records.N; i++ {
		if records.Daftar[i].Package.Name == paket {
			fmt.Printf("%10v %4v-%2v-%4v %10v %10v %20v\n", records.Daftar[i].ID, records.Daftar[i].Day, records.Daftar[i].Month, records.Daftar[i].Year, records.Daftar[i].Patient.Name, records.Daftar[i].Package.Name, records.Daftar[i].Result)
		}
	}
	fmt.Println("==================================================================")
	fmt.Print("Input anything to return : ")
	var input string
	fmt.Scanln(&input)
	search_patient_menu()
}

func search_by_period() {
	var month, year int
	var bulan string
	fmt.Print("Input month : ")
	fmt.Scanln(&month)
	fmt.Print("Input year : ")
	fmt.Scanln(&year)
	switch month {
	case 1:
		bulan = "January"
	case 2:
		bulan = "February"
	case 3:
		bulan = "March"
	case 4:
		bulan = "April"
	case 5:
		bulan = "May"
	case 6:
		bulan = "June"
	case 7:
		bulan = "July"
	case 8:
		bulan = "August"
	case 9:
		bulan = "September"
	case 10:
		bulan = "October"
	case 11:
		bulan = "November"
	case 12:
		bulan = "December"
	}
	ClearScreen()
	fmt.Println("==================================================================")
	fmt.Printf("Main menu >> Patient >> Search >> Period %v-%v\n", bulan, year)
	fmt.Println("==================================================================")

	fmt.Printf("%10v %12v %10v %10v %20v\n", "ID", "Date", "Patient", "Package", "Result")
	for i := 0; i < records.N; i++ {
		if records.Daftar[i].Month == month && records.Daftar[i].Year == year {
			fmt.Printf("%10v %4v-%2v-%4v %10v %10v %20v\n", records.Daftar[i].ID, records.Daftar[i].Day, records.Daftar[i].Month, records.Daftar[i].Year, records.Daftar[i].Patient.Name, records.Daftar[i].Package.Name, records.Daftar[i].Result)
		}
	}
	fmt.Println("==================================================================")
	fmt.Print("Input anything to return : ")
	var input string
	fmt.Scanln(&input)
	search_patient_menu()
}

func patient_header_detail() {
	ClearScreen()
	fmt.Println("===================================================================================================================================")
	fmt.Println("Main menu >> Patient >> Patient detail     ")
	fmt.Println("===================================================================================================================================")
	fmt.Println("Patient List                               ")
	fmt.Println("===================================================================================================================================")
	print_patient_header_detail()
	fmt.Println("===================================================================================================================================")
	fmt.Println("Total patient registered :", patients.N)
	fmt.Println("1. Sort by date ascending                    ")
	fmt.Println("2. Sort by date descending                   ")
	fmt.Println("3. Sort by package ascending                 ")
	fmt.Println("4. Sort by package descending                ")
	fmt.Println("0. Return                         ")
	fmt.Print("Choice : ")
}

func print_patient_header_detail() {
	fmt.Printf("%10v %17v %8v %8v %27v %10v %45v\n", "ID", "Name", "Age", "Gender", "Package", "Date", "Result")
	for i := 0; i < patients.N; i++ {
		idx := -1
		for j := 0; j < records.N; j++ {
			if patients.Daftar[i].Name == records.Daftar[j].Patient.Name {
				idx = j
			}
		}
		if idx != -1 {
			fmt.Printf("%10v %17v %8v %8v %27v %2v-%2v-%4v %45v\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Age, patients.Daftar[i].Gender, records.Daftar[idx].Package.Name, records.Daftar[idx].Day, records.Daftar[idx].Month, records.Daftar[idx].Year, records.Daftar[idx].Result)
		} else {
			fmt.Printf("%10v %17v %8v %8v %27v %10v %45v\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Age, patients.Daftar[i].Gender, "<null>", "<null>", "<null>")
		}
	}
}

func main_patient_detail() {
	ClearScreen()
	patient_header_detail()
	var choice int
	var validChoice bool = false
	var tempRecord tabRecord = records
	for !validChoice {
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			sort_by_date_asc(&tempRecord)
			patient_header_detail_sorted(tempRecord)
		case 2:
			sort_by_date_desc(&tempRecord)
			patient_header_detail_sorted(tempRecord)
		case 3:
			sort_by_package_asc(&tempRecord)
			patient_header_detail_sorted(tempRecord)
		case 4:
			sort_by_package_desc(&tempRecord)
			patient_header_detail_sorted(tempRecord)
		case 0:
			validChoice = true
			patient_main()
		default:
			fmt.Print("Invalid! Choice : ")
		}
	}
}

func patient_header_detail_sorted(A tabRecord) {
	ClearScreen()
	fmt.Println("===================================================================================================================================")
	fmt.Println("Main menu >> Patient >> Patient detail     ")
	fmt.Println("===================================================================================================================================")
	fmt.Println("Patient List                               ")
	fmt.Println("===================================================================================================================================")
	print(A)
	patient_no_record()
	fmt.Println("===================================================================================================================================")
	fmt.Println("Total patient registered :", patients.N)
	fmt.Println("1. Sort by date ascending                    ")
	fmt.Println("2. Sort by date descending                   ")
	fmt.Println("3. Sort by package ascending                 ")
	fmt.Println("4. Sort by package descending                ")
	fmt.Println("0. Return                         ")
	fmt.Print("Choice : ")
}

func print(A tabRecord) {
	fmt.Printf("%10v %17v %8v %8v %27v %10v %45v\n", "ID", "Name", "Age", "Gender", "Package", "Date", "Result")
	for i := 0; i < A.N; i++ {
		fmt.Printf("%10v %17v %8v %8v %27v %2v-%2v-%4v %45v\n", A.Daftar[i].Patient.ID, A.Daftar[i].Patient.Name, A.Daftar[i].Patient.Age, A.Daftar[i].Patient.Gender, A.Daftar[i].Package.Name, A.Daftar[i].Day, A.Daftar[i].Month, A.Daftar[i].Year, A.Daftar[i].Result)

	}
}

func patient_no_record() {
	var temp tabRecord
	var exist bool
	k := 0
	for i := 0; i < patients.N; i++ {
		exist = false
		for j := 0; j < records.N; j++ {
			if patients.Daftar[i].Name == records.Daftar[j].Patient.Name {
				exist = true
			}
		}
		if !exist {
			temp.Daftar[k].Patient = patients.Daftar[i]
			k++
		}
	}
	for i := 0; i < k; i++ {
		fmt.Printf("%10v %17v %8v %8v %27v %10v %45v\n", temp.Daftar[i].Patient.ID, temp.Daftar[i].Patient.Name, temp.Daftar[i].Patient.Age, temp.Daftar[i].Patient.Gender, "<null>", "<null>", "<null>")
	}
}

func sort_by_date_asc(A *tabRecord) {
	var temp Record
	pass := 1
	for pass < A.N {
		idx := pass - 1
		i := pass
		for i < A.N {
			if A.Daftar[idx].Year > A.Daftar[i].Year {
				idx = i
			} else if A.Daftar[idx].Year == A.Daftar[i].Year && A.Daftar[idx].Month > A.Daftar[i].Month {
				idx = i
			} else if A.Daftar[idx].Year == A.Daftar[i].Year && A.Daftar[idx].Month == A.Daftar[i].Month && A.Daftar[idx].Day > A.Daftar[i].Day {
				idx = i
			}
			i++
		}
		temp = A.Daftar[pass-1]
		A.Daftar[pass-1] = A.Daftar[idx]
		A.Daftar[idx] = temp
		pass++
	}
}

func sort_by_date_desc(A *tabRecord) {
	var temp Record
	pass := 1
	for pass < A.N {
		idx := pass - 1
		i := pass
		for i < A.N {
			if A.Daftar[idx].Year < A.Daftar[i].Year {
				idx = i
			} else if A.Daftar[idx].Year == A.Daftar[i].Year && A.Daftar[idx].Month < A.Daftar[i].Month {
				idx = i
			} else if A.Daftar[idx].Year == A.Daftar[i].Year && A.Daftar[idx].Month == A.Daftar[i].Month && A.Daftar[idx].Day < A.Daftar[i].Day {
				idx = i
			}
			i++
		}
		temp = A.Daftar[pass-1]
		A.Daftar[pass-1] = A.Daftar[idx]
		A.Daftar[idx] = temp
		pass++
	}
}

func sort_by_package_asc(A *tabRecord) {
	var temp Record
	pass := 1
	for pass < A.N {
		i := pass
		temp = A.Daftar[pass]
		for i > 0 && temp.Package.Name < A.Daftar[i-1].Package.Name {
			A.Daftar[i] = A.Daftar[i-1]
			i--
		}
		A.Daftar[i] = temp
		pass++
	}
}

func sort_by_package_desc(A *tabRecord) {
	var temp Record
	pass := 1
	for pass < A.N {
		i := pass
		temp = A.Daftar[pass]
		for i > 0 && temp.Package.Name > A.Daftar[i-1].Package.Name {
			A.Daftar[i] = A.Daftar[i-1]
			i--
		}
		A.Daftar[i] = temp
		pass++
	}
}

// // RECORD MANAGEMENT RECORD MANAGEMENT RECORD MANAGEMENT RECORD MANAGEMENT RECORD MANAGEMENT
// // RECORD MANAGEMENT RECORD MANAGEMENT RECORD MANAGEMENT RECORD MANAGEMENT RECORD MANAGEMENT
// // RECORD MANAGEMENT RECORD MANAGEMENT RECORD MANAGEMENT RECORD MANAGEMENT RECORD MANAGEMENT

func record_header2() {
	ClearScreen()
	fmt.Println("================================================================")
	fmt.Println("Main menu >> Record                        ")
	fmt.Println("================================================================")
	print_record_header2()
	fmt.Println("Record List                                ")
	fmt.Println("====================================================================================================================")
	print_record_header()
	fmt.Println("====================================================================================================================")
	fmt.Println("Total record registered :", records.N)
	fmt.Println("1. Add record                         ")
	fmt.Println("2. Edit record                        ")
	fmt.Println("3. Delete record                      ")
	fmt.Println("4. Show/Hide patients & package       ")
	fmt.Println("0. Return                            ")
	fmt.Print("Choice : ")
}

func record_header() {
	ClearScreen()
	fmt.Println("==================================================================================================================")
	fmt.Println("Main menu >> Record                        ")
	fmt.Println("==================================================================================================================")
	fmt.Println("Record List                                ")
	fmt.Println("==================================================================================================================")
	print_record_header()
	fmt.Println("==================================================================================================================")
	fmt.Println("Total record registered :", records.N)
	fmt.Println("1. Add record                         ")
	fmt.Println("2. Edit record                        ")
	fmt.Println("3. Delete record                      ")
	fmt.Println("4. Show/Hide patients & package list  ")
	fmt.Println("0. Return                            ")
	fmt.Print("Choice : ")
}

func print_record_header2() {
	var n int
	if packages.N > patients.N {
		n = packages.N
	} else {
		n = patients.N
	}
	fmt.Printf("%30v %2v %30v\n", "Patient list", "||", "Package list")
	fmt.Println("================================================================")
	for i := 0; i < n; i++ {
		fmt.Printf("%30v %2v %30v\n", patients.Daftar[i].Name, "||", packages.Daftar[i].Name)
	}
	fmt.Println("================================================================")

}

func print_record_header() {
	fmt.Printf("%10v %12v %17v %26v %45v\n", "ID", "Date", "Patient", "Package", "Result")
	for i := 0; i < records.N; i++ {
		fmt.Printf("%10v %4v-%2v-%4v %17v %26v %45v\n", records.Daftar[i].ID, records.Daftar[i].Day, records.Daftar[i].Month, records.Daftar[i].Year, records.Daftar[i].Patient.Name, records.Daftar[i].Package.Name, records.Daftar[i].Result)
	}
}

func record_main() {
	ClearScreen()
	record_header()
	var choice int
	var validChoice bool = false
	var show bool = false
	for !validChoice {
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			add_record()
			validChoice = true
		case 2:
			edit_record()
			validChoice = true
		case 3:
			delete_record()
			validChoice = true
		case 4:
			if show == true {
				record_header()
				show = false
			} else if show == false {
				record_header2()
				show = true
			}
		case 0:
			main_menu()
			validChoice = true
		default:
			fmt.Print("Invalid! Choice : ")
		}
	}
}

func add_record() {
	var record Record
	var finish bool = false
	var idxPatient, idxPackage int
	for !finish {
		fmt.Println("Type finish if you have finished inputting")
		fmt.Print("Patient name : ")
		record.Patient.Name = inputKalimat()

		switch record.Patient.Name {
		case "finish":
			finish = true
		default:
			idxPatient = search_patient(record.Patient.Name)
			switch idxPatient {
			case -1:
				fmt.Println("Patient not found!")
			default:
				fmt.Print("Package name : ")
				record.Package.Name = inputKalimat()
				idxPackage = search_package(record.Package.Name)
				switch idxPackage {
				case -1:
					fmt.Println("Package not found!")
				default:
					fmt.Print("Day (DD) : ")
					fmt.Scanln(&record.Day)
					fmt.Print("Month (MM) : ")
					fmt.Scanln(&record.Month)
					fmt.Print("Year (YYYY) : ")
					fmt.Scanln(&record.Year)
					fmt.Print("Result : ")
					record.Result = inputKalimat()

					records.Daftar[records.N] = record
					records.Daftar[records.N].ID = 56000 + records.N + 1
					records.Daftar[records.N].Patient = patients.Daftar[idxPatient]
					records.Daftar[records.N].Package = packages.Daftar[idxPackage]

					records.N++
				}
			}
		}

	}

	record_main()
}

func edit_record() {
	var record Record
	var input string
	var idxRecord int
	var idxPatient, idxPackage int
	fmt.Print("Patient name : ")
	record.Patient.Name = inputKalimat()
	fmt.Print("Package name : ")
	record.Package.Name = inputKalimat()
	idxRecord = search_record(record.Patient.Name, record.Package.Name)
	switch idxRecord {
	case -1:
		fmt.Print("Record not found! (enter anything to return)")
		fmt.Scanln(&input)
	default:
		record.ID = records.Daftar[idxRecord].ID
		fmt.Print("Input new patient name : ")
		record.Patient.Name = inputKalimat()
		idxPatient = search_patient(record.Patient.Name)
		switch idxPatient {
		case -1:
			fmt.Println("Patient not found!")
		default:
			fmt.Print("Input new package name : ")
			record.Package.Name = inputKalimat()
			idxPackage = search_package(record.Package.Name)
			switch idxPackage {
			case -1:
				fmt.Println("Package not found!")
			default:
				fmt.Print("Input new day (DD) : ")
				fmt.Scanln(&record.Day)
				fmt.Print("Input new Month (MM) : ")
				fmt.Scanln(&record.Month)
				fmt.Print("Input new year (YYYY) : ")
				fmt.Scanln(&record.Year)
				fmt.Print("input new Result : ")
				record.Result = inputKalimat()

				var accept string
				fmt.Print("Accept changes (Y/N) : ")
				fmt.Scanln(&accept)
				for accept != "Y" && accept != "N" {
					fmt.Println("Invalid")
					fmt.Print("Accept changes (Y/N) : ")
					fmt.Scanln(&accept)
				}
				if accept == "Y" {
					records.Daftar[idxRecord] = record
					records.Daftar[idxRecord].Patient = patients.Daftar[idxPatient]
					records.Daftar[idxRecord].Package = packages.Daftar[idxPackage]
				}
			}
		}
	}

	record_main()
}

func search_record(name_patient, name_package string) int {
	for i := 0; i < records.N; i++ {
		if records.Daftar[i].Patient.Name == name_patient && records.Daftar[i].Package.Name == name_package {
			return i
		}
	}
	return -1
}

func delete_record() {
	var record Record
	var input string
	var idxRecord int
	fmt.Print("Patient name : ")
	record.Patient.Name = inputKalimat()
	fmt.Print("Package name : ")
	record.Package.Name = inputKalimat()
	idxRecord = search_record(record.Patient.Name, record.Package.Name)
	switch idxRecord {
	case -1:
		fmt.Print("Record not found! (enter anything to return)")
		fmt.Scanln(&input)
	default:
		var accept string
		fmt.Print("Accept delete (Y/N) : ")
		fmt.Scanln(&accept)
		for accept != "Y" && accept != "N" {
			fmt.Println("Invalid")
			fmt.Print("Accept delete (Y/N) : ")
			fmt.Scanln(&accept)
		}
		if accept == "Y" {
			var i int
			for i = idxRecord; i < records.N; i++ {
				records.Daftar[i] = records.Daftar[i+1]
				records.Daftar[i].Patient = records.Daftar[i+1].Patient
				records.Daftar[i].Package = records.Daftar[i+1].Package
			}
			records.Daftar[i].ID = 0
			records.Daftar[i].Patient.ID = 0
			records.Daftar[i].Patient.Name = ""
			records.Daftar[i].Patient.Age = 0
			records.Daftar[i].Patient.Gender = ""
			records.Daftar[i].Package.ID = 0
			records.Daftar[i].Package.Name = ""
			records.Daftar[i].Package.Category = ""
			records.Daftar[i].Package.Price = 0
			records.Daftar[i].Year = 0
			records.Daftar[i].Month = 0
			records.Daftar[i].Day = 0
			records.Daftar[i].Result = ""
			records.N--
		}
	}

	record_main()
}

// REPORT REPORT REPORT REPORT REPORT REPORT REPORT REPORT REPORT
// REPORT REPORT REPORT REPORT REPORT REPORT REPORT REPORT REPORT
// REPORT REPORT REPORT REPORT REPORT REPORT REPORT REPORT REPORT

func report_header() {
	fmt.Println("===========================================")
	fmt.Println("Main menu >> Reports                       ")
	fmt.Println("===========================================")
	fmt.Printf("Total income : Rp.%v\n", total_income())
	fmt.Println("Total record :", records.N)
	fmt.Println("Total patient :", patients.N)
	fmt.Println("===========================================")
	fmt.Println("1. See income in a period               ")
	fmt.Println("0. Return                         ")
	fmt.Print("Choice : ")
}

func total_income() int {
	totalIncome := 0
	for i := 0; i < records.N; i++ {
		totalIncome += records.Daftar[i].Package.Price
	}
	return totalIncome
}

func report_main() {
	ClearScreen()
	report_header()
	var choice int
	var validChoice bool = false
	for !validChoice {
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			income_in_period()
			validChoice = true
		case 0:
			main_menu()
			validChoice = true
		default:
			fmt.Println("Invalid! Choice :")
		}
	}
}

func report_period_header(month, year int) {
	ClearScreen()
	var bulan string
	switch month {
	case 1:
		bulan = "January"
	case 2:
		bulan = "February"
	case 3:
		bulan = "March"
	case 4:
		bulan = "April"
	case 5:
		bulan = "May"
	case 6:
		bulan = "June"
	case 7:
		bulan = "July"
	case 8:
		bulan = "August"
	case 9:
		bulan = "September"
	case 10:
		bulan = "October"
	case 11:
		bulan = "November"
	case 12:
		bulan = "December"
	}
	totalIncome := 0
	totalRecord := 0
	total_income_in_period(&totalIncome, &totalRecord, month, year)
	fmt.Println("===========================================")
	fmt.Println("Main menu >> Reports >>", bulan, year)
	fmt.Println("===========================================")
	fmt.Printf("Total income : Rp.%v\n", totalIncome)
	fmt.Println("Total record :", totalRecord)
	fmt.Println("===========================================")
}

func total_income_in_period(tI, tR *int, month, year int) {
	for i := 0; i < records.N; i++ {
		if records.Daftar[i].Month == month && records.Daftar[i].Year == year {
			*tI += records.Daftar[i].Package.Price
			*tR++
		}
	}
}

func income_in_period() {
	var month, year int
	fmt.Print("Input month (MM) : ")
	fmt.Scanln(&month)
	fmt.Print("Input Year (YYYY) :")
	fmt.Scanln(&year)
	report_period_header(month, year)
	fmt.Println("(Input anything to return) Input : ")
	var input string
	fmt.Scanln(&input)
	report_main()
}

func main() {
	err := loadArray()
	if err != nil {
		fmt.Println("Error loading data:", err)
	}

	main_menu()

	err = saveArray()
	if err != nil {
		fmt.Println("Error saving data:", err)
	}
}

func saveArray() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := dir + "/data.json"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(patients)
	if err != nil {
		return err
	}
	err = encoder.Encode(packages)
	if err != nil {
		return err
	}
	err = encoder.Encode(records)
	if err != nil {
		return err
	}

	return nil
}

func loadArray() error {
	file, err := os.Open("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Data file doesn't exist. Starting with empty data.")
			return nil
		}
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&patients)
	if err != nil {
		return err
	}
	err = decoder.Decode(&packages)
	if err != nil {
		return err
	}
	err = decoder.Decode(&records)
	if err != nil {
		return err
	}

	fmt.Println("   Data loaded successfully.")
	return nil
}
