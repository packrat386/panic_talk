package main

func ohNo(done chan int) {
	panic("oh no")
	done <- 0
}

func recovery() {
	done := make(chan int)
	defer func() {
		if r := recover(); r != nil {
			println("recovered: ", r)
		}
	}()

	go ohNo(done)
	<-done
}

func main() {
	println("here we go")
	recovery()
}

