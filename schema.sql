-- Auto-generated PostgreSQL schema

CREATE TYPE message_status AS ENUM (
  'PENDING',
  'SENT',
  'DELIVERED',
  'READ'
);

CREATE TABLE accounts (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  user_id TEXT NOT NULL,
  type TEXT NULL,
  provider TEXT NULL,
  provider_account_id TEXT NULL,
  refresh_token TEXT NULL,
  access_token TEXT NULL,
  expires_at TIMESTAMPTZ NULL,
  token_type TEXT NULL,
  scope TEXT NULL,
  id_token TEXT NULL,
  session_state TEXT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE users (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  status SMALLINT NULL,
  approved_at TIMESTAMPTZ NULL,
  availability TEXT NULL,
  email TEXT NULL,
  username TEXT NULL,
  name TEXT NULL,
  first_name TEXT NULL,
  last_name TEXT NULL,
  password TEXT NULL,
  domain TEXT NULL,
  avatar TEXT NULL,
  phone_number TEXT NULL,
  country TEXT NULL,
  state TEXT NULL,
  city TEXT NULL,
  address TEXT NULL,
  zip_code TEXT NULL,
  gender TEXT NULL,
  date_of_birth TIMESTAMPTZ NULL,
  billing_id TEXT NULL,
  type TEXT NULL,
  email_verified_at TIMESTAMPTZ NULL,
  is_two_factor_enabled INTEGER NULL,
  two_factor_secret TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE ucodes (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  status SMALLINT NULL,
  user_id TEXT NULL,
  token TEXT NULL,
  email TEXT NULL,
  expired_at TIMESTAMPTZ NULL,
  PRIMARY KEY (id)
);

CREATE TABLE roles (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  status SMALLINT NULL,
  title TEXT NULL,
  name TEXT NULL,
  user_id TEXT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE permissions (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  status SMALLINT NULL,
  title TEXT NULL,
  action TEXT NULL,
  subject TEXT NULL,
  conditions TEXT NULL,
  fields TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE permission_roles (
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  permission_id TEXT NOT NULL,
  role_id TEXT NOT NULL,
  FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE,
  FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE TABLE role_users (
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  role_id TEXT NOT NULL,
  user_id TEXT NOT NULL,
  FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE notification_events (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  status SMALLINT NULL,
  type TEXT NULL,
  text TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE notifications (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  read_at TIMESTAMPTZ NULL,
  status SMALLINT NULL,
  sender_id TEXT NULL,
  receiver_id TEXT NULL,
  notification_event_id TEXT NULL,
  entity_id TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE user_payment_methods (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  user_id TEXT NULL,
  payment_method_id TEXT NULL,
  checkout_id TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE payment_transactions (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  store_id TEXT NULL,
  user_id TEXT NULL,
  order_id TEXT NULL,
  type TEXT NULL,
  withdraw_via TEXT NULL,
  provider TEXT NULL,
  reference_number TEXT NULL,
  status TEXT NULL,
  raw_status TEXT NULL,
  currency TEXT NULL,
  paid_currency TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE messages (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  sender_id TEXT NULL,
  receiver_id TEXT NULL,
  conversation_id TEXT NULL,
  attachment_id TEXT NULL,
  message TEXT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

CREATE TABLE attachments (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  name TEXT NULL,
  type TEXT NULL,
  size INTEGER NULL,
  file TEXT NULL,
  file_alt TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE conversations (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  creator_id TEXT NULL,
  participant_id TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE faqs (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  status SMALLINT NULL,
  sort_order INTEGER NULL,
  question TEXT NULL,
  answer TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE contacts (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  first_name TEXT NULL,
  last_name TEXT NULL,
  email TEXT NULL,
  phone_number TEXT NULL,
  message TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE social_medias (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  status SMALLINT NULL,
  sort_order INTEGER NULL,
  name TEXT NULL,
  url TEXT NULL,
  icon TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE website_infos (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  name TEXT NULL,
  phone_number TEXT NULL,
  email TEXT NULL,
  address TEXT NULL,
  logo TEXT NULL,
  favicon TEXT NULL,
  copyright TEXT NULL,
  cancellation_policy TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE settings (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  category TEXT NULL,
  label TEXT NULL,
  description TEXT NULL,
  key TEXT NULL,
  default_value TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE user_settings (
  id TEXT NOT NULL DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ NULL,
  user_id TEXT NULL,
  setting_id TEXT NULL,
  value TEXT NULL,
  PRIMARY KEY (id)
);

