variable "service" {
  type = object({
    name               = string
    cluster_id         = string
    security_group_ids = list(string)
    subnet_ids         = list(string)
  })
}

variable "task" {
  type = object({
    name               = string
    family             = string
    image              = string
    execution_role_arn = string
    port               = number
    memory             = number
    command            = list(string)
    environment = list(object({
      name  = string
      value = string
    }))
  })
}

variable "registry" {
  type = object({
    name         = string
    namespace_id = string
  })
}