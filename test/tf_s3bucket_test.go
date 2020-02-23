
package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3(t *testing.T) {
	terraformOptions := &terraform.Options {
		// The path to where your Terraform code is located
		TerraformDir: "../",
	  }
	  // At the end of the test, run `terraform destroy`
	  defer terraform.Destroy(t, terraformOptions)

	  // Run `terraform init` and `terraform apply`
	  terraform.InitAndApply(t, terraformOptions)

	  // Run `terraform output` to get the value of output variables
	  actualBucketName := terraform.Output(t, terraformOptions, "bucket_name")

	  // Verify we're getting back the outputs we expect
	  assert.Equal(t, expectedText, actualBucketName)
	}
