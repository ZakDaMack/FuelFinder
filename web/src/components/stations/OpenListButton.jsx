import { useDispatch, useSelector } from 'react-redux';

import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import Typography from '@mui/material/Typography';

import ChevronRightIcon from '@mui/icons-material/ChevronRight';

import { openMenu } from '../../slices/menuSlice';

export default function OpenListButton() {
    const dispatch = useDispatch();
    const handleOpen = () => dispatch(openMenu('stations'))
    const stationLen = useSelector((state) => state.stations.value.length)

    return (
        <Card raised onClick={handleOpen} sx={{
            position: 'absolute',
            top: 80, left: 0,
            zIndex: 500, p: 1, m: 2
        }}>
            <Box display='flex'>
                <Typography>List view ({stationLen})</Typography>
                <ChevronRightIcon />
            </Box>
        </Card>
    );
}