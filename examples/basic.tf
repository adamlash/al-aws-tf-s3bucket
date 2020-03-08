provider "aws" {
  version = "~> 2.0"
  region  = "us-east-1"
}

module "aws_s3_bucket_1" {
  source = "../"
  name = "al-testbucket-1234"
  tags = {
    tag1 = "tag1"
  }
}

