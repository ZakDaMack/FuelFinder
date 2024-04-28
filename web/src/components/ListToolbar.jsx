import { useState, cloneElement } from 'react';


import CloseIcon from '@mui/icons-material/Close';

import Toolbar from '@mui/material/Toolbar';
import useScrollTrigger from '@mui/material/useScrollTrigger';
import AppBar from '@mui/material/AppBar';
import IconButton from '@mui/material/IconButton';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';

export default function ListToolbar(props) {
  const { close, url } = props;

  return (
    <ElevationScroll>
      <AppBar position="relative">
        <Toolbar sx={{
          backgroundImage: `linear-gradient(#000a, #000a), url("${url}")`,
          backgroundRepeat: 'no-repeat',
          backgroundPosition: 'center',
          backgroundSize: 'cover'
        }}>
          <Typography variant="h6" component="h2" sx={{ flexGrow: 1, color: 'white' }}>
            {props.children}
          </Typography>
          <Box sx={{flex: 1}}></Box>
          <IconButton onClick={close} sx={{color: 'white'}}>
            <CloseIcon />
          </IconButton>
        </Toolbar>
      </AppBar>
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