import { DefaultApi } from './api-gen';

export const Api = new DefaultApi(
  undefined,
  import.meta.env.VITE_SERVER_URL || window.location.origin,
);
