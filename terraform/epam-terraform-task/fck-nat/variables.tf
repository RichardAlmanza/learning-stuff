variable "vpc_id" {
  type = string
}

variable "public_subnet_id" {
  type = string
}

variable "cidr_block" {
  type = string
}

variable "private_route_table" {
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
