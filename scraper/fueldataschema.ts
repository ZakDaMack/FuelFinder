import { Schema } from 'mongoose';
import IFuelData from './fueldata';

export default new Schema<IFuelData>({
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