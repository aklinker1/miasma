module.exports = {
  title: 'Miasma',
  tagline: 'Open source, Docker Swarm based PaaS with ARM support',
  url: 'https://aklinker1.github.io',
  baseUrl: '/miasma/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'aklinker1',
  projectName: 'miasma',
  themeConfig: {
    navbar: {
      title: 'Miasma',
      // logo: {
      //   alt: 'My Site Logo',
      //   src: 'img/logo.svg',
      // },
      items: [
        {
          to: '/docs',
          activeBasePath: '/docs',
          label: 'Documentation',
          position: 'left',
        },
        {
          to: '/docs/server',
          activeBasePath: '/docs/server',
          label: 'Server API',
          position: 'right',
        },
        {
          to: '/docs/cli',
          activeBasePath: '/docs/cli',
          label: 'CLI Usage',
          position: 'right',
        },
        {
          href: 'https://github.com/aklinker1/miasma',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'GitHub',
          items: [
            {
              label: 'Repository',
              to: 'https://github.com/aklinker1/miasma',
            },
            {
              label: 'Milestones',
              to: 'https://github.com/aklinker1/miasma/milestones',
            },
          ],
        },
        {
          title: 'Help',
          items: [
            {
              label: 'Issues',
              href: 'https://github.com/aklinker1/miasma/issues',
            },
            {
              label: 'Discussions',
              href: 'https://github.com/aklinker1/miasma/discussions',
            },
            {
              label: 'FAQ',
              href: '/docs/faq',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Aaron Klinker`,
    },
    colorMode: {
      defaultMode: "dark",
      disableSwitch: true,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl:
            'https://github.com/aklinker1/miasma/edit/main/docs/docs/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
