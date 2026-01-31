package ext

func ToDP(f float64, dp int) float64 {
	mult := float64(1)
	for i := 0; i < dp; i++ {
		mult *= 10
	}
	return float64(int(f*mult)) / mult
}
