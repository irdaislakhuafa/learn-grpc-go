package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/protobuf/generated/pb"
	"google.golang.org/protobuf/proto"
)

func main() {

	// as example response of grpc, imagine this is response from other grpc server
	response := pb.UserResponsePagination{
		Pagination: &pb.Pagination{
			TotalData:   2,
			CurrentPage: 1,
			Limit:       1,
			TotalPages:  2,
		},
		Data: []*pb.User{
			{
				Id:      uuid.NewString(),
				Name:    "irda islakhu afa",
				Age:     21,
				Hobbies: []string{"coding", "eat", "sleep"},
				Address: []*pb.Address{
					{
						Id:          uuid.NewString(),
						Country:     "indonesia",
						Province:    "east java",
						Regency:     "tuban",
						SubDistrict: "montong",
					},
				},
			},
			{
				Id:      uuid.NewString(),
				Name:    "someone",
				Age:     23,
				Hobbies: []string{"read", "fishing"},
				Address: []*pb.Address{
					{
						Id:          uuid.NewString(),
						Country:     "indonesia",
						Province:    "east java",
						Regency:     "tuban",
						SubDistrict: "montong",
					},
				},
			},
		},
	}

	// data will be marshaled to bytes before send
	dataByte, err := proto.Marshal(&response)
	if err != nil {
		log.Fatalf("failed to marshal, %v\n", err)
	}

	// and bytes data used fewer resource
	fmt.Printf("dataByte: %+v\n", dataByte)

	// imagine this is response byte of from grpc server above, will be unmarshaled in client to make it data processable
	receiveResponse := pb.UserResponsePagination{}
	if err := proto.Unmarshal(dataByte, &receiveResponse); err != nil {
		log.Fatalf("failed to unmarshal, %v\n", err)
	}

	// as example, here i loop data after receive response grpc server above
	for _, v := range receiveResponse.GetData() {
		fmt.Printf("id: '%v' name:'%v' age: '%v'\n", v.GetId(), v.GetName(), v.GetAge())
	}
}
