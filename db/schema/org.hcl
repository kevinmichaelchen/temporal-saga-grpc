schema "public" {
  comment = "standard public schema"
}

table "org" {
  schema = schema.public
  comment = "An organization"
  column "id" {
    comment = "UUID primary key"
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "name" {
    comment = "The organization's name"
    null = true
    type = varchar(100)
  }
  primary_key {
    columns = [column.id]
  }
}
