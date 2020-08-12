.PHONY: new
new: 
	@rm -rf wallet
	@cd ../devnet \
	&& ./builder.sh network default --ca \
	&& ./builder.sh channel default \
	&& ./builder.sh chaincode default

.PHONY: destroy
destroy:
	@rm -rf wallet
	@cd ../devnet && ./builder.sh destroy