package testunitario

import "testing"

func TestSuma(t *testing.T){
	total := Suma(5,2)

	if total != 7{
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 7)
	}

	//Ejecutar con go test
}

func TestSuma2(t *testing.T){
	tabla := []struct{
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{13, 24, 37},
	}

	for _, item := range tabla{
		total := Suma(item.a, item.b)

		if total != item.c{
			t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, item.c)
		}
	}
}

func TestGetMax(t *testing.T){
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1,5, 5},
		{6,12, 12},
		{24, 7, 24},
	}

	for _, item := range tabla{
		max := GetMax(item.a, item.b)

		if max != item.c{
			t.Errorf("GetMax incorrecta, tiene %d se esperaba %d", max, item.c)
		}
	}
}

//Este test demora mucho en ejecutarse
func TestFibo(t *testing.T){
	tabla := []struct {
		n int
		b int
	}{
		{1,1},
		{8,21},
		{50, 1258626925},
	}

	for _, item := range tabla{
		fib := Fibonacci(item.n)

		if fib != item.b{
			t.Errorf("Fibonacci incorrecta, tiene %d se esperaba %d", fib, item.b)
		}
	}
}