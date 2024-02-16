function fetchCryptoData() {
    const apiUrl = 'https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1';

    fetch(apiUrl)
        .then(response => response.json())
        .then(data => displayCryptoData(data))
        .catch(error => console.error('Error fetching crypto data:', error));
}

function displayCryptoData(data) {
    const cryptoTable = document.getElementById('cryptoTable');

    data.slice(0, 5).forEach(coin => {
        const row = cryptoTable.insertRow(-1);
        const cell1 = row.insertCell(0);
        const cell2 = row.insertCell(1);
        const cell3 = row.insertCell(2);

        cell1.textContent = coin.id;
        cell2.textContent = coin.symbol;
        cell3.textContent = coin.name;

        if (coin.symbol === 'usdt') {
            row.classList.add('usdt');
        } else if (data.indexOf(coin) < 5) {
            row.classList.add('highlight');
        }
    });
}

fetchCryptoData();
