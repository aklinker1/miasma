import { version } from '~~/package.json';

export default defineEventHandler(() => {
  const app = useExtendedNitroApp();
  return {
    status: 'UP',
    version,
    mode: import.meta.env.MODE,
    auth: app.auth.type,
  };
});
