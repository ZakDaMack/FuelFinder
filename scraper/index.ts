import ST from 'stjs';
import fs from 'fs';
import path from 'path';
import axios from 'axios';
import 'dotenv/config';

import dayjs from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat';
dayjs.extend(customParseFormat)

import StationDataSchema from '../models/stationdataschema';
import { model, connect } from 'mongoose';

import IFuelData from '../models/fueldata';
import ConfData from '../models/conf';

const _configDir = process.env.CONFIG_DIR;
const _connectionString = process.env.MONGO_CONN;

async function main() {
    // ensure connection to mongo
    await connect(_connectionString);

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
    const stationData = model<IFuelData>('station', StationDataSchema);

    return ST.select(data)
    .transformWith(config.template)
    .root()
    .map(x => new stationData({
        ...x,
        "created_at": dayjs(x.created_at, config.dateFormat, true).toISOString()
    }))
    .map(u => u.save());
}


// EXECUTE
main()
    .then(() => console.log("Completed"))
    .catch((err) => console.error(`CATASTROPHIC ERROR!\n${err}`));