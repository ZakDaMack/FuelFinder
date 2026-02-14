import { FC } from 'react';

import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { faKey, faUser } from '@fortawesome/free-solid-svg-icons';


const LoginPage: FC = () => {
    return (
        <div className='w-full max-w-lg mx-auto'>
            {/* form */}
            <div className='flex flex-col gap-4'>
                <Input startIcon={faUser} type='email' placeholder='Email address' className='bg-white/20 dark:bg-black/20 backdrop-blur-sm' />
                <Input startIcon={faKey} type='password' placeholder='Password' className='bg-white/20 dark:bg-black/20 backdrop-blur-sm' />
                <a className='text-sm text-right text-blue-600' href="">Forgot your password?</a>
                <div className='h-2'></div>
                <Button size='xl'>Sign up</Button>
                <Button size="xl">Back to app</Button>
            </div>
        </div>
    )
}

export default LoginPage;