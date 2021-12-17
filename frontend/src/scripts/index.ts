import '../styles/index.scss';
import sampleData from './controls/demo_data.json';
import { getStateFromHash } from './controls/hashed_fragment_state';
import { showResults } from './controls/results';

const demo = document.querySelector('#demo')!;
const hash = window.location.hash;

if (hash.startsWith('#')) {
  showResults(getStateFromHash(hash));
}

function showDemo() {
  showResults(sampleData);
}

demo.addEventListener('click', (e) => {
  e.preventDefault();
  showDemo();
});
