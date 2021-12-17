import { invokeAPI } from './api';
import sampleData from './demo_data.json';
import { displayTiles } from './results';

export interface LoginForm extends HTMLFormElement {
  username: HTMLInputElement;
  password: HTMLInputElement;
  extracurricular: HTMLInputElement;
}

function toggleSubmit() {
  document.querySelector('input[type=submit]')?.toggleAttribute('disabled');
}

function showErrors(message: string | null = null, statusCode = -1) {
  const ref = document.querySelector('#errors')!;

  switch (statusCode) {
    case 400:
    case 401:
    case 403:
      ref.textContent =
        'Wystąpił błąd podczas logowania! Sprawdź dane i spróbuj ponownie';
      break;
    case 429:
      ref.textContent =
        'Serwer jest aktualnie przeciążony lub skorzystałeś/aś z kalkulatora zbyt wiele razy. Odczekaj chwilę i spróbuj ponownie';
      break;
    case 500:
      ref.textContent = `Wystąpił błąd po stronie serwera: ${message}! Spróbuj ponownie później`;
      break;
    default:
      ref.textContent = message;
  }
}

function showResults(data: any, inclEx: boolean) {
  const semester = document.querySelector<HTMLSelectElement>('#semester')!;

  semester.addEventListener('change', () => displayTiles(data, inclEx));
  semester.style.display = 'flex';

  displayTiles(data, inclEx);

  (adsbygoogle = (window as any).adsbygoogle || []).push({});
}

export async function handleForm(target: LoginForm) {
  toggleSubmit();

  const { username, password, extracurricular } = target;
  const response = await invokeAPI({
    username: username.value,
    password: password.value
  });
  const data = await response.json();

  if (response.status == 200) {
    target.reset();
    showErrors();
    showResults(data, extracurricular.checked);

    (adsbygoogle = (window as any).adsbygoogle || []).push({});
    ga('send', 'event', 'AJAX', 'FREQ', 'FREQ0');
  } else {
    showErrors(data.Reason, response.status);
    toggleSubmit();

    ga('send', 'event', 'AJAX', 'FREQ', 'FREQ1');
  }
}

export function showDemo() {
  showResults(sampleData, true);
}
