async function sendRequest() {
  const city = document.getElementById('city').value;
  const language = document.getElementById('language').value;

  const payload = { city, language };

  const response = await fetch('/api/weather', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  });

  const data = await response.json();

  document.getElementById('result').innerHTML = `
    <strong>Ciudad:</strong> ${data.location?.name || data.city || 'N/A'}<br>
    <strong>Temperatura:</strong> ${data.current?.temp_c || data.temperature || 'N/A'} °C<br>
    <strong>Descripción:</strong> ${data.current?.condition?.text || data.description || 'N/A'}
  `;

  loadHistory();
}

async function loadHistory() {
  const response = await fetch('/api/weather/history');
  const data = await response.json();

  const historyDiv = document.getElementById('history');
  historyDiv.innerHTML = "";

  if (data.length === 0) {
    historyDiv.innerText = "No hay consultas aún.";
    return;
  }

  data.forEach(item => {
    const div = document.createElement('div');
    div.classList.add('history-item');
    div.innerHTML = `
      <strong>${item.city}</strong> — ${item.temperature}°C — ${item.description}<br>
      <small>${new Date(item.queried_at).toLocaleString()}</small>
    `;
    historyDiv.appendChild(div);
  });
}

window.onload = () => {
  loadHistory();
};
