interface Station {
    id: string;
    node_id: string;
    longitude: number;
    latitude: number;

    name: string;
    brand: string;
    address: string;
    postcode: string;
    phone?: string;

    e5?: number;
    e10?: number;
    b7_standard?: number;
    b7_premium?: number;

    prices_effective_at?: string;
    updated_at: string;

    // calculated client-side from the map location
    distance: number;
}

export type { Station };
