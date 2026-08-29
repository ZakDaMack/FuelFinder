import dayjs from "dayjs";
import { FC, useMemo, useState } from "react";

import { Station } from "@/models/station";

import { Button } from "@/components/ui/button";
import { faSort } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { AreaChart, CartesianGrid, XAxis, Area, YAxis } from "recharts";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent } from "@/components/ui/chart";
import { DropdownMenu, DropdownMenuContent, DropdownMenuGroup, DropdownMenuLabel, DropdownMenuRadioGroup, DropdownMenuRadioItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";

const chartConfig = {
  e5: {
    label: "Super (E5)",
    color: "var(--color-blue-700)",

  },
  e10: {
    label: "Petrol (E10)",
    color: "var(--color-green-700)",
  },
  b7_standard: {
    label: "Diesel (B7)",
    color: "var(--color-yellow-700)",
  },
  b7_premium: {
    label: "Super Diesel (B7 Premium)",
    color: "var(--color-orange-700)",
  },
} satisfies ChartConfig

const ranges = {
    '7d': 'Last 7 Days',
    '30d': 'Last 30 Days',
    '-1': 'All Time'
}

const fuels = {
    'e5'          : {colour: 'blue-700', bg: 'blue-300', title: 'Super'},
    'e10'         : {colour: 'green-700', bg: 'green-300', title: 'Petrol'},
    'b7_standard' : {colour: 'yellow-700', bg: 'yellow-300', title: 'Diesel'},
    'b7_premium'  : {colour: 'orange-700', bg: 'orange-300', title: 'Super Diesel'},
}

const PriceHistoryCard: FC<{stations: Station[]}> = ({ stations }) => {
    const [range, setRange] = useState<keyof typeof ranges>('7d')
    
    const data = useMemo(() => {
        // each row is one price record, timestamped by when its prices took effect
        let s = stations
            .filter(ss => !!ss.prices_effective_at)
            .sort((a,b) => a.prices_effective_at!.localeCompare(b.prices_effective_at!))
        if (s.length === 0) return [];

        const fuels = ['e5', 'e10', 'b7_standard', 'b7_premium'] as (keyof Station)[]
        const r = parseInt(range.split('d')[0])

        return s
            .filter(ss => r === -1 || Date.now() - new Date(ss.prices_effective_at!).getTime() <= r * 24 * 60 * 60 * 1000)
            .map(ss => ({
                date: new Date(ss.prices_effective_at!).getTime(),
                ...fuels.reduce((acc, fuel) => ({ ...acc, [fuel]: !!ss[fuel] ? ss[fuel] : undefined }), {})
            }))
    }, [stations, range])

    return (
        <Card className="bg-white shadow-lg">
            <CardHeader className="flex">
                <div className="grid flex-1 gap-1">
                    <CardTitle className="text-lg font-semibold">Price History</CardTitle>
                    <CardDescription>Historical price data for this station</CardDescription>
                </div>
                <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <Button variant="outline">
                        {ranges[range] as string}
                        <FontAwesomeIcon icon={faSort} className="ml-2" />
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="w-40 z-[10000]" align="end">
                    <DropdownMenuGroup>
                        <DropdownMenuLabel>Time Range</DropdownMenuLabel>
                        <DropdownMenuRadioGroup value={range} onValueChange={setRange as ((value: string) => void)}>
                            {Object.entries(ranges).map(([key, label]) => (
                                <DropdownMenuRadioItem key={key} value={key}>{label}</DropdownMenuRadioItem>
                            ))}
                        </DropdownMenuRadioGroup>
                    </DropdownMenuGroup>
                </DropdownMenuContent>
                </DropdownMenu>
            </CardHeader>
            <CardContent>
                {/* <pre>{JSON.stringify(data, null, 2)}</pre> */}
                <ChartContainer config={chartConfig}>
                    <AreaChart
                        accessibilityLayer
                        data={data}
                        margin={{
                            left: 12,
                            right: 12,
                        }}
                    >
                    <CartesianGrid vertical={false} />
                    <XAxis
                        dataKey="date"
                        type="number"
                        scale="time"
                        domain={['dataMin', 'dataMax']}
                        tickMargin={12}
                        tickFormatter={(value) => dayjs(value).toDate().toLocaleDateString()}
                    />
                    <YAxis 
                        type="number" 
                        tickMargin={8} 
                        domain={[
                            100, 
                            (dataMax) => Math.ceil(dataMax / 10) * 10
                        ]}
                    />

                    <ChartTooltip
                        cursor={false}
                        content={<ChartTooltipContent
                            labelKey="date"
                            labelFormatter={(val) => val ? dayjs(val).toDate().toLocaleString() : ''}
                        />}
                    />
                    
                    {/* axis */}
                    {Object.entries(fuels).map(([fuel, {colour, bg, title}]) => (
                         <Area
                            key={fuel}
                            dataKey={fuel}
                            name={title}
                            type="step"
                            fill={`var(--color-${bg})`}
                            fillOpacity={0.3}
                            stroke={`var(--color-${colour})`}
                        />
                    ))}
                </AreaChart>
                </ChartContainer>
            </CardContent>
        </Card>
    )
}

export default PriceHistoryCard;