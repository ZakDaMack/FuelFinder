import { useEffect } from 'react';
import { useDispatch } from 'react-redux';

import { fetchData } from '../slices/stationSlice';
import { fetchBrands } from '../slices/brandSlice';

export default function Init() {
    const dispatch = useDispatch()
    
    useEffect(() => {   
        dispatch(fetchBrands())
        navigator.geolocation.getCurrentPosition((pos) => 
            dispatch(fetchData([pos.coords.latitude, pos.coords.longitude]))
        )
    }, [])

    return null;
}