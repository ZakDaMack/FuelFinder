
import Stack from '@mui/material/Stack';
import Tooltip from '@mui/material/Tooltip';
import Button from '@mui/material/Button';

import FilterAltIcon from '@mui/icons-material/FilterAlt';
import MyLocationIcon from '@mui/icons-material/MyLocation';
import TuneIcon from '@mui/icons-material/Tune';

import styled from '@emotion/styled';
import IconButton from '@mui/material/IconButton';
import { grey } from '@mui/material/colors';

export default function Toolbar(props) {
    const { recentre } = props;

    return (
        <Stack spacing={2} direction="column" sx={{
            position: 'absolute',
            right: 0, bottom: 0,
            zIndex: 1000, m: 2, pb: 1
        }}>
            <Tooltip title="Centre on me" placement="left">
                <MapButton onClick={recentre}>
                    <MyLocationIcon />
                </MapButton>
            </Tooltip>
            <Tooltip title="Filter stations" placement="left">
                <MapButton>
                    <FilterAltIcon />
                </MapButton>
            </Tooltip>
            <Tooltip title="Tune preferences" placement="left">
                <MapButton>
                    <TuneIcon />
                </MapButton>
            </Tooltip>
        </Stack>
    );
}

const MapButton = styled(IconButton)({
    background: grey[50],
    boxShadow: 'rgba(0, 0, 0, 0.15) 1.95px 1.95px 2.6px',
    '&:hover': {
        background: grey[100]
    }
});