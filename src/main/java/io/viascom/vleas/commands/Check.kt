package io.viascom.vleas.commands

import com.github.ajalt.clikt.core.CliktCommand
import com.github.ajalt.clikt.parameters.arguments.argument
import com.github.ajalt.clikt.parameters.types.file
import io.viascom.vleas.models.mavencentral.MavenCentralResponse
import io.viascom.vleas.services.MavenCentralApi


class Check : CliktCommand(help = "Check for new dependency updates") {

    private val file by argument("--file", help = "Full path to dependency file").file(
        exists = true,
        fileOkay = true,
        readable = true
    )


    override fun run() {
        echo("Check dependencies in file: " + file.name)
        val response: MavenCentralResponse = MavenCentralApi.retrieveDependency("ch.viascom.groundwork", "foxhttp")
            .getParsedBody(MavenCentralResponse::class.java)
        println("Version: " + response.response.docs[0].v)
    }
}