export default function GetUrlParams(obj: Object) {
    return '?' + Object.entries(obj)
        .filter(([_, value]) => !!value)
        .map(([key, value]) => `${key}=${value}`)
        .join("&")
}