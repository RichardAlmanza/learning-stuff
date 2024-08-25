terraform {
  required_version = "~> 1.9"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.56"
    }
  }
}

provider "aws" {
  # Uses credentials and configuration in ~/.aws/
  profile = "default"
}

data "aws_availability_zones" "aws_region_azs" {
  state = "available"
}

locals {
  number_az            = 2
  number_subnets_by_az = 2
  newbits              = log(local.number_az * local.number_subnets_by_az, 2)

  base_tags = {
    Project     = "${var.project_name}"
    Environment = "Temporal"
  }
}

resource "aws_vpc" "vpc" {
  cidr_block       = var.cidr_block
  instance_tenancy = "default"

  tags = merge({
    Name = "${var.project_name}-vpc"
  }, local.base_tags)
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.vpc.id

  tags = merge({
    Name = "${var.project_name}-igw"
  }, local.base_tags)
}

resource "aws_route_table" "public_table" {
  vpc_id = aws_vpc.vpc.id

  route {
    cidr_block = var.cidr_block
    gateway_id = "local"
  }

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = merge({
    Name = "${var.project_name}-public-rtb"
  }, local.base_tags)
}

resource "aws_route_table" "private_table" {
  vpc_id = aws_vpc.vpc.id

  route {
    cidr_block = var.cidr_block
    gateway_id = "local"
  }

  tags = merge({
    Name = "${var.project_name}-private-rtb"
  }, local.base_tags)
}

resource "aws_subnet" "public_subnets" {
  count  = local.number_az
  vpc_id = aws_vpc.vpc.id

  availability_zone = data.aws_availability_zones.aws_region_azs.names[count.index]
  cidr_block        = cidrsubnet(var.cidr_block, local.newbits, 2 * count.index)

  tags = merge({
    Name = "${var.project_name}-subnet-public-${count.index}"
    Type = "public"
  }, local.base_tags)
}

resource "aws_subnet" "private_subnets" {
  count  = local.number_az
  vpc_id = aws_vpc.vpc.id

  availability_zone = data.aws_availability_zones.aws_region_azs.names[count.index]
  cidr_block        = cidrsubnet(var.cidr_block, local.newbits, 2 * count.index + 1)

  tags = merge({
    Name = "${var.project_name}-subnet-private-${count.index}"
    Type = "private"
  }, local.base_tags)
}

resource "aws_route_table_association" "public_table_association" {
  count = local.number_az

  route_table_id = aws_route_table.public_table.id
  subnet_id      = aws_subnet.public_subnets[count.index].id
}

resource "aws_route_table_association" "private_table_association" {
  count = local.number_az

  route_table_id = aws_route_table.private_table.id
  subnet_id      = aws_subnet.private_subnets[count.index].id
}
