import json from '@rollup/plugin-json';
import replace from '@rollup/plugin-replace';
import ts from '@rollup/plugin-typescript';
import { randomBytes } from 'crypto';
import { readFileSync } from 'fs';
import pug from 'pug';
import htmlmin from 'rollup-plugin-html-minifier';
import html2 from 'rollup-plugin-html2';
import live from 'rollup-plugin-livereload';
import polyfills from 'rollup-plugin-polyfill-node';
import pugm from 'rollup-plugin-pug';
import scss from 'rollup-plugin-scss';
import serve from 'rollup-plugin-serve';
import { terser } from 'rollup-plugin-terser';
import { parse } from 'yaml';

const hash = randomBytes(7).toString('hex');

const production = process.env.NODE_ENV == 'production';

const config = parse(readFileSync('./config.env.yaml').toString());

export default {
  input: 'src/scripts/index.ts',
  output: {
    file: `dist/index.${hash}.js`,
    format: 'iife',
    sourcemap: !production
  },
  plugins: [
    polyfills(),
    json(),
    ts(),
    pugm(),
    replace({
      values: {
        __AD_ID__: config.adId,
        __AD_SLOT__: config.adSlot
      },
      include: '**/*/results.ts'
    }),
    scss({
      sass: require('sass'),
      output: `dist/index.${hash}.css`
    }),
    html2({
      template: pug.renderFile('./src/index.pug', {
        analyticsId: config.analyticsId,
        adId: config.adId,
        adSlot: config.adSlot
      }),
      fileName: 'index.html',
      externals: { after: [{ tag: 'link', href: `index.${hash}.css` }] }
    }),
    htmlmin({
      collapseWhitespace: true,
      minifyJS: true
    }),
    !production && serve('dist'),
    !production && live(),
    production && terser()
  ]
};
