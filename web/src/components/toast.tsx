import { faWarning } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { toast } from 'sonner';

export const sendErrorToast = (message: string, header?: string) => {
    toast.error((
        <div className='text-red-900'>
            {!!header && <h3 className='text-lg font-bold'>{header}</h3>}
            <p className='text-xs'>{message[0].toLocaleUpperCase() + message.slice(1)}</p>
        </div>
    ), {
        duration: Infinity,
        richColors: false,
        closeButton: true,
        icon: <FontAwesomeIcon icon={faWarning} className="text-red-500" />,
        className: '!bg-red-50 !border-red-200',
    });
}