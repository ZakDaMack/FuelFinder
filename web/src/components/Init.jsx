import { useEffect } from 'react';
import { useDispatch } from 'react-redux';

import VERSION, { GetStoredVersion } from '../lib/getAppVersion';
import { fetchData } from '../slices/stationSlice';
import { fetchBrands } from '../slices/brandSlice';
import { openMenu } from '../slices/menuSlice';

export default function Init() {
    const dispatch = useDispatch()
    useEffect(() => {   
        dispatch(fetchBrands())
        dispatch(fetchData())

        // is a new version of the app out?, is it greater then the stored veresion?
        if (VERSION > GetStoredVersion()) {
            dispatch(openMenu('info'))
        }

    }, [])

    return null;
}