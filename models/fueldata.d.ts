import { Schema } from 'mongoose';

type Nullable<T> = T | null;

export default interface IFuelData {
    site_id: string;
    company: string;
    e5: Nullable<number>;
    e10: Nullable<number>;
    b7: Nullable<number>;
    sdv: Nullable<number>;
    address: string;
    postcode: string;  
    location: {
        type: string;
        coordinates: Array<number>
    },
    created_at: string;
}