import { Schema } from 'mongoose';
import IFuelData from './fueldata';

export default new Schema<IFuelData>({
    site_id: String,
    company: String,
    e5: { type: Number, required: false },
    e10: { type: Number, required: false },
    b7: { type: Number, required: false },
    sdv: { type: Number, required: false },
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