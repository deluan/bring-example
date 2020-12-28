
.PHONY: vnc
vnc:
	go run . vnc 10.1.0.11 5901 vnc vncpassword

.PHONY: rdp
rdp:
	go run . rdp 10.1.0.12 3389 root Docker

.PHONY: vbox
vbox:
	go run . rdp `ipconfig getifaddr en0` 3389 "" ""
