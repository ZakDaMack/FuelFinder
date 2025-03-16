interface Station {
    site_id: string;
    brand: string;
    address: string;
    postcode: string;
    location: {
        type: string;
        coordinates: [number, number];
    };
    e5: number;
    e10: number;
    b7: number;
    sdv: number;
    distance: number;
    created_at: number;
}

export type { Station };
