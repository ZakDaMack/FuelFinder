const https = require('https');
import ST from 'stjs';

// receive data from provided url
return 


function convertData(data: any, template: Object): FuelData {
    return ST.select(data)
        .transformWith(template)
        .root();
}