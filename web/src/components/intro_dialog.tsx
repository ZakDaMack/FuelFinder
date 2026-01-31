import { FC, ReactNode, useState } from 'react';
import { closeMenu } from '@/slices/menu_slice';
import { storeVersion } from '@/lib/app_version';
import { useAppDispatch, useAppSelector } from '@/store';

import { Button } from './ui/button';
import { Checkbox } from './ui/checkbox';

import { 
    Dialog,
    DialogTitle,
    DialogHeader,
    DialogContent,
} from "./ui/dialog";

const changes = [
    "Added the ability to filter by fuel type"
];

const IntroDialog: FC = () => {
    const dispatch = useAppDispatch();
    const isOpen = useAppSelector((state) => state.menus.info)

    const [dontShow, setDontShow] = useState<boolean>(false)
    const close = () => {
        if (dontShow) storeVersion()
        dispatch(closeMenu('info'))
    }

    return ( 
        <Dialog open={isOpen}>    
            <DialogContent hideClose className='z-[30000] rounded-2xl'>
                <DialogHeader>
                    <DialogTitle className='font-normal text-3xl'>
                        Welcome to <span className='font-semibold bg-clip-text text-transparent bg-linear-45 from-primary from-60% to-blue-700'>FuelFinder</span>
                    </DialogTitle>
                </DialogHeader>
                <div className='space-y-6'>
                    <p className='text-justify'>
                        FuelFinder is an open source project with the aim of saving you money.
                        The app will show you the cheapest fuel prices in your area by using data provided
                        by brands participating in the <Link href="https://www.gov.uk/guidance/access-fuel-price-data">UK gov scheme</Link> to
                        give drivers access to live fuel prices and keep prices down.
                    </p>
                    <p>
                        Set your preferences by clicking on the toolbar above, and
                        use the map or list view to view your results.
                    </p>
                    <div>
                        <h4 className='text-xl font-bold'>What's new?</h4>
                        <ul className='list-disc [&>*]:ml-4'>
                            {changes.map(x => (<li key={x}>{x}</li>))}
                        </ul>
                    </div>
                    <div>
                        <h4 className='text-xl font-bold'>Attribution</h4>
                        <ul className='list-disc [&>*]:ml-4'>
                            <li className='space-x-1'>
                                <Link href="https://leafletjs.com">Leaflet</Link>
                                <span aria-hidden="true">|</span>
                                <span>©</span>
                                <Link href="https://www.openstreetmap.org/copyright">OpenStreetMap</Link>
                                <span>contributors</span>
                            </li>
                            <li>
                                <Link href="https://github.com/ZakDaMack/FuelFinder">
                                    FuelFinder
                                </Link> by <Link href="https://zakdowsett.co.uk/">
                                    Zak Dowsett
                                </Link>
                            </li>
                            <li>
                                Report any issues <Link href="https://github.com/ZakDaMack/FuelFinder/issues/new">here</Link>
                            </li>
                        </ul>
                    </div>
                    <div className='flex justify-end gap-4'>
                        <div className="flex items-center space-x-2">
                            <Checkbox id="dontshow" checked={dontShow} onCheckedChange={(val) => setDontShow(val as boolean)} />
                            <label htmlFor="dontshow" className="text-sm">Don't show me again</label>
                        </div>
                        <Button variant='primary' onClick={close}>Ok, got it</Button>
                    </div>
                </div>
            </DialogContent>
        </Dialog>
    );
}

export default IntroDialog;

const Link: FC<{
    href: string,
    children: ReactNode|ReactNode[]
}> = ({ href, children }) => (
    <a href={href} className='text-indigo-600' target="_blank" rel="noreferrer">{children}</a>
);