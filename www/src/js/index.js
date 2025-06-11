const statTable = {};

// function jsonToObj()
function showCareerStats() {

  fetch('js/player_career.json')
  // fetch('/select') // confirmed this works - using .json for speed
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
        <table>
          <thead>
            <tr>
              <th scope="col">Points</th>
              <th scope="col">Assists</th>
              <th scope="col">Rebounds</th>
              <th scope="col">FG Made</th>
              <th scope="col">FG %</th>
              <th scope="col">3s Made</th>
              <th scope="col">FG3 %</th>
              <th scope="col">FT Made</th>
              <th scope="col">FT %</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>${player.pts}</td>
              <td>${player.ast}</td>
              <td>${player.reb}</td>
              <td>${player.fgm}</td>
              <td>${Math.round(player.fg_pct * 100)}</td>
              <td>${player.fg3m}</td>
              <td>${Math.round(player.fg3_pct * 100)}</td>
              <td>${player.ftm}</td>
              <td>${Math.round(player.ft_pct * 100)}</td>
            </tr>
          </tbody>
        </table>
        `;
        container.appendChild(playerDiv);
      });
    })
    .catch(error => {
      console.error("Error loading JSON:", error);
      document.getElementById("nba").innerText = "Failed to load player data.";
    });
  }
