import { FC, useEffect } from "react";
import { useParams } from "react-router-dom";

import { detailsSelector, fetchStations } from "@/slices/details_slice";
import { useAppDispatch, useAppSelector } from "@/store";

const StationPage: FC = () => {
    const { id } = useParams()
    const dispatch = useAppDispatch()

    useEffect(() => {
        if (id) dispatch(fetchStations(id))
    }, [id])

    const stations = useAppSelector(state => detailsSelector(state, id!))
    
    return (
        <main>
            <h1>Station Page</h1>
            <h2>{id}</h2>

            <pre>{JSON.stringify(stations, null, 2)}</pre>
        </main>
    )
}

export default StationPage