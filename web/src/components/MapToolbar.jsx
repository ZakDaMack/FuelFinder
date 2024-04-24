
import Stack from '@mui/material/Stack';
import Tooltip from '@mui/material/Tooltip';

import FilterAltIcon from '@mui/icons-material/FilterAlt';
import MyLocationIcon from '@mui/icons-material/MyLocation';
import TuneIcon from '@mui/icons-material/Tune';

import MapButton from './MapButton';

export default function MapToolbar() {

    return (
        <Stack spacing={2} direction="column" sx={{
            position: 'absolute',
            right: 0, bottom: 0,
            zIndex: 1000, m: 2, pb: 1
        }}>
            <Tooltip title="Centre on me" placement="left">
                <MapButton>
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