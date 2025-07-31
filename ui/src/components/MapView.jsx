import React, { useRef, useEffect } from 'react';
import mapboxgl from 'mapbox-gl';

mapboxgl.accessToken = 'YOUR_MAPBOX_ACCESS_TOKEN';

const MapView = () => {
  const mapContainer = useRef(null);

  useEffect(() => {
    new mapboxgl.Map({
      container: mapContainer.current,
      style: 'mapbox://styles/mapbox/streets-v11',
      center: [77.5946, 12.9716], // Bengaluru
      zoom: 12,
    });
  }, []);

  return <div ref={mapContainer} style={{ height: '500px', width: '100%' }} />;
};

export default MapView;
