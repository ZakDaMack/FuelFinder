import { useState, useEffect } from 'react';

import { MapContainer } from 'react-leaflet/MapContainer'
import { TileLayer } from 'react-leaflet/TileLayer' 
import { Marker } from 'react-leaflet/Marker' 
import { Popup } from 'react-leaflet/Popup' 
import { useMap } from 'react-leaflet/hooks'

import Toolbar from './Toolbar';
import StationList from './StationList';
import StationMarker from './StationMarker';

export default function Map(props) {
    const [location, setLocation] = useState([51.4649, -0.1596]);
    const [stations, setStations] = useState([]);

    const getCurrentLocation = () => navigator.geolocation.getCurrentPosition(
        (pos) => setLocation([pos.coords.latitude, pos.coords.longitude]),
        (err) => console.error("failed to get location"));

    // update loc on load
    useEffect(() => {
        getCurrentLocation();
        if (stations.length == 0) {
            fetch('http://localhost:3001?' + new URLSearchParams({
                lat: location[0],
                lng: location[1],
                distance: 50
            }))
            .then(res => res.json())
            .then(data => setStations(data))
        }
    }, []);

    return (
        <MapContainer center={location} zoom={13} scrollWheelZoom={false} className="Map__container">
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            {/* <Marker position={location}>
                <Popup>
                    A pretty CSS3 popup. <br /> Easily customizable.
                </Popup>
            </Marker> */}
            {stations.map(s => (<StationMarker key={s._id} company={s} />))}
            <StationList />
            <Toolbar
                recentre={getCurrentLocation}
            />
            <RecentreAutomatically location={location} />
        </MapContainer>
    );
}

const RecentreAutomatically = ({ location }) => {
    const map = useMap();
    useEffect(() => {
       map.panTo(location);
    }, [location]);
    return null;
};