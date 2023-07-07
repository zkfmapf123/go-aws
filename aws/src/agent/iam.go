package agent

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func (c *AWSconfig) SetIAM(username string, arns []string) {

	for _, arn := range arns {

		client := iam.NewFromConfig(c.cfg)
		_, err := client.AttachUserPolicy(context.TODO(), &iam.AttachUserPolicyInput{
			UserName:  &username,
			PolicyArn: &arn,
		})

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%v is attach success\n", arn)
	}
}
