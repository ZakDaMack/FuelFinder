export default function GetUrlParams(obj) {
    return '?' + Object.entries(obj)
        .filter(([key, value]) => !!value)
        .map(([key, value]) => `${key}=${value}`)
        .join("&")
}