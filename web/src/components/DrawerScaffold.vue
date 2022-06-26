<script lang="ts" setup>
defineProps<{
  navTitle: string;
}>();
</script>

<template>
  <div class="drawer">
    <input id="my-drawer-3" type="checkbox" class="drawer-toggle" />
    <div class="drawer-content flex flex-col">
      <!-- Navbar -->
      <div class="w-full navbar bg-base-300">
        <div class="w-full max-w-screen-lg mx-auto">
          <div class="flex-none sm:hidden">
            <label for="my-drawer-3" class="btn btn-circle btn-ghost">
              <i-mdi-menu class="w-6 h-6" />
            </label>
          </div>
          <div class="px-2 mx-2 font-bold text-xl">
            <router-link to="/" class="link link-hover">{{
              navTitle
            }}</router-link>
          </div>
          <div class="flex-none hidden sm:block">
            <ul class="menu menu-horizontal rounded-box p-2 space-x-2">
              <slot name="nav-items" />
            </ul>
          </div>
        </div>
      </div>

      <div class="w-full max-w-screen-lg mx-auto py-8 px-16">
        <slot name="content" />
      </div>
    </div>
    <div class="drawer-side">
      <label for="my-drawer-3" class="drawer-overlay"></label>
      <ul class="menu p-4 overflow-y-auto w-80 bg-base-100">
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
