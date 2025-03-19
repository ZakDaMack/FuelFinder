import { SyntheticEvent } from 'react';
import { useMap } from 'react-leaflet/hooks';

import { Button } from './ui/button';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons'


export default function MapZoom() {
    const map = useMap();
    const zoomIn = (e: SyntheticEvent) => {
        e.stopPropagation()
        map.zoomIn()
    }

    const zoomOut = (e: SyntheticEvent) => {
        e.stopPropagation()
        map.zoomOut()
    }

    return (
        <div className='shadow-lg rounded absolute top-4 right-4 z-[1000]'>
            <Button size='icon' variant='secondary' className='rounded-b-none size-9' onClick={zoomIn}>
                <FontAwesomeIcon size='lg' className='text-primary' icon={faPlus} />
            </Button>
            <div className='border-b' />
            <Button size='icon' variant='secondary' className='rounded-t-none size-9' onClick={zoomOut}>
                <FontAwesomeIcon size='lg' className='text-primary' icon={faMinus} />
            </Button>
        </div>
    );
}