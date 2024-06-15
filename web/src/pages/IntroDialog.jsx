import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { closeMenu } from '../slices/menuSlice';
import { StoreVersion } from '../lib/getAppVersion';

import Box from '@mui/material/Box';
import Dialog from '@mui/material/Dialog';
import Button from '@mui/material/Button';
import Checkbox from '@mui/material/Checkbox';
import Typography from '@mui/material/Typography';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import FormControlLabel from '@mui/material/FormControlLabel';


export default function IntroDialog() {
    const dispatch = useDispatch();
    const isOpen = useSelector((state) => state.menus.info)

    const changes = [
        "Added the ability to filter by fuel type"
    ];

    const [dontShow, setDontShow] = useState(false)

    const close = () => {
        if (dontShow) StoreVersion()
        dispatch(closeMenu('info'))
    }

    return ( 
        <Dialog transitionDuration={{enter: 0, exit: 500}} open={isOpen} sx={{
            zIndex: 10000,
            '& .MuiDialog-paper': {
                margin: 2,
                maxHeight: 'calc(100% - 32px)'
            }
        }}>
            <DialogContent sx={{
                '& li:not(:last-child)': { pb: 0.5 },
                '& > .MuiTypography-root:not(:first-child)': {
                    pt: 3
                }
            }}>
                <Typography component='h2' variant='h4' sx={(theme) => ({
                    '& > span': {
                        color: theme.palette.primary.main,
                        fontWeight: 600,
                        background: `linear-gradient(45deg, ${theme.palette.primary.main} 40%, ${theme.palette.primary.dark})`,
                        backgroundClip: 'text',
                        textFillColor: 'transparent',
                    }
                })}>
                    Welcome to <span>FuelFinder</span>
                </Typography>

                <Typography align='justify'>
                    FuelFinder is an open source project with the aim of saving you money.
                    The app will show you the cheapest fuel prices in your area by using data provided
                    by brands participating in the <Link href="https://www.gov.uk/guidance/access-fuel-price-data">UK gov scheme</Link> to
                    give drivers access to live fuel prices and keep prices down.
                </Typography>
                <Typography>
                    Set your preferences by clicking on the toolbar above, and
                    use the map or list view to view your results.
                </Typography>
                
                <Typography variant='h6' compoonent='h4'>What's new?</Typography>
                <Box component='ul' px={3} my={1}>
                    {changes.map(x => (
                        <Typography component='li' key={x}>{x}</Typography>
                    ))}
                </Box>

                <Typography variant='h6' compoonent='h4'>Attribution</Typography>
                <Box component='ul' px={3} my={1}>
                    <Typography component='li' sx={{'& > *': {pr:1}}}>
                        <Link href="https://leafletjs.com">
                            Leaflet
                        </Link>
                        <span aria-hidden="true">|</span>
                        <span>©</span>
                        <Link href="https://www.openstreetmap.org/copyright">OpenStreetMap</Link>
                        <span>contributors</span>
                    </Typography>
                    <Typography component='li'>
                        <Link href="https://github.com/ZakDaMack/FuelFinder">
                            FuelFinder
                        </Link> by <Link href="https://zakdowsett.co.uk/">
                            Zak Dowsett
                        </Link>
                    </Typography>
                    <Typography component='li'>
                        Report any issues <Link href="https://github.com/ZakDaMack/FuelFinder/issues/new">here</Link>
                    </Typography>
                </Box>

            </DialogContent>
            <DialogActions>
                <FormControlLabel label="Don't show again" control={
                    <Checkbox value={dontShow} onChange={(e, c) =>setDontShow(c)} />
                } />
                <Button variant='contained' onClick={close}>Ok, got it</Button>
            </DialogActions>
        </Dialog>
    );
}

const Link = ({ href, children }) => (
    <a href={href} target="_blank" rel="noreferrer">{children}</a>
);