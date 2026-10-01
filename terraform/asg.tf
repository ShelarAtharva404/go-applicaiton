# -------------------------
# Auto Scaling Group
# -------------------------
resource "aws_autoscaling_group" "app" {
  name = "${var.project_name}-asg"

  min_size         = 2
  max_size         = 2
  desired_capacity = 2

  vpc_zone_identifier = aws_subnet.app[*].id

  target_group_arns = [
    aws_lb_target_group.frontend.arn,
    aws_lb_target_group.backend.arn
  ]

  health_check_type         = "ELB"
  health_check_grace_period = 300

  launch_template {
    id      = aws_launch_template.app.id
    version = "$Latest"
  }

  tag {
    key                 = "Name"
    value               = "${var.project_name}-app"
    propagate_at_launch = true
  }

  tag {
    key                 = "Project"
    value               = var.project_name
    propagate_at_launch = true
  }
}
