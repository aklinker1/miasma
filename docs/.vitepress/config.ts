import { resolve } from 'path';
import { defineConfig } from 'vitepress';

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'Miasma',
  description: 'A cloud for your closet',

  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    logo: '/logo.svg',
    editLink: {
      pattern: 'https://github.com/aklinker1/miasma/edit/main/docs/:path',
    },

    nav: [
      { text: 'Guide', link: '/get-started' },
      { text: 'Config', link: '/config' },
    ],
    sidebar: [
      {
        text: 'Guide',
        items: [
          { text: 'Introduction', link: '/introduction' },
          { text: 'Get Started', link: '/get-started' },
          { text: 'Authentication', link: '/authentication' },
          { text: 'Plugins', link: '/plugins' },
          { text: 'Advanced Usage', link: '/advanced-usage' },
        ],
      },
      {
        text: 'Other',
        items: [
          { text: 'Contributing', link: '/contributing' },
          { text: 'Migrate to V2', link: '/v2-migration' },
        ],
      },
    ],

    socialLinks: [{ icon: 'github', link: 'https://github.com/aklinker1/miasma' }],
  },
});
