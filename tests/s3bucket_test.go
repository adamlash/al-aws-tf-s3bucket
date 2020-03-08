
package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3(t *testing.T) {
	t.Parallel()

	// Give Bucket our Name String
	expectedName := fmt.Sprintf("al-bucket-test-%s", strings.(ToLower(random.UniqueId())))
	// Choose a Random Region
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	// Begin Creating Options
	terraformOptions := &terraform.Options {
		// The path to where your Terraform code is located
		TerraformDir: "../",

		// Vars to Pass to the Module via -var. Made up from Above Strings
		Vars: map[string]interface{}{
			"name":	expectedName
		}

		// Environment Variables to run alongside Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion
		}
	  }
	  // At the end of the test, run `terraform destroy`
	  defer terraform.Destroy(t, terraformOptions)

	  // Run `terraform init` and `terraform apply`
	  terraform.InitAndApply(t, terraformOptions)

	  // Run `terraform output` to get the value of output variables
	  actualBucketName := terraform.Output(t, terraformOptions, "name")

	  // Look up Tags of the S3 Bucket

	  // Verify we're getting back the outputs we expect
	  assert.Equal(t, expectedName, actualBucketName)
	}
