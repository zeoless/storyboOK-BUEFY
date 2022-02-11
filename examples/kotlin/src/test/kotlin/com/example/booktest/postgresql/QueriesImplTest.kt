package com.example.booktest.postgresql

import com.example.dbtest.PostgresDbTestExtension
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.extension.RegisterExtension
import java.time.OffsetDateTime
import java.time.format.DateTimeFormatter

class QueriesImplTest {
    companion object {
        @JvmField @RegisterExtension val dbtest = PostgresDbTestExtension("src/main/resources/booktest/postgresql/schema.sql")
    }

    @Test
    fun testQueries() {
        val conn = dbtest.getConnection()
        val db = QueriesImpl(conn)
        val author = db.createAuthor("Unknown Master")!!

        // Start a transaction
        conn.autoCommit = false
        db.createBook(
                authorId = author.authorId,
                isbn = "1",
                title = "my book title",
                bookType = BookType.NONFICTION,
                year = 2016,
                available = OffsetDateTime.now(),
                tags = lis