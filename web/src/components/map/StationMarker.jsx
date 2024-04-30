import { Marker } from 'react-leaflet/Marker' 
import { Popup } from 'react-leaflet/Popup' 

import StationItem from '../Station';

export default function StationMarker(props) {
    const coords = [...props.company.location.coordinates].reverse();

    return (
        <Marker position={coords}>
            <Popup className="Map__popup">
                <StationItem company={props.company} />
            </Popup>
        </Marker>
    );
}
