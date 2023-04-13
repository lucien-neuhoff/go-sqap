## users

```sql
id (primary key)
uuid (unique identifier)
username
email
password (hashed)
created_at
updated_at
```

## sessions

```sql
id (primary key)
uuid (unique identifier)
user_id (foreign key to users table)
session_key (hashed)
created_at
updated_at
```

## public_keys

```sql
id (primary key)
user_id (foreign key to users table)
public_key
created_at
updated_at
```
