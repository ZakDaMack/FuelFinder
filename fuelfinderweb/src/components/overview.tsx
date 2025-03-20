import { Button } from './ui/button';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faInfo, faSliders } from '@fortawesome/free-solid-svg-icons'
import { useAppDispatch, useAppSelector } from '@/store';
import { openMenu } from '@/slices/menu_slice';

export default function Overview() {
    const dispatch = useAppDispatch();
    const openPreferences = () => dispatch(openMenu('preferences'));

    const filters = useAppSelector(s => s.stations.filters);
    const filterLength = 1

    return (
        <div className='bg-card shadow-lg flex items-center gap-3 p-2 rounded-full absolute top-4 left-4 z-[20000]'>
            <Button size='icon' className='rounded-full bg-black size-12 aspect-square' onClick={openPreferences}>
                <FontAwesomeIcon size='4x' icon={faSliders} />
            </Button>
            <div>
                <h3 className='text-lg'>Within {filters.radius} mile{filters.radius === 1 ? '' : 's'}</h3>
                <h4 className='text-neutral-500'>(+{filterLength} filter{filterLength === 1 ? '' : 's'})</h4>
            </div>
            <Button size='icon' className='rounded-full ml-1 size-5 bg-primary'>
                <FontAwesomeIcon size='sm' icon={faInfo} />
            </Button>
        </div>
    );
}