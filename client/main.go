package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/yigitaltunay/go-grpc-demo/api"
	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewWeatherServiceClient(conn)

	ctx := context.Background()

	go func() {
		for {
			resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})
			if err != nil {
				panic(err)
			}
			for _, city := range resp.Items {
				fmt.Println(city)
			}
			time.Sleep(time.Second * 10)
		}

	}()

	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{
		CityCode: "asdasd",
	})
	if err != nil {
		fmt.Println(err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
		fmt.Println(msg.GetTemperature())
	}

	select {}

}
