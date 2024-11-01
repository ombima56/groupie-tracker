document.addEventListener('DOMContentLoaded', function () {
  // Add event listeners for tabs
  const tabs = document.querySelectorAll('.tab');
  tabs.forEach((tab) => {
    tab.addEventListener('click', function () {
      switchTab(this);
    });
  });
});

function switchTab(clickedTab) {
  const tabs = document.querySelectorAll('.tab');
  tabs.forEach((tab) => {
    tab.classList.remove('active');
  });
  clickedTab.classList.add('active');
  // Add logic here to show/hide content based on the selected tab
}

document.addEventListener('DOMContentLoaded', function () {
  // Add event listeners for tabs
  const tabs = document.querySelectorAll('.tab');
  tabs.forEach((tab) => {
    tab.addEventListener('click', function () {
      switchTab(this);
    });
  });

  // Add scroll event listener
  var topbar = document.getElementById('topbar');
  var mainContent = document.getElementById('main-content');
  window.addEventListener('scroll', function () {
    const topbar = document.querySelector('.topbar');
    const mainContent = document.querySelector('.main-content');
    if (mainContent) {
      const mainContentTop = mainContent.offsetTop;
      if (window.scrollY >= mainContentTop) {
        topbar.classList.add('scrolled');
      } else {
        topbar.classList.remove('scrolled');
      }
    }
  });
});

function openTab(evt, tabName) {
  var i, tabcontent, tablinks;
  tabcontent = document.getElementsByClassName('tab-content');
  for (i = 0; i < tabcontent.length; i++) {
    tabcontent[i].style.display = 'none';
  }
  tablinks = document.getElementsByClassName('tab');
  for (i = 0; i < tablinks.length; i++) {
    tablinks[i].className = tablinks[i].className.replace(' active', '');
  }
  document.getElementById(tabName).style.display = 'block';
  evt.currentTarget.className += ' active';
}

// Initialize the map
var map = L.map('map').setView([0, 0], 2); // Default view (global)

// Load the OpenStreetMap tile layer
L.tileLayer('https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png', {
  maxZoom: 19,
  attribution: '© OpenStreetMap contributors'
}).addTo(map);

// Select all concert location elements
const locationElements = document.querySelectorAll('.relation-location');
const locations = Array.from(locationElements).map(element => element.textContent.trim());

// Array to store marker coordinates for fitting map bounds later
const markerCoords = [];

// Function to place a marker based on a location name
function placeMarker(location) {
  // Geocode location using OpenStreetMap's Nominatim API
  return fetch(`https://nominatim.openstreetmap.org/search?format=json&q=${encodeURIComponent(location)}&accept-language=en`)
    .then(response => response.json())
    .then(data => {
      if (data.length > 0) {
        const lat = parseFloat(data[0].lat);
        const lon = parseFloat(data[0].lon);

        // Store coordinates to adjust map bounds later
        markerCoords.push([lat, lon]);

        // Create a marker for each location
        L.marker([lat, lon]).addTo(map)
          .bindPopup(`${location}`);
      } else {
        console.error("No results found for:", location);
      }
    })
    .catch(error => console.error("Error fetching data:", error));
}

// Geocode and place markers for all locations
Promise.all(locations.map(location => placeMarker(location)))
  .then(() => {
    // Adjust the map view to include all markers
    if (markerCoords.length > 0) {
      map.fitBounds(markerCoords);
    } else {
      console.warn("No valid locations found to center the map.");
    }
  });
