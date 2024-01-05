schema "public" {
  comment = "standard public schema"
}

table "profile" {
  schema = schema.public
  comment = "A user profile"
  column "id" {
    comment = "Numeric, auto-incrementing primary key"
    null = false
    type = int
    identity {
      generated = ALWAYS
      start = 1
      increment = 1
    }
  }
  column "full_name" {
    comment = "The user's full name"
    null = true
    type = varchar(100)
  }
  column "org_id" {
    comment = "The user's organization"
    null = false
    type = int
  }
  primary_key {
    columns = [column.id]
  }
}
