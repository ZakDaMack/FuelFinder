import { faWarning } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { toast } from 'sonner';

export const sendErrorToast = (message: string, header?: string) => {
    toast.error((
        <div className='text-white'>
            {!!header && <h3 className='text-lg font-bold'>{header}</h3>}
            <p className='text-xs'>{message[0].toLocaleUpperCase() + message.slice(1)}</p>
        </div>
    ), { 
        duration: Infinity, 
        richColors: true, 
        closeButton: true,
        icon: <FontAwesomeIcon icon={faWarning} className="text-red-600" />,
    });
}