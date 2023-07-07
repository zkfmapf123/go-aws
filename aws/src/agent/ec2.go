package agent

import (
	"context"
	"log"

	"github.com/aws-tools/src/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EC2CreateInstanceAPI interface {
	RunInstances(ctx context.Context,
		params *ec2.RunInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error)

	CreateTags(ctx context.Context,
		params *ec2.CreateTagsInput,
		optFns ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error)
}

func (c *AWSconfig) SetEC2(ins utils.Instance) {
	client := ec2.NewFromConfig(c.cfg)
	minMaxCount := int32(1)

	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(ins.AMI),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     &minMaxCount,
		MaxCount:     &minMaxCount,
	}

	res, err := makeInstance(context.TODO(), client, input)
	if err != nil {
		log.Fatalln(err)
	}

	tag := &ec2.CreateTagsInput{
		Resources: []string{*res.Instances[0].InstanceId},
		Tags: []types.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("test-ec2"),
			},
		},
	}

	_, err = makeTags(context.TODO(), client, tag)
	if err != nil {
		log.Fatalln(err)
	}
}

func makeInstance(c context.Context, api EC2CreateInstanceAPI, input *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error) {
	return api.RunInstances(c, input)
}

func makeTags(c context.Context, api EC2CreateInstanceAPI, input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	return api.CreateTags(c, input)
}
