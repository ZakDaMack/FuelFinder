import { cn } from "@/lib/utils";

import { icon } from "@fortawesome/fontawesome-svg-core";
import { faGasPump } from "@fortawesome/free-solid-svg-icons";

const faSvg = icon(faGasPump).html[0];

export function GradientFuelIcon({ className }: { className?: string }) {
  return (
    <div
      className={cn(`
        w-8 h-8   
      `, className)}
      style={{
        WebkitMaskImage: `url("data:image/svg+xml;utf8,${encodeURIComponent(faSvg)}")`,
        maskImage: `url("data:image/svg+xml;utf8,${encodeURIComponent(faSvg)}")`,
        WebkitMaskRepeat: "no-repeat",
        maskRepeat: "no-repeat",
        WebkitMaskSize: "contain",
        maskSize: "contain",
        WebkitMaskPosition: "center",
        maskPosition: "center",
      }}
    />
  );
}