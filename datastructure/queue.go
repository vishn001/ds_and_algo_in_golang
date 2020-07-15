package main

type Item interface{}

type Queue struct {
	items []Item
	mutex  sync.Mutex
}

func (queue *Queue) Enqueue(item Item) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = append(queue.items, item)
}

func (queue *Queue) Dequeue() Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if len(queue.items) == 0 {
		return nil
	}

	lastItem := queue.items[0]
	queue.items = queue.items[1:]

	return lastItem
}

func main() {
	var queue Queue

	queue.Enqueue(5)
	queue.Enqueue(4)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)

	fmt.Println("Queue:", queue.Dump())
	fmt.Println("The last item:", queue.Peek())

	queue.Dequeue()

	fmt.Println("Queue:", queue.Dump())
	fmt.Println("Queue is empty:", queue.IsEmpty())

	queue.Reset()

	fmt.Println("Queue is empty:", queue.IsEmpty())
}

