package main
import "fmt"

func main(){
	const NMAX int =100
	var A[NMAX]int
	var n, i, jumlah, jumlahTidakLulus int
	var rata float64

	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
		jumlah+= A[i]
	}
	rata = float64(jumlah)/float64(i)

	for i = 0; i < n; i++{
		if A[i] < int(rata){
			jumlahTidakLulus++
		}
	}
	fmt.Print(rata, " ", jumlahTidakLulus)
}