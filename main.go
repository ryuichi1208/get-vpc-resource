package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	// AWS Configのロード
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	// EC2クライアントの作成
	ec2Client := ec2.NewFromConfig(cfg)

	// VPC IDの指定
	vpcID := "vpc-0123456789abcdef0"

	// VPCに紐づくリソースを取得する
	listInstances(ec2Client, vpcID)
}

func listInstances(client *ec2.Client, vpcID string) {
	// インスタンスのリストを取得
	input := &ec2.DescribeInstancesInput{
		Filters: []ec2types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{vpcID},
			},
		},
	}

	result, err := client.DescribeInstances(context.TODO(), input)
	if err != nil {
		log.Fatalf("Failed to get instances: %v", err)
	}

	fmt.Println("Instances:")
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("  ID: %s, State: %s\n", aws.ToString(instance.InstanceId), instance.State.Name)
		}
	}
}
