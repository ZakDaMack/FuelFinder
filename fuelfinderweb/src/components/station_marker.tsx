import { FC } from "react";

import { Station } from "@/models/station";

import { Icon } from "leaflet";
import StationItem from "./station_item";
import { Popup } from 'react-leaflet/Popup' 
import { Marker } from 'react-leaflet/Marker' 

const stationIcon = new Icon({
  iconUrl: "/station.png",
  iconSize: [40, 55], // size of the icon
  iconAnchor: [20, 55], // point of the icon which will correspond to marker's location
  popupAnchor: [0, -52], // point from which the popup should open relative to the iconAnchor
});

interface StationProps {
    station: Station;
}

const StationMarker: FC<StationProps> = ({ station }) => {
    const coords = [...station.location.coordinates].reverse() as [number, number];
    return (
        <Marker position={coords} icon={stationIcon}>
          <Popup minWidth={300}>
            <StationItem station={station} />
          </Popup>
        </Marker>
    )
}

export default StationMarker;