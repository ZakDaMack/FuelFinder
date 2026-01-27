import * as React from "react"

import { cn } from "@/lib/utils"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { IconDefinition } from "@fortawesome/fontawesome-svg-core"
import { faE, faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons"

type InputProps = React.InputHTMLAttributes<HTMLInputElement> & {
  startIcon?: IconDefinition
}

function Input({ className, startIcon, type, ...props }: InputProps) {
  const [hidden, setHidden] = React.useState(true)
  const toggleHidden = () => setHidden(!hidden)

  return (
    <div className="relative w-full">
      {startIcon && (
        <span className="pointer-events-none absolute left-5 top-1/2 -translate-y-1/2 text-muted-foreground">
          <FontAwesomeIcon icon={startIcon} />
        </span>
      )}

      <input
        type={type === "password" && !hidden ? "text" : type}
        data-slot="input"
        className={cn(
          "file:text-foreground placeholder:text-muted-foreground selection:bg-primary selection:text-primary-foreground dark:bg-input/30 border-input h-14 w-full min-w-0 rounded-full border bg-transparent py-3 text-base shadow-xs transition-[color,box-shadow] outline-none file:inline-flex file:h-7 file:border-0 file:bg-transparent file:text-sm file:font-medium disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
          "focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[2px]",
          "aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
          startIcon ? "pl-12 pr-5" : "px-5",
          className
        )}
        {...props}
      />

      {type === "password" && (
        <button
          type="button"
          onClick={toggleHidden}
          className={cn(
            "absolute  top-1/2 -translate-y-1/2 text-sm text-muted-foreground hover:text-foreground",
            hidden ? 'right-[20px]' : 'right-[19px]'
          )}
        >
          <FontAwesomeIcon icon={hidden ? faEye : faEyeSlash} />
        </button>
      )}
    </div>
  )
}

export { Input }
