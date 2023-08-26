<script lang="ts" setup>
import { version } from '@@/package.json';
defineProps<{
  navTitle: string;
}>();
</script>

<template>
  <div class="drawer">
    <input id="drawer-toggle" type="checkbox" class="drawer-toggle" />
    <div class="drawer-content flex flex-col">
      <!-- Navbar -->
      <div class="w-full navbar bg-base-300 shrink-0">
        <div class="w-full max-w-screen-lg mx-auto">
          <div class="flex-none sm:hidden">
            <label for="drawer-toggle" class="btn btn-circle btn-ghost">
              <div class="i-mdi-menu text-2xl" />
            </label>
          </div>
          <div class="px-2 mx-2 font-bold text-xl">
            <nuxt-link to="/" class="link link-hover" :title="navTitle">
              <img src="~/assets/img/nav-branding.svg" class="w-32 aspect-[706/108]" />
            </nuxt-link>
          </div>
          <div class="flex-none hidden sm:block ml-4">
            <ul class="menu menu-horizontal rounded-box p-2 space-x-2">
              <slot name="nav-items" />
            </ul>
          </div>
        </div>
      </div>

      <div class="w-full max-w-screen-lg mx-auto p-8 space-y-8 min-h-[80vh] shrink-0">
        <slot name="content" />
      </div>

      <footer class="footer grid-rows-1 p-10 bg-neutral-focus text-neutral-content">
        <div>
          <span class="footer-title">Miasma {{ version }}</span>
          <a
            href="https://github.com/aklinker1/miasma"
            target="_blank"
            class="link link-hover flex items-center gap-2"
          >
            <span class="i-mdi-github text-lg inline-block" />
            <span>GitHub</span>
          </a>
          <a
            href="https://miasma.aklinker1.io/"
            target="_blank"
            class="link link-hover flex items-center gap-2"
          >
            <span class="i-mdi-book-open-page-variant text-lg inline-block" />
            <span>Documentation</span>
          </a>
          <a
            href="https://github.com/aklinker1/miasma"
            target="_blank"
            class="link link-hover flex items-center gap-2"
          >
            <span class="i-mdi-star text-lg inline-block" />
            <span>Leave a star!</span>
          </a>
        </div>
        <div>
          <span class="footer-title">Help</span>
          <a
            href="https://github.com/aklinker1/miasma/releases"
            target="_blank"
            class="link link-hover flex items-center gap-2"
          >
            <span class="i-mdi-creation text-lg inline-block" />
            <span>Changelog</span>
          </a>
          <a
            href="https://github.com/aklinker1/miasma/discussions/new?category=q-a"
            target="_blank"
            class="link link-hover flex items-center gap-2"
          >
            <span class="i-mdi-help-circle text-lg inline-block" />
            <span>Have a question?</span>
          </a>
          <a
            href="https://github.com/aklinker1/miasma/issues/new"
            target="_blank"
            class="link link-hover flex items-center gap-2"
          >
            <span class="i-mdi-ladybug text-lg inline-block" />
            <span>Report a bug</span>
          </a>
        </div>
      </footer>
    </div>
    <div class="drawer-side">
      <label for="drawer-toggle" class="drawer-overlay"></label>
      <ul class="menu p-4 overflow-y-auto w-80 bg-base-100 h-full">
        <nuxt-link to="/" class="link link-hover px-4 pt-4 pb-8" :title="navTitle">
          <img src="~/assets/img/nav-branding.svg" class="w-32 aspect-[706/108]" />
        </nuxt-link>

        <slot name="drawer-items" />
      </ul>
    </div>
  </div>
</template>

<style scoped>
a.router-link-active {
  @apply active;
}
</style>
