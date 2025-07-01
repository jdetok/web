package mariadb

// this file should just be string literals of queries to pass to the Select function

type Query struct {
	Args []string // arguments to accept
	Q string // query
}

type Queries struct {
	DbQueries []Query
}


var Players = Query {
	Args: []string{},
	Q: `
	select player_id, player, lg 
	from player 
	where lg in ("NBA", "WNBA") 
	group by player_id, player, lg
	`,
}

var PlayersOld = Query{
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

var Teams = Query{
	Args: []string{},
	Q: `
	SELECT a.lg, a.team_id, a.team, a.team_name
	FROM team a
	INNER JOIN ( 
		SELECT season_id, team_id
		FROM t_box
		WHERE LEFT(season_id, 1) = '2'
		AND RIGHT(season_id, 4) >= '2000'
		GROUP BY season_id, team_id
		) b ON b.team_id = a.team_id
	WHERE a.lg in ('NBA', 'WNBA')
	GROUP BY a.lg, a.team_id, a.team, a.team_name
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