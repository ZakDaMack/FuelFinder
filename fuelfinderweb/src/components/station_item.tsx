import { FC } from "react";
import { Station } from "@/models/station";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGasPump } from "@fortawesome/free-solid-svg-icons";

interface StationProps {
    station: Station;
}

const StationItem: FC<StationProps> = ({ station }) => {
    return (
<div>
    <div>
        
    </div>
</div>
    );
}

export default StationItem;

import styled from '@emotion/styled';

import Box from '@mui/material/Box';
import List from '@mui/material/List';
import Avatar from '@mui/material/Avatar';
import ListItem from '@mui/material/ListItem';
import Typography from '@mui/material/Typography';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import useMediaQuery from '@mui/material/useMediaQuery';

import { grey } from '@mui/material/colors';

import LocalGasStationIcon from '@mui/icons-material/LocalGasStation';

import dayjs from 'dayjs';
import calendar from 'dayjs/plugin/calendar';
dayjs.extend(calendar)

export default function StationItem(props) {
    const { distance, brand, location, address, postcode, b7, e5, e10, created_at } = props.company;

    const isMobile = useMediaQuery((theme) => theme.breakpoints.down('sm'))

    const formattedDistance = distance.toLocaleString('en-GB', { maximumFractionDigits: 0 });
    const getValue = (val) => val ? `${(Math.round(val * 10) / 10).toFixed(1)} p/L` : 'N/A';

    return (
        <Box>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'end', pb: 1 }}>
                <Typography variant="h5">{brand}</Typography>
                <Typography variant="subtitle" sx={{ color: 'green', pb: 0.5, textAlign: 'right' }}>
                    {formattedDistance} metres away
                </Typography>
            </Box>
            <Typography variant="subtitle">{address}, {postcode}</Typography>
            <Typography my={1} component='div' variant="body2">
                {isMobile ? (
                    <a href={`geo:${location.coordinates[1]},${location.coordinates[0]}`} target="_blank">Open in map</a>
                ) : (
                    <a href={`https://www.google.com/maps/place/${location.coordinates[1]},${location.coordinates[0]}`} target="_blank">Open in Google Maps</a>
                )}
            </Typography>
            <List>
                <ListItem disablePadding>
                    <ListItemAvatar>
                        <Avatar sx={{ 
                            background: 'linear-gradient(45deg, rgba(33,150,243,1) 0%, rgba(69,111,145,1) 100%)'
                        }} variant="rounded">
                            <LocalGasStationIcon />
                        </Avatar>
                    </ListItemAvatar>
                    <PopupText primary="Super (E5)" secondary={getValue(e5)} />
                </ListItem>
                <ListItem disablePadding>
                    <ListItemAvatar>
                        <Avatar sx={{
                            background: 'linear-gradient(45deg, rgba(76,175,80,1) 0%, rgba(45,94,47,1) 100%)'
                        }} variant="rounded">
                            <LocalGasStationIcon />
                        </Avatar>
                    </ListItemAvatar>
                    <PopupText primary="Petrol (E10)" secondary={getValue(e10)} />
                </ListItem>
                <ListItem disablePadding>
                    <ListItemAvatar>
                        <Avatar sx={{ 
                            background: 'linear-gradient(45deg, rgba(253,216,53,1) 0%, rgba(186,166,79,1) 100%)'
                        }} variant="rounded">
                            <FontAwesomeIcon icon={faGasPump} />
                        </Avatar>
                    </ListItemAvatar>
                    <PopupText primary="Diesel (B7)" secondary={getValue(b7)} />
                </ListItem>
            </List>
            <Box sx={{textAlign: 'right'}}>
                <Typography variant="subtitle2" sx={{fontSize: 10, color: grey[600]}}>
                    Last Updated {dayjs.unix(created_at).calendar()}
                </Typography>
            </Box>
        </Box>
    );
}

const PopupText = styled(ListItemText)({
    '& .MuiListItemText-secondary': {
        margin: 0
    }
});