import tableTile from '../../templates/results/tile.pug';

type Attendance = { Value: number; Types: number[] };
type GPA = number[];

interface Results {
  Attendance: { [subject: string]: Attendance[] };
  GPAs: { [subject: string]: GPA };
  GPA: GPA;
}

interface ResultsTileData {
  [subject: string]: {
    Attendance: Attendance[];
    GPA: GPA;
  };
}

type ResultsTile = ResultsTileData;

const AD_ID = '__AD_ID__';
const AD_SLOT = '__AD_SLOT__';

function parse(data: Results): [ResultsTile, string[]] {
  const results: ResultsTile = {};

  for (const subject in data.Attendance) {
    const Attendance = data.Attendance[subject];
    const GPA = data.GPAs[subject];
    results[subject] = { Attendance, GPA };
  }

  results.Overall = { Attendance: data.Attendance.Overall, GPA: data.GPA };

  return [
    results,
    Object.keys(results).sort((a, b) => {
      if (a === 'Overall') return 0;
      else if (b === 'Overall') return 1;

      for (let i = 0; i < 3; i++) {
        if (a.startsWith('*')) return 1;
        const sum = a.charCodeAt(i) - b.charCodeAt(i);
        if (sum !== 0) return sum;
      }

      return 0;
    })
  ];
}

function displayTiles(data: Results) {
  const ctx = document.querySelector('#renderContext')!;
  const semester = document.querySelector<HTMLSelectElement>('#semester')!;

  ctx.innerHTML = tableTile({
    subjects: parse(data),
    adId: AD_ID,
    adSlot: AD_SLOT,
    semester: parseInt(semester.value)
  });

  ctx.scrollIntoView({ behavior: 'smooth' });
}

export function showResults(data: Results) {
  const semester = document.querySelector<HTMLSelectElement>('#semester')!;

  semester.addEventListener('change', () => displayTiles(data));
  semester.style.display = 'flex';

  displayTiles(data);

  (adsbygoogle = (window as any).adsbygoogle || []).push({});
}
