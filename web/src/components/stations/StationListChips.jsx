import { useDispatch, useSelector } from 'react-redux';

import Chip from '@mui/material/Chip';
import Box from '@mui/material/Box';
import ListItem from '@mui/material/ListItem';

import { updateSort } from '../../slices/stationSlice';

export default function ListSortChips() { 
    const dispatch = useDispatch()
    const sortList = [
        {text: 'Distance', value: 'distance'},
        {text: 'Super', value: 'e5'},
        {text: 'Petrol', value: 'e10'},
        {text: 'Diesel', value: 'b7'},
    ];

    const sortKey = useSelector((state) => state.stations.sortKey)
    const handleClick = (key) => dispatch(updateSort(key))

    return (
        <Box display='flex' flexWrap='wrap'>
            {sortList.map((data) => (
                <ListItem 
                    key={data.value} 
                    sx={{ m: 0.5, p: 0, width: 'unset' }}
                >
                    <Chip
                        label={data.text}
                        color={sortKey === data.value ? 'primary' : undefined}
                        onClick={() => handleClick(data.value)}
                    />
                </ListItem>
            ))}
        </Box>
    );
}