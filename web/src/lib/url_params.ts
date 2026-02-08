export default function GetUrlParams(obj: Object) {
    return '?' + Object.entries(obj)
        .filter(([_, value]) => !!value)
        .filter(([_, value]) => Array.isArray(value) ? value.length > 0 : true)
        .map(([key, value]) => `${key}=${value}`)
        .join("&")
}