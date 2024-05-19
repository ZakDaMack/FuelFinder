package database

import "go.mongodb.org/mongo-driver/bson"

/*
AGGREGATE DATA
[
	{
		$geoNear: {
			key: 'location',
			near: {
				type: 'Point',
				coordinates: [parseFloat(longitude), parseFloat(latitude)]
			},
			distanceField: 'distance',
      		query: {"brand": {"$in": ["Tesco", "Esso"]} },
			maxDistance: milesToMeters(distance),
			spherical: true
		}
	},
	{ $sort: { createdat: -1 } },
	{
		$group: {
			_id: '$siteid',
			records: { $push: '$$ROOT' }
		}
	},
	{
		$replaceRoot: {
			newRoot: { $first: '$records' }
		}
	}
],
{ maxTimeMS: 60000, allowDiskUse: true }
*/

func MakeFilter(lat, long, distRads float64) bson.D {
	return bson.D{{
		Key: "location", Value: bson.D{{
			Key: "$geoWithin", Value: bson.D{{
				Key: "$centerSphere", Value: bson.A{
					bson.A{long, lat},
					distRads,
				},
			}},
		}},
	}}
}

func MakeAggregatePipeline(lat, long float64, distMetres int, brands []string) bson.A {
	geoNear := bson.D{
		{Key: "key", Value: "location"},
		{Key: "near", Value: bson.D{
			{Key: "type", Value: "Point"},
			{Key: "coordinates", Value: bson.A{long, lat}},
		}},
		{Key: "distanceField", Value: "distance"},
		{Key: "maxDistance", Value: distMetres},
		{Key: "spherical", Value: true},
	}

	if brands != nil {
		geoNear = append(geoNear, bson.E{
			Key: "query", Value: bson.D{
				{Key: "brand", Value: bson.D{
					{Key: "$in", Value: brands},
				}},
			}},
		)
	}

	return bson.A{
		bson.D{{
			Key: "$geoNear", Value: geoNear,
		}},
		bson.D{{
			Key: "$sort", Value: bson.D{{
				Key: "createdat", Value: -1,
			}},
		}},
		bson.D{{
			Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$siteid"},
				{Key: "records", Value: bson.D{
					{Key: "$push", Value: "$$ROOT"},
				}},
			},
		}},
		bson.D{{
			Key: "$replaceRoot", Value: bson.D{{
				Key: "newRoot", Value: bson.D{{
					Key: "$first", Value: "$records",
				}},
			}},
		}},
	}
}
