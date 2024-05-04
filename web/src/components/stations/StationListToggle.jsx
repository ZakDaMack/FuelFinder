import useMediaQuery from '@mui/material/useMediaQuery';

import ClosedDrawer from './ClosedDrawer';
import OpenListButton from './OpenListButton';

export default function ListSortChips() { 
    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'))
    return isMobile ? (<ClosedDrawer />) : (<OpenListButton />)
}