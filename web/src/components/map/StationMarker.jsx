import { Marker } from 'react-leaflet/Marker' 
import { Popup } from 'react-leaflet/Popup' 
import { Icon } from 'leaflet'

import StationItem from '../Station';

const stationIcon = new Icon ({
    iconUrl : '/favicon.ico',
    iconSize : [40,40], // size of the icon
    iconAnchor : [20,40], // point of the icon which will correspond to marker's location
    popupAnchor : [0, -40] // point from which the popup should open relative to the iconAnchor
});

export default function StationMarker(props) {
    const coords = [...props.company.location.coordinates].reverse();

    return (
        <Marker position={coords} icon={stationIcon}>
            <Popup className="Map__popup">
                <StationItem company={props.company} />
            </Popup>
        </Marker>
    );
}
