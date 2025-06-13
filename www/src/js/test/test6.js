// trying to dynamically create table from json keys

// function careerStatsBtn(data) {
//     document.addEventListener('DOMContentLoaded', () => {
//     document.getElementById('fetchBtn').addEventListener('click', () => {
//         getData("https://jdeko.me/select");
//         });
//     });
// };

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
            err = `HTTP Error: ${response.status}`
            console.log(err)
            throw new Error(`HTTP Error: ${response.status}`)
        }
        console.log("getting json");
        const data = await response.json();
        statusEl.textContent = ''; 
        console.log("json received");
        tableFromJSON(data, 2, ' - ');
    }
    catch(error) {
        console.log(error);
        statusEl.textContent = "Failed to load player data.";
    };
};

// TODO - numH  as arg to determine when to start inner loop
function tableFromJSON(data, numH, hDelim) {
    const container = document.getElementById("nba");
    container.innerHTML = "";
    const keys = jsonKeys(data);
    for (const r of data) {
        const pTable = document.createElement('table');
        pTable.style.border = '1px solid black';
        pTable.style.marginBottom = '1em';
        // first two keys are player and team - separate for header
        // console.log(`${r[keys[0]]} ${hDelim} ${r[keys[1]]}`);
        
        let hdr = "";
        let h = 0;

        while (h < numH) {
            hdr += r[keys[h]];
            h++;
            if (h < numH) {
                hdr += hDelim;
            }
        };

        const caption = document.createElement('caption');
        caption.textContent = hdr;
        caption.style.fontWeight = 'bold';
        pTable.appendChild(caption);
        
        // start inner loop at 2 to start with pts
        for (let i = numH; i < keys.length; i++) {
            const row = document.createElement('tr');

            const col = document.createElement('th');
            col.textContent = keys[i];
            col.style.textAlign = 'left';

            const val = document.createElement('td');
            val.textContent = r[keys[i]];

            row.appendChild(col);
            row.appendChild(val);
            pTable.appendChild(row);
            

            // console.log(`${keys[i]}: ${r[keys[i]]}`)
        };
        container.append(pTable);
    };
};

function jsonKeys(data) {
    return Object.keys(data[0]);
};

document.addEventListener('DOMContentLoaded', () => {
    document.getElementById('fetchBtn').addEventListener('click', () => {
        getData("https://jdeko.me/select");
    });
});