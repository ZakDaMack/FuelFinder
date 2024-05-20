import { useDispatch, useSelector } from 'react-redux';

import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Slider from '@mui/material/Slider';

import RoomIcon from '@mui/icons-material/Room';

import { updateRadius } from '../../slices/stationSlice';

export default function DistanceSlider() {
    const dispatch = useDispatch();

    const radius = useSelector((state) => state.stations.filters.radius);
    const setRadius = (val) => dispatch(updateRadius(val))

    const boxProps = {
        display: 'flex',
        alignItems: 'center'
    }
    
    return (
        <Box>
            <Box {...boxProps}>
                <RoomIcon />
                <Typography ml={1} component='h3' variant='h6'>Distance (miles)</Typography>
            </Box>
            <Box {...boxProps}>
                <Slider value={radius} onChange={(e, v) => setRadius(v)} min={1} max={20} sx={{m:1, width: '95%'}} />
                <Typography ml={2}>{radius}</Typography>
            </Box>
        </Box>
    );
}