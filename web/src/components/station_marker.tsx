import { FC } from "react";
import { cn } from "@/lib/utils";
import { renderToString } from 'react-dom/server';

import { Station } from "@/models/station";

import { DivIcon } from "leaflet";
import BrandLogo from "./brand_logo";
import StationItem from "./station_item";
import { Popup } from 'react-leaflet/Popup' 
import { Marker } from 'react-leaflet/Marker' 

interface StationProps {
    station: Station;
}

const StationMarker: FC<StationProps> = ({ station }) => {
    const coords = [station.latitude, station.longitude] as [number, number];

    const stationIcon = new DivIcon({
      html: renderToString(
        <div data-site-id={station.site_id} className={cn(
          `bg-white relative border rounded-2xl grid place-items-center 
          [&_img]:max-h-[32px]! [&_img]:max-w-[50px]! w-full p-1 space-y-1

          shadow-sm hover:shadow-lg transition-transform duration-200
          hover:-translate-y-1 hover:scale-105 will-change-transform

          before:content-[""] before:absolute before:bottom-[-9px]
          before:left-1/2 before:-translate-x-1/2 before:w-0 before:h-0
          before:border-l-[9px] before:border-l-transparent
          before:border-r-[9px] before:border-r-transparent
          before:border-t-[9px] before:border-t-gray-200
          after:content-[''] after:absolute after:bottom-[-8px]
          after:left-1/2 after:-translate-x-1/2
          after:w-0 after:h-0
          after:border-l-[8px] after:border-l-transparent
          after:border-r-[8px] after:border-r-transparent
          after:border-t-[8px] after:border-t-white`
        )}>
          <BrandLogo brand={station.brand} />
          <div className="border rounded-xl w-full mx-1 p-[1px] bg-linear-to-r from-green-400 to-green-700">
            <p className="text-xs text-white text-center">{!!station.e10 && station.e10 > 0 ? station.e10 : '-'}</p>
          </div>
          <div className="border rounded-xl w-full mx-1 p-[1px] bg-linear-to-r from-yellow-400 to-yellow-700">
            <p className="text-xs text-white text-center">{!!station.b7 && station.b7 > 0 ? station.b7 : '-'}</p>
          </div>
        </div>
      ),
      className: 'flex! items-end hover:z-[10000]!', // to remove default 'leaflet-div-icon' class styles
      iconSize: [60, 100], // size of the icon
      iconAnchor: [30, 109], // point of the icon which will correspond to marker's location
      popupAnchor: [0, -100], // point from which the popup should open relative to the iconAnchor
    });

    return (
      <Marker position={coords} icon={stationIcon}>
        <Popup minWidth={300} className="[&_.leaflet-popup-content-wrapper]:rounded-2xl!">
          <StationItem station={station} />
        </Popup>
      </Marker>
    )
}

export default StationMarker;