import SvgIcon from "@mui/material/SvgIcon"
import './Draw.css'

export default function Arrow1() {
    return (
        <SvgIcon sx={{
            stroke: 'white', 
            fill: 'none',
            width: 200,
            height: 200,
            strokeWidth: 2,
            strokeDasharray: 400,
            strokeDashoffset: 400,
            animation: 'draw 2s forwards',
            '&.tail-1':{
                animationDelay: '.5s'
            },
            '&.tail-2':{
                animationDelay: '.7s'
            },
        }}>
            <svg version="1.1" xmlns="http://www.w3.org/2000/svg"x="0px" y="0px" viewBox="0 0 43.1 85.9">
            <path strokeLinecap="round" strokeLinejoin="round" className="st0 draw-arrow" d="M11.3,2.5c-5.8,5-8.7,12.7-9,20.3s2,15.1,5.3,22c6.7,14,18,25.8,31.7,33.1" />
            <path strokeLinecap="round" strokeLinejoin="round" className="draw-arrow tail-1" d="M40.6,78.1C39,71.3,37.2,64.6,35.2,58" />
            <path strokeLinecap="round" strokeLinejoin="round" className="draw-arrow tail-2" d="M39.8,78.5c-7.2,1.7-14.3,3.3-21.5,4.9" />
            </svg>
        </SvgIcon>
    )
}
