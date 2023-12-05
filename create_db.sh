docker run\
    -d\
    -p 5431:5432\
    --name skooter-postgres\
    -e POSTGRES_DB=skooterdb\
    -e POSTGRES_USER=admin\
    -e POSTGRES_PASSWORD=skooteradmin\
    postgres:latest