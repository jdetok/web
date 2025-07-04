https://jdeko.me
full stack repo
written/maintained by Justin DeKock (jdeko17@gmail.com)

/api contains the main http server/router & all endpoints

/internal contains several tools for backend/internal use
 - /db to interact with mariadb
 - /jsonops to read/write/format json files
 - /store contains cache logic
 - /env contains functions to read environment variables
 - /logs contains log functions
 - /errs contains code for standardizing error handling

 /static contains static html/css/js/png files
 /scripts contains sql and sh scripts used primarily in deve