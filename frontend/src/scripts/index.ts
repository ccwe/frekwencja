import '../styles/index.scss';
import { handleForm, LoginForm, showDemo } from './controls/form';

const form = document.querySelector<LoginForm>('form')!;
const demo = document.querySelector('#demo')!;

form.addEventListener('submit', (e) => {
  e.preventDefault();
  handleForm(form);
});

demo.addEventListener('click', (e) => {
  e.preventDefault();
  showDemo();
});
