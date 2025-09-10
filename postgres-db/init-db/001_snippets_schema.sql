CREATE TABLE IF NOT EXISTS snippets (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    expires TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_snippets_created ON snippets (created);

INSERT INTO
    snippets (title, content, created, expires)
VALUES
    (
        'Не имей сто рублей',
        'Не имей сто рублей, а имей сто друзей.',
        NOW (),
        NOW () + INTERVAL '365 days'
    ),
    (
        'Лучше один раз увидеть',
        'Лучше один раз увидеть, чем сто раз услышать.',
        NOW (),
        NOW () + INTERVAL '365 days'
    ),
    (
        'Не откладывай на завтра',
        'Не откладывай на завтра, что можешь сделать сегодня.',
        NOW (),
        NOW () + INTERVAL '7 days'
    );
