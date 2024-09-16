variable "vpc_id" {
  type = string
}

variable "public_subnet_id" {
  type = string
}

variable "public_ssh_key" {
  type = string
  default = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIM8OIkQr6NkrQLNXqz6kfzANbDPKBbiUPxrtgm6Y3Dka"
}

variable "external_ip_to_allows_ssh_connections" {
  type = string
}

variable "project_name" {
  type    = string
  default = "epam-task-tf"
}

variable "base_tags" {
  type = map(string)
  default = {
    "Enviroment" = "Temporal"
  }
  description = "Tags that will be added to all resource"
}
