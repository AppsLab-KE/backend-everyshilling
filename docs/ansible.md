- Create a VPC with public routes and security groups.
- Creates relevant IAM Roles for ECS/FARGATE.
- Push docker image to AWS ECR.
- Create an AWS ECS cluster and service and task.
- Start ECS cluster task based on the container.
- Attach the ECS FARGATE cluster to an AWS ELB for dynamic scaling.
- scale number of nodes to run by changing size_of_cluster:  in ./vars/everything.yml
- This also puts logs in to AWS Cloudwatch.