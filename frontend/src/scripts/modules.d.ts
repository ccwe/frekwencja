declare module '*.pug' {
  export default function (locals?: { [local: string]: any }): string;
}
declare function ga(...args: string[]): void;
declare let adsbygoogle: any;
