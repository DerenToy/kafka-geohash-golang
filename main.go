package main

import (
	"encoding/json"
	"fmt"
	cc "kafka-golang/consumerConnect"
	"kafka-golang/producer"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

var geolocation producer.Geolocation

func main() {
	ch := make(chan int)
	topic := "geolocations"
	worker, err := cc.ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	// Calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))

				if topic == "geolocations" {
					json.Unmarshal(msg.Value, &geolocation)
					long := geolocation.Longitude
					lat := geolocation.Latitude
					fmt.Printf("Latitude: %f , Longitude: %f  \n", lat, long)

					go cc.GetPrecisionAsync(ch)
					go cc.CreateGeohash(lat, long, ch)
					time.Sleep(time.Second * 5)

					// fmt.Println("neler oluyor")
					// var pleasewait sync.WaitGroup
					// pleasewait.Add(1)

					// go func() {
					// 	defer pleasewait.Done()
					// 	fmt.Println("içerdeyim")
					// 	cc.CreateGeohash(lat, long, cc.Channel) // doesn't work; no response :-(
					// }()
					// process(w, cr) // works, but blank response :-(

					// pleasewait.Wait()
					// go cc.CreateGeohash(lat, long, ch)
					// app.Listen(":3001")
					// time.Sleep(time.Second * 2)

				}

			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}

}

// func handler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println(req)
// 	precision, _ := strconv.Atoi(req.FormValue("precision"))
// 	fmt.Println("precision handler: ", precision)
// 	cc.Channel <- precision
// 	fmt.Println("burdaa noluyo ya")
// }
