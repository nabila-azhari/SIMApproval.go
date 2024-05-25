package main

import "fmt"

const NMAX int = 1000

type dataDriver struct {
	nama, alamat, id                             string
	tesKesehatan, tesTulis, tesPraktik, ratarata int
}

type tabDriver [NMAX]dataDriver

func main() {
	var arrDriver tabDriver
	var nCalonPengemudi int
	var pm1, pm1_2, pm2 int
	var indeks1, indeks2 int
	var nm bool

	fmt.Println("|-----------------------------|")
	fmt.Println("| WELCOME TO SIM APPROVAL APP |")
	fmt.Println("|-----------------------------|")
	fmt.Println()

	fmt.Print("How many drivers will you input?: ")
	fmt.Scan(&nCalonPengemudi)
	fmt.Println("Driver's data:")
	readDriverData(&arrDriver, &nCalonPengemudi)
	printDriverData(arrDriver, nCalonPengemudi)
	sortingbyId(&arrDriver, nCalonPengemudi)

	menu1()
	fmt.Print("Choose (1/2)?: ")
	fmt.Scan(&pm1)
	for pm1 != 1 {
		menu1_2()
		fmt.Print("Choose (1/2/3)?: ")
		fmt.Scan(&pm1_2)
		if pm1_2 == 1 {
			var idToDelete string
			fmt.Print("Enter ID to delete: ")
			fmt.Scan(&idToDelete)
			indeks1 = findIdBinarySearch(idToDelete, arrDriver, nCalonPengemudi)
			if indeks1 != -1 {
				deleteData(&arrDriver, &nCalonPengemudi, indeks1)
				printDriverData(arrDriver, nCalonPengemudi)
			} else {
				fmt.Println("ID not found.")
			}
		} else if pm1_2 == 2 {
			var idToEdit string
			fmt.Print("Enter ID to edit: ")
			fmt.Scan(&idToEdit)
			indeks2 = findIdSequential(idToEdit, arrDriver, nCalonPengemudi)
			if indeks2 != -1 {
				editData(&arrDriver, &nCalonPengemudi, indeks2)
				printDriverData(arrDriver, nCalonPengemudi)
			} else {
				fmt.Println("ID not found.")
			}
		} else if pm1_2 == 3 {
			var nameToFind string
			fmt.Print("Enter name: ")
			fmt.Scan(&nameToFind)
			nm = findNameSequential(nameToFind, arrDriver, nCalonPengemudi)
			if nm != true {
				fmt.Println("Name not found")
			} else {
				fmt.Println("Name found")
			}
		} else {
			menu1()
			fmt.Print("Choose (1/2)?: ")
			fmt.Scan(&pm1)
		}
	}

	calculateFinalScore(&arrDriver, nCalonPengemudi)

	for {
		menu2()
		fmt.Print("Choose (1/2/3/4)?: ")
		fmt.Scan(&pm2)
		if pm2 == 1 {
			cetakDataCalonFinal(arrDriver, nCalonPengemudi)
		} else if pm2 == 2 {
			sortingbyIdDescending(&arrDriver, nCalonPengemudi)
			cetakDataCalonFinal(arrDriver, nCalonPengemudi)
		} else if pm2 == 3 {
			fmt.Println("Sorted final data score driver's data:")
			sortingbyScoreDescending(&arrDriver, nCalonPengemudi)
			cetakDataCalonFinal(arrDriver, nCalonPengemudi)
		} else if pm2 == 4 {
			sortingbyScoreAscending(&arrDriver, nCalonPengemudi)
			cetakDataCalonFinal(arrDriver, nCalonPengemudi)
		} else if pm2 == 5 {
			return
		}
	}
}

func readDriverData(P *tabDriver, n *int) {
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Printf("%d) Enter name: ", i+1)
		fmt.Scan(&P[i].nama)
		fmt.Print("Enter address: ")
		fmt.Scan(&P[i].alamat)
		fmt.Print("Enter ID: ")
		fmt.Scan(&P[i].id)
		fmt.Print("Enter Health Test Score: ")
		fmt.Scan(&P[i].tesKesehatan)
		fmt.Print("Enter Written Test Score: ")
		fmt.Scan(&P[i].tesTulis)
		fmt.Print("Enter Practical Test Score: ")
		fmt.Scan(&P[i].tesPraktik)
		fmt.Println()
	}
}

func menu1() {
	fmt.Println("-----------------------")
	fmt.Println(" Is the data correct?  ")
	fmt.Println("-----------------------")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Println("-----------------------")
}

func menu1_2() {
	fmt.Println("-----------------------")
	fmt.Println("  Choose to fix data   ")
	fmt.Println("-----------------------")
	fmt.Println("1. Delete Data")
	fmt.Println("2. Modify Data")
	fmt.Println("3. Find Name")
	fmt.Println("4. Exit this page")
	fmt.Println("-----------------------")
}

func menu2() {
	fmt.Println("-----------------------")
	fmt.Println("        MENU        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Show driver's final data by Ascending ID")
	fmt.Println("2. Show driver's final data by Descending ID")
	fmt.Println("3. Show Driver's final data score from highest score to lowest score")
	fmt.Println("4. Show Driver's final data score from lowest score to highest score")
	fmt.Println("5. Exit the app")
	fmt.Println("-----------------------")
}

func deleteData(P *tabDriver, n *int, index int) {
	for i := index; i < *n-1; i++ {
		P[i] = P[i+1]
	}
	*n--
}

func editData(P *tabDriver, n *int, index int) {
	var pilihEdit int

	fmt.Println("Choose what to modify:")
	fmt.Println("1. Name")
	fmt.Println("2. Address")
	fmt.Println("3. Health Test Score")
	fmt.Println("4. Written Test Score")
	fmt.Println("5. Practical Test Score")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&pilihEdit)

	if pilihEdit == 1 {
		var newName string
		fmt.Print("Enter new name: ")
		fmt.Scan(&newName)
		P[index].nama = newName
	} else if pilihEdit == 2 {
		var newAddress string
		fmt.Print("Enter new address: ")
		fmt.Scan(&newAddress)
		P[index].alamat = newAddress
	} else if pilihEdit == 3 {
		var newHealthTest int
		fmt.Print("Enter new health test score: ")
		fmt.Scan(&newHealthTest)
		P[index].tesKesehatan = newHealthTest
	} else if pilihEdit == 4 {
		var newWrittenTest int
		fmt.Print("Enter new written test score: ")
		fmt.Scan(&newWrittenTest)
		P[index].tesTulis = newWrittenTest
	} else if pilihEdit == 5 {
		var newPracticalTest int
		fmt.Print("Enter new practical test score: ")
		fmt.Scan(&newPracticalTest)
		P[index].tesPraktik = newPracticalTest
	} else {
		fmt.Println("Invalid choice")
		return
	}
	fmt.Println("Modification successful.")
}

func calculateFinalScore(P *tabDriver, n int) {
	for i := 0; i < n; i++ {
		P[i].ratarata = (P[i].tesKesehatan + P[i].tesTulis + P[i].tesPraktik) / 3
	}
}

func printDriverData(P tabDriver, n int) {
	fmt.Println("Driver's data:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %s %d %d %d\n", P[i].nama, P[i].alamat, P[i].id, P[i].tesKesehatan, P[i].tesTulis, P[i].tesPraktik)
	}
	fmt.Println()
}

func cetakDataCalonFinal(P tabDriver, n int) {
	fmt.Println("Driver's final data:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %s %d %d %d", P[i].nama, P[i].alamat, P[i].id, P[i].tesKesehatan, P[i].tesTulis, P[i].tesPraktik)
		if P[i].tesKesehatan >= 60 && P[i].tesTulis >= 80 && P[i].tesPraktik >= 85 {
			fmt.Println(" ----> Passed with average score", P[i].ratarata)
		} else {
			fmt.Println(" ----> Didn't Pass with average score", P[i].ratarata)
		}
	}
	fmt.Println()
}

func minPosId(P tabDriver, n int) int {
	if n == 0 {
		return -1
	}

	minIndex := 0
	minId := P[0].id
	for i := 1; i < n; i++ {
		if P[i].id > minId {
			minIndex = i
			minId = P[i].id
		}
	}
	return minIndex

}

func swapId(P *tabDriver, i, j int) {
	var temp dataDriver
	temp = P[i]
	P[i] = P[j]
	P[j] = temp
}

func sortingbyId(P *tabDriver, n int) {
	for i := 0; i < n-1; i++ {
		minIndex := i

		
		for j := i + 1; j < n; j++ {
			if P[j].id < P[minIndex].id {
				minIndex = j
			}
		}
		swapId(P, i, minIndex)
	}
}
func findIdBinarySearch(DelId string, P tabDriver, n int) int {
	var left, right, mid, idx int

	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if P[mid].id > DelId {
			right = mid - 1
		} else if P[mid].id < DelId {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func findNameSequential(FindName string, P tabDriver, n int) bool {
	var ketemu bool
	var k int
	ketemu = false
	k = 0
	for ketemu == false && k < n {
		if P[k].nama == FindName {
			ketemu = true
		}
		k++
	}
	return ketemu
}

func findIdSequential(FindId string, P tabDriver, n int) int {
	var ketemu, k int

	ketemu = -1
	k = 0
	for ketemu == -1 && k < n {
		if P[k].id == FindId {
			ketemu = k
		}
		k++
	}
	return ketemu
}

func maksPosRataRata(P tabDriver, n int) int {
	if n == 0 {
		return -1
	}

	maxIndex := 0
	maxRataRata := P[0].ratarata
	for i := 1; i < n; i++ {
		if P[i].ratarata > maxRataRata {
			maxIndex = i
			maxRataRata = P[i].ratarata
		}
	}
	return maxIndex
}

func swapScore(P *tabDriver, i, j int) {
	var temp dataDriver
	temp = P[i]
	P[i] = P[j]
	P[j] = temp
}

func sortingbyScoreDescending(P *tabDriver, n int) {
	for i := 0; i < n-1; i++ {
		maxIndex := i
		for j := i + 1; j < n; j++ {
			if P[j].ratarata > P[maxIndex].ratarata {
				maxIndex = j
			}
		}
		swapScore(P, i, maxIndex)
	}
}

func sortingbyScoreAscending(P *tabDriver, n int) {
	for i := 1; i < n; i++ {
		temp := P[i]
		j := i - 1
		for j >= 0 && P[j].ratarata > temp.ratarata {
			P[j+1] = P[j]
			j = j - 1
		}
		P[j+1] = temp
	}
}

func sortingbyIdDescending(P *tabDriver, n int) {
	for i := 1; i < n; i++ {
		temp := P[i]
		j := i - 1
		for j >= 0 && P[j].id < temp.id {
			P[j+1] = P[j]
			j = j - 1
		}
		P[j+1] = temp
	}
}
