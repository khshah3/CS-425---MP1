CS425 MP1
=========


## Building

You'll need [Go](http://golang.org) to build this project.

Call `make` in root directory. 
There should be `dgrep` and `grepserver` executables afterwards.


## Configuration

Edit `serverconfig.cfg` in the project's root directory to add/remove servers.
The dgrep client reads this file each time it is executed to get the list of
servers it should query. The file has one server on each line:

    <serverAdress:port> <log location>


## Running

To run the server on a given port, run:

    ./grepserver <port>

Be sure to specify the port that is in the config file so the 
client knows where to find your server.


To run a distributed grep query, run:

    ./dgrep <keyQuery> <valQuery>

The config file is specified in dgrep, it will read the list of 
servers and ports and query the appropriate log directories on each.
