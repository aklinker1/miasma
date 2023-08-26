FROM node:18-alpine AS base
WORKDIR /app

FROM base AS build
RUN corepack enable
RUN corepack prepare pnpm@8.6.12 --activate
COPY package.json pnpm-lock.yaml ./
RUN pnpm install --shamefully-hoist --frozen-lockfile
COPY . .
RUN pnpm build

FROM base
COPY --from=build /app/.output /app/.output
EXPOSE 3000
CMD [ "node", ".output/server/index.mjs" ]
