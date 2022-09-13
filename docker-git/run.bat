@REM Dockerfile automation

@REM This new created image only contains `git`.
@REM And the entrypoint configures the container's mount-point that is exposing to be executable by the users.
docker container run --rm -it docker-git-test:auto status