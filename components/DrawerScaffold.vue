<script lang="ts" setup>
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

      <div class="bg-neutral-focus w-full p-8 -z-10">
        <footer class="footer max-w-screen-md mx-auto">
          <div>
            <span class="footer-title">Links</span>
            <a href="https://aklinker1.github.io/miasma" target="_blank" class="link link-hover"
              >Documentation</a
            >
            <a href="https://github.com/aklinker1/miasma" target="_blank" class="link link-hover"
              >GitHub</a
            >
          </div>
          <div>
            <span class="footer-title">Help</span>
            <a
              href="https://github.com/aklinker1/miasma/discussions/new?category=q-a"
              target="_blank"
              class="link link-hover"
              >Have a question?</a
            >
            <a
              href="https://github.com/aklinker1/miasma/issues/new"
              target="_blank"
              class="link link-hover"
              >Report a Bug</a
            >
          </div>
        </footer>
      </div>
    </div>
    <div class="drawer-side">
      <label for="drawer-toggle" class="drawer-overlay"></label>
      <ul class="menu p-4 overflow-y-auto w-80 bg-base-100">
        <slot name="drawer-items" />
      </ul>
    </div>
  </div>
</template>

<style scoped>
a.nuxt-link-active {
  @apply active;
}
</style>
