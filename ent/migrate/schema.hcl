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
    type = text
  }

  column "user_id" {
    null = true
    type = bigint
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "posts_user_id_fkey" {
    columns     = [column.user_id]
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
    null = true
    type = character_varying
  }

  column "body" {
    null = true
    type = text
  }

  column "price" {
    null = true
    type = integer
  }

  column "thoughts" {
    null = true
    type = text
  }

  primary_key {
    columns = [column.id]
  }
}
