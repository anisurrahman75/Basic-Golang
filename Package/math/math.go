package math

func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
func Min(slc []int) int {
	ans := slc[0]
	for _, i := range slc {
		if i < ans {
			ans = i
		}
	}
	return ans
}
func Max(slc []int) int {
	ans := slc[0]
	for _, i := range slc {
		if i > ans {
			ans = i
		}
	}
	return ans
}
