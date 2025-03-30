import { FC } from "react"
import { cn } from "@/lib/utils"
import { useAppSelector } from "@/store"

const LoadingIndicator: FC = () => {
    const isLoading = useAppSelector(s => s.stations.loading)
    return (
        <div className={cn(
            "bg-neutral-100 shadow-lg flex items-center gap-3 p-3 rounded-full absolute bottom-24 md:bottom-[unset] md:top-4 left-1/2 -translate-x-1/2 z-[10000]",
            "[&>*]:animate-dot",
            { "hidden": !isLoading },
        )}>
            <div className="size-3 rounded-full bg-neutral-900"></div>
            <div className="size-3 rounded-full bg-neutral-900 [animation-delay:_0.2s_!important]"></div>
            <div className="size-3 rounded-full bg-neutral-900 [animation-delay:_0.3s_!important]"></div>
        </div>
    )
}

export default LoadingIndicator