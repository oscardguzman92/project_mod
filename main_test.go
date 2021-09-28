package main

import "testing"

func TestSum(t *testing.T)  {
	/*total := Sum(5, 5)

	if total != 11 {
	t.Errorf("Sum was incorrect, got %d expected %d", total, 11)
	}*/
		tables := []struct {
			a int 
			b int 
			n int 
		} {
			{1,2,3},
			{2,2,4},
			{25,26,52},
		}

		for _, item := range tables { //_ si no se sabe el tamaño del arreglo, para no declarar variables que no se van a usar 
			total := Sum(item.a, item.b)
			if total != item.n {
				t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
			}
		}
}