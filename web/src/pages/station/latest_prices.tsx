import { FC } from "react";
import { cn } from "@/lib/utils";

import { Station } from "@/models/station";
import { faGasPump } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const fuels = {
    'e5'  : {colour: 'from-blue-400 to-blue-700', title: 'Super'},
    'e10' : {colour: 'from-green-400 to-green-700', title: 'Petrol'},
    'b7'  : {colour: 'from-yellow-400 to-yellow-700', title: 'Diesel'},
    'sdv' : {colour: 'from-orange-400 to-orange-700', title: 'Super Diesel'},
}

const LatestPrices: FC<{station?: Station}> = ({ station }) => {
    const getValue = (val?: number) => val ? `${(Math.round(val * 10) / 10).toFixed(1)} p/L` : 'N/A';

    if (!station) {
        return (<></>)
    }

    return (
        <div className='space-y-4'>
            {Object.entries(fuels).map(([fuel, data]) => (
                <div key={fuel} className="flex items-center gap-4">
                    <div className={cn(
                        `bg-linear-45 ${data.colour}`,
                        "rounded-lg size-11 grid"
                    )}>
                        <FontAwesomeIcon icon={faGasPump} color="white" className="place-self-center" fontSize={20} />
                    </div>
                    <div>
                        <h4 className="text-base">{data.title} ({fuel.toUpperCase()})</h4>
                        <p className="m-0! p-0 text-sm text-neutral-500">{getValue(station[fuel as keyof Station] as number)} </p>
                    </div>

                </div>
            ))}
        </div>
    )
}

export default LatestPrices;