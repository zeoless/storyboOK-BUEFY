// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package com.example.authors.mysql

import java.sql.Connection
import java.sql.SQLException
import java.sql.Statement

const val createAuthor = """-- name: createAuthor :execresult
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ? 
)
"""

const val deleteAuthor = """-- name: deleteAuthor :exec
DELETE FROM authors
WHERE id = ?
"""

const val getAuthor = """-- name: getAuthor :one
SELECT id, name, bio FROM authors
WHERE id = ? LIMIT 1
"""

const val listAuthors = """-- name: listAuthors :many
SELECT id, name, bio FROM authors
ORDER BY name
"""

class QueriesImpl(private val conn: Connection) : Queries {

  @Throws(SQLException::class)
  override fun createAuthor(name: String, bio: String?): Long {
    return conn.prepareStatement(createAuthor, Statement.RETURN_GENERATED_KEYS).use { stmt ->
      stmt.setString(1, name)
          stmt.setString(2, bio)

      stmt.execute()

      val results = stmt.generatedKeys
      if (!results.next()) {
          throw SQLException("no generated key returned")
      }
	  results.getLong(1)
    }
  }

  @Throws(SQLException::class)
  override fun deleteAuthor(id: Long) {
    conn.prepareStatement(deleteAuthor).use { stmt ->
      stmt.