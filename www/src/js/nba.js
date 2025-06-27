//  JS FOR NBA SITES PAGE
let url = "https://jdeko.me/bball";

document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('playerForm');
    form.addEventListener('submit', async (event) => {
        event.preventDefault();
        
        // PARAMETERS PASSED
        let player = encodeURIComponent(
            document.getElementById('playerInput').value.trim()
        ).toLowerCase();
        const lg = encodeURIComponent(
            document.getElementById('league').value.trim()
        );
        const sType = encodeURIComponent(
            document.getElementById('statType').value.trim()
        );
        
        // CHECK IF USER SPECIFIED A PLAYER IN THE SEARCH BOX
        if (player.length < 1) { // EMPTY SEARCH BOX -> player=all
            player = 'all';
        } else {
            console.log(player);
            let playerId = await getPlayerId(url, player);
            console.log(playerId);
            let imgUrl = `https://cdn.${lg}.com/headshots/${lg}/latest/1040x760/${playerId}.png`;
            console.log(imgUrl);
            const container = document.getElementById('hs');
            container.innerHTML = '';
            const img = document.createElement('img');
        // img.src = imgSrc.path;
            img.src = imgUrl;
            img.alt = "image not found"
            img.style.maxWidth = '50%';
            img.style.height = 'auto';
            img.style.marginLeft = 'auto';
            container.append(img);
            // div.append(img);

        }
        // CONSTRUCT THE QUERY STRING
        const qUrl = (url + `/players?lg=${lg}&stype=${sType}&player=${player}`)
        getData(qUrl, 2, ' - ');

    }); 
});

async function getPlayerId(url, player) {
    const idUrl = url + `/players/id?player=${player}`;
    const response = await fetch(idUrl);
    if (!response.ok) {
        throw new Error(`HTTP Error getting player id: ${response.status}`)
    }
    const jsonResp = await response.json();
    // if (!jsonResp.ok) {
    //     throw new Error(`HTTP Error getting player id json: ${jsonResp.status}`)
    // }
    const playerId = jsonResp.playerId;
    return String(playerId);
};

// REQUEST JSON FROM API
async function getData(url, numCapFlds, capDelim) {
    // LOADING MESSAGE
    const loadMsg = document.getElementById('loadmsg');
    loadMsg.textContent = 'Requesting data from API...';

    // DIV TO CREATE STATS ELEMENTS
    const nbaEl = document.getElementById('nba');
    nbaEl.innerHTML = ''; 
    
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

// LIST OF KEYS IN EACH OF THE JSON OBECTS
function jsonKeys(data) { // all objects will be the same - just get keys from the first
    return Object.keys(data[0]);
};

// // REQUEST IMG FROM API
async function getImg(data, keys) {
    const key = keys[0];
    
    player = encodeURIComponent(data[0][key])
    console.log(player);
    // API CALL HERE
    // let pUrl = (url + `/players/headshot?lg=${lg}player=${player}`)
    let pUrl = (url + `/players/headshot?player=${player}`)
    console.log(pUrl);
    const response = await fetch(pUrl);
    if (!response.ok) { 
        throw new Error(`HTTP Error: ${response.status}`)
    } // CONVERT SUCCESSFUL RESPONSE TO JSON & CLEAR LOADMSG
    
    const imgSrc = await response.json();
    
    console.log(imgSrc.path)
    return imgSrc
};

// dynamically create HTML table element with caption
// fields used for MUST be the first #numCapFlds fields of each json object
// numCapFlds - number of fields in each json object to be used for the caption
// capDelim - string delimiter used in caption
async function tableFromJSON(data, numCapFlds, capDelim) {
    const div = document.getElementById("nba");

    div.innerHTML = ""; // clear the current nba container
    // GET KEYS
    const keys = jsonKeys(data);

        // TODO - IF SINGLE PLAYER REQUEST, GET THE PLAYER'S PICTURE
    // if (data.length == 1 && keys[0] === 'player') {
        // const imgSrc = await getImg(data, keys)
        // console.log(imgSrc.path)
    //     const img = document.createElement('img');
    //     // img.src = imgSrc.path;
    //     img.src = imgSrc.path;
    //     img.alt = "image not found"
    //     img.style.maxWidth = '40%';
    //     img.style.height = 'auto';
    //     img.style.marginLeft = 'auto';
    //     div.append(img);
    // }

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