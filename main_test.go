	package main

	import (
		"testing"
	)

func TestQueue_Empty(t *testing.T) {
	queue := Queue{}

	_, err := queue.Dequeue()
	if err == nil {
		t.Error("no error when dequeue")
	}

	_, err = queue.First()
	if err == nil {
		t.Error("no error to first element")
	}

	_, err = queue.Last()
	if err == nil {
		t.Error("no error to last element")
	}

	size := queue.Len()
	if size != 0 {
		t.Error("size is not 0, size got: ", size)
	}
}

func TestQueue_OneElement(t *testing.T) {
	queue := Queue{}
	queue.Enqueue("1")

	value, err := queue.First()
	if err != nil {
		t.Error("shouldn't be any error, got: ", err)
	}

	if value != "1" {
		t.Error("should be 1, got: ", value)
	}

	value, err = queue.Last()
	if err != nil {
		t.Error("shouldn't be any error, got: ")
	}

	size := queue.Len()
	if size != 1 {
		t.Error("size is not 1, size got: ", size)
	}

	value, err = queue.Dequeue()
	if err != nil {
		t.Error("error when dequeue")
	}

	if value != "1" {
		t.Error("expects 1, got: ", value)
	}

	size = queue.Len()
	if size != 0 {
		t.Error("size is not 0, size got: ", size)
	}
}


func TestQueue_ManyElements(t *testing.T) {
	queue := Queue{}
	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	_, err := queue.First()
	if err != nil {
		t.Error("shouldn't be any error, got: ", err)
	}
	_, err = queue.Last()
	if err != nil {
		t.Error("shouldn't be any error, got: ", err)
	}

	size := queue.Len()
	if size != 3 {
		t.Error("size is not 3, size got: ", size)
	}

	value, err := queue.Dequeue()
	if err != nil {
		t.Error("Error is not null, got: ", err)
	}
	if value != "1"{
		t.Error("size is not 1, got: ", value)
	}

	value, err = queue.Dequeue()
	if err != nil {
		t.Error("Error is not null, got: ", err)

	}
	if value != "2"{
		t.Error("size is not 2, got: ", value)
	}



	value, err = queue.Dequeue()
	if err != nil {
		t.Error("Error is not null, got: ", err)
	}
	if value != "3"{
		t.Error("size is not 3, got:", value)
	}
}