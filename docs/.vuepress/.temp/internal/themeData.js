export const themeData = JSON.parse("{\"logo\":\"/nav-branding.svg\",\"docsRepo\":\"https://github.com/aklinker1/miasma\",\"docsDir\":\"docs\",\"editLink\":true,\"repo\":\"https://github.com/aklinker1/miasma\",\"navbar\":[{\"text\":\"Guide\",\"link\":\"/\"},{\"text\":\"Reference\",\"children\":[{\"text\":\"Miasma\",\"children\":[\"/reference/server-config.md\",\"/reference/cli.md\",\"/reference/graphql.md\"]}]}],\"sidebar\":{\"/\":[\"/index.md\",\"/guide/installation.md\",\"/guide/first-app.md\",\"/guide/port-management.md\",\"/guide/app-communication.md\",\"/guide/examples.md\",\"/authorization.md\",\"/guide/private-registries.md\",{\"text\":\"Plugins\",\"children\":[\"/plugins/index.md\",\"/plugins/traefik.md\"]},\"/guide/troubleshooting.md\",\"/faq.md\",\"/contributing.md\"],\"/reference/\":[\"/reference/server-config.md\",\"/reference/cli.md\",\"/reference/graphql.md\"]},\"locales\":{\"/\":{\"selectLanguageName\":\"English\"}},\"colorMode\":\"auto\",\"colorModeSwitch\":true,\"selectLanguageText\":\"Languages\",\"selectLanguageAriaLabel\":\"Select language\",\"sidebarDepth\":2,\"editLinkText\":\"Edit this page\",\"lastUpdated\":true,\"lastUpdatedText\":\"Last Updated\",\"contributors\":true,\"contributorsText\":\"Contributors\",\"notFound\":[\"There's nothing here.\",\"How did we get here?\",\"That's a Four-Oh-Four.\",\"Looks like we've got some broken links.\"],\"backToHome\":\"Take me home\",\"openInNewWindow\":\"open in new window\",\"toggleColorMode\":\"toggle color mode\",\"toggleSidebar\":\"toggle sidebar\"}")

if (import.meta.webpackHot) {
  import.meta.webpackHot.accept()
  if (__VUE_HMR_RUNTIME__.updateThemeData) {
    __VUE_HMR_RUNTIME__.updateThemeData(themeData)
  }
}

if (import.meta.hot) {
  import.meta.hot.accept(({ themeData }) => {
    __VUE_HMR_RUNTIME__.updateThemeData(themeData)
  })
}
