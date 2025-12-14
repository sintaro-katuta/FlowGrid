FROM node:22-bookworm-slim

# Enable pnpm
RUN corepack enable

WORKDIR /app

# Copy package files
COPY package.json pnpm-lock.yaml ./

# Install dependencies
RUN pnpm install

# Copy source code (though volume mount usually overrides this in dev, good for building context)
COPY . .

EXPOSE 3000

CMD ["pnpm", "run", "dev", "--host"]
