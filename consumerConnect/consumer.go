package consumer

import (
	"fmt"
	"kafka-golang/api"
	"kafka-golang/geohash"
	"kafka-golang/producer"
	"time"

	"github.com/Shopify/sarama"
)

var Channel chan int

type Precision struct {
	newPrecision int
}

func ConnectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// var myClient = &http.Client{Timeout: 3 * time.Second}

// func getJson(url string, target interface{}) error {
// 	r, err := myClient.Get(url)
// 	r.
// 	if err != nil {
// 		return err
// 	}
// 	defer r.Body.Close()

// 	return json.NewDecoder(r.Body).Decode(target)
// }

func GetPrecisionAsync(channel chan int) {
	fmt.Println("afasd")
	go api.Api.Post("/precision", api.AssignPrecision)
	time.Sleep(time.Second * 10)
	//g.Api.Get("/precision", g.GetPrecision)

	// apiUrl := "0.0.0.0:3000/api/v1/"
	// resource := "precision"
	// data := url.Values{}
	// data.Set("newPrecision", "5")

	// u, _ := url.ParseRequestURI(apiUrl)
	// u.Path = resource
	// u.RawQuery = data.Encode()
	// urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/?name=foo&surname=bar"

	// client := &http.Client{}
	// r, _ := http.NewRequest("POST", urlStr, nil)

	// resp, _ := client.Do(r)
	// fmt.Println(resp.Status)

	// req, _ := http.NewRequest("POST", "0.0.0.0:3000/api/v1/precision", 5)
	// req.Header.Add("Accept", "application/json")
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Errored when sending request to the server")
	// 	return
	// }

	// defer resp.Body.Close()
	// resp_body, _ := ioutil.ReadAll(resp.Body)
	// responseBody := string(resp_body)

	// precision, _ := strconv.Atoi(responseBody)
	// fmt.Println(resp.Status)
	// fmt.Println(string(resp_body))

	// fmt.Println("çalıştım")
	// precision := new(Precision)

	// var pleasewait sync.WaitGroup
	// pleasewait.Add(1)

	// go func() {
	// 	defer pleasewait.Done()
	// 	fmt.Println("içerdeyim")
	// 	go getJson("0.0.0.0:3000/api/v1/precision", precision) // doesn't work; no response :-(
	// }()
	// //process(w, cr) // works, but blank response :-(

	// pleasewait.Wait()
	// //go getJson("0.0.0.0:3000/api/v1/precision", precision)

	// //time.Sleep(time.Second * 5)

	//fmt.Println("precision değeri getJson", precision)
	channel <- api.Precision
	fmt.Println("getprecisionasync ", api.Precision)

}

func CreateGeohash(longitude, latitude float64, channel chan int) {
	precision := <-channel
	geohash := geohash.Encode(longitude, latitude, precision)
	fmt.Printf("Geohash: %s \n", geohash)
	producer.PushToQueue("geohashes", []byte(geohash))

}
