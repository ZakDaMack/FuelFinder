import { FC } from 'react';
import { Outlet } from 'react-router-dom';

const AuthLayout: FC = () => {
    return (
        <main className='w-full h-screen bg-[url(/stations_bg.png)] bg-cover relative'>
            <div className='absolute top-0 bottom-0 left-0 right-0 backdrop-blur-xs' />
            <div className='bg-white/20 dark:bg-black/20 backdrop-blur-md border h-screen absolute right-0 w-full md:w-1/2 grid p-6 z-10'>
                {/* logo */}
                <div className='flex gap-2 items-center h-10'>
                    <img height={50} width={50} src='/favicon.ico' />
                    <h1 className='text-4xl font-bold bg-gradient-to-r from-green-600 to-blue-700 from-70% bg-clip-text text-transparent [text-shadow:0_1px_2px_rgba(0,0,0,0.15)]'>FuelFinder</h1>
                </div>
                {/* main content */}
                <Outlet />
            </div>
        </main>
    )
}

export default AuthLayout;