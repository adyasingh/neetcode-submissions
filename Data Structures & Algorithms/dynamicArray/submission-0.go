type DynamicArray struct {
    arr []int
    length int
    capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
    arr := make([]int, capacity)

    return &DynamicArray{
        arr: arr,
        length:0,
        capacity: capacity,
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.arr[i]=n
}

func (da *DynamicArray) Pushback(n int) {
   if da.length==da.capacity{
    da.resize()
   }
   da.arr[da.length] = n
   da.length++
}

func (da *DynamicArray) Popback() int {
    last := da.arr[da.length-1]
    da.length--
    return last
}

func (da *DynamicArray) resize() {
    size:=da.capacity*2
    newArr := make([]int, size)
    for i:=0;i<da.length;i++{
        newArr[i] = da.arr[i]
    }
    da.arr = newArr
    da.capacity = size
}

func (da *DynamicArray) GetSize() int {
    return da.length
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
