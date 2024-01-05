schema "public" {
  comment = "standard public schema"
}

table "profile" {
  schema = schema.public
  comment = "A user profile"
  column "id" {
    comment = "UUID primary key"
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "full_name" {
    comment = "The user's full name"
    null = true
    type = varchar(100)
  }
  column "org_id" {
    comment = "The user's organization"
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
}
