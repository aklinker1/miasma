import { VueQueryPlugin, VueQueryPluginOptions, QueryClient, hydrate, dehydrate } from 'vue-query';

export default defineNuxtPlugin(nuxt => {
  // Modify your Vue Query global settings here
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        // Stale time is only for the background to prevent double fetching
        // https://vue-query-next.vercel.app/#/guides/ssr?id=staleness-is-measured-from-when-the-query-was-fetched-on-the-server
        staleTime: 1e3,
      },
    },
  });
  const options: VueQueryPluginOptions = { queryClient };

  nuxt.vueApp.use(VueQueryPlugin, options);

  if (process.server) {
    nuxt.hooks.hook('app:rendered', () => {
      nuxt.payload['vue-query'] = dehydrate(queryClient);
    });
  }

  if (process.client) {
    nuxt.hooks.hook('app:created', () => {
      hydrate(queryClient, nuxt.payload['vue-query']);
    });
  }
});
