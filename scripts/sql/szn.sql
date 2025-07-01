select season_id, season_desc
from season;

select 
	a.player, 
	b.team, 
	sum(c.pts) as pts, 
	sum(c.ast) as ast,
	sum(c.reb) as reb,
	sum(d.fgm) as fgm,
	sum(d.fg3m) as fg3m,
	sum(d.ftm) as ftm
from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join p_shtg d 
		on d.player_id = a.player_id and d.game_id = c.game_id
	inner join season e on e.season_id = c.season_id
	where a.active = 1
	and a.lg = "NBA"
	and e.season_id = 22024
	group by a.player, b.team	
	order by pts desc;
	
    