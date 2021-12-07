package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/yigitaltunay/go-grpc-demo/api"
	"google.golang.org/grpc"
)

type server struct {
	api.UnimplementedWeatherServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterWeatherServiceServer(srv, &server{})
	fmt.Println("Server started on port 8080")
	
	panic(srv.Serve(lis))
}

func (s *server) ListCities(ctx context.Context,
	req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{CityCode: "2", CityName: "Denizli"},
			&api.CityEntry{CityCode: "1", CityName: "Istanbul"},
		},
	}, nil
}

func (s *server) QueryWeather(req *api.WeatherRequest, resp api.WeatherService_QueryWeatherServer) error {
	for {
		err := resp.Send(&api.WeatherResponse{
			Temperature: rand.Float32()*10 + 10})
		if err != nil {
			break
		}
		time.Sleep(time.Second * 10)
	}
	return nil
}
