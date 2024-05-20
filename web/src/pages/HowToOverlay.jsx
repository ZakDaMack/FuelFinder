import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import '../Animate.css'
import Box from '@mui/material/Box';
import Backdrop from '@mui/material/Backdrop';
import Typography from '@mui/material/Typography';

import { Button, Checkbox, FormControlLabel  } from '@mui/material';
import Arrow1 from '../components/arrows/Arrow1';

export default function HowToOverlay() {
    const [open, setOpen] = useState(true)
    const [dontShow, setDontShow] = useState(false)

    const close = () => {
        setOpen(false)
    }

    return ( 
        <Backdrop
            sx={{color: '#fff', zIndex: 5000}}
            open={open}
        >
            <Box width='100%' height='100%'>
                <Box sx={{position: 'absolute', bottom: 100, left: '10%'}}>
                    <Typography sx={{maxWidth: '16em'}}>Drag up or tap below to view the fuel stations nearest to you</Typography>
                    <Arrow1></Arrow1>
                </Box>

                {/* close box */}
                <Box sx={{display: 'block', color: 'white'}}>
                    <FormControlLabel label="Don't show me again" control={
                        <Checkbox value={dontShow} onChange={(e, c) =>setDontShow(c)} />
                    } />
                    <Button variant='contained' onClick={close}>Ok, got it</Button>
                </Box>
            </Box>
        </Backdrop>
    );
}