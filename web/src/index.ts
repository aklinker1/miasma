import { createApp } from 'vue';
import { RouterView } from 'vue-router';
import { router } from './router';
import './tailwind.css';

createApp(RouterView).use(router).mount('#app');
