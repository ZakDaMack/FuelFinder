import { FC, useEffect } from 'react';
import { useSelector } from 'react-redux';

import { MapContainer } from 'react-leaflet/MapContainer'
import { TileLayer } from 'react-leaflet/TileLayer' 
import { useMap } from 'react-leaflet/hooks'
import { Marker } from 'react-leaflet/Marker' 
import { Icon, LatLngExpression } from 'leaflet'

// import StationMarker from '../components/map/StationMarker';
// import PreferencesList from '../components/preferences/PreferencesList';
// import SummaryBar from '../components/preferences/SummaryBar';
import MapZoom from './map_zoom';
import Overview from './overview';
import { useAppSelector } from '@/store';
// import StationListView from '../components/stations/StationListView';
// import StationListToggle from '../components/stations/StationListToggle';

export default function Map() {
    
    // const stations = useSelector((state) => state.stations.value)
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
            {/* {stations?.map(s => (
                <StationMarker key={s.site_id} company={s} />
            ))} */}
            <MapZoom />
            <Overview />
            {/* <SummaryBar />
            <PreferencesList />
            <StationListView />
            <StationListToggle /> */}
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