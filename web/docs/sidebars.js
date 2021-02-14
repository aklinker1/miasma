module.exports = {
  'docs': [
    "introduction",
    "get-started",
    "first-app",
    {
      type: 'category',
      label: "Plugins",
      items: ["plugins/plugins", "plugins/traefik"],
      collapsed: false,
    },
    {
      type: 'category',
      label: "Common App Templates",
      items: ["templates/docker-registry", "templates/postgres", "templates/mongo"],
      collapsed: false,
    },
    "faq",
    "contributing",
  ],
  // 'docs/cli': [],
  // 'docs/server': [],
};
