# waitserver

Simple Go/Chi Webserver that just waits and returns how long it waited.

Requests to `/wait` will return after a random time between 2 and 30 seconds with a 200 status.

Requests to `/wait?time=23` will return after 23 sec.

Requests to  `/wait?min=4&max=20` will return after some random time between 4 and 20 sec.

There is a docker image at `ghcr.io/100days/waitserver:main` which contains this server.
I can be configured with the following ENV Vars to controll chi-router's "throttle" middleware:
- THROTTLE_LIMIT : max number of concurrent requests
- THROTTLE_BACKLOG_LIMIT: max number of requests waiting for a free slot
- THROTTLE_BACKLOG_TIMEOUT:  max number of seconds a request stays in the backlog

Example can be seen in /docker-compose.yaml

