import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';

import TuneIcon from '@mui/icons-material/Tune';

import { useSelector, useDispatch } from 'react-redux';
import { closeAll, updateMenu } from '../../slices/menuSlice';
import { grey } from '@mui/material/colors';

export default function SummaryBar() {
    const dispatch = useDispatch();
    const open = () => dispatch(updateMenu('preferences'));

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
            <IconButton size='large' sx={{bgcolor: 'black'}} onClick={open}>
                <TuneIcon sx={{color: 'white'}} />
            </IconButton>
            <Box sx={{mx: 2}}>
                <Typography variant='body1'>
                    Within {radius} mile{radius == 1 ? '' : 's'}
                </Typography>
                <Typography variant='body2' color={grey[600]}>
                    (+{filterNos} filter{filterNos == 1 ? '' : 's'})
                </Typography>
            </Box>
        </Card>
    );
}