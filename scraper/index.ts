import ST from 'stjs';
import dayjs from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat';
import express from 'express';
import axios from 'axios';
import { model, connect } from 'mongoose';

import { IFuelData, FuelDataSchema } from './fueldata';

const port = process.env.PORT ?? 3000;
const url = 'https://www.bp.com/en_gb/united-kingdom/home/fuelprices/fuel_prices_data.json';
dayjs.extend(customParseFormat)

const dateFormat = "DD/MM/YYYY HH:mm:ss"
const template = {
    "{{#each stations}}": {
        "site_id": "{{site_id}}",
        "company": "applegreen",
        "e5": "{{#? prices.E5}}",
        "e10": "{{#? prices.E10}}",
        "b7": "{{#? prices.B7}}",
        "address": "{{address}}",
        "postcode": "{{postcode}}",
        "location": {
            "type": "Point",
            "coordinates": [
                "{{location.longitude}}",
                "{{location.latitude}}"
            ] 
        },
        "created_at": "{{$root.last_updated}}"
    }
};

const fuelData = model<IFuelData>('fuel_data', FuelDataSchema);

const app = express();
app.get('/', async (req, res) => {
    await connect('mongodb://database:27017/openfueldata');
    const data = await fetchData(url);
    convertData(data);
    res.send("Success");
});

app.listen(port, () => console.log(`Scraper API listening on port ${port}!`));

async function fetchData(url: string): Promise<any> {
    const res = await axios(url);
    return res.data;
}

function convertData(data: any) {
    return ST.select(data)
    .transformWith(template)
    .root()
    .map(x => new fuelData({
        ...x,
        "created_at": dayjs(x.created_at, dateFormat, true).toISOString()
    }))
    .map(u => u.save());
}