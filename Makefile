flash:
	@tinygo flash -size short -target lorae5 -ldflags="-X main.appEUI='$(shell cat ./keys/appeui)' -X main.devEUI='$(shell cat ./keys/deveui)' -X main.appKey='$(shell cat ./keys/appkey)'" .
