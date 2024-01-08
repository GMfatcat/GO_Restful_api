
function showMap() {
    console.log("Show Map")
        // get map from leaflet
        const mapContainer = L.DomUtil.get('map');
        // map center set as Taichung coords, zoom level = 5
        const map = L.map(mapContainer).setView([23.0, 115.5], 5);
        // add OpenstreetMap layer
        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© OpenStreetMap contributors'}).addTo(map);
        const jsonData = JSON.parse(document.getElementById('jsonResult').innerText);
        // console.log(jsonData);

        // loop through one array, but use array index to access data from other arrays
        jsonData.lats.forEach((lat,index) => {
            const lon = jsonData.lons[index];
            const count = jsonData.counts[index];

            if (lat && lon && count != undefined) {
                const radius = getRadius(count);
                L.circleMarker([lat, lon], { radius: radius, fillColor: '#ff7f50',
                    color: '#ff7f50', weight: 1, opacity: 1, fillOpacity: 0.8 })
                .addTo(map);
            }
        });

        function getRadius(count){
            return count * 2; // 2 time of count set to radius
        }
    }
    function saveJSON() {

        const jsonData = document.getElementById('jsonResult').innerText;

        if (jsonData !== '') {

            console.log("Save JSON");
            const blob = new Blob([jsonData], { type: 'application/json' });
            const url = URL.createObjectURL(blob);
            // 創建一個動態的 <a> 元素
            const a = document.createElement('a');
            a.href = url;
            // 使用者可選下載位置/名稱
            a.download = '';
            document.body.appendChild(a);
            // 觸發點擊事件，開始下載
            a.click();
            // 刪除動態創建的 <a> 元素
            document.body.removeChild(a);
            // 釋放 Blob 的 URL
            URL.revokeObjectURL(url);

        } else {
            console.log("No JSON data, nothing to save");
        }
    }

    function goBack() {
        window.history.back();
    }

	// Load JSON Data immediately after page load
    document.addEventListener('DOMContentLoaded', function () {
    // get value from localStorage
    var storedSearchInputValue = localStorage.getItem('searchInputValue');
    console.log('Stored Search Input:', storedSearchInputValue);

	// show in the page
	var searchResultElement = document.getElementById('searchResult');
	searchResultElement.innerHTML = 'Search Input: ' + storedSearchInputValue;

    // Send search input to backend and recieve results
    // 使用 Fetch API 發送 POST 請求
    fetch('/receiveInput', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({inputText: storedSearchInputValue}),
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('jsonResult').innerText = JSON.stringify(data, null, 2);
        console.log('Response json from server');
    })
    .catch(error => {
        console.error('Error fetching JSON:', error);
        console.error('Response text:', error.responseText);
    });

});
