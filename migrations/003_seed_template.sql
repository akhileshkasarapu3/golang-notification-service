INSERT INTO templates (name, subject_template, body_template)
VALUES (
    'welcome_email',
    'Welcome, {{.name}}',
    'Hello {{.name}}, your {{.product}} is ready'
)
ON CONFLICT (name) DO NOTHING;