terraform {
  experiments = [module_variable_optional_attrs]
}

variable "service" {
  type = object({
    name               = string
    cluster_id         = string
    security_group_ids = list(string)
    subnet_ids         = list(string)
    load_balancer = optional(object({
      container_name   = string
      container_port   = number
      target_group_arn = string
    }))
  })
}

variable "task" {
  type = object({
    name               = string
    volume             = optional(string)
    family             = string
    image              = string
    execution_role_arn = string
    portMappings = optional(list(object({
      containerPort = number
      hostPort      = number
    })))
    memory  = optional(number)
    command = optional(list(string))
    environment = optional(list(object({
      name  = string
      value = string
    })))
    dockerLabels  = optional(map(string))
    task_role_arn = optional(string)
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
