import { useDispatch, useSelector } from 'react-redux';
import { useState } from 'react'

import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';
import Chip from '@mui/material/Chip';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import SwipeableDrawer from '@mui/material/SwipeableDrawer';
import useMediaQuery from '@mui/material/useMediaQuery';

import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';

import StationItem from './Station';
import ListToolbar from './ListToolbar';
import { updateRadius } from '../slices/stationSlice';
import { closeAll, updateMenu } from '../slices/menuSlice';
import { Button, IconButton, Typography, Slider, TextField } from '@mui/material';

export default function PreferencesList() {
    const dispatch = useDispatch();

    const radius = useSelector((state) => state.stations.filters.radius);
    const setRadius = (val) => dispatch(updateRadius(val))
    const [brands, setBrands] = useState({
        'BP': true,
        'AppleGreen': true,
        'Tesco': true,
        'Sainsburys': true,
        'Asda': true,
        'Esso': true,
        'Costco': true,
    });

    const updateVals = (val) => {
        let n = {};
        Object.keys(brands).forEach((k) => n[k] = val)
        setBrands(n)
    }

    const isOpen = useSelector((state) => state.menus.preferences)
    const close = () => dispatch(closeAll());
    const open = () => dispatch(updateMenu('preferences'));
    
    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'))
    
    return (
        <>
            <SwipeableDrawer
                anchor="right"
                open={isOpen}
                onOpen={open}
                onClose={close}
            >
                <Box sx={{position: 'relative', width: isMobile ? '100%' : 400}}>
                    <Box sx={{position: 'absolute', width: '100%', zIndex: 2000}}>
                        <ListToolbar close={close} url="/cardash.jpg">
                            Preferences
                        </ListToolbar>
                    </Box>
                    <Box sx={{overflow: 'scroll', p: 2}}>
                        <Toolbar />

                        <Typography component='h3' variant='h6'>Distance (miles)</Typography>
                        <Slider color="info" value={radius} onChange={(e, v) => setRadius(v)} min={1} max={20} sx={{m:1, width: '95%'}} />
                        <Box sx={{display: 'flex'}}>
                            <IconButton onClick={() => setRadius(radius - 1)} disabled={radius <= 1}>
                                <ChevronLeftIcon fontSize='large' />
                            </IconButton>
                            <TextField type="number" fullWidth variant="outlined" value={radius} onChange={(e, v) => setRadius(v)} />
                            <IconButton onClick={() => setRadius(radius + 1)} disabled={radius >= 20}>
                                <ChevronRightIcon fontSize='large' />
                            </IconButton>
                        </Box>

                        <Box sx={{m:8}}></Box>
                        <Typography component='h3' variant='h6'>Stations</Typography>
                        <Box sx={{display: 'flex', alignItems: 'baseline'}}>
                            <Button color="info" onClick={() => updateVals(true)}>All</Button>
                            <Typography>|</Typography>
                            <Button color="info" onClick={() => updateVals(false)}>None</Button>
                        </Box>
                        <List sx={{display: 'flex', flexWrap: 'wrap', px: 1}}>
                            {Object.entries(brands).map((b) => (
                                <ListItem key={b[0]} sx={{ m: 0.5, p: 0, width: 'unset' }}>
                                    <Chip 
                                        label={b[0]}
                                        color={b[1] ? 'primary' : undefined}
                                        onClick={() => setBrands({
                                            ...brands,
                                            [b[0]]: !b[1]
                                        })}
                                    />
                                </ListItem>
                            ))}
                        </List>
                    </Box>
                </Box>
            </SwipeableDrawer>
        </>
    );
}