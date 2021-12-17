import { unescape } from 'querystring';
import { gunzipSync } from 'zlib';

export function getStateFromHash(hash: string) {
  return JSON.parse(
    gunzipSync(Buffer.from(unescape(hash), 'base64url')).toString()
  );
}
