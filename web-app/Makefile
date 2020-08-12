.PHONY: new
new: 
	@rm -rf wallet
	@cd ../devnet \
	&& ./builder.sh network default --ca \
	&& ./builder.sh channel default \
	&& ./builder.sh chaincode default \
		--ccp /opt/gopath/src/github.com/hyperledger/chaincode/fabcar/go \
		--ccl fabcar \
		--cci '{"Args":["initLedger"]}' \
		--ccq '{"Args":["queryAllCars"]}'

.PHONY: destroy
destroy:
	@rm -rf wallet
	@cd ../devnet && ./builder.sh destroy