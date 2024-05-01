import { useDispatch, useSelector } from 'react-redux';
import { useEffect } from 'react';

import Stack from '@mui/material/Stack';
import Tooltip from '@mui/material/Tooltip';
import useMediaQuery from '@mui/material/useMediaQuery';
import ClickAwayListener from '@mui/material/ClickAwayListener';

import TuneIcon from '@mui/icons-material/Tune';
import MyLocationIcon from '@mui/icons-material/MyLocation';
import MenuIcon from '@mui/icons-material/List';

import MapButton from './MapButton';
import { closeAll, updateMenu } from '../../slices/menuSlice';
import { fetchData } from '../../slices/stationSlice';


export default function MapToolbar() {
    const dispatch = useDispatch();

    const handleRefresh = () => dispatch(fetchData())

    const close = () => dispatch(closeAll())
    const openStationList = () => dispatch(updateMenu('stations'))
    const openPreferencesList = () => dispatch(updateMenu('preferences'))

    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'))

    // on first load, auto show tooltips so user knows what they do
    useEffect(() => {
        dispatch(updateMenu('tooltips'))
    }, [])

    return (
        <ClickAwayListener onClickAway={close}>
            <Stack spacing={2} direction="column" sx={{
                position: 'absolute',
                right: 0, bottom: isMobile ? '5%' : 0,
                zIndex: 1000, m: 2, pb: 1
            }}>
                <FixableTooltip title="Centre on me">
                    <MapButton>
                        <MyLocationIcon onClick={handleRefresh}/>
                    </MapButton>
                </FixableTooltip>
                <FixableTooltip title="Filter stations">
                    <MapButton onClick={openPreferencesList}>
                        <TuneIcon />
                    </MapButton>
                </FixableTooltip>
                <FixableTooltip title="View station list">
                    <MapButton onClick={openStationList}>
                        <MenuIcon />
                    </MapButton>
                </FixableTooltip>
            </Stack>   
        </ClickAwayListener>
    );
}

function FixableTooltip(props) {
    const dispatch = useDispatch();

    const tooltipsOn = useSelector((state) => state.menus.tooltips && !state.stations.loading)
    const close = () => dispatch(closeAll())

    return (
        <Tooltip open={tooltipsOn} onClose={close} title={props.title} placement="left" enterDelay={500} leaveDelay={200}>
            <Tooltip title={props.title}  placement="left">
                {props.children}
            </Tooltip>
        </Tooltip>
    )
}