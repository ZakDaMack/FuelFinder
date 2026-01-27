import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { faKey, faUser } from '@fortawesome/free-solid-svg-icons';
import { FC, ReactNode } from 'react';

interface LoginProps {
    children? : ReactNode
}

const LoginPage: FC<LoginProps> = ({ children }) => {
    return (
        <main className='w-full h-screen bg-[url(/map_bg.png)] bg-cover relative'>
            <div className='absolute top-0 bottom-0 left-0 right-0 backdrop-blur-xs' />
            <div className='bg-white/20 dark:bg-black/20 backdrop-blur-md border h-screen absolute right-0 w-1/2 grid p-6 z-10'>
                {/* logo */}
                <div className='flex gap-2 items-center h-10'>
                    <img height={50} width={50} src='/favicon.ico' />
                    <h1 className='text-4xl font-bold bg-gradient-to-r from-green-600 to-blue-700 from-70% bg-clip-text text-transparent [text-shadow:0_1px_2px_rgba(0,0,0,0.15)]'>FuelFinder</h1>
                </div>
                {/* main content */}
                <div className='w-lg mx-auto'>
                    {/* form */}
                    <div className='flex flex-col gap-4'>
                        <Input startIcon={faUser} type='email' placeholder='Email address' className='bg-white/20 dark:bg-black/20 backdrop-blur-sm' />
                        <Input startIcon={faKey} type='password' placeholder='Password' className='bg-white/20 dark:bg-black/20 backdrop-blur-sm' />
                        <a className='text-sm text-right text-blue-600' href="">Forgot your password?</a>
                        <div className='h-2'></div>
                        <Button size='xl'>Sign up</Button>
                        <Button size="xl">Back to app</Button>
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
            </div>
        </main>
    )
}

export default LoginPage;