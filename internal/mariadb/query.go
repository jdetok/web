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
	limit 1
	`,
}

var Seasons = Query{
	Args: []string{},
	Q: `
	select season_id, season_desc, wseason_desc
	from season
	where left(season_id, 1) in ('2', '4')
	and right(season_id, 4) >= 2000
	-- and right(season_id, 4) >= year(sysdate()) - 15
	order by right(season_id, 4) desc, left(season_id, 1)
	`,
}


// -- and a.lg = ?
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

var LgPlayerAvg = Query {
	Args: []string{"lg", "player"},
	Q: `
	select 
	a.player,
	b.team, 
	round(avg(c.pts), 2) as pts,
	round(avg(c.ast), 2) as ast,
	round(avg(c.reb), 2) as reb,
	round(avg(d.fgm), 2) as fgm,
	round(avg(d.fg3m), 2) as fg3m,
	round(avg(d.ftm), 2) as ftm,
	round(avg(d.fg_pct), 2) as fg_pct,
	round(avg(d.fg3_pct), 2) as fg3_pct,
	round(avg(d.ft_pct), 2) as ft_pct
	
	from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join p_shtg d on d.player_id = a.player_id and d.game_id = c.game_id
	inner join season e on e.season_id = c.season_id
	
	where a.active = 1
	and e.season like "%RS"
	and a.lg = ?
	and a.player_id = ?
	
	group by a.player, b.team	
	order by pts desc;
	`,
}