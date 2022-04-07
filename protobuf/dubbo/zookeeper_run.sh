docker-compose -f ./zookeeper.yml up -d

docker run --rm -p 2181:2181 zookeeper