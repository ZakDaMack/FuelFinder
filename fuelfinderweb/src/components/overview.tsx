import { Button } from './ui/button';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faInfo, faSliders } from '@fortawesome/free-solid-svg-icons'

export default function Overview() {
    return (
        <div className='bg-white shadow-lg flex items-center gap-3 p-2 rounded-full absolute top-4 left-4 z-[1000]'>
            <Button size='icon' className='rounded-full bg-black p-4 aspect-square'>
                <FontAwesomeIcon size='4x' className='text-white' icon={faSliders} />
            </Button>
            <div>
                <h3 className='text-lg'>Within 3 miles</h3>
                <h4 className='text-neutral-500'>(+0 filters)</h4>
            </div>
            <Button size='icon' className='rounded-full size-4'>
                <FontAwesomeIcon size='sm' className='text-primary' icon={faInfo} />
            </Button>
        </div>
    );
}