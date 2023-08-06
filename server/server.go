package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "grpc-test/grpctest"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedPersonServiceServer
}

type person struct {
	Id   int64
	Name string
}

func loadPeople(filename string) (*[]person, error) {
	// JSONファイルを読み込む
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ファイルをオープンできませんでした: %v", err)
	}
	defer file.Close()

	// ファイルの中身をバイト列として読み込む
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ファイルを読み込めませんでした: %v", err)
	}

	// JSONデコードして構造体に入れる
	var people []person
	err = json.Unmarshal(byteValue, &people)
	if err != nil {
		return nil, fmt.Errorf("JSONデコードエラー: %v", err)
	}

	return &people, nil
}

func find(people *[]person, id int64) *person {
	for _, person := range *people {
		if person.Id == id {
			return &person
		}
	}

	return nil
}

func (s *server) GetFeature(ctx context.Context, in *pb.Person) (*pb.Feature, error) {
	people, err := loadPeople("../data.json")
	if err != nil {
		return nil, err
	}

	person := find(people, int64(in.Id))
	if person == nil {
		return nil, fmt.Errorf("person not found")
	}

	return &pb.Feature{Name: person.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{})

	// grep_cli 用
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
