import { useMap } from 'react-leaflet/hooks';

import Stack from '@mui/material/Stack';
import Card from '@mui/material/Card';
import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';

import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';



export default function MapZoom() {
    const map = useMap();
    const zoomIn = () => map.zoomIn()
    const zoomOut = () => map.zoomOut()

    return (
        <Card raised sx={{
            position: 'absolute',
            top: 0, right: 0,
            display:'flex',
            zIndex: 1000, m: 2
        }}>
            <Stack>
                <IconButton size='small' onClick={zoomIn}>
                    <KeyboardArrowUpIcon />
                </IconButton>
                <Divider />
                <IconButton size='small' onClick={zoomOut}>
                    <KeyboardArrowDownIcon />
                </IconButton>
            </Stack>
        </Card>
    );
}