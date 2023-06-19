import { UseQueryOptions, useQuery } from 'vue-query';

export default function <TData, TError extends H3Error<any>>(
  options: UseQueryOptions<TData, TError, TData>,
) {
  return useQuery({
    retry(failureCount, error) {
      return error.statusCode >= 500 && failureCount < 3;
    },
    ...options,
  });
}
