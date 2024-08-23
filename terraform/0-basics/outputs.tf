output "vpc_id" {
  description = "VPC ID"
  value       = try(aws_vpc.simple_vpc.id, "")
}

output "vpc_name" {
  description = "VPC Name"
  value       = try(aws_vpc.simple_vpc.tags.Name, "")
}
