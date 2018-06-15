package weather

import (
	"fmt"
	_ "learngo/weather/backends"
	"learngo/weather/iface"
)

// Run app
func Run() {
	fmt.Println("Start weather app.")

	for _, be := range iface.AllBackends {
		be.Setup()
	}

	// selectBackend := "owm"
	// owm, ok := iface.AllBackends[selectBackend]
	// if !ok {
	// 	log.Fatalf("Could not find selected backend \"%s\"", selectBackend)
	// }
	// data := owm.Fetch("New York", 5)
	// fmt.Println(data)
}
