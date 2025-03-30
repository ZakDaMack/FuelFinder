import { FC } from "react";
import { openMenu } from "@/slices/menu_slice";
import { useAppDispatch, useAppSelector } from "@/store";
import { stationsSelector } from "@/slices/station_slice";

import { Button } from "./ui/button";

const StationSummary: FC = () => {
    const dispatch = useAppDispatch()
    const handleClick = () => dispatch(openMenu('stations'))
    
    const stations = useAppSelector(stationsSelector)

    return (
        <Button 
            variant='secondary'
            className="hidden md:block absolute left-5 top-24 shadow-lg z-1000"
            onClick={handleClick}
        >
            View Stations ({stations.length})
        </Button>
    )
}

export default StationSummary;