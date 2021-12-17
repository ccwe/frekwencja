import { Buffer } from 'buffer';
import gz from 'zlib';

export function getStateFromHash(hash: string) {
  return JSON.parse(
    gz.gunzipSync(Buffer.from(decodeURI(hash), 'base64')).toString()
  );
}
