import { useDispatch, useSelector } from 'react-redux';
import { useState } from 'react'

import Box from '@mui/material/Box';
import Divider from '@mui/material/Divider';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import SwipeableDrawer from '@mui/material/SwipeableDrawer';
import useMediaQuery from '@mui/material/useMediaQuery';

import { fetchData } from '../../slices/stationSlice';
import { closeAll, updateMenu } from '../../slices/menuSlice';
import DistanceSlider from './DistanceSlider';
import BrandSelectionChips from './BrandSelectionChips';

export default function PreferencesList() {
    const dispatch = useDispatch();

    const isOpen = useSelector((state) => state.menus.preferences)
    const close = () => dispatch(closeAll());
    const open = () => dispatch(updateMenu('preferences'));
    
    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'))

    const handleClick = () => {
        close()
        dispatch(fetchData())
    }
    
    return (
        <SwipeableDrawer
            anchor="bottom"
            open={isOpen}
            onOpen={open}
            onClose={close}
            className='BottomSheet__rounded'
        >
            <Typography component='h2' variant='h6' m={2}>Preferences</Typography>
            <Divider sx={{borderWidth: 1}}></Divider>
        
            <Box m={2} sx={{overflow: 'scroll'}}>
                <DistanceSlider />
                <Box m={3}></Box>
                <BrandSelectionChips />
            </Box>
            <Box sx={{display: 'flex', justifyContent: 'center', p:2}}>
                {/* <Button variant='outlined' color='info' onClick={handleClick}>Cancel</Button> */}
                {/* <Box sx={{width: 30}}></Box> */}
                <Button variant='contained' color='info' onClick={handleClick}>Update</Button>
            </Box>
        </SwipeableDrawer>
    );
}