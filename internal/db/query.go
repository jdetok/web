package db

// this file should just be string literals of queries to pass to the Select function

type Query struct {
	Args []string // arguments to accept
	Q string // query
}

type Queries struct {
	DbQueries []Query
}

var Players = Query{
	Args: []string{},
	Q: `
	select a.player_id
	from player a
	where a.player = ?
	and a.lg = ?
	`,
}

var AllPlayerStats = Query{
	Args: []string{},
	Q: `
	select a.player, b.team, 
		sum(c.pts) as pts, 
		sum(c.ast) as ast,
		sum(c.reb) as reb,
		sum(d.fgm) as fgm,
		sum(d.fg3m) as fg3m,
		sum(d.ftm) as ftm,
		avg(d.fg_pct) as fg_pct,
		avg(d.fg3_pct) as fg3_pct,
		avg(d.ft_pct) as ft_pct
		
	from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join p_shtg d 
		on d.player_id = a.player_id and d.game_id = c.game_id
	inner join season e on e.season_id = c.season_id
	where a.active = 1
	and a.lg = "NBA"
	and e.season like "%RS"
	group by a.player, b.team	
	order by pts desc
	`,
}

var LgPlayersStat = Query{
	Args: []string{"lg"},
	Q:`
	select a.player, b.team, 
		sum(c.pts) as pts, 
		sum(c.ast) as ast,
		sum(c.reb) as reb,
		sum(d.fgm) as fgm,
		sum(d.fg3m) as fg3m,
		sum(d.ftm) as ftm--,
		--avg(d.fg_pct) as fg_pct,
		--avg(d.fg3_pct) as fg3_pct,
		--avg(d.ft_pct) as ft_pct
		
	from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join p_shtg d 
		on d.player_id = a.player_id and d.game_id = c.game_id
	inner join season e on e.season_id = c.season_id
	where a.active = 1
	and a.lg = ?
	and e.season like "%RS"
	group by a.player, b.team	
	order by pts desc
`,
}

var LgPlayerStat = Query{
	Args: []string{"lg", "player"},
	Q:`
	select a.player, b.team, 
		sum(c.pts) as pts, 
		sum(c.ast) as ast,
		sum(c.reb) as reb,
		sum(d.fgm) as fgm,
		sum(d.fg3m) as fg3m,
		sum(d.ftm) as ftm-- ,
		-- avg(d.fg_pct) as fg_pct,
		-- avg(d.fg3_pct) as fg3_pct,
		-- avg(d.ft_pct) as ft_pct
		
	from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join p_shtg d 
		on d.player_id = a.player_id and d.game_id = c.game_id
	inner join season e on e.season_id = c.season_id
	where a.active = 1
	and a.lg = ?
	and a.player_id = ?
	and e.season like "%RS"
	group by a.player, b.team	
	order by pts desc
	`,
}


var CarrerStats string = 
`
	select a.player, b.team, 
		sum(c.pts) as pts, 
		sum(c.ast) as ast,
		sum(c.reb) as reb,
		sum(d.fgm) as fgm,
		sum(d.fg3m) as fg3m,
		sum(d.ftm) as ftm,
		avg(d.fg_pct) as fg_pct,
		avg(d.fg3_pct) as fg3_pct,
		avg(d.ft_pct) as ft_pct
		
	from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join p_shtg d 
		on d.player_id = a.player_id and d.game_id = c.game_id
	inner join season e on e.season_id = c.season_id
	where a.active = 1
	and a.lg = "NBA"
	and e.season like "%RS"
	-- and b.team = "LAL"
	group by a.player, b.team	
	order by pts desc
	
`
var CarrerStatsByLg string = 
`
	select a.player, b.team, 
		sum(c.pts) as pts, 
		sum(c.ast) as ast,
		sum(c.reb) as reb,
		sum(d.fgm) as fgm,
		sum(d.fg3m) as fg3m,
		sum(d.ftm) as ftm,
		avg(d.fg_pct) as fg_pct,
		avg(d.fg3_pct) as fg3_pct,
		avg(d.ft_pct) as ft_pct
		
	from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join p_shtg d 
		on d.player_id = a.player_id and d.game_id = c.game_id
	inner join season e on e.season_id = c.season_id
	where a.active = 1
	and a.lg = ?
	and e.season like "%RS"
	group by a.player, b.team	
	order by pts desc
`

var CarrerStatsByPlayer string = 
`
	select a.player, b.team, 
		sum(c.pts) as pts, 
		sum(c.ast) as ast,
		sum(c.reb) as reb,
		sum(d.fgm) as fgm,
		sum(d.fg3m) as fg3m,
		sum(d.ftm) as ftm,
		avg(d.fg_pct) as fg_pct,
		avg(d.fg3_pct) as fg3_pct,
		avg(d.ft_pct) as ft_pct
		
	from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join p_shtg d 
		on d.player_id = a.player_id and d.game_id = c.game_id
	inner join season e on e.season_id = c.season_id
	where a.active = 1
	and a.lg = ?
	and a.player = ?
	and e.season like "%RS"
	group by a.player, b.team	
	order by pts desc
`
var Games string = 
`
	select 
		b.team,
		c.season,
		a.game_date,
		a.diff,
		d.final,
		a.loc,
		a.ot
	
	from t_game a
	inner join team b on a.team_id = b.team_id
	inner join season c on a.season_id = c.season_id
	inner join game d on a.game_id = d.game_id

	where b.lg = "NBA"
	and a.season_id = 22024
`