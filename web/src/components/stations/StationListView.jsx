import { useSelector, useDispatch } from 'react-redux';
import { useEffect, useRef, useState } from 'react';

import Box from '@mui/material/Box';
import AppBar from '@mui/material/AppBar';
import List from '@mui/material/List';
import Paper from '@mui/material/Paper';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import Divider from '@mui/material/Divider';
import useMediaQuery from '@mui/material/useMediaQuery';
import SwipeableDrawer from '@mui/material/SwipeableDrawer';

import CloseIcon from '@mui/icons-material/Close';

import StationItem from '../Station';
import StationListChips from './StationListChips';
import { closeAll, openMenu } from '../../slices/menuSlice';

export default function StationListView() {
    const dispatch = useDispatch();

    const [height, setHeight] = useState(0)
    const chipToolbar = useRef(null)

    const isOpen = useSelector((state) => state.menus.stations)
    const close = () => dispatch(closeAll());
    const open = () => dispatch(openMenu('stations'));

    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'))
    const stations = useSelector((state) => state.stations.value)
    const sortKey = useSelector((state) => state.stations.sortKey)

    const filteredStations = stations
        .filter(s => !!s[sortKey])
        .sort((a,b) => (a[sortKey] - b[sortKey]) || (a.distance - b.distance)) // sort by specified key, if equal, sort by dist

    useEffect(() => {
        setHeight(chipToolbar.current?.offsetHeight ?? 0)
    }, [chipToolbar])

    return (
        <SwipeableDrawer
            anchor={isMobile ? 'bottom' : 'left'}
            open={isOpen}
            onOpen={open}
            onClose={close}
        >
            <Paper elevation={0} sx={{
                position: 'relative', 
                height: '100vh',
                width: isMobile ? '100%' : 400    
            }}>
                <AppBar position='fixed' sx={{
                    left: 0, bgcolor: 'white',
                    width: isMobile ? '100%' : 400
                }}>
                    {/* close button */}
                    <Toolbar sx={{
                        px: 2, pt: 3, pb: 1, display: 'flex',
                        justifyContent: 'right'
                    }}>
                        <IconButton size='large' onClick={close} sx={{
                            border: 'grey solid 1px'
                        }}>
                            <CloseIcon />
                        </IconButton>
                    </Toolbar>
                    {/* sort buttons */}
                    <Toolbar ref={chipToolbar}>
                        <Typography mr={1}>Sort</Typography>
                        <StationListChips />
                    </Toolbar>
                </AppBar>
                <List>
                    <Toolbar sx={{p: 5}}></Toolbar>
                    <Toolbar sx={{height: height}} />
                    {filteredStations.length === 0 && (
                        <Typography p={2} mt={4} textAlign='center'>
                            There are no stations nearby that match your search
                        </Typography>
                    )}
                    {filteredStations.map((s, i) => (
                        <>
                            <Box component='li' key={s._id} sx={{py:2, px:1}}>
                                <StationItem company={s} />
                            </Box>    
                            {i !== filteredStations.length - 1 && 
                            (<Divider sx={{borderWidth: 2}} key={`${s._id}-divider`} />)}
                        </>      
                    ))}
                </List>    
            </Paper>
        </SwipeableDrawer>
    );
}