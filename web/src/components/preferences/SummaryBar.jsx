import { useSelector, useDispatch } from 'react-redux';
import { openMenu } from '../../slices/menuSlice';

import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import { grey } from '@mui/material/colors';

import TuneIcon from '@mui/icons-material/Tune';
import InfoIcon from '@mui/icons-material/Info';

export default function SummaryBar() {
    const dispatch = useDispatch();
    const openPrefs = (e) => {
        e.stopPropagation()
        dispatch(openMenu('preferences'))
    }
    
    const openInfo = (e) => {
        e.stopPropagation()
        dispatch(openMenu('info'))
    }

    const filterNos = useSelector((state) => state.stations.filters.brands?.length ?? 0)
    const radius = useSelector((state) => state.stations.filters.radius)

    return (
        <Card raised sx={{
            position: 'absolute',
            top: 0, left: 0,
            display:'flex',
            alignItems: 'center',
            borderRadius: '5rem',
            zIndex: 2000, m: 2, p:1
        }}>
            <IconButton size='large' sx={{background: 'black !important'}} onClick={openPrefs}>
                <TuneIcon sx={{fill: 'white'}} />
            </IconButton>
            <Box sx={{mx: 2}}>
                <Typography variant='body1'>
                    Within {radius} mile{radius === 1 ? '' : 's'}
                </Typography>
                <Typography variant='body2' color={grey[600]}>
                    (+{filterNos} filter{filterNos === 1 ? '' : 's'})
                </Typography>
            </Box>
            <IconButton color='primary' size='small' onClick={openInfo}>
                <InfoIcon />
            </IconButton>
        </Card>
    );
}