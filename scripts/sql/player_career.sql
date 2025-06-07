select a.player, b.team, sum(c.pts) as career_pts
from player a
inner join team b on b.team_id = a.team_id
inner join p_box c on c.player_id = a.player_id
inner join season d on d.season_id = c.season_id
where a.active = 1
and a.lg = "NBA"
and d.season like "%RS"
and player = ?
group by a.player, b.team