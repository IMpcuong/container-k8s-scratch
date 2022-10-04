FROM pahud/aws-sam-cli:latest
USER root

WORKDIR /app

# COPY . .
COPY handler/ .
COPY template.yaml ./template.yaml
COPY samconfig.toml ./samconfig.toml

RUN aws configure set aws_access_key_id "test" && \
  aws configure set aws_secret_access_key "test" && \
  aws configure set default.region ap-southeast-1
RUN find . -type f -iname "*.go"
RUN aws --version && sam --version
RUN cat /etc/*release
RUN export SAM_CLI_TELEMETRY=0; \
  sam build -t /app/template.yaml \
  --build-dir=.aws-sam/build \
  --base-dir=handler --debug

CMD ["sam", "deploy", "--config-file", "samconfig.toml", "--config-env", "test", "--debug"]
