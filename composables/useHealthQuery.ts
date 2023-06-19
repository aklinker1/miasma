import { useQuery } from 'vue-query';

export default function () {
  return useQuery(QueryKeys.Health, () => $fetch('/api/health'));
}
