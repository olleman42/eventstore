package server

/*func main() {
	waiter := make(chan int)

	ess, err := eventstore.NewEventStore()
	if err != nil {
		log.Fatal(err)
	}
	defer ess.Close()

	go startgRPCServer(ess) // TODO: Let consumer spec host/port
	go startNatsComponent(ess)

	<-waiter
}
*/
