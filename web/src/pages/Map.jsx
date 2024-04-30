import { useEffect } from 'react';
import { useSelector } from 'react-redux';

import { MapContainer } from 'react-leaflet/MapContainer'
import { TileLayer } from 'react-leaflet/TileLayer' 
import { useMap } from 'react-leaflet/hooks'
import { Marker } from 'react-leaflet/Marker' 
import { Icon } from 'leaflet'

import OfdBanner from '../components/map/Banner';
import StationList from '../components/stations/StationList';
import StationMarker from '../components/map/StationMarker';
import MapToolbar from '../components/map/MapToolbar';
import PreferencesList from '../components/preferences/PreferencesList';

export default function Map() {
    
    const stations = useSelector((state) => state.stations.value)
    const location = useSelector((state) => state.stations.location)

    const carIcon = new Icon ({
        iconUrl : '/car.png',
        iconSize : [40,40], // size of the icon
        iconAnchor : [20,20], // point of the icon which will correspond to marker's location
        popupAnchor : [-3, -76] // point from which the popup should open relative to the iconAnchor
    });

    return (
        <MapContainer center={location} zoom={13} scrollWheelZoom={false} className="Map__container">
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            <Marker position={location} icon={carIcon}>
            </Marker>
            {stations?.map(s => (
                <StationMarker key={s.site_id} company={s} />
            ))}
            <StationList />
            <PreferencesList />
            <MapToolbar />
            <RecentreAutomatically location={location} />
            {/* <OfdBanner /> */}
        </MapContainer>
    );
}

const RecentreAutomatically = ({ location }) => {
    const map = useMap();
    useEffect(() => {
       map.flyTo(location, 13);
    //    map.setView(location, 13);
    }, [location]);
    return null;
};