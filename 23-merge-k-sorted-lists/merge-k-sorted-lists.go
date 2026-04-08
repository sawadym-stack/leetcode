/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ----------- Min Heap -----------

type ListMinHeap struct {
	arr []*ListNode
}

func (h *ListMinHeap) Insert(node *ListNode) {
	h.arr = append(h.arr, node)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *ListMinHeap) ExtractMin() *ListNode {
	if len(h.arr) == 0 {
		return nil
	}

	min := h.arr[0]
	last := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	if len(h.arr) > 0 {
		h.arr[0] = last
		h.heapifyDown(0)
	}

	return min
}

func (h *ListMinHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.arr[parent].Val <= h.arr[i].Val {
			break
		}
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		i = parent
	}
}

func (h *ListMinHeap) heapifyDown(i int) {
	n := len(h.arr)
	for {
		smallest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && h.arr[left].Val < h.arr[smallest].Val {
			smallest = left
		}
		if right < n && h.arr[right].Val < h.arr[smallest].Val {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		i = smallest
	}
}

// ----------- Merge K Lists -----------

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListMinHeap{}

	for _, l := range lists {
		if l != nil {
			h.Insert(l)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for len(h.arr) > 0 {
		node := h.ExtractMin()
		curr.Next = node
		curr = curr.Next

		if node.Next != nil {
			h.Insert(node.Next)
		}
	}

	return dummy.Next
}