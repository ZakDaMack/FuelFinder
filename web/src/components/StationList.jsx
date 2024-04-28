import { useDispatch, useSelector } from 'react-redux';

import Divider from '@mui/material/Divider';
import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';
import List from '@mui/material/List';
import SwipeableDrawer from '@mui/material/SwipeableDrawer';
import useMediaQuery from '@mui/material/useMediaQuery';

import StationItem from './Station';
import ListToolbar from './ListToolbar';
import ListSortChip from './StationListChips';
import { closeAll, updateMenu } from '../slices/menuSlice';

export default function StationList() {
    const dispatch = useDispatch();

    const isOpen = useSelector((state) => state.menus.stations)
    const close = () => dispatch(closeAll());
    const open = () => dispatch(updateMenu('stations'));
    
    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'))
    const stations = useSelector((state) => state.stations.value)
    const sortKey = useSelector((state) => state.stations.sortKey)

    const filteredStations = stations
        .filter(s => !!s[sortKey])
        .sort((a,b) => a[sortKey] - b[sortKey])

    return (
        <>
            <SwipeableDrawer
                anchor="left"
                open={isOpen}
                onOpen={open}
                onClose={close}
            >
                <Box sx={{position: 'relative', width: isMobile ? '100%' : 400}}>
                    <Box sx={{position: 'absolute', width: '100%', zIndex: 2000}}>
                        <ListToolbar close={close} url="fuelstation.jpg">
                            Stations
                        </ListToolbar>
                        <ListSortChip />
                    </Box>
                    <Box sx={{overflow: 'scroll'}}>
                        <Toolbar />
                        <Toolbar />
                        <List>
                            {filteredStations.length === 0 && (<Box sx={{p:2, textAlign: 'center'}}>
                                There are no stations nearby that match your search
                            </Box>)}
                            {filteredStations
                            .map((s, i) => (
                                <>
                                    <Box key={s._id} sx={{py:2, px:1}}>
                                        <StationItem company={s} />
                                    </Box>
                                    {i !== filteredStations.length - 1 && 
                                    (<Divider key={`${s._id}-divider`} />)}
                                </>
                            ))}
                        </List>
                    </Box>
                </Box>
            </SwipeableDrawer>
        </>
    );
}