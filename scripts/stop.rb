#!/usr/bin/env/ruby

require 'aws-sdk-rds'
require 'aws-sdk-ecs'

rds = Aws::RDS::Client.new

rds.describe_db_instances.db_instances.each do |db|
  # resp = rds.stop_db_instance(db_instance_identifier: db.db_instance_identifier)
end

ecs = Aws::ECS::Client.new
ecs.list_tasks(cluster: 'Sock-Shop').task_arns.each do |task|
  ecs.stop_task({
    cluster: 'Sock-Shop',
    task: task,
  })
end
