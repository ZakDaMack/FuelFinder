import { useDispatch } from 'react-redux';

import Stack from '@mui/material/Stack';
import Tooltip from '@mui/material/Tooltip';

import TuneIcon from '@mui/icons-material/Tune';
import MyLocationIcon from '@mui/icons-material/MyLocation';
import MenuIcon from '@mui/icons-material/List';

import MapButton from './MapButton';
import { updateMenu } from '../../slices/menuSlice';
import { fetchData } from '../../slices/stationSlice';

export default function MapToolbar() {
    const dispatch = useDispatch();

    const openStationList = () => dispatch(updateMenu('stations'))
    const openPreferencesList = () => dispatch(updateMenu('preferences'))

    const handleRefresh = () => {
        navigator.geolocation.getCurrentPosition((pos) => 
            dispatch(fetchData([pos.coords.latitude, pos.coords.longitude]))
        )
    }

    return (
        <Stack spacing={2} direction="column" sx={{
            position: 'absolute',
            right: 0, bottom: 0,
            zIndex: 1000, m: 2, pb: 1
        }}>
            <Tooltip title="Centre on me" placement="left">
                <MapButton >
                    <MyLocationIcon onClick={handleRefresh}/>
                </MapButton>
            </Tooltip>
            <Tooltip title="Filter stations" placement="left">
                <MapButton onClick={openPreferencesList}>
                    <TuneIcon />
                </MapButton>
            </Tooltip>
            <Tooltip title="View list" placement="left">
                <MapButton onClick={openStationList}>
                    <MenuIcon />
                </MapButton>
            </Tooltip>
        </Stack>
    );
}