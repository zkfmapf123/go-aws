package main

import (
	"fmt"

	"github.com/aws-tools/src/agent"
	"github.com/aws-tools/src/utils"
)

func main() {

	configure, policy, instance := utils.GetAgentConfigure(), utils.GetPolicyConfigure(), utils.GetEC2Configure()
	cfg := agent.NewAWS(configure.Region)

	setIam(cfg, configure, policy)
	setEC2(cfg, instance)

}

// 1. set ec2, s3 policy
func setIam(cfg *agent.AWSconfig, configure utils.Configure, policy utils.Policy) {
	fmt.Println("Init IAM")
	cfg.SetIAM(configure.User, []string{policy.EC2, policy.S3})
}

// 2. init ec2
func setEC2(cfg *agent.AWSconfig, ec2 utils.Instance) {
	fmt.Println("Init EC2")
	cfg.SetEC2(ec2)
}
