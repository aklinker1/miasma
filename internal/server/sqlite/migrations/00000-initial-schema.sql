CREATE TABLE apps(
    "id" text PRIMARY KEY NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "name" text NOT NULL,
    "group" text,
    "image" text NOT NULL,
    "image_digest" text NOT NULL,
    "hidden" boolean NOT NULL,
    "routing" blob,
    "target_ports"  blob,
    "published_ports" blob,
    "placement" blob,
    "volumes" blob,
    "networks" blob,
    "command" text
);

CREATE UNIQUE INDEX idx_apps_name ON apps (name);
