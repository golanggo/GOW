package DataStruct
import "fmt"

const MAXSIZE = 4

type Elem int

type List struct {
	Data   [MAXSIZE] Elem
	Lenght int
}

type ArrayList interface {
	Insert(index int, elem Elem) bool
	Get(index int) Elem
	Del(index int) bool
	Sort()
}

func (l *List) Insert(index int, elem Elem) bool {

	// 判断当前的容量是否已满
	if l.Lenght == MAXSIZE {
		fmt.Println("overflow MAXSIZE")
		return false
	}
	//检查插入的位置是否超出范围
	if index < 0 || index > l.Lenght {
		fmt.Println("index out of lenght")
		return false
	}

	//从最后的元素开始，向后移动一位，知道插入的位置
	for k := l.Lenght; k >= index; k-- {
		l.Data[k+1] = l.Data[k];
	}
	l.Data[index] = elem
	l.Lenght++

	return true
}

func (l *List) Del(index int) bool {
	if l.Lenght == 0 || index > l.Lenght-1 || index < 0 {
		fmt.Println("del fail")
		return false
	}

	for k := index; k <= l.Lenght-1; k++ {
		l.Data[k] = l.Data[k+1]
	}
	l.Lenght--
	return true
}

func (l *List) Get(index int) Elem {
	if l.Lenght == 0 {
		fmt.Println("not data")
		return 0
	}
	if index < 0 || index >= l.Lenght {
		fmt.Println("index over out")
		return 0
	}

	return l.Data[index];
}

func (l *List) Sort() {

	for i := 0;i<l.Lenght;i++{
		for j := i; j <= l.Lenght-1-i; j++ {
			if(l.Data[j] > l.Data[j+1]){
				tmp := l.Data[j]
				l.Data[j] =l.Data[j+1]
				l.Data[j+1]=tmp
			}
		}
	}
}
