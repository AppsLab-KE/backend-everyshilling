# Docker Service for Spotlight Prism

### To test the mock api server Run

 N/B: Docker compose might generate errors...if the everyshilling.yml file is not placed under the root folder.

- Under the **root directory** on your terminal run:


```
docker-compose up
```

- In either a new terminal or your browser test out your endpoint using curl command.
Replace currencies with your endpoint

```
curl http://127.0.0.1:4010/currencies
```
# Postgres DB Service

```
To connect to a PostgreSQL database running inside a Docker container, you can follow these steps:

Start a PostgreSQL container using the official PostgreSQL image from Docker Hub:
docker run --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
This will start a PostgreSQL container with the name my-postgres, set the password for the postgres user to mysecretpassword, and run the container in detached mode (-d).



Once the container is running after doing docker compose up, you can connect to it using the psql command-line tool:

```
docker exec -it (your container name) psql -U postgres
```
To know what your contaner name you can run:
```docker ps ```

This command starts an interactive session with PostgreSQL databse running inside the container. Once connected to the databse, you can run SQL commands as usual.

Command to run multiple docker compose files
```
docker-compose -f docker-compose.yml docker-compose.dev.yml up
```

.......Pushing images to docker hub..........
docker ps
docker image tag <service_name>:<tag> <your_dockerhub_username>/<repository_name>:<tag>
docker image push <your_dockerhub_username>/<repository_name>:<tag>

#Dummy writing to see if my mckserver autodeployment works