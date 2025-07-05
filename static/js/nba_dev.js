/*
TODO: NBA TEAM LOGOS
https://cdn.nba.com/logos/nba/1610612738/primary/L/logo.svg
https://cdn.wnba.com/logos/wnba/1611661319/primary/L/logo.svg
*/

// BASE URLS
const base = "https://jdeko.me/bball";
const nbaHsBase = "https://cdn.nba.com/headshots/nba/latest/1040x760";
const wnbaHsBase = "https://cdn.wnba.com/headshots/wnba/latest/1040x760";

// CALL ANY LISTENERS HERE
document.addEventListener('DOMContentLoaded', () => {
    resetListener();
    searchListener();
    lgChangeListener();
    tmChangeListener
    loadSeasonOpts();
    loadTeamOpts();
});

async function search() {
    let lg = encodeURIComponent(
        document.getElementById('league').value.trim()
    );

    let szn = encodeURIComponent(
        document.getElementById('season').value.trim()
    );

    let tm = encodeURIComponent(
        document.getElementById('team').value.trim()
    );

    console.log(lg)
    console.log(tm)
    console.log(szn)
    let url = base + `/leaders?league=${lg}&season=${szn}&team=${tm}`
    const response = await fetch(url);
    if (!response.ok) { 
        throw new Error(`HTTP Error: ${response.status}`)
    } // CONVERT SUCCESSFUL RESPONSE TO JSON & CLEAR LOADMSG
    const data = await response.json();
    const d = document.getElementById('nbasumtest')
    d.innerHTML = ""
    // d.innerHTML = data[0].team;
    const keys = Object.keys(data[0]);
    for (const obj of data) { 
        for (let i = 0; i < keys.length; i++) {
            let k = document.createElement('p');
            k.textContent = `${keys[i]}: ${obj[keys[i]]}\n`;
            d.append(k)
        }
    }
}

function searchListener() {
    const btn = document.getElementById('searchBtn');
    btn.addEventListener('click', async (event) => {
        event.preventDefault();
        await search();
    });
};


// RESET FORM
function resetListener() {
    const btn = document.getElementById('rstBtn');
    btn.addEventListener('click', async (event) => {
        event.preventDefault();
        document.getElementById('playerForm').reset();
    });
};

async function lgChangeListener() {
    const slct = document.getElementById('league');
    slct.addEventListener('change', async (event) => {
        event.preventDefault();
        await loadTeamOpts();
    });
};

async function tmChangeListener() {
    const slct = document.getElementById('team');
    slct.addEventListener('change', async (event) => {
        event.preventDefault();
        await loadTeamOpts();
    });
};


// LOAD OPTIONS FOR SEASON SELECTOR
async function loadSeasonOpts() {
    try {
        const response = await fetch(base + '/seasons');
        if (!response.ok) { 
                throw new Error(`HTTP Error: ${response.status}`);
            } // CONVERT SUCCESSFUL RESPONSE TO JSON
        const data = await response.json();
        if (data[0] == '') {
            console.log('empty json');
        }

        const slct = document.getElementById('season');
        const defaultOpt = document.createElement('option');
        defaultOpt.textContent = `2024-2025 Regular Season`;
        defaultOpt.value = '22024'
        slct.appendChild(defaultOpt);
        // each player
        let i;
        for (i=0; i<data.length; i++){
            let opt = document.createElement('option');
            opt.textContent = data[i].Season;
            opt.value = data[i].SeasonId;
            slct.appendChild(opt);
            // console.log(data[i].Season);
        }   
    } catch (error) {
        console.error("failed to load seasons")
    }
};

// LOAD OPTIONS FOR TEAM SELECTOR
async function loadTeamOpts() {
    try {
        const response = await fetch(base + '/teams');
        if (!response.ok) { 
                throw new Error(`HTTP Error: ${response.status}`);
            } // CONVERT SUCCESSFUL RESPONSE TO JSON
        const data = await response.json();
        if (data[0] == '') {
            console.log('empty json');
        }

        let lg = document.getElementById('league').value.trim()
        
        const slct = document.getElementById('team');
        slct.innerHTML = ``;

        // default all teams option
        const defaultOpt = document.createElement('option');
        defaultOpt.textContent = `All ${lg.toUpperCase()} Teams`;
        defaultOpt.value = 'all'
        slct.appendChild(defaultOpt);
        // each player
        let i;
        for (i=0; i<data.length; i++){
            // TODO: only if team matches league selector
            if ((data[i].League).toLowerCase() === 
                    document.getElementById('league').value.trim()) {
                let opt = document.createElement('option');
                opt.textContent = data[i].CityTeam;
                opt.value = data[i].TeamAbbr;
                slct.appendChild(opt);
            } else {
                console.log('team not in league');
            }
            // console.log(data[i].Season);
        }   
    } catch (error) {
        console.error("failed to load seasons")
    }
};