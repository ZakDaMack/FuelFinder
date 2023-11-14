import { UUID } from "crypto";
import { Schema } from 'mongoose';

export interface IFuelData {
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

// 2. Create a Schema corresponding to the document interface.
export const FuelDataSchema = new Schema<IFuelData>({
    site_id: String,
    company: String,
    e5: Number,
    e10: Number,
    b7: Number,
    address: String,
    postcode: String,
    location: {
        type: {
            type: String,
            enum: ['Point'],
            required: true
        },
        coordinates: {
            type: [Number],
            required: true
        }
    },
    created_at: Date
});