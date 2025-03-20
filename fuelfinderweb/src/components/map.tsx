import { FC, useEffect } from 'react';
import { useAppDispatch, useAppSelector } from '@/store';

import { useMap } from 'react-leaflet/hooks'
import { Marker } from 'react-leaflet/Marker' 
import { Icon, LatLngExpression } from 'leaflet'
import { TileLayer } from 'react-leaflet/TileLayer' 
import { MapContainer } from 'react-leaflet/MapContainer'

// import StationMarker from '../components/map/StationMarker';
// import PreferencesList from '../components/preferences/PreferencesList';
// import SummaryBar from '../components/preferences/SummaryBar';
import MapZoom from './map_zoom';
import Overview from './overview';
import StationMarker from './station_marker';
import StationList from './station_list';
import { Button } from './ui/button';
import { openMenu } from '@/slices/menu_slice';
// import StationListView from '../components/stations/StationListView';
// import StationListToggle from '../components/stations/StationListToggle';

export default function Map() {
    const stations = useAppSelector((state) => state.stations.value)
    const location = useAppSelector((state) => state.stations.location)

    const dispatch = useAppDispatch();
    const open = () => dispatch(openMenu('stations'))

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
            
        <Button onClick={open} className="absolute bottom-4 left-4 z-[1000]">test</Button>
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