import { useEffect, useRef, useState } from 'react';

import Button from '@mui/material/Button';
import Divider from '@mui/material/Divider';
import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';
import List from '@mui/material/List';
import SwipeableDrawer from '@mui/material/SwipeableDrawer';
import useMediaQuery from '@mui/material/useMediaQuery';

import ChevronRightIcon from '@mui/icons-material/ChevronRight';

import styled from '@emotion/styled';
import { grey } from '@mui/material/colors';

import StationItem from './StationItem';
import StationListToolbar from './StationListToolbar';

export default function StationList(props) {
    const { stations } = props;
    const [ sortKey, setSortKey ] = useState('e5');
    const [open, setOpen] = useState(false);
    
    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'));

    const [val, setVal] = useState({});
    useEffect(() => {
        fetchData();
    }, [])

    const fetchData = async () => {
        const res = await fetch('http://localhost:3001/gcn7sj2hxk2c');
        const data = await res.json();
        setVal({
            _id: data[0]._id,
            site_id: data[0].site_id,
            company: data[0].company,
            address: data[0].address,
            postcode: data[0].postcode,
            history: data.map(s => ({
                date: s.created_at,
                sdv: s.sdv,
                b7: s.b7,
                e5: s.e5,
                e10: s.e10
            }))
        });
    };

    return (
        <>
            <SwipeableDrawer
                anchor="left"
                open={open}
                onClose={()=>setOpen(false)}
                onOpen={()=>setOpen(true)}
            >
                <Box sx={{position: 'relative', width: isMobile ? '100%' : 400}}>
                    <StationListToolbar sortKey={sortKey} setSortKey={setSortKey} close={()=>setOpen(false)} />
                    <Toolbar />
                    {/* <Box sx={{height: 'calc(100vh - 64px)', overflow: 'scroll'}}>
                        <List>
                            {stations
                            .filter(s => !!s[sortKey])
                            .sort((a,b) => a[sortKey] - b[sortKey])
                            .map((s, i) => (
                                <>
                                    <Box key={s._id} sx={{py:2, px:1}}>
                                        <StationItem company={s} />
                                    </Box>
                                    {i !== stations.length - 1 && <Divider key={`${s._id}-divider`} />}
                                </>
                            ))}
                        </List>
                    </Box> */}
                    <pre>
                        {JSON.stringify(val, null, 4)}
                    </pre>
                </Box>
            </SwipeableDrawer>
            <OpenButton onClick={()=>setOpen(true)}>
                <ChevronRightIcon />
            </OpenButton>
        </>
    );
}

const OpenButton = styled(Button)({
    position: 'absolute',
    left: 0, top: '50%',
    zIndex: 1000,
    transform: 'translateY(-50%)',
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