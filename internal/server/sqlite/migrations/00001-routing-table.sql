CREATE TABLE routes(
    "app_id" text PRIMARY KEY NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "host" text,
    "path" text,
    "traefik_rule" text
);
