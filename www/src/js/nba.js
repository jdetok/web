// BASE URLS
const base = "https://jdeko.me/bball";
const nbaHsBase = "https://cdn.nba.com/headshots/nba/latest/1040x760";
const wnbaHsBase = "https://cdn.wnba.com/headshots/wnba/latest/1040x760";

// CALL ANY LISTENERS HERE
document.addEventListener('DOMContentLoaded', () => {
    searchListener();
    randomListener();
});

// listen for the search button
function searchListener() {
    const btn = document.getElementById('searchBtn');
    btn.addEventListener('click', async (event) => {
        event.preventDefault();
        await search();
    });
};

// listen fro the random player button
function randomListener() {
    const btn = document.getElementById('randBtn');
    btn.addEventListener('click', async (event) => {
        event.preventDefault();
        // get random player json from api
        json = await getRandomPlayer();

        // search with input & league selector as random player vals
        document.getElementById('playerInput').value = json.player;
        document.getElementById('league').value = json.league;
        await search();
    });
}

// get headshot & stats based on element values
async function search() {
    // PARAMETERS PASSED
    let player = encodeURIComponent(
        document.getElementById('playerInput').value.trim()
    ).toLowerCase();
    let lg = encodeURIComponent(
        document.getElementById('league').value.trim()
    );
    let sType = encodeURIComponent(
        document.getElementById('statType').value.trim()
    );
    
    // CHECK IF USER SPECIFIED A PLAYER IN THE SEARCH BOX
    data = await getData(lg, sType, player)
    if (data[1] != '') {
        const hdrEl = document.getElementById('playerteam');
        hdr = await getHeader(data[0], 2);
        console.log(hdr);
        await appendHdr(hdr, hdrEl);
        // hdrEl.textContent = hdr
        await appendImg(data[1], 'hs');
    }
    await tableFromJSON(data[0], 2, ' - ');
    // doing this to get player name and team as header

    document.getElementById('playerInput').value = '';
}

async function appendHdr(playerTeam, el) {
    const hdr = document.createElement('h3');
    hdr.innerHTML = playerTeam;
    el.innerHTML = ''
    el.appendChild(hdr);
}

async function getData(lg, sType, player) {
    // empty player search makes player = all
    if (player.length < 1) { // EMPTY SEARCH BOX -> player=all
        player = 'all';
    } 
    // get stats data
    const url = (base + `/players?lg=${lg}&stype=${sType}&player=${player}`)
    stats = await getStats(url);

    // // derive player/team header from data
    // hdr = await getHeader(stats, 2); 
    // console.log(hdr);
    if (player != 'all') {
        hs = await getHeadshot(lg, player);
    } else {
        hs = ''
    }
    return [stats, hs]
}


// get stats with built url
async function getStats(url) {
    const loadMsg = document.getElementById('loadmsg');
    loadMsg.textContent = 'Requesting data from API...';
    try { // WAIT FOR API RESPONSE
        const response = await fetch(url);
        if (!response.ok) { 
            throw new Error(`HTTP Error: ${response.status}`)
        } // CONVERT SUCCESSFUL RESPONSE TO JSON & CLEAR LOADMSG
        const data = await response.json();
        loadMsg.textContent = ''; 
        // CONVERT JSON RESPONSE TO HTML TABLE ELEMENTS
        return data
    }
    catch(error) {
        console.log(error);
        loadMsg.textContent = "Failed to load player data";
    };
};

async function getHeader(data, numF) {
    const keys = Object.keys(data[0]);
    let hdr = "";
    console.log(data[0]);
    console.log(keys);
    
    for (i=0; i<numF; i++) {
        hdr += data[0][keys[i]];
        console.log(data[0][keys[i]])
        if (i < numF - 1) {
            hdr += " - ";
        }
    }
    console.log(hdr);
    return hdr;
}

// get player's headshot
async function getHeadshot(lg, player) {
    let playerId = await getPlayerId(base, player);
    let url = `https://cdn.${lg}.com/headshots/${lg}/latest/1040x760/${playerId}.png`
    return makeImg(url);
}



 

// make html image with built url
async function makeImg(url) {
    const img = document.createElement('img');
    img.src = url;
    img.alt = "image not found"
    img.style.maxWidth = '50%';
    img.style.height = 'auto';
    img.style.marginLeft = 'auto';
    return img
}

// append image to document with src url
async function appendImg(img, el) {
    const container = document.getElementById(el);
    container.innerHTML = '';
    container.appendChild(img);
}

// returns random player json from api
async function getRandomPlayer() {
    const url = base + `/players/random`;
    const response = await fetch(url)
    if (!response.ok) {
        throw new Error(`HTTP Error getting player id: ${response.status}`);
    }
    const json = await response.json();
    return json;
}


// pass encoded player name to /players/id to get the player id
async function getPlayerId(url, player) {
    const idUrl = url + `/players/id?player=${player}`;
    const response = await fetch(idUrl);
    if (!response.ok) {
        throw new Error(`HTTP Error getting player id: ${response.status}`)
    }
    const jsonResp = await response.json();
    const playerId = jsonResp.playerId;
    return String(playerId);
};



// dynamically create HTML table element with caption
// fields used for MUST be the first #numCapFlds fields of each json object
// numCapFlds - number of fields in each json object to be used for the caption
// capDelim - string delimiter used in caption
async function tableFromJSON(data, numCapFlds, capDelim) {
    // DIV TO CREATE STATS ELEMENTS
    const statsEl = document.getElementById('stats');
    statsEl.innerHTML = ''; 

    const keys = Object.keys(data[0]);
    for (const obj of data) { 
        const objTbl = document.createElement('table');

        // FIRST numCapFlds OBJECTS WILL CONSTRUCT CAPTION STRING
        let cap = "";
        let capVal = 0;
        while (capVal < numCapFlds) {
            cap += obj[keys[capVal]];
            capVal++; //  e.g LeBron James - LAL
            if (capVal < numCapFlds) { // ignore last value so no delim at end
                cap += capDelim; // concat capDelim between capVals
            }
        };

        // CREATE HTML CAPTION ELEMENT FROM cap 
        const caption = document.createElement('caption');
        caption.textContent = cap;
        objTbl.appendChild(caption); // APPEND CAPTION TO HTML TABLE
        
        // LOOP THROUGH FIELDS > numCapFlds, EACH LOOP APPENDS A ROW TO TABLE
        for (let i = numCapFlds; i < keys.length; i++) {
            const row = document.createElement('tr');
            const label = document.createElement('th');
            const val = document.createElement('td');

            // FIELD NAME IN LEFT COLUMN OF TABLE (RIGHT ALIGNED)
            label.textContent = keys[i];
            label.style.textAlign = 'right';

            // VALUE IN RIGHT COLUMN OF TABLE (LEFT ALIGNED)
            val.textContent = obj[keys[i]];
            val.style.textAlign = 'left';
            
            row.appendChild(label); // APPEND LABEL TO ROW
            row.appendChild(val); // APPEND VALUE TO ROW
            objTbl.appendChild(row); // APPEND ROW TO TABLE
        };
        
        statsEl.append(objTbl);
        // div.append(objTbl); // APPEND TABLE TO DIV
    };
};
