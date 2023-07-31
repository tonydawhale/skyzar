FROM node:19

RUN npm install -g pnpm
WORKDIR /app

COPY package.json ./
COPY pnpm-lock.yaml ./

RUN pnpm install

COPY . .

RUN pnpm run build

EXPOSE 23572

CMD ["pnpm", "run", "deploy"]