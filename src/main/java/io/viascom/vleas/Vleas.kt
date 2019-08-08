package io.viascom.vleas

import com.github.ajalt.clikt.core.CliktCommand
import com.github.ajalt.clikt.core.subcommands
import io.viascom.vleas.commands.Check
import io.viascom.vleas.commands.Update

class Vleas : CliktCommand() {
    override fun run() = Unit
}

fun main(args: Array<String>) = Vleas().subcommands(Check(), Update()).main(args)