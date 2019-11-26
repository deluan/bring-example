Here are the steps to run the app:

1) You'll need a working `guacd` server in your machine. The easiest way is using docker 
and docker-compose. Just call `docker-compose up -d` in the root of this project. It 
starts the `guacd` server and a sample headless linux with a VNC server

2) Run the sample app with `make run`. It will connect to the linux container started by docker.
