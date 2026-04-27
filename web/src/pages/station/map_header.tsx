import { FC, useEffect, useRef } from "react";

import { Icon, Map } from "leaflet";
import { Marker } from "react-leaflet/Marker";
import { TileLayer } from "react-leaflet/TileLayer";
import { MapContainer } from "react-leaflet/MapContainer";

const MapHeader: FC<{coords: [number, number]}> = ({ coords }) => {
  const mapRef = useRef<Map>(null);
  useEffect(() => {
    mapRef.current?.setView(coords, 13);
  }, [coords]);

  const stationIcon = new Icon({
    iconUrl: '/station.png',
    iconSize: [40, 50],
    iconAnchor: [20, 50],
    className: 'pointer-events-none cursor-default' // to allow clicks to pass through
  });

  return (
    <MapContainer 
      ref={mapRef} 
      center={coords} 
      zoom={13} 
      className="w-screen h-[200px] lg:h-[300px] pointer-events-none [&>.leaflet-control-container]:hidden"
      doubleClickZoom={false}
    >
      <TileLayer
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      <Marker position={coords} icon={stationIcon} />
    </MapContainer>
  );
}

export default MapHeader;