package com.example.dbtest

import org.junit.jupiter.api.extension.AfterEachCallback
import org.junit.jupiter.api.extension.BeforeEachCallback
import org.junit.jupiter.api.extension.ExtensionContext
import java.nio.file.Files
import java.nio.file.Paths
import java.sql.Connection
import java.sql.DriverManager
import kotlin.streams.toList


class MysqlDbTestExtension(private val migrationsPath: String) : BeforeEachCallback, AfterEachCallback {
    val user = System.getenv("MYSQL_USER") ?: "root"
    val pass = System.getenv("MYSQL_ROOT_PASSWORD") ?: "mysecretpassword"
    val host = System.getenv("MYSQL_HOST") ?: "127.0.0.1"
    val port = System.getenv("MYSQL_PORT") ?: "3306"
    val mainDb = System.getenv("MYSQL_DATABASE") ?: "dinotest"
    val testDb = "sqltest_mysql"

    override fun beforeEach(context: ExtensionContext) {
        getConnection(mainDb).createStatement().execute("CREATE DATABASE $testDb")
        val path = Paths.get(migrationsPath)
        val migrations = if (F