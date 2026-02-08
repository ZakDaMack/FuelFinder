import { LatLng } from "leaflet";

export const calculateDistance = (pointA: LatLng, pointB: LatLng): number => {
    // earth radius in metres
    const radius = 6371e3;

    // latitude in radians
    const latARad = pointA.lat * Math.PI/180; 
    const latBRad = pointB.lat * Math.PI/180;

    const distanceLatRad = (pointB.lat - pointA.lat) * Math.PI/180;
    const distanceLngRad = (pointB.lng - pointA.lng) * Math.PI/180;

    // calc azimuth
    const azimuth = Math.sin(distanceLatRad/2) * Math.sin(distanceLatRad/2) +
            Math.cos(latARad) * Math.cos(latBRad) *
            Math.sin(distanceLngRad/2) * Math.sin(distanceLngRad/2);

    // calc distance
    const c = 2 * Math.atan2(Math.sqrt(azimuth), Math.sqrt(1-azimuth));
    const d = radius * c; // in metres
    return d;
}