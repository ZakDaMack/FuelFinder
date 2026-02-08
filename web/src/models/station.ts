interface Station {
    id: string;
    site_id: string;
    longitude: number;
    latitude: number;

    brand: string;
    address: string;
    postcode: string;

    e5?: number;
    e10?: number;
    b7?: number;
    sdv?: number;

    distance: number;
    created_at: number;
}

export type { Station };
