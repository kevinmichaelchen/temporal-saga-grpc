schema "public" {
  comment = "standard public schema"
}

table "license" {
  schema = schema.public
  comment = "A license indicating some kind of purchase, valid for a specific window of time"
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
  column "start_time" {
    null = false
    type = timestamptz
  }
  column "end_time" {
    null = false
    type = timestamptz
  }
  column "user_id" {
    comment = "The license's user"
    null = false
    type = int
  }
  primary_key {
    columns = [column.id]
  }
}
