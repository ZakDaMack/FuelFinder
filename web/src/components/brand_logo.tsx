import { FC, useMemo } from "react";
import { cn } from "@/lib/utils";

import ApplegreenLogo from '@/assets/logos/applegreen.png';
import AsdaLogo from '@/assets/logos/asda.png';
import BpLogo from '@/assets/logos/bp.png';
import CoopLogo from '@/assets/logos/coop.png';
import CostcoLogo from '@/assets/logos/costco.png';
import EssarLogo from '@/assets/logos/essar.png';
import EssoLogo from '@/assets/logos/esso.png';
import EGOnTheMoveLogo from '@/assets/logos/eg-on-the-move.png';
import GulfLogo from '@/assets/logos/gulf.png';
import HarvestEnergyLogo from '@/assets/logos/harvest-energy.png';
import JetLogo from '@/assets/logos/jet.png';
import MorrisonsLogo from '@/assets/logos/morrisons.png';
import MurcoLogo from '@/assets/logos/murco.png';
import SainsburysLogo from '@/assets/logos/sainsburys.png';
import ShellLogo from '@/assets/logos/shell.png';
import TescoLogo from '@/assets/logos/tesco.png';
import TexacoLogo from '@/assets/logos/texaco.png';
import ValeroLogo from '@/assets/logos/valero.png';
import WelcomebreakLogo from '@/assets/logos/welcome-break.png'

const BRAND_LOGOS: [string, string][] = [
    ['APPLEGREEN', ApplegreenLogo],
    ['ASDA', AsdaLogo],
    ['BP', BpLogo],
    ['COSTCO', CostcoLogo],
    ['ESSAR', EssarLogo],
    ['GULF', GulfLogo],
    ['JET', JetLogo],
    ['MURCO', MurcoLogo],
    ['SHELL', ShellLogo],
    ['TEXACO', TexacoLogo],
    ['COOP', CoopLogo],
    ['ESSO', EssoLogo],
    ['EGONTHEMOVE', EGOnTheMoveLogo],
    ['HARVESTENERGY', HarvestEnergyLogo],
    ['MORRISONS', MorrisonsLogo],
    ['SAINSBURYS', SainsburysLogo],
    ['TESCO', TescoLogo],
    ['VALERO', ValeroLogo],
    ['WELCOMEBREAK', WelcomebreakLogo]
];

const DEFAULT_LOGO = '/station.png';

const BrandLogo: FC<{
    brand: string, 
    className?: string 
}> = ({ 
    brand, 
    className 
}) => {
    const logo =  useMemo(() => {
        const normalised = brand?.toUpperCase().replace(/[^A-Z]/g, '') ?? '';
        return BRAND_LOGOS.find(([name]) => normalised.includes(name))?.[1] ?? DEFAULT_LOGO;
    }, [brand]);

    return <img src={logo} alt={`${brand} logo`} className={cn("max-h-[50px] max-w-[150px]", className)} />
}

export default BrandLogo;