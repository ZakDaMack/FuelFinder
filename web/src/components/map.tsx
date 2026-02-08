import { FC, useEffect } from 'react';
import { useAppSelector, useAppDispatch } from '@/store';

import { LatLngExpression } from 'leaflet'
import { TileLayer } from 'react-leaflet/TileLayer' 
import { MapContainer } from 'react-leaflet/MapContainer'
import { useMap, useMapEvents } from 'react-leaflet/hooks'

import { fetchData } from '@/slices/station_slice';

import MapZoom from './map_zoom';
import Overview from './overview';
import CarMarker from './car_marker';
import StationMarker from './station_marker';
import StationSummary from './station_summary';
import LoadingIndicator from './loading_indicator';

export default function Map() {
    const stations = useAppSelector((state) => state.stations.value)
    const location = useAppSelector((state) => state.stations.location)

    return (
        <MapContainer center={location} zoom={13} scrollWheelZoom={false} className="w-screen h-screen [&>.leaflet-control-container]:hidden">
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            {stations?.map(s => (
                <StationMarker key={s.site_id} station={s} />
            ))}
            <CarMarker coords={location} />
            <MapZoom />
            <Overview />
            <StationSummary />
            <LoadingIndicator />
            <BoundsListener />
        
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

const BoundsListener: FC = () => {
    const dispatch = useAppDispatch();
    const map = useMapEvents({
        moveend: () => {
            dispatch(fetchData(map.getBounds()));
        },
        zoomend: () => {
            dispatch(fetchData(map.getBounds()));
        }
    });
    return null;
}