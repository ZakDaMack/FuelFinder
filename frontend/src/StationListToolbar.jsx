import { useEffect, useRef, useState, cloneElement } from 'react';

import Toolbar from '@mui/material/Toolbar';

import FilterAltIcon from '@mui/icons-material/FilterAlt';
import TuneIcon from '@mui/icons-material/Tune';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';

import Button from '@mui/material/Button';
import useScrollTrigger from '@mui/material/useScrollTrigger';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';

export default function StationListToolbar(props) {

  const [anchorEl, setAnchorEl] = useState(null);
  const open = Boolean(anchorEl);

  const handleClick = (event) => setAnchorEl(event.currentTarget);
  const handleClose = () => setAnchorEl(null);

  const sortList = [
    {text: 'Distance', value: 'distance'},
    {text: 'Super', value: 'e5'},
    {text: 'Petrol', value: 'e10'},
    {text: 'Diesel', value: 'b7'},
  ];

  return (
    <ElevationScroll>
        <Toolbar sx={{position: 'absolute', zIndex: 2000, bgcolor: 'white', w: '100%'}}>
            <Button variant="contained"
                startIcon={<FilterAltIcon />}
                endIcon={<ExpandMoreIcon />}
                id="menu-button"
                aria-controls={open ? 'menu' : undefined}
                aria-haspopup="true"
                aria-expanded={open ? 'true' : undefined}
                onClick={handleClick}
                sx={{textTransform: 'none', letterSpacing: 'normal'}}
            >
                { sortList[0].text }
            </Button>
            <Menu
                id="basic-menu"
                anchorEl={anchorEl}
                open={open}
                onClose={handleClose}
                MenuListProps={{
                    'aria-labelledby': 'menu-button',
                }}
            >
                {sortList.map(s => (<MenuItem key={s.value} onClick={handleClose}>{s.text}</MenuItem>))}
            </Menu>
        </Toolbar>
    </ElevationScroll>
  );
}

function ElevationScroll(props) {
    const { children } = props;
    // Note that you normally won't need to set the window ref as useScrollTrigger
    // will default to window.
    // This is only being set here because the demo is in an iframe.
    const trigger = useScrollTrigger({
      disableHysteresis: true,
      threshold: 0
    });
  
    return cloneElement(children, {
      sx: {
        borderBottom: trigger ? '1px solid black' : undefined
      }
    });
  }