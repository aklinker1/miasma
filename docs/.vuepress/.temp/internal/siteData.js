export const siteData = JSON.parse("{\"base\":\"/\",\"lang\":\"en-US\",\"title\":\"Miasma\",\"description\":\"A Heroku-like, docker based PaaS with cluster and ARM support\",\"head\":[[\"link\",{\"rel\":\"icon\",\"href\":\"/favicon.svg\"}],[\"script\",{\"async\":true,\"defer\":true,\"data-website-id\":\"b29bfbde-9a3d-4550-92af-03f64684c08a\",\"src\":\"https://stats.aklinker1.io/umami.js\"}]],\"locales\":{}}")

if (import.meta.webpackHot) {
  import.meta.webpackHot.accept()
  if (__VUE_HMR_RUNTIME__.updateSiteData) {
    __VUE_HMR_RUNTIME__.updateSiteData(siteData)
  }
}

if (import.meta.hot) {
  import.meta.hot.accept(({ siteData }) => {
    __VUE_HMR_RUNTIME__.updateSiteData(siteData)
  })
}
