import { useDispatch, useSelector } from 'react-redux';

import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Chip from '@mui/material/Chip';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';

import LocalGasStationIcon from '@mui/icons-material/LocalGasStation';

import { updateFilters, } from '../../slices/stationSlice';

export default function BrandSelectionChips() {
    const dispatch = useDispatch();

    const brands = useSelector((state) => state.brands.value);
    const activeFilter = useSelector((state) => state.stations.filters.brands);
    
    const update = (newArr) => dispatch(updateFilters({ brands: newArr }))
    const toggleBrand = item => {
        const arr = activeFilter ?? [];
        let newArr = arr.includes(item) ? arr.filter(i => i !== item) : [ ...arr, item ];
        if (brands.length == newArr.length) newArr = null
        console.log(newArr);
        update(newArr)
    }

    const boxProps = {
        display: 'flex',
        alignItems: 'center'
    }
    
    return (
        <Box>
            <Box {...boxProps}>
                <LocalGasStationIcon />
                <Typography ml={1} component='h3' variant='h6'>Stations</Typography>
            </Box>
            <Box sx={{display: 'flex', alignItems: 'baseline'}}>
                <Button color="info" onClick={() => update(null)}>All</Button>
                <Typography>|</Typography>
                <Button color="info" onClick={() => update([])}>None</Button>
            </Box>
            <List sx={{display: 'flex', flexWrap: 'wrap',}}>
                {brands.map((b) => (
                    <ListItem key={b} sx={{ m: 0.5, p: 0, width: 'unset' }}>
                        <Chip 
                            label={b}
                            color={(activeFilter?.includes(b) ?? true) ? 'primary' : undefined}
                            onClick={() => toggleBrand(b)}
                        />
                    </ListItem>
                ))}
            </List>
        </Box>
    );
}