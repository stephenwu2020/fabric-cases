CHAINCODE_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/health
CHAINCODE_LABEL=health
CHAINCODE_INVOKE_OPTION=""
CHAINCODE_QUERY_OPTION=""

.PHONY: new
new: 
	@rm -rf ./wallet
	@cd ../devnet \
	&& ./builder.sh network default --ca --couch \
	&& ./builder.sh channel default \
	&& ./builder.sh chaincode default \
		--ccp ${CHAINCODE_PATH} \
		--ccl ${CHAINCODE_LABEL} \
		--cci ${CHAINCODE_INVOKE_OPTION} \
		--ccq ${CHAINCODE_QUERY_OPTION} \

.PHONY: upgradeCC
upgradeCC:
	@cd ../devnet \
	&& ./builder.sh chaincode upgrade \
		--ccp ${CHAINCODE_PATH} \
		--ccl ${CHAINCODE_LABEL} \
		--cci ${CHAINCODE_INVOKE_OPTION} \
		--ccq ${CHAINCODE_QUERY_OPTION} \

.PHONY: destroy
destroy:
	@rm -rf ./wallet
	@cd ../devnet && ./builder.sh destroy

.PHONY: down
down:
	@cd ../devnet \
	&& ./builder.sh network down

.PHONY: up
up:
	@cd ../devnet \
	&& ./builder.sh network up --ca --couch