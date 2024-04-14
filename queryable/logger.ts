export default function logger(req, res, next): void {
    console.log(`${req.method.toUpperCase()} ${req.path} ${res.status} [${req.socket.remoteAddress}]`)
    next();
}