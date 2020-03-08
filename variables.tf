variable "name" {
  description = "Name of Bucket"
  type        = string
  default     = ""
}
variable "tags" {
  description = "Array of Tags"
  type        = map(string)
  default     = {}
}