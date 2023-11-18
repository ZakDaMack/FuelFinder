import dayjs from 'dayjs';
import express from 'express';
import cors from 'cors';
import 'dotenv/config';

import StationDataSchema from '../models/stationdataschema';
import { model, connect } from 'mongoose';

import IFuelData from '../models/fueldata';
import logger from './logger';

const _port = process.env.PORT;
const _connectionString = process.env.MONGO_CONN;
const _stationData = model<IFuelData>('station', StationDataSchema);

const app = express();
app.use(cors());

app.get('/ping',(req, res) => res.send({ status: "PONG", time: dayjs().toISOString() }));

app.get('/', async (req, res) => {
  console.log(`${req.method.toUpperCase()} ${req.originalUrl}`);
    await connect(_connectionString);
    
    const distance = parseInt(req.query.distance ?? 5);
    const latitude = req.query.lat;
    const longitude = req.query.lng;

    if (!latitude || !longitude) {
        res.status = 400;
        return res.send({ message: "lng or lat is missing" });
    }

    const results = await _stationData.aggregate([
        {
          $geoNear: {
            key: 'location',
            near: {
              type: 'Point',
              coordinates: [parseFloat(longitude), parseFloat(latitude)]
            },
            distanceField: 'distance',
            maxDistance: milesToMeters(distance),
            spherical: true
          }
        },
        { $sort: { created_at: -1 } },
        {
          $group: {
            _id: '$site_id',
            records: { $push: '$$ROOT' }
          }
        },
        {
          $replaceRoot: {
            newRoot: { $first: '$records' }
          }
        }
      ],
      { maxTimeMS: 60000, allowDiskUse: true })
    .exec();
      console.log("query got " + results.length);

    res.send(results);

    // const t = await _stationData.find({ company: "BP" }).exec();
    // res.send(t);
});

app.use(logger);
app.listen(_port, () => console.log(`Queryable API listening on port ${_port}!`));

const milesToRadians = (miles: number): number => miles / 3963.2;
const kilometersToRadians = (kms: number): number => kms / 6378.1;

const milesToKilometers = (miles: number): number => miles * 1.6093;
const milesToMeters = (miles: number): number => milesToKilometers(miles) * 1000;