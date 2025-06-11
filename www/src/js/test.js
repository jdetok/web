async function getData(url) {
    const statusEl = document.getElementById('status');
    const outputEl = document.getElementById('output');
    const nbaEl = document.getElementById('nba');
    
    statusEl.textContent = 'Loading...';
    nbaEl.innerHTML = ''; 
    outputEl.innerHTML = '';

    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`HTTP Error: ${response.status}`)
        }
        
        const data = await response.json()
        statusEl.textContent = ''; 

        // outputEl.textContent = JSON.stringify(data, null, 2);
        data.forEach(player => {
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
        `
        nbaEl.appendChild(playerDiv);
        })
    }
    catch(error) {
        console.log(error);
        statusEl.textContent = 'Failed to load player data.';
    };
};

function careerStatsBtn() {
    document.addEventListener('DOMContentLoaded', () => {
    document.getElementById('fetchBtn').addEventListener('click', () => {
        getData("http://192.168.0.8:3002/select");
        });
    });
};

careerStatsBtn();

    

// async function main() {
//   const data = await getData("http://192.168.0.8:3002/select");
//   console.log(data);
// }

// main(); 

