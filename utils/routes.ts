export const routes = {
  login: '/login',

  services: '/services',
  service: (id: string | undefined) => `/services/${id}`,

  nodes: '/nodes',
  node: (id: string | undefined) => `/nodes/${id}`,

  plugins: '/plugins',
  traefikPlugin: '/plugins/traefik',
};
