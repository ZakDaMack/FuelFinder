import { useDispatch, useSelector } from 'react-redux';

import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Chip from '@mui/material/Chip';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';

import LocalGasStationIcon from '@mui/icons-material/LocalGasStation';

import { updateFilters, } from '../../slices/stationSlice';

const boxProps = {
    display: 'flex',
    alignItems: 'center'
}

const types = {
    Super: "e5",
    Petrol: "e10",
    Diesel: "b7",
    SDV: "sdv"
}

export default function FuelTypeSelectionChips() {
    const dispatch = useDispatch();

    const activeFilter = useSelector((state) => state.stations.filters.fueltypes);
    const update = (newArr) => dispatch(updateFilters({ fueltypes: newArr }))

    const toggleType = item => {
        const arr = activeFilter ?? [];
        let newArr = arr.includes(item) ? arr.filter(i => i !== item) : [ ...arr, item ];
        if (types.length === newArr.length) newArr = null
        update(newArr)
    }
    
    return (
        <Box>
            <Box {...boxProps}>
                <LocalGasStationIcon />
                <Typography ml={1} component='h3' variant='h6'>Fuel Types</Typography>
            </Box>
            <Box sx={{display: 'flex', alignItems: 'baseline'}}>
                <Button color="info" onClick={() => update(null)}>All</Button>
                <Typography>|</Typography>
                <Button color="info" onClick={() => update([])}>None</Button>
            </Box>
            <List sx={{display: 'flex', flexWrap: 'wrap',}}>
                {Object.entries(types).map(([key, value]) => {
                    const isActive = activeFilter?.includes(value) ?? true
                    return (
                        <ListItem key={value} sx={{ m: 0.5, p: 0, width: 'unset' }}>
                            <Chip 
                                label={key}
                                sx={{boxShadow: isActive ? 2 : 0}}
                                color={isActive ? 'primary' : undefined}
                                onClick={() => toggleType(value)}
                            />
                        </ListItem>
                    )
                })}
            </List>
        </Box>
    );
}