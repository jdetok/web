https://jdeko.me
full stack repo
written/maintained by Justin DeKock (jdeko17@gmail.com)

# TODO: random player button -> use random number within len(players) to pick a random player

/api contains the main http server/router & all endpoints
/internal contains several tools for backend/internal use
 - /db to interact with mariadb
 - /jsonops to read/write/format json files
 - /store contains cache logic
 - /env contains functions to read environment variables
 - /logs contains log functions
 - /errs contains code for standardizing error handling
 /external contains code for making external http requests
 - /get contains requests to external sports APIs
 - /clean contains structs for cleaning those responses
 - /pics contains functions to get player headshot .png files using player IDs
 /www contains static html/css/js/png files
 /scripts contains sql and sh scripts used primarily in development
 /docker contains Dockerfile files and docker compose files
 /json contains misc. json files from external/get responses