package indicators

import "testing"
import "reflect"

func TestSMA(t *testing.T) {

	var v []float64

	vTest := mfloat{1, 2, 3, 4, 5, 6, 7, 8, 9}
	vResult := []float64{2.5,3.5,4.5,5.5,6.5,7.5}

	v = vTest.SMA(4)

	if !reflect.DeepEqual(v, vResult) {
		t.Error("Expected [2.5,3.5,4.5,5.5,6.5,7.5], got", v)
	}
}

func TestEMA(t *testing.T) {

	var v []float64

	vTest := mfloat{1, 2, 3, 4, 5, 6, 7, 8, 9}
	vResult := []float64{1, 1.4, 2.04, 2.824, 3.6944, 4.61664, 
						5.569984, 6.5419903999999995, 
						7.525194239999999}

	v = vTest.EMA(4)

	if !reflect.DeepEqual(v, vResult) {
		t.Error("Expected [1, 1.4, 2.04, 2.824, 3.6944, 4.61664, 5.569984, 6.5419903999999995, 7.525194239999999], got", v)
	}
}

//func BollingerBands(slice mfloat, period int, nStd float64) ([]float64, []float64, []float64) {

//func MACD(data mfloat, ema ...int) ([]float64, []float64) {

//func OBV(priceData, volumeData mfloat) []float64 {

//func IchimokuCloud(priceData, lowData, highData mfloat, configs []int) ([]float64, []float64, []float64,[]float64, []float64) {
