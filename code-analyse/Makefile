CRYPTOGEN=../bin/cryptogen
CONFIGTXGEN=../bin/configtxgen
CRYPTO_CONFIG=./examplecom-config.yaml
CONFIGTX_CONFIG=./configtx.yaml
CRYPTO_PATH=crypto
GENESIS_BLOCK=./genesisblock

.PHONY: crypto
crypto: clean
	@echo "Generate materials..."
	@${CRYPTOGEN} generate --config=${CRYPTO_CONFIG} --output=${CRYPTO_PATH}
	@${CONFIGTXGEN} -profile SampleSoloSystemChannel -channelID testchannelid -outputBlock ${GENESIS_BLOCK}
	@echo "Done!"

.PHONY: clean
clean:
	@echo "Clean materials..."
	@rm -rf ${CRYPTO_PATH} ${GENESIS_BLOCK}
	@echo "Done"

.PHONY: commit
commit:
	@echo "Commiting..."
	@git add -A; git commit -am "Commit after code modified."
	@echo "Done!"

.PHONY: listener
listener:
	@lsof -nP  | grep LISTEN