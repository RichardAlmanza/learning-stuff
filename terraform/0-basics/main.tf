# AWS VPC Definition
resource "aws_vpc" "simple_vpc" {
  cidr_block = "10.0.0.0/16"
  tags       = var.vpc_tags
}
