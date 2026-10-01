output "vpc_id" {
  value = aws_vpc.main.id
}

output "public_subnet_ids" {
  value = aws_subnet.public[*].id
}

output "app_subnet_ids" {
  value = aws_subnet.app[*].id
}

output "database_subnet_ids" {
  value = aws_subnet.database[*].id
}
output "rds_endpoint" {
  value = aws_db_instance.postgres.address
}

output "rds_port" {
  value = aws_db_instance.postgres.port
}
output "alb_dns_name" {
  value = aws_lb.main.dns_name
}
output "autoscaling_group_name" {
  value = aws_autoscaling_group.app.name
}
