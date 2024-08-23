variable "vpc_tags" {
  type        = map(string)
  description = "Tags for the VPC"
  default = {
    Name        = "0-basics-vpc"
    Environment = "Temporal"
  }
}
