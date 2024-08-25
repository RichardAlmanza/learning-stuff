data "aws_availability_zones" "aws_region_azs" {
  state = "available"

  lifecycle {
    precondition {
      condition     = parseint(split("/", var.cidr_block)[1], 10) + local.newbits <= 28
      error_message = "The current cidr block cannot support ${var.number_az * local.number_subnets_by_az} subnets"
    }

    postcondition {
      condition     = length(self.names) >= var.number_az
      error_message = "There are less AZs available than requested"
    }
  }
}

locals {
  number_subnets_by_az = 2
  newbits              = ceil(log(var.number_az * local.number_subnets_by_az, 2))

  base_tags = merge({
    Project = "${var.project_name}"
  }, var.base_tags)
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
  count  = var.number_az
  vpc_id = aws_vpc.vpc.id

  availability_zone = data.aws_availability_zones.aws_region_azs.names[count.index]
  cidr_block        = cidrsubnet(var.cidr_block, local.newbits, 2 * count.index)

  tags = merge({
    Name = "${var.project_name}-subnet-public-${count.index}"
    Type = "public"
  }, local.base_tags)
}

resource "aws_subnet" "private_subnets" {
  count  = var.number_az
  vpc_id = aws_vpc.vpc.id

  availability_zone = data.aws_availability_zones.aws_region_azs.names[count.index]
  cidr_block        = cidrsubnet(var.cidr_block, local.newbits, 2 * count.index + 1)

  tags = merge({
    Name = "${var.project_name}-subnet-private-${count.index}"
    Type = "private"
  }, local.base_tags)
}

resource "aws_route_table_association" "public_table_association" {
  count = var.number_az

  route_table_id = aws_route_table.public_table.id
  subnet_id      = aws_subnet.public_subnets[count.index].id
}

resource "aws_route_table_association" "private_table_association" {
  count = var.number_az

  route_table_id = aws_route_table.private_table.id
  subnet_id      = aws_subnet.private_subnets[count.index].id
}
