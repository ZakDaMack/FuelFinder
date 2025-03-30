import { useEffect } from 'react';
import { useAppDispatch } from '@/store';
import VERSION, { getStoredVersion } from '@/lib/app_version';

import { fetchData } from '@/slices/station_slice';
import { fetchBrands } from '@/slices/brand_slice';
import { openMenu } from '@/slices/menu_slice';

export default function Init() {
    const dispatch = useAppDispatch()
    useEffect(() => {   
        dispatch(fetchBrands())
        dispatch(fetchData())

        // is a new version of the app out?, is it greater then the stored version?
        const v = getStoredVersion() ?? 0
        if (VERSION > v) {
            dispatch(openMenu('info'))
        }
    }, [])

    return null;
}