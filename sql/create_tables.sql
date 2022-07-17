-- Create posts table
CREATE TABLE IF NOT EXISTS posts (
    id int GENERATED BY DEFAULT AS IDENTITY,
    user_id integer NOT NULL,
    title text NOT NULL,
    body text NOT NULL
);