import { Schema } from 'mongoose';

export default interface IFuelData {
    site_id: string;
    company: string;
    e5: number;
    e10: number;
    b7: number;
    address: string;
    postcode: string;  
    location: {
        type: string;
        coordinates: Array<number>
    },
    created_at: string;
}