// Load Mapbox CSS
const mapboxCSS = document.createElement('link');
mapboxCSS.rel = 'stylesheet';
mapboxCSS.href = 'https://api.mapbox.com/mapbox-gl-js/v2.15.0/mapbox-gl.css';
document.head.appendChild(mapboxCSS);

// Import Mapbox GL JS from CDN
const script = document.createElement('script');
script.src = 'https://api.mapbox.com/mapbox-gl-js/v2.15.0/mapbox-gl.js';
document.head.appendChild(script);

// Initialize map after script loads
script.onload = () => {
    initializeMap();
};

function initializeMap() {
    // Initialize Mapbox
    mapboxgl.accessToken = 'pk.eyJ1IjoiYnJhdm4iLCJhIjoiY20ydmJpN3dvMGRjdTJpcXl4bGd6bHJpeiJ9.mlEfzqY52K9hUCr9KvyIGQ';
    
    const map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/streets-v11',
        center: [0, 0],
        zoom: 2  // Starting with a wider view
    });

    // Add navigation controls
    map.addControl(new mapboxgl.NavigationControl());

    // Wait for map to load before fetching coordinates
    map.on('load', function() {
        console.log('Map loaded');
        fetchAndDisplayMarkers(map);
    });
}

function fetchAndDisplayMarkers(map) {
    fetch('/coordinates')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log('Received data:', data);
            addMarkersToMap(map, data);
        })
        .catch(error => {
            console.error('Error:', error);
            const mapContainer = document.getElementById('map');
            const errorDiv = document.createElement('div');
            errorDiv.className = 'map-error';
            errorDiv.textContent = 'Failed to load location data. Please try again later.';
            mapContainer.appendChild(errorDiv);
        });
}

function addMarkersToMap(map, locationData) {
    // Create bounds object to fit all markers
    const bounds = new mapboxgl.LngLatBounds();

    // Iterate through the location data
    Object.entries(locationData).forEach(([location, coords]) => {
        try {
            // Validate coordinates
            if (!coords || coords.length !== 2 || 
                !isFinite(coords[0]) || !isFinite(coords[1])) {
                console.warn(`Invalid coordinates for location: ${location}`);
                return;
            }

            // Format location name (remove underscores and capitalize)
            const formattedLocation = location
                .replace(/_/g, ' ')
                .replace(/-/g, ', ')
                .split(' ')
                .map(word => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
                .join(' ');

            // Create and add marker
            new mapboxgl.Marker()
                .setLngLat([coords[1], coords[0]]) // [longitude, latitude]
                .setPopup(
                    new mapboxgl.Popup({ offset: 25 })
                        .setHTML(`<h3>${formattedLocation}</h3>`)
                )
                .addTo(map);

            // Extend bounds to include this location
            bounds.extend([coords[1], coords[0]]);

        } catch (error) {
            console.warn(`Error adding marker for ${location}:`, error);
        }
    });

    // Fit map to bounds if we have any valid coordinates
    if (!bounds.isEmpty()) {
        map.fitBounds(bounds, {
            padding: 50,
            maxZoom: 15
        });
    }
}

// Add some basic styling
const style = document.createElement('style');
style.textContent = `
    #map {
        width: 100%;
        height: 400px;
        position: relative;
    }
    .map-error {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        background: rgba(255, 255, 255, 0.9);
        padding: 1rem;
        border-radius: 4px;
        color: #dc3545;
        text-align: center;
    }
    .mapboxgl-popup-content h3 {
        margin: 0 0 0.5rem 0;
        font-size: 1rem;
        color: black;
    }
  
`;
document.head.appendChild(style);