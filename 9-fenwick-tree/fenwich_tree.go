package fenwicktree

func LSB(num int) (lsb int) {
	return num & -num
}

type FenwickTree struct {
	arr []int
}

func (tree *FenwickTree) Query(start, end int) int {
	return tree.PrefixSum(end) - tree.PrefixSum(start-1)
}

func (tree *FenwickTree) PrefixSum(i int) int {
	sum := 0
	for i != 0 {
		sum += tree.arr[i-1]

		i -= LSB(i)
	}

	return sum
}

func (tree *FenwickTree) Add(i int, value int) {
	for i <= len(tree.arr) {
		tree.arr[i-1] += value

		i += LSB(i)
	}
}

func New(initialData []int) *FenwickTree {
	arr := make([]int, len(initialData))
	copy(arr, initialData)

	for i, value := range arr {
		immediateParent := i + LSB(i+1) + 1

		if immediateParent <= len(arr) {
			arr[immediateParent-1] += value
		}
	}

	return &FenwickTree{arr}
}
