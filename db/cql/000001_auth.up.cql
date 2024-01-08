CREATE TABLE IF NOT EXISTS auth.token (
    user_id UUID,
    device_id UUID,
    user_agent TEXT,
    token TEXT,
    remote_ip TEXT,
    expired_date DATE,
    status TEXT,
    created_date DATE,
    PRIMARY KEY (user_id, device_id, user_agent)
)