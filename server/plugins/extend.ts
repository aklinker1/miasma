export default defineNitroPlugin(app => {
  const extended = useExtendedNitroApp(app);

  extended.auth = defineAuth(import.meta.env.VITE_AUTH);
});
