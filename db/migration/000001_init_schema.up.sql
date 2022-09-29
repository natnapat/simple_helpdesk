CREATE TABLE "tickets" (
    "id" bigserial PRIMARY KEY,
    "title" varchar NOT NULL,
    "description" varchar NOT NULL,
    "contact" varchar NOT NULL,
    "status" integer NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE FUNCTION tickets_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER tickets_updated
    BEFORE UPDATE
    ON
        tickets
    FOR EACH ROW
EXECUTE PROCEDURE tickets_updated_at();