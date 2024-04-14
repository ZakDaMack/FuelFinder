import { useState, useEffect } from 'react';

import { MapContainer } from 'react-leaflet/MapContainer'
import { TileLayer } from 'react-leaflet/TileLayer' 
import { useMap } from 'react-leaflet/hooks'

import Toolbar from './Toolbar';
import OfdBanner from './Banner';
import StationList from './StationList';
import StationMarker from './StationMarker';

export default function Map(props) {
    const [location, setLocation] = useState([51.4649, -0.1596]);
    const [loaded, setLoaded] = useState(false);

    const [stations, setStations] = useState([]);

    const getCurrentLocation = () => navigator.geolocation.getCurrentPosition(
        (pos) => {
            setLocation([pos.coords.latitude, pos.coords.longitude])
            setLoaded(true);
        },
        (err) => console.error("Failed to get location"));

    const fetchData = async () => {
        const res = await fetch(process.env.REACT_APP_API_URL + '?' + getQueryParams());
        const data = await res.json();
        setStations(data);
    }

    const getQueryParams = () => new URLSearchParams({
        lat: location[0],
        lng: location[1],
        distance: 10
    });

    // update loc on load
    useEffect(() => {
        getCurrentLocation();
    }, []);

    // update data on loc change
    useEffect(() => {
        if (!loaded) return;
        fetchData()
            .catch(err => console.error(err));
    }, [location, loaded]);


    return (
        <MapContainer center={location} zoom={13} scrollWheelZoom={false} className="Map__container">
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            {stations.map(s => (<StationMarker key={s._id} company={s} />))}
            <StationList stations={stations} />
            <Toolbar
                recentre={getCurrentLocation}
            />
            <RecentreAutomatically location={location} />
            <OfdBanner />
        </MapContainer>
    );
}

const RecentreAutomatically = ({ location }) => {
    const map = useMap();
    useEffect(() => {
       map.setView(location, 13);
    }, [location]);
    return null;
};