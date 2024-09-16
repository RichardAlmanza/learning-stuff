output "vpc_id" {
  value = aws_vpc.vpc.id
}

output "cidr_block" {
  value = aws_vpc.vpc.cidr_block
}

output "public_subnets_ids" {
  value = aws_subnet.public_subnets[*].id
}

output "private_subnets_ids" {
  value = aws_subnet.private_subnets[*].id
}

output "public_table_id" {
  value = aws_route_table.public_table.id
}

output "private_table_id" {
  value = aws_route_table.private_table.id
}
