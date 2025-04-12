package heap

type MinHeap struct {
	arr []int
}

func NewHeap() *MinHeap {
	return &MinHeap{
		arr: make([]int, 0, 5),
	}
}

func (h MinHeap) GetArr() []int {
	return h.arr[:len(h.arr)]
}

func (h *MinHeap) Insert(val int) {
	i := uint(len(h.arr))
	h.arr = append(h.arr, val)
	parentI := (i - 1) / 2
	for i != 0 && h.arr[i] < h.arr[parentI] {
		h.arr[i], h.arr[parentI] = h.arr[parentI], h.arr[i]
		i = parentI
		parentI = (i - 1) / 2
	}
}

func (h *MinHeap) DeleteI(i uint) {
	h.arr[i] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	ci := h.getSmallestChild(i)
	for ci != -1 && h.arr[i] > h.arr[ci] {
		h.arr[i], h.arr[ci] = h.arr[ci], h.arr[i]
		i = uint(ci)
		ci = h.getSmallestChild(i)
	}
}

func (h MinHeap) getSmallestChild(i uint) int {
	lci := (2 * i) + 1
	rci := (2 * i) + 2
	if lci >= uint(len(h.arr)) {
		return -1
	}
	if rci >= uint(len(h.arr)) {
		return int(lci)
	}
	var ci uint
	if h.arr[lci] <= h.arr[rci] {
		ci = lci
	} else {
		ci = rci
	}

	return int(ci)
}
