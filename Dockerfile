FROM partlab/ubuntu-golang

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies

RUN apt update
RUN apt install -y file gcc

WORKDIR /shared
COPY . /shared