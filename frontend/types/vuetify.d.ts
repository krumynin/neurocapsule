type Color = {
  base: string
  lighten5: string
  lighten4: string
  lighten3: string
  lighten2: string
  lighten1: string
  darken1: string
  darken2: string
  darken3: string
  darken4: string
  accent1: string
  accent2: string
  accent3: string
  accent4: string
}

declare module 'vuetify/lib/util/colors.mjs' {
  const colors: {
    red: Readonly<Color>
    pink: Readonly<Color>
    purple: Readonly<Color>
    deepPurple: Readonly<Color>
    indigo: Readonly<Color>
    blue: Readonly<Color>
    lightBlue: Readonly<Color>
    cyan: Readonly<Color>
    teal: Readonly<Color>
    green: Readonly<Color>
    lightGreen: Readonly<Color>
    lime: Readonly<Color>
    yellow: Readonly<Color>
    amber: Readonly<Color>
    orange: Readonly<Color>
    deepOrange: Readonly<Color>
    brown: Readonly<Color>
    blueGrey: Readonly<Color>
    grey: Readonly<Color>
    shades: Readonly<{
      black: string
      white: string
      transparent: string
    }>
  };

  export default colors;
}
