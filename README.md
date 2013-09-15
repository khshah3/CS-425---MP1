CS425 MP1
=========

## Building
Call 'make' in root directory. 
There should be 'dgrep' and 'grepserver' executables.

## Running

To run the server on a given port, run:
    ./grepserver <port>

Be sure to specify the port that is in the config file so the 
client knows where to find your server.


To run a distributed grep query, run:
    ./dgrep <keyQuery> <valQuery>

The config file is specified in dgrep, it will read the list of 
servers and ports and query the appropriate log directories on each.
