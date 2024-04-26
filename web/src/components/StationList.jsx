import { useState } from 'react';
import { useSelector } from 'react-redux';

import Divider from '@mui/material/Divider';
import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';
import List from '@mui/material/List';
import SwipeableDrawer from '@mui/material/SwipeableDrawer';
import useMediaQuery from '@mui/material/useMediaQuery';
import MenuButton from './MenuButton';

import StationItem from './Station';
import StationListToolbar from './StationListToolbar';
import ListSortChip from './StationListChips';

export default function StationList() {
    const [open, setOpen] = useState(false);
    
    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'));
    const stations = useSelector((state) => state.stations.value);
    const sortKey = useSelector((state) => state.stations.sortKey)

    return (
        <>
            <SwipeableDrawer
                anchor="left"
                open={open}
                onClose={()=>setOpen(false)}
                onOpen={()=>setOpen(true)}
            >
                <Box sx={{position: 'relative', width: isMobile ? '100%' : 400}}>
                    <Box sx={{position: 'absolute', width: '100%', zIndex: 2000}}>
                        <StationListToolbar close={()=>setOpen(false)} />
                        <ListSortChip />
                    </Box>
                    <Box sx={{overflow: 'scroll'}}>
                        <Toolbar />
                        <Toolbar />
                        <List>
                            {stations
                            .filter(s => !!s[sortKey])
                            .sort((a,b) => a[sortKey] - b[sortKey])
                            .map((s, i) => (
                                <>
                                    <Box key={s._id} sx={{py:2, px:1}}>
                                        <StationItem company={s} />
                                    </Box>
                                    {i !== stations.length - 1 && 
                                    (<Divider key={`${s._id}-divider`} />)}
                                </>
                            ))}
                        </List>
                    </Box>
                </Box>
            </SwipeableDrawer>
            <MenuButton click={() => setOpen(!open)}>
            </MenuButton>
        </>
    );
}