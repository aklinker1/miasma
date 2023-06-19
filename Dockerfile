# Setup
FROM node:18-alpine
RUN corepack enable
RUN corepack prepare pnpm@7.15.0 --activate

# Install
WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile --shamefully-hoist

# Build
COPY . .
RUN pnpm build
EXPOSE 3000
CMD [ "node", ".output/server/index.mjs" ]
