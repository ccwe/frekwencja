import '../styles/index.scss';
import { showDemo } from './controls/form';

const demo = document.querySelector('#demo')!;

demo.addEventListener('click', (e) => {
  e.preventDefault();
  showDemo();
});
