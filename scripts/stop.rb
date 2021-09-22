#!/usr/bin/env/ruby

require 'aws-sdk-rds'
require 'aws-sdk-ecs'
require 'aws-sdk-ec2'

rds = Aws::RDS::Client.new

rds.describe_db_instances.db_instances.each do |db|
  rds.stop_db_instance(db_instance_identifier: db.db_instance_identifier)
end

ecs = Aws::ECS::Client.new
ecs.list_tasks(cluster: 'Sock-Shop').task_arns.each do |task|
  ecs.stop_task({
    cluster: 'Sock-Shop',
    task: task,
  })
end

ec2 = Aws::EC2::Client.new
ec2.describe_instances.reservations.each do |r|
  r.instances.each do |i|
    ec2.stop_instances({ instance_ids: [i.instance_id] })
  end
end
