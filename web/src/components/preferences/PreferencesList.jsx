import { forwardRef, useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import Box from '@mui/material/Box';
import Slide from '@mui/material/Slide';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import Divider from '@mui/material/Divider';
import Typography from '@mui/material/Typography';
import DialogActions from '@mui/material/DialogActions';
import useMediaQuery from '@mui/material/useMediaQuery';

import { fetchData, updateFilters } from '../../slices/stationSlice';
import { closeMenu } from '../../slices/menuSlice';
import DistanceSlider from './DistanceSlider';
import BrandSelectionChips from './BrandSelectionChips';

export default function PreferencesList() {
    const dispatch = useDispatch();

    const isOpen = useSelector((state) => state.menus.preferences)
    const close = () => dispatch(closeMenu("preferences"));
    
    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'))

    const handleClose = () => {
        close()
        dispatch(updateFilters(priorVals))
    }

    const handleUpdate = () => {
        close()
        dispatch(fetchData())
    }

    const currentVals = useSelector((state) => state.stations.filters)
    const [priorVals, setPriorVals] = useState(null)
    useEffect(() => {
        setPriorVals(currentVals)
    }, [isOpen])
    
    return (
        <Dialog
            open={isOpen}
            onClose={handleClose}
            TransitionComponent={Transition}
            className='PreferencesList__dialog'
            PaperProps={isMobile ? {sx: {
                bottom: 0,
                width: '100%',
                position: 'absolute',
                borderRadius: '4px 4px 0 0',
            }}: undefined}
        >
            <Typography component='h2' variant='h6' m={2}>Preferences</Typography>
            <Divider sx={{borderWidth: 1}}></Divider>
        
            <Box m={2} sx={{overflow: 'scroll'}}>
                <DistanceSlider />
                <Box m={3}></Box>
                <BrandSelectionChips />
            </Box>
            <DialogActions sx={{display: 'flex', justifyContent: 'center', p:2}}>
                <Button variant='text' color='info' onClick={handleClose}>Cancel</Button>
                <Box sx={{width: 10}}></Box>
                <Button variant='contained' color='info' onClick={handleUpdate}>Update</Button>
            </DialogActions>
        </Dialog>
    );
}

const Transition = forwardRef(function Transition(props, ref) {
    return <Slide direction="up" ref={ref} {...props} />;
});