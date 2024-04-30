import { useDispatch, useSelector } from 'react-redux';
import { useState } from 'react'

import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';
import Chip from '@mui/material/Chip';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Slider from '@mui/material/Slider';
import TextField from '@mui/material/TextField';
import SwipeableDrawer from '@mui/material/SwipeableDrawer';
import useMediaQuery from '@mui/material/useMediaQuery';

import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';

import ListToolbar from '../ListToolbar';
import { updateRadius, fetchData } from '../../slices/stationSlice';
import { closeAll, updateMenu } from '../../slices/menuSlice';

export default function PreferencesList() {
    const dispatch = useDispatch();

    const radius = useSelector((state) => state.stations.filters.radius);
    const brands = useSelector((state) => state.brands.value);
    const setRadius = (val) => dispatch(updateRadius(val))
    const [activeBrands, setBrands] = useState([
        // 'BP'
    ]);

    const updateVals = (val) => {
        let n = {};
        Object.keys(brands).forEach((k) => n[k] = val)
        setBrands(n)
    }

    const isOpen = useSelector((state) => state.menus.preferences)
    const close = () => dispatch(closeAll());
    const open = () => dispatch(updateMenu('preferences'));
    
    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'))

    const handleClick = () => {
        close()
        navigator.geolocation.getCurrentPosition((pos) => 
            dispatch(fetchData([pos.coords.latitude, pos.coords.longitude]))
        )
    }
    
    return (
        <>
            <SwipeableDrawer
                anchor="right"
                open={isOpen}
                onOpen={open}
                onClose={close}
            >
                <Box sx={{position: 'relative', width: isMobile ? '100%' : 400, height: '100%'}}>
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
                        <Typography component='p' variant='caption' color='error'>(Station filter coming soon)</Typography>
                        <Box sx={{display: 'flex', alignItems: 'baseline'}}>
                            <Button disabled color="info" onClick={() => updateVals(true)}>All</Button>
                            <Typography>|</Typography>
                            <Button disabled color="info" onClick={() => updateVals(false)}>None</Button>
                        </Box>
                        <List sx={{display: 'flex', flexWrap: 'wrap', px: 1}}>
                            {brands.map((b) => (
                                <ListItem key={b} sx={{ m: 0.5, p: 0, width: 'unset' }}>
                                    <Chip 
                                        label={b}
                                        // color={activeBrands.includes(b) ? 'primary' : undefined}
                                        // onClick={() => setBrands([
                                        //     ...activeBrands,
                                        //     ...(activeBrands.includes(b) ? [] : [b])
                                        // ])}
                                    />
                                </ListItem>
                            ))}
                        </List>
                    </Box>
                    <Box sx={{display: 'flex', justifyContent: 'center', p:2}}>
                        <Button variant='contained' color='info' onClick={handleClick}>Update</Button>
                    </Box>
                </Box>
            </SwipeableDrawer>
        </>
    );
}