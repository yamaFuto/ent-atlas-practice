schema "public" {
  comment = "standard public schema"
}

table "users" {
  schema = schema.public

  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }

  column "name" {
    null = false
    type = character_varying
  }

  column "age" {
    null = false
    type = bigint
  }

  column "email" {
    null = false
    type = character_varying
  }

  column "description" {
    null = false
    type = character_varying
  }

  primary_key {
    columns = [column.id]
  }

  index "users_email_key" {
    unique = true
    columns = [column.email]
  }
}

table "posts" {
  schema = schema.public

  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }

  column "title" {
    null = false
    type = character_varying
  }

  column "content" {
    null = false
    type = character_varying
  }

  column "user_posts" {
    null = true
    type = bigint
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "posts_users_posts" {
    columns     = [column.user_posts]
    ref_columns = [table.users.column.id]
    on_delete   = SET_NULL
  }
}

table "books" {
  schema = schema.public

  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }

  column "title" {
    null = false
    type = character_varying
  }

  column "body" {
    null = false
    type = text
  }

  column "price" {
    null = false
    type = int8
  }

  column "thoughts" {
    null = false
    type = text
  }

  primary_key {
    columns = [column.id]
  }
}
