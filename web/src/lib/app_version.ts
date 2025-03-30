const VERSION = 3;

function getStoredVersion() {
    const v = localStorage.getItem("vers")
    return v == null ? null : Number(v)
}

function storeVersion() {
    localStorage.setItem("vers", VERSION.toString())
}

export {getStoredVersion, storeVersion};
export default VERSION;