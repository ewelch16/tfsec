package mq

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/aws/mq"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "aws_mq_broker" "bad_example" {
   broker_name = "example"
 
   configuration {
     id       = aws_mq_configuration.test.id
     revision = aws_mq_configuration.test.latest_revision
   }
 
   engine_type        = "ActiveMQ"
   engine_version     = "5.15.0"
   host_instance_type = "mq.t2.micro"
   security_groups    = [aws_security_group.test.id]
 
   user {
     username = "ExampleUser"
     password = "MindTheGap"
   }
   publicly_accessible = true
 }
 `},
		GoodExample: []string{`
 resource "aws_mq_broker" "good_example" {
   broker_name = "example"
 
   configuration {
     id       = aws_mq_configuration.test.id
     revision = aws_mq_configuration.test.latest_revision
   }
 
   engine_type        = "ActiveMQ"
   engine_version     = "5.15.0"
   host_instance_type = "mq.t2.micro"
   security_groups    = [aws_security_group.test.id]
 
   user {
     username = "ExampleUser"
     password = "MindTheGap"
   }
   publicly_accessible = false
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/mq_broker#publicly_accessible",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"aws_mq_broker",
		},
		Base: mq.CheckNoPublicAccess,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if publiclyAccessibleAttr := resourceBlock.GetAttribute("publicly_accessible"); publiclyAccessibleAttr.IsNil() { // alert on use of default value
				results.Add("Resource uses default value for publicly_accessible", resourceBlock)
			} else if publiclyAccessibleAttr.IsTrue() {
				results.Add("Resource does not have publicly_accessible set to false", publiclyAccessibleAttr)
			}
			return results
		},
	})
}
