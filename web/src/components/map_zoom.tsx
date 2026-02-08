import { SyntheticEvent } from 'react';
import { useAppSelector } from '@/store';
import { useMap } from 'react-leaflet/hooks';

import { Button } from './ui/button';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlus, faMinus, faLocationArrow } from '@fortawesome/free-solid-svg-icons'


export default function MapZoom() {
    const map = useMap();
    const location = useAppSelector((state) => state.stations.location)

    const zoomIn = (e: SyntheticEvent) => {
        e.stopPropagation()
        map.zoomIn()
    }

    const zoomOut = (e: SyntheticEvent) => {
        e.stopPropagation()
        map.zoomOut()
    }

    const recentre = (e: SyntheticEvent) => {
        e.stopPropagation()
        map.flyTo(location, 13);
    }

    return (
        <div className='shadow-lg rounded-3xl overflow-hidden absolute top-4 right-4 z-[1000]' onDoubleClick={(e) => e.stopPropagation()} onClick={(e) => e.stopPropagation()}>
            <Button size='icon' variant='secondary' className='rounded-b-none size-9' onClick={recentre} onDoubleClick={(e) => e.stopPropagation()}>
                <FontAwesomeIcon size='lg' className='text-indigo-600' icon={faLocationArrow} />
            </Button>
            <div className='border-b' />
            <Button size='icon' variant='secondary' className='rounded-none size-9' onClick={zoomIn} onDoubleClick={(e) => e.stopPropagation()}>
                <FontAwesomeIcon size='lg' className='text-neutral-900' icon={faPlus} />
            </Button>
            <div className='border-b' />
            <Button size='icon' variant='secondary' className='rounded-t-none size-9' onClick={zoomOut} onDoubleClick={(e) => e.stopPropagation()}>
                <FontAwesomeIcon size='lg' className='text-neutral-900' icon={faMinus} />
            </Button>
        </div>
    );
}