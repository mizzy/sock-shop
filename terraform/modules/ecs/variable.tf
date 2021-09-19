terraform {
  experiments = [module_variable_optional_attrs]
}

variable "service" {
  type = object({
    name               = string
    cluster_id         = string
    security_group_ids = list(string)
    subnet_ids         = list(string)
    load_balancer      = any
  })
}

variable "task" {
  type = object({
    name               = string
    volume             = optional(string)
    family             = string
    image              = string
    execution_role_arn = string
    port               = number
    memory             = optional(number)
    command            = list(string)
    environment = list(object({
      name  = string
      value = string
    }))
    dockerLabels  = map(string)
    task_role_arn = string
    mountPoints = optional(list(object({
      containerPath = string
      sourceVolume  = string
    })))
  })
}

variable "registry" {
  type = object({
    name         = string
    namespace_id = string
  })
}
