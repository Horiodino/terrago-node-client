package main

import (
	"strconv"

	"github.com/Horiodino/terrago-node-client/network"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "template.html")
	// })

	// http.HandleFunc("/hello", helloHandler)

	// fmt.Println("Server is running at http://localhost:8000")
	// http.ListenAndServe(":8000", nil)

	// 	go networks.AcepptRequest("Node 1", "localhost:8010")
	// 	go networks.AcepptRequest("Node 2", "localhost:8011")

	// 	// Send data to a specific node
	// 	networks.SendNicInfo("localhost:8011", "hiii i am node 1")
	// go network.AcepptRequest("Node 2", "localhost:8011")

	// 	// Send data to a specific node

	// Natchannel := make(chan int)

	// go network.Outbound_Traffic()

	// go network.SendNicInfo("localhost:8011", strconv.Itoa(network.Total_Incoming_Packets))

	// Call the second function in a loop to keep it running
	// go network.IncomingTraffic()
	network.IncomingTraffic()
	network.SendNicInfo("localhost:8011", strconv.Itoa(network.Total_Incoming_Packets))

}

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	status := r.URL.Query().Get("status")
// 	if status == "true" {
// 		cmd, err := exec.Command("kubectl").Output()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(string(cmd))
// 	}
// }
