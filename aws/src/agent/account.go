package agent

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AWSconfig struct {
	cfg    aws.Config
	region string
}

func NewAWS(region string) *AWSconfig {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	return &AWSconfig{
		cfg:    cfg,
		region: region,
	}
}
