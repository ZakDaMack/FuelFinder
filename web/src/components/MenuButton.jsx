import MapButton from './MapButton';

import Tooltip from '@mui/material/Tooltip';
import Box from '@mui/material/Box';

import MenuIcon from '@mui/icons-material/Menu';

export default function MenuButton(props) {
    return (
        <Box sx={{
            position: 'absolute',
            left: 35, top: 0,
            zIndex: 1000, m: 2
        }}>
            <Tooltip title="View stations" placement="right">
                <MapButton onClick={props.click}>
                    <MenuIcon fontSize="large" />
                </MapButton>
            </Tooltip>
        </Box>
    )
}