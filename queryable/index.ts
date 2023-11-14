import dayjs from 'dayjs';
import express from 'express';
import cors from 'cors';
import { model, connect } from 'mongoose';

import { IFuelData, FuelDataSchema } from './fueldata';

const port = process.env.PORT ?? 3000;
const fuelData = model<IFuelData>('fuel_data', FuelDataSchema);

const app = express();
app.use(cors());
app.get('/', async (req, res) => {
    await connect('mongodb://database:27017/openfueldata');
    
    const distance = parseInt(req.query.distance ?? 5);
    const latitude = parseFloat(req.query.lat);
    const longitude = parseFloat(req.query.lng);

    const results = await fuelData.find({
        location: {
            $geoWithin: {
                $centerSphere: [
                    [ longitude, latitude ],
                    milesToRadians(distance)
                ]
            }
        }
    }).exec();

    res.send(results);
});

app.listen(port, () => console.log(`Scraper API listening on port ${port}!`));

const milesToRadians = (miles: number): number => miles / 3963.2;
const kilometersToRadians = (kms: number): number => kms / 6378.1;