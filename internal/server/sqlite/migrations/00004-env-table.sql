CREATE TABLE env(
    "app_id" text NOT NULL,
    "key" text NOT NULL,
    "value" text NOT NULL
);

CREATE UNIQUE INDEX "pk_env_app_id_key" ON env ("app_id", "key");
