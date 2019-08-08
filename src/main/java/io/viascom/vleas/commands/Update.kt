package io.viascom.vleas.commands

import com.github.ajalt.clikt.core.CliktCommand
import com.github.ajalt.clikt.parameters.arguments.argument
import com.github.ajalt.clikt.parameters.types.file

class Update : CliktCommand(help = "Update all dependencies") {

    private val file by argument("--file", help = "Full path to dependency file").file(
        exists = true,
        fileOkay = true,
        readable = true
    )

    override fun run() {
        echo("Update dependencies.")
    }
}