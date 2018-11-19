package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type TheSearch struct {
	SearchBy  string `json:"What to search"`
	SearchFor string `json:"What to for"`
}

func Search(Q TheSearch) {
	svc := ec2.New(session.New(&aws.Config{Region: aws.String("us-east-2")}))
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(Q.SearchBy),
				Values: []*string{
					aws.String(Q.SearchFor),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println("instance id " + *instance.InstanceId)
			fmt.Println("current State " + *instance.State.Name)
		}
	}
}

func main() {
	lambda.Start(Search)
}
