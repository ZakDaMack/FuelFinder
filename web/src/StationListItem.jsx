import { useState, useEffect } from 'react';

import { Marker } from 'react-leaflet/Marker' 
import { Popup } from 'react-leaflet/Popup' 

import Typography from '@mui/material/Typography';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import Avatar from '@mui/material/Avatar';

import { blue, yellow, green } from '@mui/material/colors';

import LocalGasStationIcon from '@mui/icons-material/LocalGasStation';
import styled from '@emotion/styled';

export default function StationListItem(props) {
    const { location, company, address, postcode, b7, e5, e10 } = props.company;
    // const coords = [...location.coordinates].reverse();


    const getValue = (val) => val ? `${(Math.round(val * 10) / 10).toFixed(1)} p/L` : 'N/A';

    return (
        <ListItem sx={{flexDirection:'column', alignItems: 'start'}}>
            <Typography variant="h5">{company}</Typography>
        <Typography variant="subtitle">{address}, {postcode}</Typography>
        <List>
            <ListItem disablePadding>
                <ListItemAvatar>
                    <Avatar sx={{ bgcolor: blue[500] }} variant="rounded">
                        <LocalGasStationIcon />
                    </Avatar>
                </ListItemAvatar>
                <PopupText primary="Super (E5)" secondary={getValue(e5)} />
            </ListItem>
            <ListItem disablePadding>
                <ListItemAvatar>
                    <Avatar sx={{ bgcolor: green[500] }} variant="rounded">
                        <LocalGasStationIcon />
                    </Avatar>
                </ListItemAvatar>
                <PopupText primary="Petrol (E10)" secondary={getValue(e10)} />
            </ListItem>
            <ListItem disablePadding>
                <ListItemAvatar>
                    <Avatar sx={{ bgcolor: yellow[600] }} variant="rounded">
                        <LocalGasStationIcon />
                    </Avatar>
                </ListItemAvatar>
                <PopupText primary="Diesel (B7)" secondary={getValue(b7)} />
            </ListItem>
        </List>

        </ListItem>
    );
}

const PopupText = styled(ListItemText)({
    '& .MuiListItemText-secondary': {
        margin: 0
    }
});