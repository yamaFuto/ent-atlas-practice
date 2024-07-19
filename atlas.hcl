# ローカル環境の設定
env "local" {
  url = "postgres://postgres:postgres@:5432/ent-atlas-test?search_path=public&sslmode=disable"

  dev = "docker://postgres/15/dev?search_path=public"

  migration {
    dir = "file://ent/migrate/migrations"
  }

  src = "file://ent/migrate/schema.hcl"
}

# 開発環境の設定
env "dev" {
  url = "docker://postgres/15/dev?search_path=public"

  migration {
    dir = "file://ent/migrate/migrations"
  }

  src = "file://ent/migrate/schema.hcl"
}
