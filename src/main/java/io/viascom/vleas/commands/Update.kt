package io.viascom.vleas.commands

import com.github.ajalt.clikt.core.CliktCommand

class Update: CliktCommand(help="Update all dependencies") {
    override fun run() {

        echo("Update dependencies.")
    }
}