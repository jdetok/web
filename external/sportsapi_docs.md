The external folder contains code to call SportRadar's Sports API to retrieve NBA data as JSON

# GETTING TEAMS/PLAYERS
- Teams endpoint can be called to return all active
-- https://api.sportradar.com/nba/trial/v8/en/league/teams.json
 
- Put team ids in an array and loop through it to call team profile endpoint
-- this will return the specified team's players and other info
-- https://api.sportradar.com/nba/trial/v8/en/teams/<TEAM_ID>/profile.json
