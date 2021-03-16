package main


func FindNb(m int) int {
	n := parseInt(Math.sqrt(2*(Math.sqrt(m)))) - 1
	for ((n*(n+1))/2)**2 < m {
		n++
	}
	if ((n*(n+1))/2)**2 > m {
		-1
	} else {
		n
	}
}

func main() {
	FindNb(5)
}