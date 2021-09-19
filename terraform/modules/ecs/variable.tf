variable "task" {
  type = object({
    name               = string
    family             = string
    image              = string
    execution_role_arn = string
    port               = number
    memory             = number
  })
}
