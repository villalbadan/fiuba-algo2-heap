package cola_prioridad

const (
	_TAMANIO_INICIAL    = 10
	_FACTOR_REDIMENSION = 2
	_AGRANDAR           = 4
)

type colaPrioridad[T comparable] struct {
	datos    []T
	cantidad int
	cmp      funcCmp[T]
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(colaPrioridad[T])
	heap.cmp = funcion_cmp
	return heap
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {

}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {

}

// AUXILIARES ---------------------------------------------------------------------------------------------------------

func (heap *colaPrioridad[T]) redimensionar(nuevoTamanio int) {
	temp := heap.datos
	heap.datos = make([]T, nuevoTamanio)
	copy(heap.datos, temp)
}

// SwapIndices intercambia los valores de los indices indicados
func (heap colaPrioridad[T]) SwapIndices(i, j int) {
	heap.datos[i], heap.datos[j] = heap.datos[j], heap.datos[i]
}

//abs devuelve el valor absoluto de un integer
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Calculo de padres e hijos >>>>>>>>>>>>>>>>>>>>>>>>>>>

func padre(indice int) int {
	return abs((indice - 1) / 2)
}

func hijoIzq(indice int) int {
	return 2*indice + 1
}

func hijoDer(indice int) int {
	return 2*indice + 2
}

func (heap colaPrioridad[T]) condicionHeapPadre (hijo, padre T) bool {
	return heap.cmp(hijo, padre)
}

func (heap colaPrioridad[T]) upheap (elem T, pos int){
	if !heap.condicionHeapPadre(elem, heap.datos[padre(pos)]){
		heap.SwapIndices(pos, padre(pos))
		heap.upheap(elem, padre(pos))
	}
}

func (heap colaPrioridad[T]) downheap {

}

// PRIMITIVAS COLA DE PRIORIDAD -------------------------------------------------------------------------------------

// EstaVacia devuelve true si la cola se encuentra vacía, false en caso contrario.
func (heap colaPrioridad[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

// Encolar Agrega un elemento al heap.
func (heap *colaPrioridad[T]) Encolar(elem T) {
	if len(heap.datos) == heap.cantidad {
		heap.redimensionar(heap.cantidad * _FACTOR_REDIMENSION)
	}

	heap.datos[heap.cantidad] = elem
	heap.upheap(elem, heap.cantidad)
	heap.cantidad++

}

// VerMax devuelve el elemento con máxima prioridad. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (heap colaPrioridad[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

// Desencolar elimina el elemento con máxima prioridad, y lo devuelve. Si está vacía, entra en pánico con un
// mensaje "La cola esta vacia"
func (heap *colaPrioridad[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}

	heap.cantidad--
	heap.SwapIndices(0, heap.cantidad)
	heap.downheap(heap.datos[0], 0)

	if (len(heap.datos) >= heap.cantidad*_AGRANDAR) && (len(heap.datos)/_FACTOR_REDIMENSION >= _TAMANIO_INICIAL) {
		heap.redimensionar(len(heap.datos) / _FACTOR_REDIMENSION)
	}

	return heap.datos[heap.cantidad]
}

// Cantidad devuelve la cantidad de elementos que hay en la cola de prioridad.
func (heap colaPrioridad[T]) Cantidad() int {
	return heap.cantidad
}
