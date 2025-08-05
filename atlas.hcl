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

# 段階的マイグレーションのマイグレーションファイル作成(schema.goを元に)
# atlas migrate diff <Migration名> \
#   --dir "file://ent/migrate/migrations" \
#   --to "ent://ent/schema" \
#   --dev-url "docker://postgres/15/test?search_path=public"

# 段階的マイグレーションのマイグレーションファイル作成(schema.hclを元に)
# atlas migrate diff <Migration名> --env local

# DBに適用
# atlas migrate apply --env local
