import { version } from '~~/package.json';

export default defineEventHandler(() => {
  return {
    status: 'UP',
    version,
    mode: import.meta.env.MODE,
  };
});
