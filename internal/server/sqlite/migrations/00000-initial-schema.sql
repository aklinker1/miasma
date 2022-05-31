CREATE TABLE apps(
    "id" text PRIMARY KEY NOT NULL,
    "name" text NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "group" text,
    "image" text NOT NULL,
    "image_digest" text NOT NULL,
    "hidden" integer,
    "target_ports" blob,
    "published_ports" blob,
    "placement" blob,
    "volumes" blob,
    "networks" blob,
    "routing" blob,
    "command" text
);
