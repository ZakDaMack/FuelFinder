import { FC, useEffect } from 'react';
import { useAppSelector } from '@/store';

import { useMap } from 'react-leaflet/hooks'
import { Marker } from 'react-leaflet/Marker' 
import { Icon, LatLngExpression } from 'leaflet'
import { TileLayer } from 'react-leaflet/TileLayer' 
import { MapContainer } from 'react-leaflet/MapContainer'

import MapZoom from './map_zoom';
import Overview from './overview';
import StationMarker from './station_marker';
import StationSummary from './station_summary';
import LoadingIndicator from './loading_indicator';

export default function Map() {
    const stations = useAppSelector((state) => state.stations.value)
    const location = useAppSelector((state) => state.stations.location)

    const carIcon = new Icon({
        iconUrl: '/car.png',
        iconSize: [40,30], // size of the icon
        iconAnchor: [20,15], // point of the icon which will correspond to marker's location
        popupAnchor: [-3, -76] // point from which the popup should open relative to the iconAnchor
    });

    return (
        <MapContainer center={location} zoom={13} scrollWheelZoom={false} className="w-screen h-screen [&>.leaflet-control-container]:hidden">
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            <Marker position={location} icon={carIcon} />
            {stations?.map(s => (
                <StationMarker key={s.site_id} station={s} />
            ))}
            <MapZoom />
            <Overview />
            <StationSummary />
            <LoadingIndicator />
         
            <RecentreAutomatically location={location} />
        </MapContainer>
    );
}

interface RecentreProps {
    location: LatLngExpression
}

const RecentreAutomatically: FC<RecentreProps> = ({ location }) => {
    const map = useMap();
    useEffect(() => {
       map.flyTo(location, 13);
    //    map.setView(location, 13);
    }, [location]);
    return null;
};