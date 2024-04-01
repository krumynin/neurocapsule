export const useGetRouteQuery = () => {
  const route = useRoute();

  const getQuery = <T, >(query: string) => {
    const _query = route.query[query];

    if (typeof _query === 'string') {
      return route.query[query] as T;
    }

    return undefined;
  };

  return getQuery;
};
