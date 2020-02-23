variable "name" {
  description = "Name of Bucket"
  type        = string
  default     = ""
}

variable "kms_key_arn" {
  description = "(Optional) ARN for KMS Key"
  type        = string
  default     = ""
}

variable "logging_bucket" {
  description = "Name of Bucket to Log to"
  type        = string
  default     = ""
}

variable "tags" {
  description = "Array of Tags"
  type        = map(string)
  default     = {}
}