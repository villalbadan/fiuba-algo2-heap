package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
)

const (
	_ELEM_TEST_COMUN = 15
)

var (
	TAMS_VOLUMEN    = []int{12500, 25000, 50000, 100000, 200000, 400000}
	ARREGLO_STRINGS = []string{"G", "K", "M", "B", "C", "W", "O", "A", "V", "F"}
	ARREGLO_INTS    = []int{4, 5, 6, 1, 2, 9, 7, 0, 8, 3}
)

//FUNC CMP -----------------------------------------------------------------------------------------------------------

func mayorEntreInts(clave1, clave2 int) int {
	return clave1 - clave2
}

func mayorEntreStrings(clave1, clave2 string) int {
	return strings.Compare(clave1, clave2)
}

//--------------------------------------------------------------------------------------------------------------------

func colaVacia[T comparable](heap TDAHeap.ColaPrioridad[T], t *testing.T) {
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, heap.EstaVacia(), true)
	require.EqualValues(t, heap.Cantidad(), 0)
}

func TestColaVacia(t *testing.T) {
	t.Log("Cola de prioridad recién creada se comporta como vacia")
	heap := TDAHeap.CrearHeap[int](mayorEntreInts)
	colaVacia(heap, t)
}

func TestColaDeUnElemento(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](mayorEntreInts)
	t.Log("Guardar y borrar un solo elemento en dicc recién creado")
	heap.Encolar(0)
	require.EqualValues(t, 0, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Desencolar())
	colaVacia(heap, t)
}

func TestEncolarVariosElementosOrdenados(t *testing.T) {
	t.Log("Al encolar varios elementos en orden ascendente, VerMax y Cantidad devuelven lo deseado")
	heap := TDAHeap.CrearHeap[int](mayorEntreInts)
	for i := 0; i <= _ELEM_TEST_COMUN; i++ {
		heap.Encolar(i)
		require.EqualValues(t, heap.VerMax(), i)
		require.EqualValues(t, heap.Cantidad(), i+1)
	}
}

func TestDesencolarAlgunosElementos(t *testing.T) {
	t.Log("Prueba de desencolar con algunos enteros arbitrarios")
	heap := TDAHeap.CrearHeap[int](mayorEntreInts)
	heap.Encolar(15)
	heap.Encolar(5)
	require.EqualValues(t, 15, heap.VerMax())
	require.EqualValues(t, 15, heap.Desencolar())
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 5, heap.Desencolar())
	heap.Encolar(7)
	heap.Encolar(100)
	heap.Encolar(0)
	require.EqualValues(t, 100, heap.VerMax())
	heap.Encolar(200)
	heap.Encolar(300)
	require.EqualValues(t, 300, heap.Desencolar())
	require.EqualValues(t, 200, heap.VerMax())
	require.EqualValues(t, 200, heap.Desencolar())
	require.EqualValues(t, 100, heap.VerMax())
	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 7, heap.VerMax())
	require.EqualValues(t, 7, heap.Desencolar())
	require.EqualValues(t, 0, heap.VerMax())
	require.EqualValues(t, 0, heap.Desencolar())
	colaVacia(heap, t)
}

func TestDesencolarHastaVacia(t *testing.T) {
	t.Log("Al desencolar hasta que está vacía hace que la cola se comporte como recién creada")
	heap := TDAHeap.CrearHeap[int](mayorEntreInts)
	for i := 0; i <= _ELEM_TEST_COMUN; i++ {
		heap.Encolar(i)
	}
	for i := 0; i <= _ELEM_TEST_COMUN; i++ {
		heap.Desencolar()
	}
	colaVacia(heap, t)

}

func TestColaDeStrings(t *testing.T) {
	t.Log("Prueba con primitivas con cadenas")
	nombres := TDAHeap.CrearHeap[string](mayorEntreStrings)
	nombres.Encolar("Daniela")
	require.EqualValues(t, "Daniela", nombres.VerMax())
	nombres.Encolar("Anibal")
	require.EqualValues(t, "Daniela", nombres.VerMax())
	nombres.Encolar("Matias")
	require.EqualValues(t, "Matias", nombres.VerMax())
	require.EqualValues(t, "Matias", nombres.Desencolar())
	require.EqualValues(t, "Daniela", nombres.VerMax())
	require.EqualValues(t, "Daniela", nombres.Desencolar())
	require.EqualValues(t, "Anibal", nombres.VerMax())
	require.EqualValues(t, "Anibal", nombres.Desencolar())
	colaVacia(nombres, t)

}

func TestHeapify(t *testing.T) {
	t.Log("Heapify funciona como un heap")
	heapArr := TDAHeap.CrearHeapArr(ARREGLO_INTS, mayorEntreInts)
	for i := 9; i >= 0; i-- {
		require.EqualValues(t, heapArr.Cantidad(), i+1)
		require.EqualValues(t, heapArr.VerMax(), i)
		require.EqualValues(t, heapArr.Desencolar(), i)
	}
	colaVacia(heapArr, t)
}

func TestHeapsort(t *testing.T) {
	t.Log("Heapsort devuelve un arreglo ordenado")
	arr := TDAHeap.HeapSort(ARREGLO_INTS, mayorEntreInts)
	for i := 0; i < 10; i++ {
		require.EqualValues(t, i, arr[i])
	}
}

// PRUEBAS DE VOLUMEN -----------------------------------------------------------------------------------------------
//Basadas en las pruebas de benchmark de hash de la cátedra con algunas modificaciones para que se ingresen elementos
//desordenados y se apliquen las primitivas correspondientes

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func listaNumerosRandoms(n int) []int {

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}

	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		swap(&nums[i], &nums[j])
	}

	return nums
}

func ejecutarPruebaVolumenHeap(b *testing.B, n int) {
	heap := TDAHeap.CrearHeap[int](mayorEntreInts)
	nums := listaNumerosRandoms(n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		heap.Encolar(nums[i])
		require.EqualValues(b, i+1, heap.Cantidad())
	}

	require.EqualValues(b, n, heap.Cantidad(), "La cantidad de elementos es incorrecta")

	for i := n - 1; i >= 0; i-- {
		/* Verifica que borre y devuelva los valores correctos */
		require.EqualValues(b, i, heap.VerMax())
		require.EqualValues(b, i, heap.Desencolar())
		require.EqualValues(b, i, heap.Cantidad())

	}

	require.EqualValues(b, 0, heap.Cantidad())
}

func BenchmarkHeap(b *testing.B) {
	b.Log("Prueba de stress del Heap. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. ")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeap(b, n)
			}
		})
	}
}

func ejecutarPruebaVolumenHeapify(b *testing.B, n int) {

	nums := listaNumerosRandoms(n)
	heap := TDAHeap.CrearHeapArr[int](nums, mayorEntreInts)
	require.EqualValues(b, n, heap.Cantidad(), "La cantidad de elementos es incorrecta")

	for i := n - 1; i >= 0; i-- {
		/* Verifica que borre y devuelva los valores correctos */
		require.EqualValues(b, i, heap.VerMax())
		require.EqualValues(b, i, heap.Desencolar())
		require.EqualValues(b, i, heap.Cantidad())

	}

	require.EqualValues(b, 0, heap.Cantidad())
}

func BenchmarkHeapify(b *testing.B) {
	b.Log("Prueba de stress del Heapify. Prueba transformando distintos arreglos (muy grandes) en heaps, " +
		"ejecutando muchas veces las pruebas para generar un benchmark. ")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeapify(b, n)
			}
		})
	}
}

func ejecutarPruebaVolumenHeapsort(b *testing.B, n int) {
	nums := listaNumerosRandoms(n)
	TDAHeap.HeapSort(nums, mayorEntreInts)
	for i := 0; i < 10; i++ {
		require.EqualValues(b, i, nums[i])
	}

}

func BenchmarkHeapsort(b *testing.B) {
	b.Log("Prueba de stress del Heapsort. Prueba ordenando arreglos de distintos tamaños (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. ")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeapsort(b, n)
			}
		})
	}
}
