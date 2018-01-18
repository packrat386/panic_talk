package main

func ohNo() {
	panic("oh no")
}

func recovery() {
	defer func() {
		if r := recover(); r != nil {
			println("recovered: ", r)
		}
	}()

	ohNo()
}

func main() {
	println("here we go")
	recovery()
}

