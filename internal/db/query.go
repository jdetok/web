package db

// this file should just be string literals of queries to pass to the Select function

type Queries struct {
	CarrerStats string
	CarrerStatsByPlayer string
	Games string
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
	group by a.player, b.team	
	order by pts desc
	-- limit 30
`
var CarrerStatsByPlayer string = 
`
	select a.player, b.team, sum(c.pts) as pts, avg(c.pts) as pts_pg 
	from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join season d on d.season_id = c.season_id
	where a.active = 1
	and a.lg = "NBA"
	and d.season like "%RS"
	and player = ?
	group by a.player, b.team	
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