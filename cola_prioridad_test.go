package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const (
	_ELEM_TEST_VOLUMEN = 1000000
	_ELEM_TEST_COMUN   = 15
)

var ARREGLO_INTS = []int{4, 5, 6, 1, 2, 9, 7, 0, 8, 3}

//FUNC CMP -----------------------------------------------------------------------------------------------------------

func mayorEntreInts(clave1, clave2 int) int {
	return clave1 - clave2
}

func mayorEntreStrings(clave1, clave2 string) int {
	return strings.Compare(clave1, clave2)
}

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
	t.Log("Prueba con cadenas")
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
	t.Log("")
	heapArr := TDAHeap.CrearHeapArr(ARREGLO_INTS, mayorEntreInts)
	for i := 9; i >= 0; i-- {
		require.EqualValues(t, heapArr.Cantidad(), i+1)
		require.EqualValues(t, heapArr.VerMax(), i)
		require.EqualValues(t, heapArr.Desencolar(), i)
	}
}
