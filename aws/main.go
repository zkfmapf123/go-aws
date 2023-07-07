package main

import (
	"github.com/aws-tools/src/agent"
	"github.com/aws-tools/src/utils"
)

func main() {

	configure, policy := utils.GetAgentConfigure(), utils.GetPolicyConfigure()
	cfg := agent.NewAWS(configure.Region)

	setIam(cfg, configure, policy)

}

// 1. set ec2, s3 policy
func setIam(cfg *agent.AWSconfig, configure utils.Configure, policy utils.Policy) {
	cfg.SetIAM(configure.User, []string{policy.EC2, policy.S3})
}
