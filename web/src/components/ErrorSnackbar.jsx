import { useDispatch, useSelector } from 'react-redux';

import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';
import Snackbar from '@mui/material/Snackbar';
import Button from '@mui/material/Button';

import { fetchData } from '../slices/stationSlice';

export default function ErrorSnackbar() {
    const dispatch = useDispatch()
    const retry = () => dispatch(fetchData())

    const error = useSelector((state) => state.stations.error)

    return (
        <Snackbar
            anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
            open={error != null}
        >
             <Alert severity='error' color='error' action={
                <Button color='error' onClick={retry}>Retry</Button>
             }>
                <AlertTitle>Error</AlertTitle>
                {error}
             </Alert>
        </Snackbar>
    );
}