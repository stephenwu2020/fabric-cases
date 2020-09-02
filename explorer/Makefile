.PHONY: download
download:
	wget https://raw.githubusercontent.com/hyperledger/blockchain-explorer/master/examples/net1/config.json
	wget https://raw.githubusercontent.com/hyperledger/blockchain-explorer/master/examples/net1/connection-profile/first-network.json -P connection-profile
	wget https://raw.githubusercontent.com/hyperledger/blockchain-explorer/master/docker-compose.yaml

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down -v

.PHONY: status
status:
	docker ps -a | grep mynetwork

.PHONY: logs
logs:
	docker logs explorer.mynetwork.com