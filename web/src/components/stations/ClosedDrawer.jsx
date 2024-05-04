import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import Typography from '@mui/material/Typography';

import { grey } from '@mui/material/colors';
import { openMenu } from '../../slices/menuSlice';

const SWIPE_DISTANCE = 50;

export default function ClosedDrawer() {
    const dispatch = useDispatch();
    const handleOpen = () => dispatch(openMenu('stations'))
    const stationLen = useSelector((state) => state.stations.value.length)

    const [touchStart, setTouchStart] = useState(null)

    const onTouchStart = (e) => setTouchStart(e.targetTouches[0].clientY)
    const onTouchMove = (e) => {
        const touchEnd = e.targetTouches[0].clientY;
        if ((touchStart - touchEnd) < SWIPE_DISTANCE) return
        handleOpen()
    }

    return (
        <Card onClick={handleOpen} onTouchStart={onTouchStart} onTouchMove={onTouchMove}
            raised className='BottomSheet__rounded' sx={{
            position: 'absolute',
            bottom: 0, width: '100%',
            zIndex: 500, p:1
        }}>
            <Box sx={{
                bgcolor: grey[300],
                borderRadius: '5rem',
                margin: '0 auto',
                height: 10,
                width: 100
            }}></Box> 
            <Box mt={2} mb={6} textAlign='center'>
                <Typography>
                    {stationLen} station{stationLen === 1 ? '' : 's'} nearby
                </Typography>
            </Box>
        </Card>
    );
}