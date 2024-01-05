schema "public" {
  comment = "standard public schema"
}

table "org" {
  schema = schema.public
  comment = "An organization"
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
  column "name" {
    comment = "The organization's name"
    null = true
    type = varchar(100)
  }
  primary_key {
    columns = [column.id]
  }
}
