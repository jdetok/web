//  JS FOR NBA SITES PAGE
const url = "https://jdeko.me/bball";

// CALL ANY LISTENERS HERE
document.addEventListener('DOMContentLoaded', () => {
     formListener();
     randomListener();
});

function formListener() {
    const btn = document.getElementById('fetchBtn');
    btn.addEventListener('click', async (event) => {
        event.preventDefault();
        
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
        if (player.length < 1) { // EMPTY SEARCH BOX -> player=all
            player = 'all';
        } else {
            getHeadshot(lg, player);
        }
        // CONSTRUCT THE QUERY STRING
        const qUrl = (url + `/players?lg=${lg}&stype=${sType}&player=${player}`)
        getStats(qUrl, 2, ' - ');

    });
};

function randomListener() {
    const btn = document.getElementById('randBtn');
    btn.addEventListener('click', async (event) => {
        // ON CLICK, HIT ENDPOINT THAT PROVIDES RANDOM PLAYER
        // FILL THE SEARCH BOX WITH THAT PLAYER AND SUBMIT FORM
        // event.preventDefault();
        num = await getRandomPlayer();
        console.log(num);
        // const jsResp = response.json()

    })
}

async function getRandomPlayer() {
    console.log('random button pressed');
    const rUrl = url + `/players/random`;
    const response = await fetch(rUrl);
    if (!response.ok) {
        throw new Error(`HTTP Error getting player id: ${response.status}`);
    }
    const json = await response.json();
    return json.num;
}

async function getHeadshot(lg, player) {
    let playerId = await getPlayerId(url, player);
    let imgUrl = `https://cdn.${lg}.com/headshots/${lg}/latest/1040x760/${playerId}.png`;
    
    const container = document.getElementById('hs');
    container.innerHTML = '';
    const img = document.createElement('img');

    img.src = imgUrl;
    img.alt = "image not found"
    img.style.maxWidth = '50%';
    img.style.height = 'auto';
    img.style.marginLeft = 'auto';
    container.append(img);
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

// REQUEST STATS JSON FROM API
async function getStats(url, numCapFlds, capDelim) {
    // LOADING MESSAGE
    const loadMsg = document.getElementById('loadmsg');
    loadMsg.textContent = 'Requesting data from API...';

    // DIV TO CREATE STATS ELEMENTS
    const statsEl = document.getElementById('stats');
    statsEl.innerHTML = ''; 
    
    try { // WAIT FOR API RESPONSE
        const response = await fetch(url);
        if (!response.ok) { 
            throw new Error(`HTTP Error: ${response.status}`)
        } // CONVERT SUCCESSFUL RESPONSE TO JSON & CLEAR LOADMSG
        const data = await response.json();
        loadMsg.textContent = ''; 
        
        // CONVERT JSON RESPONSE TO HTML TABLE ELEMENTS
        tableFromJSON(data, numCapFlds, capDelim);
    }
    catch(error) {
        console.log(error);
        loadMsg.textContent = "Failed to load player data";
    };
};

// dynamically create HTML table element with caption
// fields used for MUST be the first #numCapFlds fields of each json object
// numCapFlds - number of fields in each json object to be used for the caption
// capDelim - string delimiter used in caption
async function tableFromJSON(data, numCapFlds, capDelim) {
    const div = document.getElementById("stats");

    div.innerHTML = ""; // clear the current nba container
    // GET KEYS
    // const keys = jsonKeys(data);
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
        
        div.append(objTbl);
        // div.append(objTbl); // APPEND TABLE TO DIV
    };
};