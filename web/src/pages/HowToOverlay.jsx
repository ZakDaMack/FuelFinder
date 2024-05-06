import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import '../Animate.css'
import Box from '@mui/material/Box';
import Backdrop from '@mui/material/Backdrop';
import Typography from '@mui/material/Typography';

import LocalGasStationIcon from '@mui/icons-material/LocalGasStation';

export default function HowToOverlay() {
    const isLoading = useSelector((state) => state.stations.loading)

    return ( 
        <Backdrop
            sx={{ color: '#fff', zIndex: 5000 }}
            open={isLoading}
        >
            <Box sx={{textAlign: 'center'}}>
                <LocalGasStationIcon sx={{fontSize: 100}} className='animate' />
                <LoadingTicker />
                <Typography variant='h5' px={1}>
                    {puns[random]}
                </Typography>
                <Typography variant='h5' px={1}>
                    {puns[random]}
                </Typography>
            </Box>
        </Backdrop>
    );
}