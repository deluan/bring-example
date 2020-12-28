# Bring sample desktop app

This is a sample app showing off the [Bring Guacamole client](https://github.com/deluan/bring) capabilities.
It implements a simple VNC/RDP client, using the [Fyne GUI toolkit](https://github.com/fyne-io/fyne)

Here are the steps to run the app:

1) You'll need a working `guacd` server in your machine. The easiest way is using docker 
and docker-compose. Just call `docker-compose up -d` in the root of this project. It 
starts the `guacd` server and two sample headless linux: one with a VNC server and another with RDP

2) Run the sample app with `make rdp` or `make vnc`. It will connect to the appropriate linux container started by docker.


## TODO
- Handle Caps Lock ([waiting Fyne support](https://github.com/fyne-io/fyne/issues/552))
- Turn off mouse cursor (waiting Fyne support)
- Window resizing
- Dialog to open connection
