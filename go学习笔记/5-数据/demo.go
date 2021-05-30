package main

func main(){
	exit := make(chan struct{})

	go func() {
		println("hello, world!")
		exit <- struct{}{}
	}()

	<-exit
	println("end.")
}
