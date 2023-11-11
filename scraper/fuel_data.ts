interface FuelData {
    id: number;
    site_id: string;
    company: string;
    e5: number;
    e10: number;
    b7: number;
    address: string;
    postcode: string;
    latitude: number;
    longitude: number;
    created_at: DeconstructedDate;
}