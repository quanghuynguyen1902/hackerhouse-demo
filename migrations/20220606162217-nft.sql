-- +migrate Up
create table "nft_collection" (
  "id" serial primary key,
  "created_time" timestamp not null default now(),
  "last_updated_time" timestamp not null default now(),
  "address" text not null,
  "name" text,
  "symbol" text,
  "chain_id" integer,
  "supply" integer,
  "erc_format" text,
  unique("address", "chain_id")
);

create table "nft_token" (
  "id" serial primary key,
  "token_id" text,
  "created_time" timestamp not null default now(),
  "last_updated_time" timestamp not null default now(),
  "collection_address" text,
  "name" text,
  "description" text,
  "amount" integer,
  "image" text,
  "image_cdn" text,
  "thumbnail_cdn" text,
  "image_content_type" text,
  "rarity_rank" integer,
  "rarity_score" text,
  "ranking" integer,
	unique ("collection_address", "token_id")
);

create table "nft_token_attribute" (
  "id" serial primary key,
  "created_time" timestamp not null default now(),
  "last_updated_time" timestamp not null default now(),
  "collection_address" text,
  "token_id" text,
  "trait_type" text,
  "value" text,
  "count" integer,
  "rarity" text,
  "frequency" text,
	unique ("collection_address", "token_id", "trait_type")
);

create table "nft_owner" (
  "id" serial primary key,
  "created_time" timestamp not null default now(),
  "last_updated_time" timestamp not null default now(),
  "owner_address" text,
  "collection_address" text,
  "token_id" integer
);

alter table nft_token add column rarity_tier text;

alter table nft_token drop column ranking;

ALTER TABLE
	nft_token
ADD
	COLUMN is_self_hosted BOOL DEFAULT FALSE;

ALTER TABLE
	nft_collection
ADD
	COLUMN IF NOT EXISTS image TEXT;

alter table nft_token_attribute add column chain_id integer;

ALTER TABLE
	nft_collection
ADD
	COLUMN description TEXT,
ADD
	COLUMN contract_scan TEXT,
ADD
	COLUMN discord TEXT,
ADD
	COLUMN twitter TEXT,
ADD
	COLUMN website TEXT;

CREATE TABLE IF NOT EXISTS solana_nft_mapping (
  id serial,
  token_id text NOT NULL,
  collection_address text NOT NULL,
  mint_id text NOT NULL,
  PRIMARY KEY (id)
);

create index idx_solana_nft_mapping_token_id on solana_nft_mapping (token_id);
create index idx_solana_nft_mapping_collection_address on solana_nft_mapping (collection_address);
create index idx_solana_nft_mapping_mint_id on solana_nft_mapping (mint_id);

-- +migrate Down
drop table if exists nft_token_attribute;
drop table if exists nft_token;
drop table if exists nft_collection;
drop table if exists nft_owner;
