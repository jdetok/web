fetch('/select')
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP Error: ${response.status}`)
    }
    return response.json()
  }) 
  .then(players => {
    const container = document.getElementById("nba");
    container.innerHTML = ""; // clear in case of previous content

    players.forEach(player => {
      const playerDiv = document.createElement("div");
      playerDiv.innerHTML = `
        <h3>${player.player} (${player.team})</h3>
        <ul class="horiz">
        <li>Points: ${player.pts}</li>
        <li>Assists: ${player.ast}</li>
        <li>Rebounds: ${player.reb}</li>
        <li>FGM: ${player.fgm}</li>
        <li>FG%: ${Math.round(player.fg_pct * 100)}%</li>
        <li>3PM: ${player.fg3m}</li>
        <li>3P%: ${Math.round(player.fg3_pct * 100)}%</li>
        <li>FTM: ${player.ftm}</li>
        <li>FT%: ${Math.round(player.ft_pct * 100)}%</li>
        </ul>
        <hr>
      `;
      container.appendChild(playerDiv);
    });
  })
  .catch(error => {
    console.error("Error loading JSON:", error);
    document.getElementById("nba").innerText = "Failed to load player data.";
  });
