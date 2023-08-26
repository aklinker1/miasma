FROM node:18-alpine AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
RUN corepack prepare pnpm@8.6.12 --activate
WORKDIR /app
COPY package.json pnpm-lock.yaml ./

# Setup
FROM base AS build
RUN pnpm install --shamefully-hoist --frozen-lockfile
COPY . .
RUN pnpm build

FROM base
COPY --from=build /app/.output /app/.output
RUN ls node_modules && sleep 5s
EXPOSE 3000
CMD [ "node", ".output/server/index.mjs" ]
