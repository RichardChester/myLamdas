package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Person struct {
	ID string
}

type Confermation struct {
	Message string
}

func StartInsti(Q Person) (Confermation, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(Q.ID),
		},
		DryRun: aws.Bool(true),
	}
	result, error := svc.StopInstances(input)
	awsErr, ok := error.(awserr.Error)
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, error = svc.StopInstances(input)
		if error != nil {
			fmt.Println("Error", error)
		} else {
			fmt.Println("Success", result.StoppingInstances)
		}
	} else {
		fmt.Println("Error", error)
	}
	return Confermation{Message: fmt.Sprintf("instance %s has now started", Q.ID)}, error
}

func main() {
	lambda.Start(StartInsti)
}
