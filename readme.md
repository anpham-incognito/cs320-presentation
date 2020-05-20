docker build -t myubuntu .
docker run -t -v ~/go/src/cs320:/shared -i myubuntu /bin/bash
