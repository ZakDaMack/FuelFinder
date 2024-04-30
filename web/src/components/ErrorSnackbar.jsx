import { useSelector } from 'react-redux';

import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';
import Snackbar from '@mui/material/Snackbar';

export default function ErrorSnackbar() {
    const error = useSelector((state) => state.stations.error)

    return (
        <Snackbar
            anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
            open={error != null}
        >
             <Alert severity='error' color='error'>
                <AlertTitle>Error</AlertTitle>
                {error}
             </Alert>
        </Snackbar>
    );
}