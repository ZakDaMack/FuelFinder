import { useState } from 'react';
import { useSelector } from 'react-redux';

import Box from '@mui/material/Box';
import Dialog from '@mui/material/Dialog';
import Button from '@mui/material/Button';
import Checkbox from '@mui/material/Checkbox';
import Typography from '@mui/material/Typography';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import FormControlLabel from '@mui/material/FormControlLabel';

export default function IntroDialog() {

    const vers = 1
    const changes = [
        "Filter by range and brands",
        "Sort list view by distance/price",
        "New, fresh mobile-first design"
    ];

    const storedVersion = localStorage.getItem("vers") ?? 0
    const [open, setOpen] = useState(storedVersion < vers)
    const [dontShow, setDontShow] = useState(false)

    const close = () => {
        if (dontShow) localStorage.setItem("vers", vers)
        setOpen(false)
    }

    return ( 
        <Dialog sx={{zIndex: 10000}} transitionDuration={{enter: 0, exit: 500}} open={open}>
            <DialogContent sx={{
                '& li': { pb: 0.5 },
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
                    FuelFinder is an open source project developed by <a href="https://zakdowsett.co.uk/" target="_blank">Zak Dowsett</a>.
                    The app will show you the cheapest fuel prices in your area by using data provided
                    by brands participating in the <a href="https://www.gov.uk/guidance/access-fuel-price-data" target="_blank">UK gov scheme</a> to
                    give drivers access to live fuel prices and keep prices down.
                </Typography>
                <Typography>
                    Set your preferences by clicking on the toolbar above, and
                    use the map or list view to view your results.
                </Typography>
                
                <Typography variant='h6' compoonent='h4'>Latest Changes</Typography>
                <Box component='ul' px={3} my={1}>
                    {changes.map(x => (
                        <Typography component='li'>{x}</Typography>
                    ))}
                </Box>

                <Typography variant='h6' compoonent='h4'>Attribution</Typography>
                <Box component='ul' px={3} my={1}>
                    <Typography component='li' sx={{'& > *': {pr:1}}}>
                        <a href="https://leafletjs.com" target="_blank" title="A JavaScript library for interactive maps">
                            Leaflet
                        </a>
                        <span aria-hidden="true">|</span>
                        <span>©</span>
                        <a href="https://www.openstreetmap.org/copyright" target="_blank">OpenStreetMap</a>
                        <span>contributors</span>
                    </Typography>
                    <Typography component='li'>
                        <a href="https://github.com/ZakDaMack/FuelFinder" target="_blank">
                            FuelFinder
                        </a> by <a href="https://zakdowsett.co.uk/" target="_blank">
                            Zak Dowsett
                        </a>
                    </Typography>
                    <Typography component='li'>
                        Report any issues <a href="https://github.com/ZakDaMack/FuelFinder/issues/new" target="_blank">here</a>
                    </Typography>
                </Box>

            </DialogContent>
            <DialogActions>
                <FormControlLabel label="Don't show this again" control={
                    <Checkbox value={dontShow} onChange={(e, c) =>setDontShow(c)} />
                } />
                <Button variant='contained' onClick={close}>Ok, got it</Button>
            </DialogActions>
        </Dialog>
    );
}