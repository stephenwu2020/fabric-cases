CHAINCODE_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/roster
CHAINCODE_LABEL=roster
CHAINCODE_INVOKE_OPTION='{"Args":["Instantiate"]}'
CHAINCODE_QUERY_OPTION='{"Args":["Version"]}'

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