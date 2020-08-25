.PHONY: new
new: 
	@cd ../devnet \
	&& ./builder.sh network default --ca \
	&& ./builder.sh channel default \
	&& ./builder.sh chaincode default

.PHONY: destroy
destroy:
	@cd ../devnet && ./builder.sh destroy