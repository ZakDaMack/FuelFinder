
import styled from '@emotion/styled';
import IconButton from '@mui/material/IconButton';
import { grey } from '@mui/material/colors';

const MapButton = styled(IconButton)({
    background: grey[50],
    boxShadow: 'rgba(0, 0, 0, 0.15) 1.95px 1.95px 2.6px',
    '&:hover': {
        background: grey[100]
    },
    '.MuiSvgIcon-root': {
        color: 'black'
    }
});

export default MapButton;