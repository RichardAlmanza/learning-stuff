output "ssh_server_id" {
  value = aws_instance.ssh_server.id
}

output "ssh_server_sg_id" {
  value = aws_security_group.ssh_server.id
}
