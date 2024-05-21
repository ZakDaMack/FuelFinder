const VERSION = 1;

function GetStoredVersion() {
    return localStorage.getItem("vers")
}

function StoreVersion() {
    localStorage.setItem("vers", VERSION)
}

export {GetStoredVersion, StoreVersion};
export default VERSION;