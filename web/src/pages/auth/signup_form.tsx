import { FC } from 'react';

import { Link } from 'react-router-dom';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';

const SignupPage: FC = () => {
    return (
        <div className='w-full max-w-lg mx-auto'>
            {/* form */}
            <div className='flex flex-col gap-4'>
                <Input type='email' placeholder='Email address' className='bg-white/20 dark:bg-black/20 backdrop-blur-sm' />
                <Input type='password' placeholder='Password' className='bg-white/20 dark:bg-black/20 backdrop-blur-sm' />
                <a className='text-sm text-right text-blue-600' href="">Forgot your password?</a>
                <div className='h-2'></div>
                <Button size='xl'>Sign up</Button>
                <Button size="xl" asChild>
                    <Link to='/'>Back to app</Link>
                </Button>
            </div>
            <div className='pt-8'>
                <p>If you sign up, you'll get access to:</p>
                <ul className='list-disc pl-4 my-2'>
                    <li>Your saved data, anywhere</li>
                    <li>Historical prices</li>
                    <li>Weekly alerts <span className='text-red-500'>[Coming soon]</span></li>
                    <li>Mobile access <span className='text-red-500'>[Coming soon]</span></li>
                    <li>And much more, all for free</li>
                </ul>
            </div>
        </div>
    )
}

export default SignupPage;