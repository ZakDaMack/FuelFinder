import { useDispatch, useSelector } from 'react-redux';
import { useState } from 'react'

import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Chip from '@mui/material/Chip';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';

import LocalGasStationIcon from '@mui/icons-material/LocalGasStation';

import { updateRadius } from '../../slices/stationSlice';

export default function BrandSelectionChips() {
    const dispatch = useDispatch();

    // const brands = useSelector((state) => state.brands.value);
    const brands = [];
    const [activeBrands, setBrands] = useState([
        // 'BP'
    ]);

    const updateVals = (val) => {
        let n = {};
        Object.keys(brands).forEach((k) => n[k] = val)
        setBrands(n)
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
            <Typography component='p' variant='caption' color='error'>(Station filter coming soon)</Typography>
            <Box sx={{display: 'flex', alignItems: 'baseline'}}>
                <Button disabled color="info" onClick={() => updateVals(true)}>All</Button>
                <Typography>|</Typography>
                <Button disabled color="info" onClick={() => updateVals(false)}>None</Button>
            </Box>
            <List sx={{display: 'flex', flexWrap: 'wrap',}}>
                {brands.map((b) => (
                    <ListItem key={b} sx={{ m: 0.5, p: 0, width: 'unset' }}>
                        <Chip 
                            label={b}
                            // color={activeBrands.includes(b) ? 'primary' : undefined}
                            // onClick={() => setBrands([
                            //     ...activeBrands,
                            //     ...(activeBrands.includes(b) ? [] : [b])
                            // ])}
                        />
                    </ListItem>
                ))}
            </List>
        </Box>
    );
}