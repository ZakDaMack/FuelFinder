
import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';

import FilterAltIcon from '@mui/icons-material/FilterAlt';
import TuneIcon from '@mui/icons-material/Tune';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';

import styled from '@emotion/styled';
import { grey } from '@mui/material/colors';
import { Box, Divider, Toolbar } from '@mui/material';
import List from '@mui/material/List';
import { useEffect, useRef, useState } from 'react';
import StationListItem from './StationListItem';
import StationListToolbar from './StationListToolbar';

export default function StationList(props) {
    const { stations } = props;
    const [ sortKey, setSortKey ] = useState('e5');

    const [open, setOpen] = useState(false);
    const toggleOpen = () => setOpen(!open);
    const drawerRef = useRef(null);

    useEffect(() => {
        const size = open ? '400' : '0';
        drawerRef.current.style.width = `${size}px`
    }, [open]);

    return (
        <Box ref={drawerRef} sx={{
            height: '100vh',
            position: 'absolute',
            left: 0, top: 0,
            zIndex: 2000,
            display: 'flex',
            alignItems: 'center',
            transition: 'width 0.2s ease-in'
        }}>
            <Box sx={{
                bgcolor: 'white', 
                overflow: 'scroll',
                position: 'relative',
                width: 400, 
                height: '100%'
            }}>
                <StationListToolbar />
                <Toolbar />
                <List>
                    {stations
                    .filter(s => !!s[sortKey])
                    .sort((a,b) => a[sortKey] - b[sortKey])
                    .map((s, i) => (
                        <>
                            <StationListItem key={s._id} company={s} />
                            {i !== stations.length - 1 && <Divider key={`${s._id}-divider`} />}
                        </>
                    ))}
                </List>
            </Box>
            <OpenButton onClick={toggleOpen}>
                {open ? (<ChevronLeftIcon />) : (<ChevronRightIcon />)}
            </OpenButton>
        </Box>
    );
}

const OpenButton = styled(Button)({
    height: 70,
    width: 40,
    minWidth: 40,
    borderTopLeftRadius: 0,
    borderBottomLeftRadius: 0,
    boxShadow: 'rgba(0, 0, 0, 0.15) 1.95px 1.95px 2.6px',
    background: 'white',
    '&:hover': {
        background: grey[100]
    },
    svg: {
        color: 'black'
    }
});