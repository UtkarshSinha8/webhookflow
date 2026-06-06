CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    merchant_id UUID NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE outbox_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    aggregate_id UUID NOT NULL,
    event_type TEXT NOT NULL,
    payload JSONB NOT NULL,

    status TEXT NOT NULL DEFAULT 'pending',

    attempt_count INTEGER NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_attempted_at TIMESTAMP
);

CREATE TABLE webhook_endpoints (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    merchant_id UUID NOT NULL,

    url TEXT NOT NULL,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    failure_count INTEGER NOT NULL DEFAULT 0,

    circuit_open BOOLEAN NOT NULL DEFAULT FALSE,

    circuit_opened_at TIMESTAMP
);

CREATE TABLE delivery_attempts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    outbox_event_id UUID NOT NULL,

    endpoint_id UUID NOT NULL,

    attempted_at TIMESTAMP NOT NULL DEFAULT NOW(),

    response_status INTEGER,

    success BOOLEAN NOT NULL,

    error_message TEXT
);
CREATE INDEX idx_outbox_pending
ON outbox_events(status, created_at);

CREATE INDEX idx_outbox_attempts
ON outbox_events(attempt_count);

CREATE INDEX idx_webhook_merchant
ON webhook_endpoints(merchant_id);

CREATE INDEX idx_delivery_event
ON delivery_attempts(outbox_event_id);