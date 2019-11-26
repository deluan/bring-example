
.PHONY: run
run:
	go run . vnc 10.1.0.11 5901

.PHONY: rdp
rdp:
	go run . rdp `ipconfig getifaddr en0` 3389
