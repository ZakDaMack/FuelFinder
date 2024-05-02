import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import '../Animate.css'
import Box from '@mui/material/Box';
import Backdrop from '@mui/material/Backdrop';
import Typography from '@mui/material/Typography';

import LocalGasStationIcon from '@mui/icons-material/LocalGasStation';

export default function LoadingOverlay() {
    const puns = [
        "What's petrols favourite rock band? Unleaded Zeppelin!",
        "A car that can go underwater is a Scuba-ru",
        "Filling up the tank",
        "Did you know? A cats favourite car is a Fur-rari",
        "How did Diesel get the job? It performed well under pressure!",
        "Turning on the ignition",
        "Car puns are exhausting...",
        "RV there yet?",
        "H-Audi, partner."
    ]
    const [random, setRandom] = useState(Math.floor(Math.random() * puns.length))

    const isLoading = useSelector((state) => state.stations.loading)

    return ( 
        <Backdrop
            sx={{ color: '#fff', zIndex: 5000 }}
            open={isLoading}
        >
            <Box sx={{textAlign: 'center'}}>
                <LocalGasStationIcon sx={{fontSize: 100}} className='animate' />
                <LoadingTicker />
                <Typography variant='h5'>
                    {puns[random]}
                </Typography>
            </Box>
        </Backdrop>
    );
}

function LoadingTicker() {

    const [tick, setTick] = useState(0)

    useEffect(() => {
        const timer = setInterval(() => {
            setTick(tick >= 3 ? 0 : tick + 1)
            updateDots();
        }, 350);

        return () => clearInterval(timer);
    });
    
    const updateDots = () => {
        const els = document.getElementsByClassName('dot')
        for (let el of els) {
            const i = el.getAttribute('aria-rowindex')
            el.style.color = tick >= parseInt(i) ? 'inherit' : 'transparent'
        }
    }

    return (
        <Typography variant='h2' sx={{
            letterSpacing: 2,
            pb: 4
        }}>
            <span>LOADING</span>
            <span className='dot' aria-rowindex='1'>.</span>
            <span className='dot' aria-rowindex='2'>.</span>
            <span className='dot' aria-rowindex='3'>.</span>
        </Typography>
    )
}