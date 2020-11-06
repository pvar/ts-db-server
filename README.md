ts-db-server
============

#### Description

ts-db-server offers timezone data to HTTP clients who send a GET request on root. It listens gets the IP of the client,
sends it to IP-API and retrieves the client's location. This location is then used to find all timezone related information,
from a specially prepared database (see ts-db-generator). The information provided by the server includes timezone name,
current and next zones, transition times, location description, an so on. All data are encoded in JSON.

<b>NOTE:</b> It listens on localhost:8086.
