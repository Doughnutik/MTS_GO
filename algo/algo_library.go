package main

import (
	"fmt"

	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
	"github.com/emirpasic/gods/sets/hashset"
	avl "github.com/emirpasic/gods/trees/avltree"
	rb "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
	dq "github.com/gammazero/deque"
)

type myStruct struct {
	a int
	b string
	c [2]int
}

func MyComparator(a, b interface{}) int {
	first := a.(myStruct)
	second := b.(myStruct)
	if first.a < second.a {
		return -1
	}
	if first.a > second.a {
		return 1
	}
	if first.b < second.b {
		return -1
	}
	if first.b > second.b {
		return 1
	}
	for i := 0; i < 2; i++ {
		if first.c[i] < second.c[i] {
			return -1
		}
		if first.c[i] > second.c[i] {
			return 1
		}
	}
	return 0
}

func TestingLinkedList() {
	// Двусвязный список. Также есть односвязный, но нет смысла его юзать
	list := dll.New()
	list.Add("a")
	list.Add(5)
	list.Insert(0, [2]int{1, 2})
	list.Insert(1, []int{1, 2, 3})
	list.Remove(1)
	it := list.Iterator()
	for it.Next() {
		index, value := it.Index(), it.Value()
		fmt.Println(index, value)
	}
}

func TestingHashSet() {
	hset := hashset.New()
	// Можно добавлять любые элементы, для которых определена операция сравнения (аналогично с map)
	// Также очень круто, что нет явной типизации для сета!
	// Но реализация чисто на map из GO. Аналогично с hashmap
	hset.Add(5)
	hset.Add("mama", "papa")
	hset.Remove("aboba") // Не кидает panic, если элемента нет в сете
	hset.Remove("mama")
	hset.Add(myStruct{4, "a", [2]int{5, 7}})
	values := hset.Values()
	fmt.Println(values)
	fmt.Println(hset.Contains(4), hset.Contains("mama"), hset.Contains("aboba"))
	fmt.Println(hset.Contains(myStruct{4, "a", [2]int{5, 7}}), hset.Contains(myStruct{3, "b", [2]int{5, 8}}))
}

func TestingTrees() {
	// Сразу объясню разницу между AVL и RB деревьями. Первое основывается на том, что глубины детей отличаются не больше, чем на 1.
	// У второго вершины красятся в 2 цвета и вводятся условия на связи вершин. Асимптотика деревьев O(log n), НО!
	// Если у нас преобладают операции поиска, то стоит использовать AVL дерево - у него меньше высота. Если же операции вставки, удаления, то RB - у AVL дерева много времени будут занимать операции перестроения.

	fmt.Println("RB tree")
	rbtree := rb.NewWithIntComparator()

	rbtree.Put(0, "mama")
	rbtree.Put(5, 100)
	rbtree.Put(3, myStruct{4, "a", [2]int{5, 7}})
	fmt.Println(rbtree.Values())
	fmt.Println(rbtree.Keys())
	fmt.Println(rbtree)
	rbtree.Remove(3)
	fmt.Println(rbtree)

	fmt.Println(rbtree.Get(0))

	// И главные операции - поиск первого >= и последнего <= нашего значения
	fmt.Println(rbtree.Floor(1))
	fmt.Println(rbtree.Ceiling(1))

	fmt.Println("AVL tree")

	avltree := avl.NewWithIntComparator()

	avltree.Put(0, "mama")
	avltree.Put(5, 100)
	avltree.Put(3, myStruct{4, "a", [2]int{5, 7}})
	fmt.Println(avltree.Values())
	fmt.Println(avltree.Keys())
	fmt.Println(avltree)
	avltree.Remove(3)
	fmt.Println(avltree)

	fmt.Println(avltree.Get(0))

	// И главные операции - поиск первого >= и последнего <= нашего значения
	fmt.Println(avltree.Floor(1))
	fmt.Println(avltree.Ceiling(1))
}

func TestingDeque() {
	var q dq.Deque[string]
	q.PushBack("foo")
	q.PushBack("bar")
	q.PushBack("baz")

	fmt.Println(q.Len())   // Prints: 3
	fmt.Println(q.Front()) // Prints: foo
	fmt.Println(q.Back())  // Prints: baz

	q.PopFront()             // remove "foo"
	fmt.Println(q.PopBack()) // remove "baz"

	q.PushFront("hello")
	q.PushBack("world")

	// Consume deque and print elements.
	for q.Len() != 0 {
		fmt.Println(q.PopFront()) // как видно, Pop возвращает элемент
	}
}

func main() {
	fmt.Println("### TestingLinkedList ###")
	TestingLinkedList()
	fmt.Println("### TestingHashSet ###")
	TestingHashSet()
	fmt.Println("### TestingTrees ###")
	TestingTrees()
	fmt.Println("### TestingDeque ###")
	TestingDeque()
}
