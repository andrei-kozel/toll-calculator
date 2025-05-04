package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"

	"github.com/andrei-kozel/toll-calculator/types"
)

const (
	wsEndpoint = "ws://127.0.0.1:30000/ws"
)

var sendInterval = 3 * time.Second

func main() {
	obuIDS := generateOBUIDS(1)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := range obuIDS {
			lat, long := genLocation()
			data := types.OBUData{
				OBUID: obuIDS[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()

	return n + f
}

func genLocation() (float64, float64) {
	return genCoord(), genCoord()
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := range n {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
