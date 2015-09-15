ccron
-----

ccron is a cron system for clusters or clouds.

# Goals

ccron aims to be a replacement for the Heroku Scheduler.  I wanted it to be:

* best effort - jobs might run twice or not at all (but mostly they run as expected)
* self-service - developers can easily edit the schedule
* simple - reuse time-tested components like crond, cronlock, postgres and redis
* platform agnostic - ccron is easily deployed to any Docker-friendly cluster (but Docker is not required)

# Components

ccron has a few main components:

* an API server (run one or more instances)
* a cron daemon (run many instances across your cluster to ensure HA)
* a CLI to talk to the API server

Additionally, ccron requires a PostgreSQL server for persistence and a Redis server for locking.

# Installation

You'll need to fill in your postgres and redis details where appropriate.

    # create a database for ccron to use
    $ psql << EOF
    CREATE DATABASE ccron;
    CREATE ROLE ccron WITH LOGIN PASSWORD 'ccron';
    GRANT ALL ON DATABASE ccron to ccron;
    EOF

    # run the API server
    $ docker run -d -n ccron-api -e POSTGRES_URL="postgres://ccron:ccron@172.17.42.1/ccron" -e REDIS_URL="redis://x:pass@172.17.42.1" -p 5000:5000 goodeggs/ccron-api

    # run as many ccrond containers as it takes for you to sleep at night
    $ for i in 1 2; do docker run -d -n ccrond$i -e CCRON_SERVER="http://172.17.42.1:5000" goodeggs/ccrond; done

    # download ccron
    $ export CCRON_SERVER="http://172.17.42.1:5000"
    $ ccron create "* * * * * echo hello world"
    Id	Schedule	Command
    1  * * * * * echo hello world

# TODOs

* TESTS, TESTS, TESTS
* authentication

# Acknowledgements

* we outsource locking to [cronlock](https://github.com/kvz/cronlock)
* everything I know about Go I learned from [Convox](https://github.com/convox)

