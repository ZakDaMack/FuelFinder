import dayjs from "dayjs";
import { FC, useEffect } from "react";
import { useParams } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "@/store";

import { detailsSelector, fetchStations } from "@/slices/details_slice";

import MapHeader from "./map_header";
import BrandLogo from "@/components/brand_logo";
import PriceHistoryCard from "./price_history_card";

const StationPage: FC = () => {
    const { id } = useParams()
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(fetchStations(id!))
    }, [id])

    const stations = useAppSelector(state => detailsSelector(state, id!))
    
    return (
        <main className="min-h-screen relative">
            <MapHeader coords={[stations[0]?.latitude ?? 0, stations[0]?.longitude ?? 0]} />
            <div className="z-[1000] absolute top-[120px] left-0 right-0 bg-linear-to-b from-transparent to-background to-10%">
                <div className="container mx-auto px-4 my-8 space-y-4">
                    <BrandLogo brand={stations[0]?.brand} className="max-w-[300px] max-h-[100px]" />
                    <h1 className="text-3xl">{stations[0]?.address}, {stations[0]?.postcode}</h1>
                    <p className="text-lg">Prices last updated: {stations[0]?.created_at ? dayjs(stations[0].created_at).calendar() : 'N/A'}</p>

                    {/* <LatestPrices station={stations[0]!} /> */}
                    <PriceHistoryCard stations={stations} />
                </div>
            </div>
            {/* <pre className="text-xs">{JSON.stringify(stations, null, 2)}</pre> */}
        </main>
    )
}

export default StationPage