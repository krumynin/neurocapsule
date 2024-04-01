export const useSizes = () => {
  const outerwearClothesList = [{
    id: 0,
    title: '46/48',
  }, {
    id: 1,
    title: '48/50',
  }, {
    id: 3,
    title: '50/52',
  }, {
    id: 4,
    title: '52/54',
  }];
  const shirtSizeList = [{
    id: 0,
    title: '46/48',
  }, {
    id: 1,
    title: '48/50',
  }, {
    id: 3,
    title: '50/52',
  }, {
    id: 4,
    title: '52/54',
  }];
  const trousersSizeList = [{
    id: 0,
    title: '46',
  }, {
    id: 1,
    title: '48',
  }, {
    id: 3,
    title: '50',
  }, {
    id: 1,
    title: '52',
  }, {
    id: 3,
    title: '54',
  }];
  const jeansClothesList = [{
    id: 0,
    title: '30/L32',
  }, {
    id: 1,
    title: '31/L32 ',
  }, {
    id: 3,
    title: '32/L32 ',
  }, {
    id: 4,
    title: '33/L32 ',
  }];

  const outerwearSize = ref(null);
  const jeansSize = ref(null);
  const trousersSize = ref(null);
  const shirtSize = ref(null);

  return {
    outerwearClothesList,
    shirtSizeList,
    trousersSizeList,
    jeansClothesList,
    outerwearSize,
    jeansSize,
    trousersSize,
    shirtSize,
  };
};
