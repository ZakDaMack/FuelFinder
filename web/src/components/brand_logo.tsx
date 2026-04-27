import { FC, useMemo } from "react";
import { cn } from "@/lib/utils";

import ApplegreenLogo from '@/assets/logos/applegreen.png';
import BpLogo from '@/assets/logos/bp.png';
import EssarLogo from '@/assets/logos/essar.png'; 
import GulfLogo from '@/assets/logos/gulf.png';
import JetLogo from '@/assets/logos/jet.png';
import MurcoLogo from '@/assets/logos/murco.png'; 
import ShellLogo from '@/assets/logos/shell.png';  
import TexacoLogo from '@/assets/logos/texaco.png'; 
import AsdaLogo from '@/assets/logos/asda.png';        
import CoopLogo from '@/assets/logos/coop.png';  
import EssoLogo from '@/assets/logos/esso.png';   
import HarvestEnergyLogo from '@/assets/logos/harvest-energy.png';  
import MorrisonsLogo from '@/assets/logos/morrisons.png';  
import SainsburysLogo from '@/assets/logos/sainsburys.png';  
import TescoLogo from '@/assets/logos/tesco.png';  
import ValeroLogo from '@/assets/logos/valero.png';

const BrandLogo: FC<{ 
    brand: string, 
    className?: string 
}> = ({ 
    brand, 
    className 
}) => {
    const logo =  useMemo(() => {
        switch (brand?.toUpperCase()) {
            case 'APPLEGREEN' : return ApplegreenLogo
            case 'ASDA' : return AsdaLogo
            case 'ASDA EXPRESS' : return AsdaLogo
            case 'BP': return BpLogo
            case 'ESSAR': return EssarLogo
            case 'GULF': return GulfLogo
            case 'JET': return JetLogo
            case 'MURCO': return MurcoLogo
            case 'SHELL': return ShellLogo
            case 'TEXACO': return TexacoLogo
            case 'COOP': return CoopLogo
            case 'CO-OP': return CoopLogo
            case 'ESSO': return EssoLogo
            case 'HARVEST ENERGY': return HarvestEnergyLogo
            case 'MORRISONS': return MorrisonsLogo
            case 'SAINSBURYS': return SainsburysLogo
            case 'SAINSBURY\'S': return SainsburysLogo
            case 'TESCO': return TescoLogo
            case 'VALERO': return ValeroLogo
            default: return null
        }
    }, [brand]);

    return logo ? 
        <img src={logo} alt={`${brand} logo`} className={cn("max-h-[50px] max-w-[150px]", className)} /> : 
        <p className={cn("uppercase text-2xl", className)}>{brand?.toUpperCase()}</p>
}

export default BrandLogo;