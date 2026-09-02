-- Add some dummy records
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Old Silent Pond',
    'E: An old silent pond...
    A frog jumps into the pond,
    splash! Silence again. 
    
    - Matsuo Bashō',
    NOW() AT TIME ZONE 'utc',
    NOW() AT TIME ZONE 'utc' + INTERVAL '365 days'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over The Wintry Forest',
    'Over the wintry
    forest, winds howl in rage
    with no leaves to blow. 
    - Natsume Soseki',
    NOW() AT TIME ZONE 'utc',
    NOW() AT TIME ZONE 'utc' + INTERVAL '365 days'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First Autumn Morning',
    'First autumn morning
    the mirror I stare into
    shows my father''s face.

    - Murakami Kijo',
    NOW() AT TIME ZONE 'utc',
    NOW() AT TIME ZONE 'utc' + INTERVAL '7 days'
);