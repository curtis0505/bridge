package math

import (
	"fmt"
	"testing"
)

func TestCalcAvgAndMedian(t *testing.T) {

	list := []float64{1, 2, 3, 4, 5}
	avg, median := CalcAvgAndMedian(list) // 예전 코드로도 이런 케이스는 pass
	fmt.Printf("%v \n avg: %f, median: %f \n", list, avg, median)
	if avg != 3 {
		t.Errorf("CalcAvgAndMedian(list) = %f, want %f", avg, 3.0)
	}
	if median != 3 {
		t.Errorf("CalcAvgAndMedian(list) = %f, want %f", median, 3.0)
	}

	list = []float64{1, 2, 3, 4, 10} // 예전 코드로는 이런 케이스 fail
	avg, median = CalcAvgAndMedian(list)
	fmt.Printf("%v \n avg: %f, median: %f \n", list, avg, median)

	if avg != 4 {
		t.Errorf("CalcAvgAndMedian(list) = %f, want %f", avg, 4.0)
	}
	if median != 3 {
		t.Errorf("CalcAvgAndMedian(list) = %f, want %f", median, 3.0)
	}

	list = []float64{} // empty list
	avg, median = CalcAvgAndMedian(list)
	fmt.Printf("%v \n avg: %f, median: %f \n", list, avg, median)

}

//=== RUN   TestCalcAvgAndMedian
//[1 2 3 4 5]
//avg: 3.000000, median: 3.000000
//[1 2 3 4 10]
//avg: 4.000000, median: 3.000000
//--- PASS: TestCalcAvgAndMedian (0.00s)
//PASS
