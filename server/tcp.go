package server

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/olleman42/dinnerbot/eventstore"
)

const (
	getAllHistoryCommand       string = "GetFullHistory"
	getTypeHistoryCommand      string = "GetTypeHistory"
	getAggregateHistoryCommand string = "GetAggregateHistory"
)

// StartQueryServer - WIP function that exposts query API over pure TCP
// Do not use as paramaters can not be passed properly currentl
func StartQueryServer(store *eventstore.EventStore) {
	fmt.Println("starting query server")
	listener, err := net.Listen("tcp", ":3010")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			//fmt.Println("got new connection")
			scanner := bufio.NewScanner(c)
			defer c.Close()
			//defer fmt.Println("closed connection")

			for scanner.Scan() {
				switch scanner.Text() {
				case getAllHistoryCommand:
					fmt.Println("Getting all history")
					store.GetHistory(c)
				case getTypeHistoryCommand:
					fmt.Println("Getting type history")
					scanner.Scan()
					err := store.GetTypeHistory(scanner.Text(), c)
					if err != nil {
						log.Println(err)
					}
				case getAggregateHistoryCommand:
					fmt.Println("Getting aggregate history")
					scanner.Scan()
					aggType := scanner.Text()
					scanner.Scan()
					aggID := scanner.Text()
					store.GetAggregateHistory(aggType, aggID, c)
				default:
					fmt.Println("Command not recognized")
				}
				break
			}

		}(conn)
	}
}
