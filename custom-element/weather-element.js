class weather_element extends HTMLElement {
  connectedCallback() {
    this.innerHTML = `
         <div class="container mx-auto p-4">
          <h1 class="p-10 text-center text-3xl font-bold mb-4">Weather Status</h1>
          <div class="grid grid-cols-2 gap-4 text-center">
            <div class="bg-white p-4 rounded shadow">
              <p class="text-xl font-bold mb-2 text-blue-400">Water Status</p>
              <p class="text-lg mb-2" id="water"></p>
              <p class="text-sm text-gray-500" id="water-status"></p>
            </div>
            <div class="bg-white p-4 rounded shadow text-center">
              <p class="text-xl font-bold mb-2 text-green-400">Wind Status</p>
              <p class="text-lg mb-2" id="wind"></p>
              <p class="text-sm text-gray-500" id="wind-status"></p>
            </div>
          </div>
        </div>
        `;

    fetch('status.json')
      .then(response => response.json())
      .then(data => {
        console.log(data)

        // Mendapatkan referensi ke elemen-elemen dalam template
        const waterStatusElement = this.querySelector('#water-status');
        const water = this.querySelector('#water');
        const windStatusElement = this.querySelector('#wind-status');
        const wind = this.querySelector('#wind');

        // Mengatur konten elemen berdasarkan data
        water.textContent = `${data.status.water}`
        waterStatusElement.textContent = `${data.status.waterRes}`;
        wind.textContent = `${data.status.wind}`
        windStatusElement.textContent = `${data.status.windRes}`;
      })
      .catch(error => {
        console.error('Error fetching data:', error);
      });
  }
}

customElements.define('weather-element', weather_element);
