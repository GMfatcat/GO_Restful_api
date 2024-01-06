// main.js

// fetch single value from backend
function fetchData() {
    fetch('/getNumber')
        .then(response => response.json())
        .then(data => {
            document.getElementById('numberResult').innerText = 'Number: ' + data.number;
        })
        .catch(error => {
            console.error('Error fetching data:', error);
        });
}

// fetch json file from backend
function fetchJSON() {
    fetch('/getJSON')
        .then(response => response.json())
        .then(data => {
            // 將整個 JSON 內容印在網頁上
            document.getElementById('jsonResult').innerText = JSON.stringify(data, null, 2);
        })
        .catch(error => {
            console.error('Error fetching JSON:', error);
        });
}

// Download json file
function saveJSON() {
    fetch('/getJSON')
        .then(response => response.json())
        .then(data => {
            const jsonString = JSON.stringify(data, null, 2);
            const blob = new Blob([jsonString], { type: 'application/json' });
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
        })
        .catch(error => {
            console.error('Error fetching JSON:', error);
        });
}

function clearResults() {
    // 清除顯示的數字和 JSON 內容
    document.getElementById('jsonResult').innerText = '';
    document.getElementById('numberResult').innerText = 'Number: ';
}

// 新增函數，處理按鈕點擊事件，發送輸入框內的字串給後端
function sendInput() {
    const inputText = document.getElementById('inputText').value;

    // 使用 Fetch API 發送 POST 請求
    fetch('/receiveInput', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ inputText }),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Response from server:', data);
    })
    .catch(error => {
        console.error('Error sending input to server:', error);
    });
}