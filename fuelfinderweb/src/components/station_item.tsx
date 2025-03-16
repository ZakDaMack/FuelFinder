import { FC } from "react";
import { cn } from "@/lib/utils";
import { useMediaQuery } from "@uidotdev/usehooks";

import { Station } from "@/models/station";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGasPump } from "@fortawesome/free-solid-svg-icons";

import dayjs from 'dayjs';
import calendar from 'dayjs/plugin/calendar';
dayjs.extend(calendar)

interface StationProps {
    station: Station;
}

const fuels = {
    'e5' : {colour: 'from-blue-400 to-blue-900', title: 'Super'},
    'e10':{colour: 'from-green-400 to-green-900', title: 'Petrol'},
    'b7':{colour: 'from-yellow-400 to-yellow-900', title: 'Diesel'},
    // 'sdv':{colour: '', title: 'SDV'},
}

const StationItem: FC<StationProps> = ({ station }) => {
    const formattedDistance = station.distance.toLocaleString('en-GB', { maximumFractionDigits: 0 });
    const getValue = (val?: number) => val ? `${(Math.round(val * 10) / 10).toFixed(1)} p/L` : 'N/A';

    const isMobile = useMediaQuery("(max-width: 768px)")  

    return (
        <div>
            {/* header */}
            <div>
                <h3>{station.brand}</h3>
                <p>{formattedDistance} metres away</p>
            </div>

            {/* address */}
            <div>
                <p>{station.address}, {station.postcode}</p>
                {isMobile ? (
                    <a href={`geo:${station.location.coordinates[1]},${station.location.coordinates[0]}`} target="_blank">Open in map</a>
                ) : (
                    <a href={`https://www.google.com/maps/place/${station.location.coordinates[1]},${station.location.coordinates[0]}`} target="_blank">Open in Google Maps</a>
                )}
            </div>

            {/* icons */}
            <div className='space-y-4'>
                {Object.entries(fuels).map(([fuel, data]) => (
                    <div key={fuel}>
                        <div className={cn(`bg-linear-45 ${data.colour}`,"rounded")}>
                            <FontAwesomeIcon icon={faGasPump} />
                        </div>
                        <div>
                            <h4>{data.title} ({fuel.toUpperCase()})</h4>
                            <p>{getValue(station[fuel as keyof Station] as number)} </p>
                        </div>

                    </div>
                ))}
            </div>

            {/* footer */}
            <div className='text-right'>
                <p>Last Updated {dayjs.unix(station.created_at).calendar()}</p>
            </div>
        </div>
    );
}

export default StationItem;