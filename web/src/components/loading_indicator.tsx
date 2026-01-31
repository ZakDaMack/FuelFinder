import { FC } from "react"
import { cn } from "@/lib/utils"
import { useAppSelector } from "@/store"

const LoadingIndicator: FC = () => {
    const isLoading = useAppSelector(s => s.stations.loading)
    return (
        <div className={cn(
            "bg-neutral-100 shadow-lg flex items-center gap-2 p-3 rounded-full absolute bottom-24 md:bottom-[unset] md:top-4 left-1/2 -translate-x-1/2 z-[10000]",
            { "hidden": !isLoading },
        )}>
            <span className="size-2 animate-ping rounded-full bg-neutral-900"></span>
            <span className="size-2 animate-ping rounded-full bg-neutral-900 [animation-delay:0.2s]"></span>
            <span className="size-2 animate-ping rounded-full bg-neutral-900 [animation-delay:0.4s]"></span>
        </div>
    )
}

export default LoadingIndicator