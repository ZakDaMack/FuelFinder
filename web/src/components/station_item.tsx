import { FC } from "react";
import { cn } from "@/lib/utils";
import { useMediaQuery } from "@uidotdev/usehooks";

import { Station } from "@/models/station";

import BrandLogo from "./brand_logo";
import { faGasPump } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import dayjs from 'dayjs';
import calendar from 'dayjs/plugin/calendar';
dayjs.extend(calendar)

interface StationProps {
    station: Station;
    onHover?: (enter: boolean) => void;
}

const fuels = {
    'e5'  : {colour: 'from-blue-400 to-blue-700', title: 'Super'},
    'e10' : {colour: 'from-green-400 to-green-700', title: 'Petrol'},
    'b7'  : {colour: 'from-yellow-400 to-yellow-700', title: 'Diesel'},
    // 'sdv':{colour: '', title: 'SDV'},
}

const StationItem: FC<StationProps> = ({ station, onHover }) => {
    const coords = `${station.location.coordinates[1]},${station.location.coordinates[0]}`
    const formattedDistance = station.distance >= 1000 ? (station.distance/1000).toLocaleString('en-GB', { maximumFractionDigits: 1 }) : station.distance.toLocaleString('en-GB', { maximumFractionDigits: 0 });

    const getValue = (val?: number) => val ? `${(Math.round(val * 10) / 10).toFixed(1)} p/L` : 'N/A';
    const isMobile = useMediaQuery("(max-width: 768px)")  

    return (
        <div 
            onMouseEnter={() => onHover && onHover(true)} 
            onMouseLeave={() => onHover && onHover(false)}
            className={cn("space-y-4", { 'hover:bg-neutral-100': !!onHover })} 
        >
            {/* header */}
            <div className="flex justify-between items-center">
                {/* <h3 className="uppercase text-2xl">{station.brand}</h3> */}
                <BrandLogo brand={station.brand} />
                <span className="text-sm m-0 p-0">
                    {formattedDistance}{station.distance >= 1000 ? 'km' : ' metres'} away
                </span>
            </div>

            {/* address */}
            <p className="text-sm">{station.address}{station.address.endsWith(station.postcode) ? '' : `, ${station.postcode}`}</p>

            {/* open in maps */}
            <div className="text-sm text-blue-600 cursor-pointer">
                {isMobile ? (
                    <a href={`geo:${coords}`} target="_blank">Open in map</a>
                ) : (
                    <a href={`https://www.google.com/maps/place/${coords}`} target="_blank">Open in Google Maps</a>
                )}
            </div>            

            {/* icons */}
            <div className='space-y-4'>
                {Object.entries(fuels).map(([fuel, data]) => (
                    <div key={fuel} className="flex gap-4">
                        <div className={cn(
                            `bg-linear-45 ${data.colour}`,
                            "rounded size-12 grid"
                        )}>
                            <FontAwesomeIcon icon={faGasPump} color="white" className="place-self-center" fontSize={20} />
                        </div>
                        <div>
                            <h4 className="text-lg">{data.title} ({fuel.toUpperCase()})</h4>
                            <p className="m-0! p-0 text-sm text-neutral-500">{getValue(station[fuel as keyof Station] as number)} </p>
                        </div>

                    </div>
                ))}
            </div>

            {/* footer */}
            <div className='text-right text-xs text-neutral-400'>
                <p>Last Updated {dayjs.unix(station.created_at).calendar()}</p>
            </div>
        </div>
    );
}

export default StationItem;