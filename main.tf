resource "aws_s3_bucket" "default" {
  bucket = var.name
  acl    = "private"

  logging {
    target_bucket = var.logging_bucket
    target_prefix = "${var.name}/"
  }
  
  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = "${aws_kms_key.mykey.arn}"
        sse_algorithm     = "aws:kms"
      }  
  tags = var.tags
}
