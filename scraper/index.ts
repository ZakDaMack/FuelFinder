import ST from 'stjs';
import fs from 'fs';
import path from 'path';
import axios from 'axios';

import dayjs from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat';
dayjs.extend(customParseFormat)

import FuelDataSchema from './fueldataschema';
import { model, connect } from 'mongoose';

import IFuelData from './fueldata';
import ConfData from './conf';

const _configDir = process.env.CONFIG_DIR ?? '../data_configs';

async function main() {
    // ensure connection to mongo
    await connect('mongodb://database:27017/openfueldata');

    // get the config dir and files within
    const files = await fs.promises.readdir(_configDir);
    files.map(async file => {
        try {
            const conf = await readConfigFile(path.join(_configDir, file));
            const resData = await fetchData(conf.url);
            convertAndSaveData(resData, conf);
            console.log(`Added ${file.split('.')[0]} to db`);
        } catch (e) {
            console.error("Error handling " + file);
            console.error(e);
        }
    }); 
}

async function readConfigFile(path: string): Promise<ConfData> {
    const file = await fs.promises.readFile(path, { encoding: "utf8"});
    return JSON.parse(file || '{}');
}

async function fetchData(url: string): Promise<any> {
    const res = await axios(url);
    return res.data;
}

function convertAndSaveData(data: any, config: ConfData): void {
    const fuelData = model<IFuelData>('fuel_data', FuelDataSchema);

    return ST.select(data)
    .transformWith(config.template)
    .root()
    .map(x => new fuelData({
        ...x,
        "created_at": dayjs(x.created_at, config.dateFormat, true).toISOString()
    }))
    .map(u => u.save());
}


// EXECUTE
main()
    .then(() => console.log("Completed"))
    .catch((err) => console.error(`CATASTROPHIC ERROR!\n${err}`));