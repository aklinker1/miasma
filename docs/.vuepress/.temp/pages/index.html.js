export const data = JSON.parse("{\"key\":\"v-8daa1a0e\",\"path\":\"/\",\"title\":\"Introduction\",\"lang\":\"en-US\",\"frontmatter\":{\"title\":\"Introduction\"},\"excerpt\":\"\",\"headers\":[{\"level\":2,\"title\":\"Features\",\"slug\":\"features\",\"children\":[{\"level\":3,\"title\":\"Future Work\",\"slug\":\"future-work\",\"children\":[]}]},{\"level\":2,\"title\":\"Not Features\",\"slug\":\"not-features\",\"children\":[]}],\"git\":{\"updatedTime\":1662527263000,\"contributors\":[{\"name\":\"Aaron Klinker\",\"email\":\"aaronklinker1@gmail.com\",\"commits\":2}]},\"filePathRelative\":\"index.md\"}")

if (import.meta.webpackHot) {
  import.meta.webpackHot.accept()
  if (__VUE_HMR_RUNTIME__.updatePageData) {
    __VUE_HMR_RUNTIME__.updatePageData(data)
  }
}

if (import.meta.hot) {
  import.meta.hot.accept(({ data }) => {
    __VUE_HMR_RUNTIME__.updatePageData(data)
  })
}
