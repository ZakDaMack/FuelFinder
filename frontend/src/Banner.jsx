import styled from "@emotion/styled";
import Box from "@mui/material/Box";

export default function OfdBanner(props) {
    return (
        <Box sx={{
            // bgcolor: 'white',
            // borderRadius: '0.375rem',
            position: 'absolute',
            m:1, top:0, right: 0,
            zIndex: 1000
        }}>
            <Image src="/banner.png" alt="OpenFuelData banner"/>
        </Box>
    );
}

const Image = styled('img')({
    borderRadius: '0.375rem',
    maxWidth: 200,
    margin: 4
});