import tableTile from '../../templates/results/tile.pug';

type Results = { Name: string; Attendance: number[][] }[];

interface ResultTileData {
  subjectName: string;
  attendance: [O: number, U: number, NU: number][];
  percentage: number[];
  isBelowHalf: boolean[];
}

const AD_ID = '__AD_ID__';
const AD_SLOT = '__AD_SLOT__';

function parse(data: Results): ResultTileData[] {
  const results: ResultTileData[] = [];

  for (const lesson of data) {
    const attendance = lesson.Attendance;
    const tile: ResultTileData = {
      subjectName: lesson.Name,
      attendance: [],
      percentage: [],
      isBelowHalf: []
    };

    attendance.push(Array<number>(5).fill(0));
    for (let i = 0; i < 5; i += 1)
      attendance[2][i] = attendance[0][i] + attendance[1][i];
    attendance.forEach((item) => {
      const present = item[0] + item[2] + item[4];
      const all = present + item[3] + item[1];
      const percentage = Math.floor((present / all) * 1000) / 10;

      tile.attendance.push([present, item[3], item[1]]);
      tile.percentage.push(percentage);
      tile.isBelowHalf.push(percentage < 50);
    });

    results.push(tile);
  }

  return results.sort((a, b) => {
    for (let i = 0; i < 3; i++) {
      if (a.subjectName.startsWith('*')) return 1;
      const sum = a.subjectName.charCodeAt(i) - b.subjectName.charCodeAt(i);
      if (sum !== 0) return sum;
    }
    return 0;
  });
}

export function displayTiles(data: Results, includeExtracurricular: boolean) {
  const ctx = document.querySelector('#renderContext')!;
  const semester = document.querySelector<HTMLSelectElement>('#semester')!;

  ctx.innerHTML = tableTile({
    subjects: parse(
      includeExtracurricular
        ? data
        : data.filter(({ Name }) => !Name.startsWith('*'))
    ),
    id: AD_ID,
    slot: AD_SLOT,
    semester: parseInt(semester.value)
  });

  ctx.scrollIntoView({ behavior: 'smooth' });
}
