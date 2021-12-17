import { stringify } from 'querystring';

type APIPayload = {
  username: string;
  password: string;
};

export function invokeAPI(payload: APIPayload) {
  return fetch('https://api.ccwe.pl/frekwencja/', {
    method: 'POST',
    body: stringify(payload),
    headers: {
      'content-type': 'application/x-www-form-urlencoded'
    }
  });
}
