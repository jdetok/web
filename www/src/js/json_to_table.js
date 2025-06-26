// trying to dynamically create table from json

// pass no param, get all
document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('playerForm');
    form.addEventListener('submit', (event) => {
        event.preventDefault();
        let url = "https://jdeko.me/select";
        const player = encodeURIComponent(document.getElementById('playerInput').value.trim());
        const lg = encodeURIComponent(document.getElementById('league').value.trim());
        
        if (player.length > 1) {
            url += `/player?lg=${lg}&player=${player}`
            let imgUrl = `https://cdn.${lg}.com/headshots/${lg}/latest/1040x760/${player}.png`
            console.log(imgUrl)
            getData(url, 2, ' - ');
        } else {
            getData((url + `?lg=${lg}`), 2, ' - ');
        }
        
    });
});

async function getData(url, numH, hDelim) {
    const statusEl = document.getElementById('status');
    statusEl.textContent = 'Requesting data from API...';

    const outputEl = document.getElementById('output');
    outputEl.textContent = '';
    // changed these from innerHTML to textContent
    const nbaEl = document.getElementById('nba');
    nbaEl.innerHTML = ''; 
    
    // try to fetch JSON from API
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`HTTP Error: ${response.status}`)
        }
        // make data into json object and clear status message
        const data = await response.json();
        statusEl.textContent = ''; 
        
        // pass data 
        tableFromJSON(data, numH, hDelim);
    }
    catch(error) {
        console.log(error);
        statusEl.textContent = "Failed to load player data.";
    };
};

// dynamically create HTML table element with caption
// elements for caption MUST BE at the beginning of the json response
// numH is the number of json objects that will be used in the dynamic caption
// hDelim is the delimiter that separates the objects
// all objects after the first numH will go into the table
function tableFromJSON(data, numH, hDelim) {
    // clear the current nba container
    const container = document.getElementById("nba");
    container.innerHTML = "";

    const keys = jsonKeys(data);
    for (const r of data) {
        const pTable = document.createElement('table');

        // append keys together with delimiter based off numH and hDelim
        let hdr = "";
        let h = 0;
        while (h < numH) {
            hdr += r[keys[h]];
            h++;
            if (h < numH) {
                hdr += hDelim;
            }
        };

        // create a caption element with the hdr string
        const caption = document.createElement('caption');
        caption.textContent = hdr;
        pTable.appendChild(caption);
        
        // after creating header, create table with the data items 
        for (let i = numH; i < keys.length; i++) {
            const row = document.createElement('tr');

            const label = document.createElement('th');
            label.textContent = keys[i];
            label.style.textAlign = 'right';
            row.appendChild(label);

            const val = document.createElement('td');
            val.textContent = r[keys[i]];
            val.style.textAlign = 'left';
            row.appendChild(val);

            pTable.appendChild(row);
        };
        container.append(pTable);
    };
};

function jsonKeys(data) {
    return Object.keys(data[0]);
};