import { FC } from "react";
import { renderToString } from 'react-dom/server';

import { DivIcon } from "leaflet";
import { Marker } from 'react-leaflet/Marker' 
import { faCar } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const CarMarker: FC<{
  coords: [number, number];
}> = ({ coords }) => {
    // const carIcon = new Icon({
    //   iconUrl: '/car.png',
    //   iconSize: [40,30], // size of the icon
    //   iconAnchor: [20,15], // point of the icon which will correspond to marker's location
    //   popupAnchor: [-3, -76] // point from which the popup should open relative to the iconAnchor
    // });
  
    const carIcon = new DivIcon({
      html: renderToString(
        <div className='relative'>
          {/* <div className='animate-ping duration-[3000ms] absolute inset-0 rounded-full bg-neutral-500 size-[28px]'></div> */}
          <FontAwesomeIcon fontSize={28} icon={faCar} className="z-[1000]"/>
        </div>
      ),
      className: '', // to remove default 'leaflet-div-icon' class styles
      iconSize: [28, 28], // size of the icon
      iconAnchor: [14, 14], // point of the icon which will correspond to marker's location
      popupAnchor: [0, -40], // point from which the popup should open relative to the iconAnchor
    });

    return (
      <Marker position={coords} icon={carIcon} />
    )
}

export default CarMarker;